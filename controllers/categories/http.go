package categories

import (
	"miniproject/business/categories"
	"miniproject/controllers"
	"miniproject/controllers/categories/requests"
	"miniproject/controllers/categories/responses"
	"miniproject/helpers"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CategoryController struct {
	CategoryUseCase categories.UseCase
}

func NewCategoryController(categoryUseCase categories.UseCase) *CategoryController {
	return &CategoryController{
		CategoryUseCase: categoryUseCase,
	}
}

func (categoryController *CategoryController) InsertCategory(c echo.Context) error {
	request :=requests.InsertCategory{}
	var err error
	err = c.Bind(&request)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	context := c.Request().Context()
	var data categories.Domain
	data, err = categoryController.CategoryUseCase.InsertCategory(context, *request.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, responses.FromDomain(data))
}

func (categoryController *CategoryController) GetAllCategory(c echo.Context) error {
	context := c.Request().Context()
	search := c.QueryParam("q")
	data, err := categoryController.CategoryUseCase.GetAllCategory(context, search)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, responses.FromCategoriesAll(data))
}

func (categoryController *CategoryController) GetCategoryById(c echo.Context) error {
	request := c.Request().Context()
	id := c.Param("id")
	convInt, errConvInt := strconv.Atoi(id)
	if errConvInt != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, errConvInt)
	}
	data, err := categoryController.CategoryUseCase.GetCategoryById(request, uint(convInt))
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, responses.FromDomain(data))
}

func (categoryController *CategoryController) UpdateCategory(c echo.Context) error {
	id := c.Param("id")
	convId, err := helpers.StringToUint(id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	request := requests.InsertCategory{}
	err = c.Bind(&request)
	if err != nil {
		return err
	}
	context := c.Request().Context()
	data, err := categoryController.CategoryUseCase.Update(context, *request.ToDomain(), convId)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, responses.FromDomain(data))
}

func (categoryController *CategoryController) DeleteCategory(c echo.Context) error {
	id := c.Param("id")
	idUint, err := helpers.StringToUint(id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	context := c.Request().Context()
	err = categoryController.CategoryUseCase.Delete(context, idUint)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, responses.CategoryResponse{Id: idUint})
}