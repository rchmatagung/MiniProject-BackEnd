package transactiondetails

import (
	transactiondetails "miniproject/business/transaction_details"
	"miniproject/controllers"
	"miniproject/controllers/transaction_details/requests"
	"miniproject/controllers/transaction_details/responses"
	"miniproject/helpers"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Transaction_DetailController struct {
	Transaction_DetailUseCase transactiondetails.Usecase
}

func NewTransaction_DetailController(transaction_detailUseCase transactiondetails.Usecase) *Transaction_DetailController {
	return &Transaction_DetailController{
		Transaction_DetailUseCase: transaction_detailUseCase,
	}
}

func (transaction_detailController *Transaction_DetailController) GetAllTransaction_Detail(c echo.Context) error {
	ctx := c.Request().Context()
	data, err := transaction_detailController.Transaction_DetailUseCase.GetAllTransaction_Detail(ctx, 1, 1)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, responses.FromTransaction_DetailsAll(data))
}

func (transaction_detailController *Transaction_DetailController) GetTransaction_DetailsById(c echo.Context) error {
	request := c.Request().Context()
	id := c.Param("id")
	convInt, errConvInt := strconv.Atoi(id)
	if errConvInt != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, errConvInt)
	}
	data, err := transaction_detailController.Transaction_DetailUseCase.GetTransaction_DetailById(request, uint(convInt))

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomain(data))
}

func (transaction_detailController *Transaction_DetailController) UpdateTransaction_Detail(c echo.Context) error {
	id := c.Param("id")
	convId, err := helpers.StringToUint(id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	req := requests.InsertTransaction_detail{}
	err = c.Bind(&req)
	if err != nil {
		return err
	}
	ctx := c.Request().Context()
	data, err := transaction_detailController.Transaction_DetailUseCase.Update(ctx, *req.ToDomain(), convId)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomain(data))
}

func (transaction_detailController *Transaction_DetailController) InsertTransaction_Detail(c echo.Context) error {
	request := requests.InsertTransaction_detail{}
	var err error
	err = c.Bind(&request)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	ctx := c.Request().Context()
	var data transactiondetails.Domain
	data, err = transaction_detailController.Transaction_DetailUseCase.InsertTransaction_Detail(ctx, request.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, responses.FromDomain(data))
}

func (transaction_detailController *Transaction_DetailController) DeleteTransaction_Details(c echo.Context) error {
	id := c.Param("id")
	idUint, err := helpers.StringToUint(id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	ctx := c.Request().Context()
	err = transaction_detailController.Transaction_DetailUseCase.Delete(ctx, idUint)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, responses.Transaction_detail_response{
		Id: idUint,
	})
}
