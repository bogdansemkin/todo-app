package main

import (
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
	"os"
	todo_app "todo-app"
	"todo-app/pkg/handler"
	"todo-app/pkg/repository"
	"todo-app/pkg/service"
)

func main(){
	if err := initConfig(); err != nil{
		logrus.Fatalf("error initializating configs: %s", err.Error())
	}

	if err := gotenv.Load(); err != nil{
		logrus.Fatalf("error loading env variables : %s",err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host: viper.GetString("db.host"),
		Port: viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBname: viper.GetString("db.dbname"),
		SSLMode: viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo_app.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil{
		logrus.Fatalf("error occured while running http server %d", err.Error())
	}
}

func initConfig() error{
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}