package app

import (
	"log"
	"miniproject/app/middleware"
	"miniproject/app/routes"
	UserUseCase "miniproject/business/users"
	UserController "miniproject/controllers/users"
	"miniproject/driver/database/users"
	UserRepository "miniproject/driver/database/users"
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
	// db.AutoMigrate(&books.Books{})
	// db.AutoMigrate(&categories.Categories{})
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

	userRepoInterface := UserRepository.NewUserRepository(db)
	userUseCaseInterface := UserUseCase.NewUseCase(userRepoInterface, timeoutContext, &configJWT)
	userControllerInterface := UserController.NewUserController(userUseCaseInterface)

	routesInit := routes.RouteControllerList {
		UserController: *userControllerInterface,
	}

	routesInit.RouteRegister(e)
	log.Fatal(e.Start(viper.GetString("server.address")))
}