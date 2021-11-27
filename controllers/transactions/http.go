package transactions

import (
	"miniproject/business/transactions"
	"miniproject/controllers"
	"miniproject/controllers/transactions/requests"
	"miniproject/controllers/transactions/responses"
	"miniproject/helpers"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TransactionController struct {
	TransactionUseCase transactions.Usecase
}

func NewTransactionController(transactionUseCase transactions.Usecase) *TransactionController {
	return &TransactionController{
		TransactionUseCase: transactionUseCase,
	}
}

func (transactionController *TransactionController) GetAllTransaction(c echo.Context) error {
	ctx := c.Request().Context()
	data, err := transactionController.TransactionUseCase.GetAllTransaction(ctx, 1, 1)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, responses.FromTransactionsAll(data))
}

func (transactionController *TransactionController) GetTransactionById(c echo.Context) error {
	request := c.Request().Context()
	id := c.Param("id")
	convInt, errConvInt := strconv.Atoi(id)
	if errConvInt != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, errConvInt)
	}
	data, err := transactionController.TransactionUseCase.GetTransactionById(request, uint(convInt))

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomain(data))
}

func (transactionController *TransactionController) UpdateTransaction(c echo.Context) error {
	id := c.Param("id")
	convId, err := helpers.StringToUint(id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	req := requests.InsertTransaction{}
	err = c.Bind(&req)
	if err != nil {
		return err
	}
	ctx := c.Request().Context()
	data, err := transactionController.TransactionUseCase.Update(ctx, *req.ToDomain(), convId)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomain(data))
}

func (transactionController *TransactionController) InsertTransaction(c echo.Context) error {
	request := requests.InsertTransaction{}
	var err error
	err = c.Bind(&request)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	ctx := c.Request().Context()
	var data transactions.Domain
	data, err = transactionController.TransactionUseCase.InsertTransaction(ctx, request.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, responses.FromDomain(data))
}

func (transactionController *TransactionController) DeleteTransactions(c echo.Context) error {
	id := c.Param("id")
	idUint, err := helpers.StringToUint(id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	ctx := c.Request().Context()
	err = transactionController.TransactionUseCase.Delete(ctx, idUint)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, responses.TransactionResponse{
		Id: idUint,
	})
}