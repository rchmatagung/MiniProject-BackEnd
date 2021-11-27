package descriptions

import (
	"miniproject/business/descriptions"
	"miniproject/controllers"
	"miniproject/controllers/descriptions/requests"
	"miniproject/controllers/descriptions/responses"
	"miniproject/helpers"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type DescriptionController struct {
	DescriptionUseCase descriptions.UseCase
}

func NewDescriptionController(descriptonUseCase descriptions.UseCase) *DescriptionController {
	return &DescriptionController{
		DescriptionUseCase: descriptonUseCase,
	}
}

func (descriptionController *DescriptionController) InsertDescription(c echo.Context) error {
	request := requests.InsertDescription{}
	var err error
	err = c.Bind(&request)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	ctx := c.Request().Context()
	var data descriptions.Domain
	data, err = descriptionController.DescriptionUseCase.InsertDescription(ctx, *request.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, responses.FromDomain(data))
}

func (descriptionController *DescriptionController) GetAllDescriptions(c echo.Context) error {
	ctx := c.Request().Context()
	search := c.QueryParam("q")
	data, err := descriptionController.DescriptionUseCase.GetAllDescription(ctx, search)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, responses.FromDescriptionsAll(data))
}

func (descriptionController *DescriptionController) GetDescriptionById(c echo.Context) error {
	request := c.Request().Context()
	id := c.Param("id")
	convInt, errConvInt := strconv.Atoi(id)
	if errConvInt != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, errConvInt)
	}
	data, err := descriptionController.DescriptionUseCase.GetDescriptionById(request, uint(convInt))

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomain(data))
}

func (descriptionController *DescriptionController) UpdateDescription(c echo.Context) error {
	id := c.Param("id")
	convId, err := helpers.StringToUint(id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	req := requests.InsertDescription{}
	err = c.Bind(&req)
	if err != nil {
		return err
	}
	ctx := c.Request().Context()
	data, err := descriptionController.DescriptionUseCase.Update(ctx, *req.ToDomain(), convId)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomain(data))
}

func (descriptionController *DescriptionController) DeleteDescription(c echo.Context) error {
	id := c.Param("id")
	idUint, err := helpers.StringToUint(id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	ctx := c.Request().Context()
	err = descriptionController.DescriptionUseCase.Delete(ctx, idUint)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, responses.DescriptionResponse{
		Id: idUint,
	})
}