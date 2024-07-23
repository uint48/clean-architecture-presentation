package main

import (
	"log"
	"myapp/internal/delivery"
	"myapp/internal/infra/database"
	"myapp/internal/repository/mysql"
	"myapp/internal/service/userservice"
)

func main() {
	// Hardcoded MySQL connection info because this code is only for learning purposes
	appMySQLInfo := &database.ConnInfo{
		Host:   "localhost",
		Port:   "3306",
		User:   "root",
		Pass:   "root",
		DBName: "myapp",
	}

	appDbConn, nErr := database.NewMySQLConnection(appMySQLInfo)
	if nErr != nil {
		log.Fatalln(nErr)
	}
	defer database.CloseMySQLConnection(appDbConn)

	userRepo, err := mysql.NewUserRepository(appDbConn)

	if err != nil {
		log.Fatalln(err)
	}

	userService := userservice.NewService(userRepo)
	apiController := delivery.NewRESTController(userService)

	// run api server
	apiController.Run(":8080")
}
