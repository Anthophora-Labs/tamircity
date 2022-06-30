package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mustafakocatepe/Tamircity/handler/api"
	"github.com/mustafakocatepe/Tamircity/pkg/service"
	"github.com/mustafakocatepe/Tamircity/pkg/store/repositories"
	"log"

	"github.com/joho/godotenv"
	dbModels "github.com/mustafakocatepe/Tamircity/pkg/models/db"
	postgres "github.com/mustafakocatepe/Tamircity/pkg/store/shared/db"
)

func main() {

	// Set enviroment variables
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db, err := postgres.NewPsqlDB()
	if err != nil {
		log.Fatal("Postgres cannot init", err)
	}
	log.Println("Postgres connected")
	/**/
	db.AutoMigrate(&dbModels.TechnicalService{}, &dbModels.Brand{}, &dbModels.Model{}, &dbModels.FixType{}, &dbModels.DeviceType{})
	/*
		//Add Seed data
		roleRepo := role.NewRoleRepository(db)
		roleRepo.Seed()
		userRepo := user.NewUserRepository(db)
		userRepo.Seed()
		statusRepo := status.NewStatusRepository(db)
		statusRepo.Seed()
		users, _ := userRepo.FindAll()
		roles, _ := roleRepo.FindAll()
		userrolemapRepo := userrolemap.NewUserRoleMapRepository(db)
		userrolemapRepo.Seed(users, roles)
	*/
	log.Println("Migrations done")

	// Store
	technicalServiceStore := repositories.NewTechnicalServiceStore(db)
	serviceTypeStore := repositories.NewServiceTypeStore(db)
	extraServiceStore := repositories.NewExtraServiceStore(db)
	brandStore := repositories.NewBrandStore(db)
	modelStore := repositories.NewModelStore(db)
	fixTypeStore := repositories.NewFixTypeStore(db)
	deviceTypeStore := repositories.NewDeviceTypeStore(db)

	// Clients
	// This one need to be integrated systems

	// Service
	technicalServiceService := service.NewTechnicalServiceService(technicalServiceStore)
	serviceTypeService := service.NewServiceTypeService(serviceTypeStore)
	extraServiceService := service.ExtraServiceService(extraServiceStore)
	brandService := service.NewBrandService(brandStore)
	modelService := service.NewModelService(modelStore)
	fixTypeService := service.NewFixTypeService(fixTypeStore)
	deviceTypeService := service.NewDeviceTypeService(deviceTypeStore)

	// Handler
	technicalServiceHandler := api.NewTechnicalServiceHandler(technicalServiceService)
	serviceTypeHandler := api.NewServiceTypeHandler(serviceTypeService)
	extraServiceHandler := api.NewExtraServiceHandler(extraServiceService)
	brandHandler := api.NewBrandHandler(brandService)
	modelHandler := api.NewModelHandler(modelService)
	fixTypeHandler := api.NewFixTypeHandler(fixTypeService)
	deviceTypeHandler := api.NewDeviceTypeHandler(deviceTypeService)

	// gin server
	router := gin.Default()
	router.Use(gin.Logger())

	// CORS Policy
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowMethods("OPTIONS", "GET", "PUT", "PATCH", "POST")
	router.Use(cors.New(corsConfig))

	route := router.Group("api/v1")
	{
		// Get All Models APIs
		route.GET("/technical-service", technicalServiceHandler.GetAll)
		route.GET("/service-type", serviceTypeHandler.GetAll)
		route.GET("/extra-service", extraServiceHandler.GetAll)
		route.GET("/brand", brandHandler.GetAll)
		route.GET("/model", modelHandler.GetAll)
		route.GET("/fix-type", fixTypeHandler.GetAll)
		route.GET("/device-type", deviceTypeHandler.GetAll)

		// Create Model APIs
		route.POST("/technical-service", technicalServiceHandler.Create)
	}

	router.Run(":8080")
}
