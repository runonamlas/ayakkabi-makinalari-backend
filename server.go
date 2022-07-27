package main

import (
	"github.com/gin-gonic/gin"
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
			println("hereerre")
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
func main() {
	defer config.CloseDatabaseConnection(db)
	//gin.SetMode(gin.ReleaseMode)
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

	err := r.Run()
	if err != nil {
		println(err)
		return
	}
}
