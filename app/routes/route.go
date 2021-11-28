package routes

import (
	"miniproject/controllers/books"
	"miniproject/controllers/categories"
	"miniproject/controllers/descriptions"
	paymentmethods "miniproject/controllers/payment_methods"
	transactiondetails "miniproject/controllers/transaction_details"
	"miniproject/controllers/transactions"
	"miniproject/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type RouteControllerList struct {
	JWTMiddleware 				middleware.JWTConfig
	UserController 				users.UserController
	BookController 				books.BookController
	CategoryController 			categories.CategoryController
	DescriptionController 		descriptions.DescriptionController
	PaymentMethodController 	paymentmethods.Payment_MethodController
	TransactionDetailController transactiondetails.Transaction_DetailController
	TransactionController 		transactions.TransactionController
}

func (controller *RouteControllerList) RouteRegister(e *echo.Echo) {
	//users
	e.POST("users/login", controller.UserController.Login)
	e.POST("users/register", controller.UserController.Register)
	e.GET("users", controller.UserController.GetAllUsers)
	e.DELETE("users/:id", controller.UserController.DeleteUser)
	e.PUT("users/:id", controller.UserController.UpdateUser)

	//books
	e.POST("books/insertbook", controller.BookController.InsertBook)
	e.GET("books", controller.BookController.GetAllBooks)
	e.GET("books/:id", controller.BookController.GetBookById)
	e.DELETE("books/:id", controller.BookController.DeleteBook)
	e.PUT("books/:id", controller.BookController.UpdateBook)

	//category
	e.POST("category/insertcategory", controller.CategoryController.InsertCategory)
	e.GET("category", controller.CategoryController.GetAllCategory)
	e.GET("category/:id", controller.CategoryController.GetCategoryById)
	e.DELETE("category/:id", controller.CategoryController.DeleteCategory)
	e.PUT("category/:id", controller.CategoryController.UpdateCategory)
	
	//description
	e.POST("descriptions/insertdescription", controller.DescriptionController.InsertDescription)
	e.GET("descriptions", controller.DescriptionController.GetAllDescriptions)
	e.GET("descriptions/:id", controller.DescriptionController.GetDescriptionById)
	e.DELETE("descriptions/:id", controller.DescriptionController.DeleteDescription)
	e.PUT("descriptions/:id", controller.DescriptionController.UpdateDescription)
	
	//payment_methods
	e.POST("payment_methods/insertpayment_methods", controller.PaymentMethodController.InsertPayment_Method)
	e.GET("payment_methods", controller.PaymentMethodController.GetAllPayment_Methods)
	e.GET("payment_methods/:id", controller.PaymentMethodController.GetPayment_MethodById)
	e.DELETE("payment_methods/:id", controller.PaymentMethodController.DeletePayment_Method)
	e.PUT("payment_methods/:id", controller.PaymentMethodController.UpdatePayment_Method)
	
	//transactions
	e.POST("transactions/inserttransactions", controller.TransactionController.InsertTransaction)
	e.GET("transactions", controller.TransactionController.GetAllTransaction)
	e.GET("transactions/:id", controller.TransactionController.GetTransactionById)
	e.DELETE("transactions/:id", controller.TransactionController.DeleteTransactions)
	e.PUT("transactions/:id", controller.TransactionController.UpdateTransaction)
	
	//transaction_details
	e.POST("transaction_details/inserttransaction_details", controller.TransactionDetailController.InsertTransaction_Detail)
	e.GET("transaction_details", controller.TransactionDetailController.GetAllTransaction_Detail)
	e.GET("transaction_details/:id", controller.TransactionDetailController.GetTransaction_DetailsById)
	e.DELETE("transaction_details/:id", controller.TransactionDetailController.DeleteTransaction_Details)
	e.PUT("transaction_details/:id", controller.TransactionDetailController.UpdateTransaction_Detail)
}