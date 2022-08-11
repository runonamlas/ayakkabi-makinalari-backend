package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/runonamlas/ayakkabi-makinalari-backend/config"
	"github.com/runonamlas/ayakkabi-makinalari-backend/controller"
	"github.com/runonamlas/ayakkabi-makinalari-backend/middleware"
	"github.com/runonamlas/ayakkabi-makinalari-backend/repository"
	"github.com/runonamlas/ayakkabi-makinalari-backend/service"
)

var (
	db         = config.SetupDatabaseConnection()
	jwtService = service.NewJWTService()

	authService    = service.NewAuthService(userRepository)
	authController = controller.NewAuthController(authService, jwtService)

	userRepository = repository.NewUserRepository(db)
	userService    = service.NewUserService(userRepository)
	userController = controller.NewUserController(userService, authService, jwtService)

	adminRepository = repository.NewAdminRepository(db)
	adminService    = service.NewAdminService(adminRepository)
	adminController = controller.NewAdminController(adminService, jwtService)

	messageRepository = repository.NewMessageRepository(db)
	messageService    = service.NewMessageService(messageRepository)
	messageController = controller.NewMessageController(messageService, jwtService)

	productRepository = repository.NewProductRepository(db)
	productService    = service.NewProductService(productRepository)
	productController = controller.NewProductController(productService, jwtService)

	productCategoryRepository = repository.NewProductCategoryRepository(db)
	productCategoryService    = service.NewProductCategoryService(productCategoryRepository)
	productCategoryController = controller.NewProductCategoryController(productCategoryService, jwtService)
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 4096
)

type message struct {
	data []byte
	room string
}

type subscription struct {
	conn *connection
	room string
}
type connection struct {
	ws   *websocket.Conn
	send chan []byte
}

type hub struct {
	rooms      map[string]map[*connection]bool
	broadcast  chan message
	register   chan subscription
	unregister chan subscription
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var h = hub{
	broadcast:  make(chan message),
	register:   make(chan subscription),
	unregister: make(chan subscription),
	rooms:      make(map[string]map[*connection]bool),
}

func (h *hub) run() {
	for {
		select {
		case s := <-h.register:
			connections := h.rooms[s.room]
			if connections == nil {
				connections = make(map[*connection]bool)
				h.rooms[s.room] = connections
			}
			h.rooms[s.room][s.conn] = true
		case s := <-h.unregister:
			connections := h.rooms[s.room]
			if connections != nil {
				if _, ok := connections[s.conn]; ok {
					delete(connections, s.conn)
					close(s.conn.send)
					if len(connections) == 0 {
						delete(h.rooms, s.room)
					}
				}
			}
		case m := <-h.broadcast:
			connections := h.rooms[m.room]
			for c := range connections {
				select {
				case c.send <- m.data:
				default:
					close(c.send)
					delete(connections, c)
					if len(connections) == 0 {
						delete(h.rooms, m.room)
					}
				}
			}
		}
	}
}

func (s subscription) readPump() {
	c := s.conn
	defer func() {
		h.unregister <- s
		c.ws.Close()
	}()
	c.ws.SetReadLimit(maxMessageSize)
	c.ws.SetReadDeadline(time.Now().Add(pongWait))
	c.ws.SetPongHandler(func(string) error { c.ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, msg, err := c.ws.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				fmt.Printf("error: %v", err)
			}
			break
		}
		m := message{msg, s.room}
		h.broadcast <- m
	}
}

func (s *subscription) writePump() {
	c := s.conn
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.ws.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.write(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.write(websocket.TextMessage, message); err != nil {
				return
			}
		case <-ticker.C:
			if err := c.write(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}
func (c *connection) write(mt int, payload []byte) error {
	c.ws.SetWriteDeadline(time.Now().Add(writeWait))
	return c.ws.WriteMessage(mt, payload)
}

func serveWs(w http.ResponseWriter, r *http.Request, roomId string) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	c := &connection{send: make(chan []byte, 256), ws: ws}
	s := subscription{c, roomId}
	h.register <- s
	go s.writePump()
	go s.readPump()
}

func main() {
	go h.run()
	defer config.CloseDatabaseConnection(db)
	if os.Getenv("CLOUD") == "true" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(CORSMiddleware())
	r.Use(gin.Logger())

	authRoutes := r.Group("auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
		authRoutes.POST("/forget", authController.Forget)
		authRoutes.POST("/change", authController.Change)
	}

	adminRoutes := r.Group("admin")
	{
		adminRoutes.POST("/login", adminController.Login)
		adminRoutes.POST("/register", adminController.Register)
		adminRoutes.GET("/profile", adminController.Profile, middleware.AuthorizeJWT(jwtService))
		adminRoutes.PUT("/profile", adminController.Update, middleware.AuthorizeJWT(jwtService))
		adminRoutes.GET("/users", adminController.Users, middleware.AuthorizeJWT(jwtService))
	}

	userRoutes := r.Group("user")
	{
		userRoutes.GET("/profile/:id", userController.Profile)
		userRoutes.PUT("/profile", userController.Update, middleware.AuthorizeJWT(jwtService))
		userRoutes.GET("/statistic", userController.Statistic, middleware.AuthorizeJWT(jwtService))
		userRoutes.GET("/products", userController.GetProducts, middleware.AuthorizeJWT(jwtService))
		userRoutes.GET("/messages", userController.GetMessages, middleware.AuthorizeJWT(jwtService))
		//userRoutes.POST("/favourite/:id", userController.AddFavourite)
		//userRoutes.DELETE("/favourite/:id", userController.DeleteFavourite)
	}

	productRoutes := r.Group("products")
	{
		productRoutes.GET("/", productController.All)
		productRoutes.POST("/add", productController.Insert, middleware.AuthorizeJWT(jwtService))
		productRoutes.GET("/:id", productController.FindByID)
		productRoutes.GET("/category/:id", productController.FindByCategory)
		productRoutes.PUT("/", productController.Update, middleware.AuthorizeJWT(jwtService))
		productRoutes.DELETE("/:id", productController.Delete, middleware.AuthorizeJWT(jwtService))
	}

	productCategoryRoutes := r.Group("product-categories")
	{
		productCategoryRoutes.GET("/", productCategoryController.All)
		productCategoryRoutes.POST("/", productCategoryController.Insert)
		productCategoryRoutes.GET("/:id", productCategoryController.FindByID)
		//productCategoryRoutes.PUT("/", productCategoryController.Update)
		//productCategoryRoutes.DELETE("/:id", productCategoryController.Delete)
	}

	messageRoutes := r.Group("messages", middleware.AuthorizeJWT(jwtService))
	{
		messageRoutes.GET("/:id", messageController.All)
		messageRoutes.POST("/", messageController.Insert)
		//messageRoutes.GET("/:id", messageController.FindByID)
		messageRoutes.PUT("/", messageController.Update)
		messageRoutes.DELETE("/:id", messageController.Delete)
	}

	r.GET("/ws/:roomId", func(c *gin.Context) {
		roomId := c.Param("roomId")
		serveWs(c.Writer, c.Request, roomId)
	})

	err := r.Run()
	if err != nil {
		println(err)
		return
	}
}
