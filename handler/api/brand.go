package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mustafakocatepe/Tamircity/pkg/service"
	"github.com/mustafakocatepe/Tamircity/pkg/utils"
)

type brandHandler struct {
	brandService service.BrandService
}

type BrandHandler interface {
	GetAll(ctx *gin.Context)
}

func NewBrandHandler(brandService service.BrandService) BrandHandler {
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
