package books

import (
	"miniproject/business/books"
	"miniproject/controllers"
	"miniproject/controllers/books/requests"
	"miniproject/controllers/books/responses"
	"miniproject/helpers"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BookController struct {
	BookUseCase books.UseCase
}

func NewBookController(bookUseCase books.UseCase) *BookController {
	return &BookController{
		BookUseCase: bookUseCase,
	}
}

func (bookController *BookController) GetAllBooks(c echo.Context) error {
	request := c.Request().Context()
	search := c.QueryParam("q")
	data, err := bookController.BookUseCase.GetAllBook(request, search)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, responses.FromBookAll(data))
}

func (bookController *BookController) GetBookById(c echo.Context) error {
	request := c.Request().Context()
	id := c.Param("id")
	convInt, errConvInt := strconv.Atoi(id)
	if errConvInt != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, errConvInt)
	}
	data, err := bookController.BookUseCase.GetBookById(request, uint(convInt))
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, responses.FromDomain(data))
}

func (bookController *BookController) InsertBook(c echo.Context) error {
	request := requests.InsertBook{}
	err := c.Bind(&request)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	context := c.Request().Context()
	var data books.Domain
	data, err = bookController.BookUseCase.InsertBook(context, request.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, responses.FromDomain(data))
}

func (bookController *BookController) UpdateBook(c echo.Context) error {
	id := c.Param("id")
	convId, err := helpers.StringToUint(id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	request := requests.InsertBook{}
	err = c.Bind(&request)
	if err != nil {
		return err
	} 
	context := c.Request().Context()
	data, err := bookController.BookUseCase.Update(context, *request.ToDomain(), convId)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, responses.FromDomain(data))
}

func (bookController *BookController) DeleteBook(c echo.Context) error {
	id := c.Param("id")
	bookid, err := helpers.StringToUint(id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	context := c.Request().Context()
	err = bookController.BookUseCase.Delete(context, bookid)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, responses.BookResponse{Id: bookid})
}