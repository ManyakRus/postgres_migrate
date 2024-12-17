package main

import (
	"github.com/ManyakRus/postgres_migrate/internal/config"
	"github.com/ManyakRus/postgres_migrate/pkg/constants"
	"github.com/ManyakRus/postgres_migrate/pkg/crud_starter"
	"github.com/ManyakRus/postgres_migrate/pkg/version"
	"github.com/ManyakRus/starter/config_main"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"github.com/ManyakRus/starter/postgres_gorm"
	"github.com/ManyakRus/starter/stopapp"
)

// main - старт приложения
func main() {
	StartApp()
}

// StartApp - выполнение всех операций для старта приложения
func StartApp() {
	micro.Show_Version(version.Version)
	config_main.LoadEnv()
	config.FillSettings()

	//
	stopapp.StartWaitStop()

	//
	postgres_gorm.Start_SingularTableName(constants.SERVICE_NAME)
	crud_starter.InitCrudTransport_DB()

	//
	stopapp.GetWaitGroup_Main().Wait()

	//
	log.Info("Application closed.")
}
