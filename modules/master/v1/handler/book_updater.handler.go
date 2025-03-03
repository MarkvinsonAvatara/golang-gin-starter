package handler

import (
	"gin-starter/common/errors"
	"gin-starter/common/interfaces"
	"gin-starter/entity"
	"gin-starter/modules/master/v1/service"
	"gin-starter/resource"
	"gin-starter/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// MasterUpdaterHandler is a handler for master finder
type BookMasterUpdaterHandler struct {
	bookMasterUpdater service.BookMasterUpdaterUseCase
	bookMasterFinder  service.BookMasterFinderUseCase
	cloudStorage  interfaces.CloudStorageUseCase
}

// NewMasterUpdaterHandler is a constructor for MasterUpdaterHandler
func NewMasterUpdaterHandler(
	bookMasterUpdater service.BookMasterUpdaterUseCase,
	bookMasterFinder service.BookMasterFinderUseCase,
	cloudStorage interfaces.CloudStorageUseCase,
) *BookMasterUpdaterHandler {
	return &BookMasterUpdaterHandler{
		bookMasterUpdater: bookMasterUpdater,
		bookMasterFinder:  bookMasterFinder,
		cloudStorage:  cloudStorage,
	}
}

func (masterUpdater *BookMasterUpdaterHandler) UpdateBook(c *gin.Context) {
	var request resource.UpdateBookRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}
	bookIDstr := c.Param("id")
	bookID, err := uuid.Parse(bookIDstr)
	if err != nil {
		c.JSON(errors.ErrInvalidArgument.Code, response.ErrorAPIResponse(errors.ErrInvalidArgument.Code, errors.ErrInvalidArgument.Message))
		c.Abort()
		return
	}

	_, err = masterUpdater.bookMasterFinder.GetBookByID(c, bookID)

	if err != nil {
		parseError := errors.ParseError(err)
		c.JSON(parseError.Code, response.ErrorAPIResponse(parseError.Code, parseError.Message))
		c.Abort()
		return
	}

	book := entity.UpdateBook(
		bookID,
		request.Isbn, 
		request.Title, 
		request.Author, 
		request.Genre, 
		request.Publisher, 
		request.Edition, 
		request.Description,
		"System",
	)

	if err := masterUpdater.bookMasterUpdater.UpdateBook(c, book); err != nil {
		parseError := errors.ParseError(err)
		c.JSON(parseError.Code, response.ErrorAPIResponse(parseError.Code, parseError.Message))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "Update success", nil))
}
