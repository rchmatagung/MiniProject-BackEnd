package paymentmethods

import (
	paymentmethods "miniproject/business/payment_methods"
	"miniproject/controllers"
	"miniproject/controllers/payment_methods/requests"
	"miniproject/controllers/payment_methods/responses"
	"miniproject/helpers"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Payment_MethodController struct {
	Payment_MethodUseCase paymentmethods.Usecase
}

func NewPayment_MethodController(payment_methodUseCase paymentmethods.Usecase) *Payment_MethodController {
	return &Payment_MethodController{
		Payment_MethodUseCase: payment_methodUseCase,
	}
}

func (payment_methodController *Payment_MethodController) GetAllPayment_Methods(c echo.Context) error {
	ctx := c.Request().Context()
	search := c.QueryParam("q")
	data, err := payment_methodController.Payment_MethodUseCase.GetAllPayment_Method(ctx, search)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, responses.FromPayment_MethodsAll(data))
}

func (payment_methodController *Payment_MethodController) GetPayment_MethodById(c echo.Context) error {
	request := c.Request().Context()
	id := c.Param("id")
	convInt, errConvInt := strconv.Atoi(id)
	if errConvInt != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, errConvInt)
	}
	data, err := payment_methodController.Payment_MethodUseCase.GetPayment_MethodById(request, uint(convInt))

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomain(data))
}

func (payment_methodController *Payment_MethodController) UpdatePayment_Method(c echo.Context) error {
	id := c.Param("id")
	convId, err := helpers.StringToUint(id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	req := requests.InsertPayment_Method{}
	err = c.Bind(&req)
	if err != nil {
		return err
	}
	ctx := c.Request().Context()
	data, err := payment_methodController.Payment_MethodUseCase.Update(ctx, *req.ToDomain(), convId)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomain(data))
}

func (payment_methodController *Payment_MethodController) InsertPayment_Method(c echo.Context) error {
	request := requests.InsertPayment_Method{}
	var err error
	err = c.Bind(&request)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	ctx := c.Request().Context()
	var data paymentmethods.Domain
	data, err = payment_methodController.Payment_MethodUseCase.InsertPayment_Method(ctx, *request.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, responses.FromDomain(data))
}

func (payment_methodController *Payment_MethodController) DeletePayment_Method(c echo.Context) error {
	id := c.Param("id")
	idUint, err := helpers.StringToUint(id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	ctx := c.Request().Context()
	err = payment_methodController.Payment_MethodUseCase.Delete(ctx, idUint)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, responses.Payment_MethodResponse{
		Id: idUint,
	})
}