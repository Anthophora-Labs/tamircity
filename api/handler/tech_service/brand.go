package tech_service

import (
	"github.com/anthophora/tamircity/pkg/models/db/tech_service"
	tech_service3 "github.com/anthophora/tamircity/pkg/models/web/tech_service"
	tech_service2 "github.com/anthophora/tamircity/pkg/service/tech_service"
	"net/http"
	"strconv"

	"github.com/anthophora/tamircity/pkg/utils"
	"github.com/gin-gonic/gin"
)

type brandHandler struct {
	brandService tech_service2.BrandService
}

type BrandHandler interface {
	GetAll(ctx *gin.Context)
	GetAllByDeviceTypeId(ctx *gin.Context)
	Create(ctx *gin.Context)
}

func NewBrandHandler(brandService tech_service2.BrandService) BrandHandler {
	return &brandHandler{
		brandService: brandService,
	}
}

func (b *brandHandler) GetAll(ctx *gin.Context) {
	brands, err := b.brandService.FindAll()
	if err != nil {
		responseErr := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}
	response := utils.HandleResponseModel(true, "", nil, brands)
	ctx.JSON(http.StatusOK, response)
}

func (b *brandHandler) GetAllByDeviceTypeId(ctx *gin.Context) {

	deviceTypeId, err := strconv.Atoi(ctx.Query("device_type_id"))
	if err != nil {
		responseErr := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}

	brands, err := b.brandService.FindByDeviceTypeId(deviceTypeId)
	if err != nil {
		responseErr := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}
	response := utils.HandleResponseModel(true, "", nil, brands)
	ctx.JSON(http.StatusOK, response)
}

func (b *brandHandler) Create(ctx *gin.Context) {
	var brand tech_service3.BrandRequest
	if err := ctx.ShouldBindJSON(&brand); err != nil {
		response := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var brandModel tech_service.Brand
	brandModel.Name = brand.Name
	brandModel.IsActive = brand.IsActive

	err := b.brandService.Create(&brandModel)
	if err != nil {
		response := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, response)
	}
	response := utils.HandleResponseModel(true, "", nil, brandModel)
	ctx.JSON(http.StatusOK, response)
}
