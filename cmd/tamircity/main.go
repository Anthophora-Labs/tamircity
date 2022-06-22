package main

import (
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
	//Set enviroment variables
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db, err := postgres.NewPsqlDB()
	if err != nil {
		log.Fatal("Postgres cannot init", err)
	}
	log.Println("Postgres connected")

	db.AutoMigrate(&dbModels.Brand{}, &dbModels.Model{}, &dbModels.FixType{}, &dbModels.DeviceType{})
	log.Println("Migrations done")

	//Store
	technicalServiceStore := repositories.NewTechnicalServiceStore(db)
	serviceTypeStore := repositories.NewServiceTypeStore(db)
	extraServiceStore := repositories.NewExtraServiceStore(db)

	//Service
	technicalServiceService := service.NewTechnicalServiceService(technicalServiceStore)
	serviceTypeService := service.NewServiceTypeService(serviceTypeStore)
	extraServiceService := service.ExtraServiceService(extraServiceStore)

	//Handler
	serviceTypeHandler := api.NewServiceTypeHandler(serviceTypeService)
	extraServiceHandler := api.NewExtraServiceHandler(extraServiceService)

	//gin server
	router := gin.Default()
	router.Use(gin.Logger())

	route := router.Group("api/v1")
	{
		route.GET("/technical-service", serviceTypeHandler.GetAll)
		route.GET("/service-type", serviceTypeHandler.GetAll)
		route.GET("/extra-service", extraServiceHandler.GetAll)
	}

	router.Run(":8080")
}
