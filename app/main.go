package main

import (
	"log"
	"miniproject/app/middleware"
	"miniproject/app/routes"

	userUseCase "miniproject/business/users"
	userController "miniproject/controllers/users"
	userRepository "miniproject/driver/database/users"

	bookUseCase "miniproject/business/books"
	bookController "miniproject/controllers/books"
	bookRepository "miniproject/driver/database/books"

	categoryUseCase "miniproject/business/categories"
	categoryController "miniproject/controllers/categories"
	categoryRepository "miniproject/driver/database/categories"

	descriptionUseCase "miniproject/business/descriptions"
	descriptionController "miniproject/controllers/descriptions"
	descriptionRepository "miniproject/driver/database/descriptions"

	paymentmethodUseCase "miniproject/business/payment_methods"
	paymentmethodController "miniproject/controllers/payment_methods"
	paymentmethodsRepository "miniproject/driver/database/payment_methods"

	transactiondetailsUseCase "miniproject/business/transaction_details"
	transactiondetailsController "miniproject/controllers/transaction_details"
	transactiondetailsRepository "miniproject/driver/database/transaction_details"

	transactionUseCase "miniproject/business/transactions"
	transactionController "miniproject/controllers/transactions"
	transactionRepository "miniproject/driver/database/transactions"

	"miniproject/driver/mysql"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	if viper.GetBool(`debug`) {
		log.Println("Service Run on Debug Mode")
	}
}

func DBMigrate(db *gorm.DB) {
	db.AutoMigrate(&userRepository.Users{})
	db.AutoMigrate(&bookRepository.Books{})
	db.AutoMigrate(&categoryRepository.Categories{})
	db.AutoMigrate(&descriptionRepository.Descriptions{})
	db.AutoMigrate(&paymentmethodsRepository.Payment_Methods{})
	db.AutoMigrate(&transactiondetailsRepository.Transaction_Detail{})
	db.AutoMigrate(&transactionRepository.Transaction{})
}

func main() {
	configDB := mysql.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}
	configJWT := middleware.ConfigJWT{
		SecretJWT: viper.GetString(`jwt.secret`),
		ExpiresDuration: viper.GetInt(`jwt.expired`),
	}

	db := configDB.InitialDB()
	DBMigrate(db)
	e := echo.New()
	timeoutContext := time.Duration(viper.GetInt("context.timeout"))* time.Second

	userRepoInterface := userRepository.NewUserRepository(db)
	userUseCaseInterface := userUseCase.NewUseCase(userRepoInterface, timeoutContext, &configJWT)
	userControllerInterface := userController.NewUserController(userUseCaseInterface)

	bookRepoInterface := bookRepository.NewUserRepository(db)
	bookUseCaseInterface := bookUseCase.NewUseCase(bookRepoInterface, timeoutContext)
	bookControllerInterface := bookController.NewBookController(bookUseCaseInterface)

	categoryRepoInterface := categoryRepository.NewCategoryRepository(db)
	categoryUseCaseInterface := categoryUseCase.NewCategoryUseCase(categoryRepoInterface, timeoutContext)
	categoryControllerInterface := categoryController.NewCategoryController(categoryUseCaseInterface)

	descriptionRepoInterface := descriptionRepository.NewUserRepository(db)
	descriptionUseCaseInterface := descriptionUseCase.NewUseCase(descriptionRepoInterface, timeoutContext)
	descriptionControllerInterface := descriptionController.NewDescriptionController(descriptionUseCaseInterface)
	
	payment_methodRepoInterface := paymentmethodsRepository.NewUserRepository(db)
	payment_methodUseCaseInterface := paymentmethodUseCase.NewUseCase(payment_methodRepoInterface, timeoutContext)
	payment_methodControllerInterface := paymentmethodController.NewPayment_MethodController(payment_methodUseCaseInterface)
	
	transaction_detailsRepoInterface := transactiondetailsRepository.NewUserRepository(db)
	transaction_detailsUseCaseInterface := transactiondetailsUseCase.NewUseCase(transaction_detailsRepoInterface, timeoutContext)
	transaction_detailsControllerInterface := transactiondetailsController.NewTransaction_DetailController(transaction_detailsUseCaseInterface)
	
	transactionRepoInterface := transactionRepository.NewUserRepository(db)
	transactionUseCaseInterface := transactionUseCase.NewUseCase(transactionRepoInterface, timeoutContext)
	transactionControllerInterface := transactionController.NewTransactionController(transactionUseCaseInterface)
	
	routesInit := routes.RouteControllerList {
		UserController: *userControllerInterface,
		BookController: *bookControllerInterface,
		CategoryController: *categoryControllerInterface,
		DescriptionController: *descriptionControllerInterface,
		PaymentMethodController: *payment_methodControllerInterface,
		TransactionDetailController: *transaction_detailsControllerInterface,
		TransactionController: *transactionControllerInterface,
	}

	routesInit.RouteRegister(e)
	log.Fatal(e.Start(viper.GetString("server.address")))
}