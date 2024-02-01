package api

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/uiansol/product-follow-up/internal/adapters/api/handlers"
	"github.com/uiansol/product-follow-up/internal/adapters/db/mysql"
	"github.com/uiansol/product-follow-up/internal/application/usecases"
	gormmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type RestServer struct {
	router     *echo.Echo
	appHandler *AppHandlers
}

type AppRepositories struct {
	productRepository *mysql.ProductRepository
}

type AppUseCases struct {
	productCreateUseCase  *usecases.ProductCreateUseCase
	productReadUseCase    *usecases.ProductReadUseCase
	productUpdateUseCase  *usecases.ProductUpdateUseCase
	productDeleteUseCase  *usecases.ProductDeleteUseCase
	productReadAllUseCase *usecases.ProductReadAllUseCase
}

type AppHandlers struct {
	pingHandler           *handlers.PingHandler
	productCreateHandler  *handlers.ProductCreateHandler
	productReadHandler    *handlers.ProductReadHandler
	productUpdateHandler  *handlers.ProductUpdateHandler
	productDeleteHandler  *handlers.ProductDeleteHandler
	productReadAllHandler *handlers.ProductReadAllHandler
}

type EnvVariables struct {
	MYSQL_ROOT_HOST     string
	MYSQL_ROOT_PASSWORD string
	MYSQL_DATABASE      string
}

func NewRestService(router *echo.Echo, appHandler *AppHandlers) *RestServer {
	return &RestServer{
		router:     router,
		appHandler: appHandler,
	}
}

func RunServer() {
	env := getEnv()

	db, err := gorm.Open(gormmysql.Open(fmt.Sprintf("root:%s@tcp(%s:3306)/%s?parseTime=true", env.MYSQL_ROOT_PASSWORD, env.MYSQL_ROOT_HOST, env.MYSQL_DATABASE)), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	repositories := configRepositories(db)
	useCases := configUseCases(repositories)
	handlers := configHandlers(useCases)

	router := echo.New()
	server := NewRestService(router, handlers)
	server.SetUpRoutes()

	server.router.Logger.Fatal(router.Start(":8080"))
}

func getEnv() EnvVariables {
	return EnvVariables{
		MYSQL_ROOT_HOST:     os.Getenv("MYSQL_ROOT_HOST"),
		MYSQL_ROOT_PASSWORD: os.Getenv("MYSQL_ROOT_PASSWORD"),
		MYSQL_DATABASE:      os.Getenv("MYSQL_DATABASE"),
	}
}

func MigrateDB() {
	fmt.Println("Migrating MySQL Models")

	env := getEnv()

	db, err := gorm.Open(gormmysql.Open(fmt.Sprintf("root:%s@tcp(%s:3306)/%s?parseTime=true", env.MYSQL_ROOT_PASSWORD, env.MYSQL_ROOT_HOST, env.MYSQL_DATABASE)), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&mysql.ProductDB{})

	fmt.Println("Migration finished")
}
