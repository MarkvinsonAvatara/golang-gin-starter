package handler

import (
	"gin-starter/common/interfaces"
	"gin-starter/common/errors"
	"gin-starter/modules/master/v1/service"
	"gin-starter/resource"
	"gin-starter/response"
	"net/http"
	"github.com/gin-gonic/gin"
	
)

// MasterCreatorHandler is a handler for master finder
type MasterCreatorHandler struct {
	
	masterCreator service.MasterCreatorUseCase
	cloudStorage  interfaces.CloudStorageUseCase
}

// NewMasterCreatorHandler is a constructor for MasterCreatorHandler
func NewMasterCreatorHandler(
	masterCreator service.MasterCreatorUseCase,
	cloudStorage interfaces.CloudStorageUseCase,
) *MasterCreatorHandler {
	return &MasterCreatorHandler{
		masterCreator: masterCreator,
		cloudStorage:  cloudStorage,
	}
}

func (masterCreator *MasterCreatorHandler) CreateBook(c *gin.Context) {
	var request resource.CreateBookRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}


	book, err := masterCreator.masterCreator.CreateBook(
		c.Request.Context(), 
		request.Isbn, 
		request.Title, 
		request.Author, 
		request.Genre, 
		request.Publisher, 
		request.Edition, 
		request.Description,
	)

	if err != nil {
		parseError := errors.ParseError(err)
		c.JSON(parseError.Code, response.ErrorAPIResponse(parseError.Code, parseError.Message))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "Add New Buku Success", resource.NewBookResponse(book)))
}

