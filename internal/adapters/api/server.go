package api

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
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
	productCreateUseCase *usecases.ProductCreateUseCase
	productReadUseCase   *usecases.ProductReadUseCase
	productUpdateUseCase *usecases.ProductUpdateUseCase
}

type AppHandlers struct {
	pingHandler          *handlers.PingHandler
	productCreateHandler *handlers.ProductCreateHandler
	productReadHandler   *handlers.ProductReadHandler
	productUpdateHandler *handlers.ProductUpdateHandler
}

type EnvVariables struct {
	MYSQL_HOST string
	MYSQL_PORT string
	MYSQL_USER string
	MYSQL_PASS string
	MYSQL_DB   string
}

func NewRestService(router *echo.Echo, appHandler *AppHandlers) *RestServer {
	return &RestServer{
		router:     router,
		appHandler: appHandler,
	}
}

func RunServer() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	env := getEnv()

	db, err := gorm.Open(gormmysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", env.MYSQL_USER, env.MYSQL_PASS, env.MYSQL_HOST, env.MYSQL_PORT, env.MYSQL_DB)), &gorm.Config{})
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
		MYSQL_HOST: os.Getenv("MYSQL_HOST"),
		MYSQL_PORT: os.Getenv("MYSQL_PORT"),
		MYSQL_USER: os.Getenv("MYSQL_USER"),
		MYSQL_PASS: os.Getenv("MYSQL_PASS"),
		MYSQL_DB:   os.Getenv("MYSQL_DB"),
	}
}
