package handler

import (
	"gin-starter/common/errors"
	"gin-starter/common/interfaces"
	"gin-starter/modules/master/v1/service"
	"gin-starter/resource"
	"gin-starter/response"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

)

// MasterDeleterHandler is a handler for master finder
type MasterDeleterHandler struct {
	masterDeleter service.MasterDeleterUseCase
	cloudStorage  interfaces.CloudStorageUseCase
}

// NewMasterDeleterHandler is a constructor for MasterDeleterHandler
func NewMasterDeleterHandler(
	masterDeleter service.MasterDeleterUseCase,
	cloudStorage interfaces.CloudStorageUseCase,
) *MasterDeleterHandler {
	return &MasterDeleterHandler{
		masterDeleter: masterDeleter,
		cloudStorage:  cloudStorage,
	}
}

func (masterDeleter *MasterDeleterHandler) DeleteBook(c *gin.Context) {
	var request resource.DeleteBookRequest

	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	reqID, err := uuid.Parse(request.ID)

	if err != nil {
		c.JSON(errors.ErrInvalidArgument.Code, response.ErrorAPIResponse(errors.ErrInvalidArgument.Code, errors.ErrInvalidArgument.Message))
		c.Abort()
		return
	}

	
	if err := masterDeleter.masterDeleter.DeleteBook(c, reqID); err != nil { 
		parseError := errors.ParseError(err)
		c.JSON(parseError.Code, response.ErrorAPIResponse(parseError.Code, parseError.Message))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "Delete success", nil))
}