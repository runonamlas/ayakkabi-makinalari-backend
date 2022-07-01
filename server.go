package main

import (
	"github.com/gin-gonic/gin"
	"github.com/my-way-teams/my_way_backend/config"
	"github.com/my-way-teams/my_way_backend/controller"
	"github.com/my-way-teams/my_way_backend/middleware"
	"github.com/my-way-teams/my_way_backend/repository"
	"github.com/my-way-teams/my_way_backend/service"
)

var (
	db = config.SetupDatabaseConnection()
	jwtService = service.NewJWTService()

	userRepository = repository.NewUserRepository(db)
	userService = service.NewUserService(userRepository)
	userController = controller.NewUserController(userService, jwtService)

	adminRepository = repository.NewAdminRepository(db)
	adminService = service.NewAdminService(adminRepository)
	adminController = controller.NewAdminController(adminService, jwtService)

	authService = service.NewAuthService(userRepository)
	authController = controller.NewAuthController(authService, jwtService)

	countryRepository = repository.NewCountryRepository(db)
	countryService = service.NewCountryService(countryRepository)
	countryController = controller.NewCountryController(countryService, jwtService)

	cityRepository = repository.NewCityRepository(db)
	cityService = service.NewCityService(cityRepository)
	cityController = controller.NewCityController(cityService, jwtService)

	placeRepository = repository.NewPlaceRepository(db)
	placeService = service.NewPlaceService(placeRepository)
	placeController = controller.NewPlaceController(placeService, jwtService)

	placeCategoryRepository = repository.NewPlaceCategoryRepository(db)
	placeCategoryService = service.NewPlaceCategoryService(placeCategoryRepository)
	placeCategoryController = controller.NewPlaceCategoryController(placeCategoryService, jwtService)

	routeRepository = repository.NewRouteRepository(db)
	routeService = service.NewRouteService(routeRepository)
	routeController = controller.NewRouteController(routeService, jwtService)
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
func main() {
	defer config.CloseDatabaseConnection(db)
	//gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(CORSMiddleware())
	r.Use(gin.Logger())

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	adminRoutes := r.Group("api/admin")
	{
		adminRoutes.POST("/login", adminController.Login)
		adminRoutes.POST("/register", adminController.Register)
		adminRoutes.GET("/profile", adminController.Profile, middleware.AuthorizeJWT(jwtService))
		adminRoutes.PUT("/profile", adminController.Update, middleware.AuthorizeJWT(jwtService))
		adminRoutes.GET("/users", adminController.Users, middleware.AuthorizeJWT(jwtService))
	}

	userRoutes := r.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.PUT("/profile", userController.Update)
		userRoutes.GET("/favourite", userController.GetFavourites)
		userRoutes.POST("/favourite/:id", userController.AddFavourite)
		userRoutes.DELETE("/favourite/:id", userController.DeleteFavourite)
		userRoutes.GET("/saved", userController.GetSaved)
		userRoutes.POST("/saved/:id", userController.AddSaved)
		userRoutes.DELETE("/saved/:id", userController.DeleteSaved)
	}

	countryRoutes := r.Group("api/countries", middleware.AuthorizeJWT(jwtService))
	{
		countryRoutes.GET("/", countryController.All)
		countryRoutes.POST("/", countryController.Insert)
		countryRoutes.GET("/:id", countryController.FindByID)
		countryRoutes.PUT("/", countryController.Update)
		countryRoutes.DELETE("/:id", countryController.Delete)
	}

	cityRoutes := r.Group("api/cities", middleware.AuthorizeJWT(jwtService))
	{
		cityRoutes.GET("/", cityController.All)
		cityRoutes.POST("/", cityController.Insert)
		cityRoutes.GET("/:id", cityController.FindByID)
		cityRoutes.PUT("/", cityController.Update)
		cityRoutes.DELETE("/:id", cityController.Delete)
	}

	placeRoutes := r.Group("api/places", middleware.AuthorizeJWT(jwtService))
	{
		placeRoutes.GET("/", placeController.All)
		placeRoutes.POST("/", placeController.Insert)
		placeRoutes.GET("/:id", placeController.FindByID)
		placeRoutes.PUT("/", placeController.Update)
		placeRoutes.DELETE("/:id", placeController.Delete)
	}

	placeCategoryRoutes := r.Group("api/place-categories", middleware.AuthorizeJWT(jwtService))
	{
		placeCategoryRoutes.GET("/", placeCategoryController.All)
		placeCategoryRoutes.POST("/", placeCategoryController.Insert)
		placeCategoryRoutes.GET("/:id", placeCategoryController.FindByID)
		placeCategoryRoutes.PUT("/", placeCategoryController.Update)
		placeCategoryRoutes.DELETE("/:id", placeCategoryController.Delete)
	}

	routeRoutes := r.Group("api/routes", middleware.AuthorizeJWT(jwtService))
	{
		routeRoutes.GET("/", routeController.All)
		routeRoutes.POST("/", routeController.Insert)
		routeRoutes.GET("/:id", routeController.FindByID)
		routeRoutes.PUT("/", routeController.Update)
		routeRoutes.DELETE("/:id", routeController.Delete)
	}

	err := r.Run()
	if err != nil {
		return
	}
}