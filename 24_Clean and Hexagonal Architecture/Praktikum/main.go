package main

import (
	"context"
	_driverFactory "belajar-go-echo/drivers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	_userUseCase "belajar-go-echo/businesses/users"
	_userController "belajar-go-echo/controllers/users"

	_dbDriver "belajar-go-echo/drivers/mysql"

	_middleware "belajar-go-echo/app/middlewares"
	_routes "belajar-go-echo/app/routes"

	echo "github.com/labstack/echo/v4"
)

type operation func(ctx context.Context) error

func main() {

	db := _dbDriver.ConnectDB()

	_dbDriver.MigrateDB(db)

	configJWT := _middleware.JWTConfig{
		SecretKey:       "rahasia",
		ExpiresDuration: 1,
	}

	configLogger := _middleware.LoggerConfig{
		Format: "[${time_rfc3339}] ${status} ${method} ${host} ${path} ${latency_human}" + "\n",
	}

	e := echo.New()

	userRepo := _driverFactory.NewUserRepository(db)
	userUsecase := _userUseCase.NewUserUseCase(userRepo, &configJWT)
	userCtrl := _userController.NewAuthController(userUsecase)

	routesInit := _routes.ControllerList{
		LoggerMiddleware:   configLogger.Init(),
		JWTMiddleware:      configJWT.Init(),
		AuthController:     *userCtrl,
	}

	routesInit.SetUpRoutes(e)

	go func() {
		if err := e.Start(":8000"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	wait := gracefulShutdown(context.Background(), 2*time.Second, map[string]operation{
		"database": func(ctx context.Context) error {
			return _dbDriver.CloseDB(db)
		},
		"http-server": func(ctx context.Context) error {
			return e.Shutdown(context.Background())
		},
	})

	<-wait
}

func gracefulShutdown(ctx context.Context, timeout time.Duration, ops map[string]operation) <-chan struct{} {
	wait := make(chan struct{})
	go func() {
		s := make(chan os.Signal, 1)

		// add any other syscalls that you want to be notified with
		signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
		<-s

		log.Println("shutting down")

		// set timeout for the ops to be done to prevent system hang
		timeoutFunc := time.AfterFunc(timeout, func() {
			log.Printf("timeout %d ms has been elapsed, force exit", timeout.Milliseconds())
			os.Exit(0)
		})

		defer timeoutFunc.Stop()

		var wg sync.WaitGroup

		// Do the operations asynchronously to save time
		for key, op := range ops {
			wg.Add(1)
			innerOp := op
			innerKey := key
			go func() {
				defer wg.Done()

				log.Printf("cleaning up: %s", innerKey)
				if err := innerOp(ctx); err != nil {
					log.Printf("%s: clean up failed: %s", innerKey, err.Error())
					return
				}

				log.Printf("%s was shutdown gracefully", innerKey)
			}()
		}

		wg.Wait()

		close(wait)
	}()

	return wait
}