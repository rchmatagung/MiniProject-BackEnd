package app

import (
	"log"
	"miniproject/app/routes"
	bookUseCase "miniproject/business/books"
	bookController "miniproject/controllers/books"
	bookRepository "miniproject/driver/database/books"
	"miniproject/driver/database/categories"
	"miniproject/driver/database/descriptions"
	paymentmethods "miniproject/driver/database/payment_methods"
	transactiondetails "miniproject/driver/database/transaction_details"
	"miniproject/driver/database/transactions"
	"miniproject/driver/database/users"
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
	db.AutoMigrate(&users.Users{})
	db.AutoMigrate(&bookRepository.Books{})
	db.AutoMigrate(&categories.Categories{})
	db.AutoMigrate(&descriptions.Descriptions{})
	db.AutoMigrate(&paymentmethods.Payment_Methods{})
	db.AutoMigrate(&transactiondetails.Transaction_Detail{})
	db.AutoMigrate(&transactions.Transaction{})
}

func main() {
	configDB := mysql.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}
	// configJWT := middleware.ConfigJWT{
	// 	SecretJWT: viper.GetString(`jwt.secret`),
	// 	ExpiresDuration: viper.GetInt(`jwt.expired`),
	// }

	db := configDB.InitialDB()
	DBMigrate(db)
	e := echo.New()
	timeoutContext := time.Duration(viper.GetInt("context.timeout"))* time.Second

	// userRepoInterface := UserRepository.NewUserRepository(db)
	// userUseCaseInterface := UserUseCase.NewUseCase(userRepoInterface, timeoutContext, &configJWT)
	// userControllerInterface := UserController.NewUserController(userUseCaseInterface)

	bookRepoInterface := bookRepository.NewUserRepository(db)
	bookUseCaseInterface := bookUseCase.NewUseCase(bookRepoInterface, timeoutContext)
	bookControllerInterface := bookController.NewBookController(bookUseCaseInterface)

	routesInit := routes.RouteControllerList {
		BookController: *bookControllerInterface,
	}

	routesInit.RouteRegister(e)
	log.Fatal(e.Start(viper.GetString("server.address")))
}