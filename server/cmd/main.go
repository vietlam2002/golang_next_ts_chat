package main

import (
	"fmt"
	"log"
	"server/db"
	"server/internal/user"
	"server/router"
)

func main() {
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("could not initialize database connection: %s", err)
	} else {
		fmt.Printf("connected database!")
		dbConn.GetDB()
	}

	// Handle sẽ xử lí yêu cầu từ router, sau đó handle gọi service -> repo -> database
	userRep := user.NewRepository(dbConn.GetDB())
	userSvc := user.NewService(userRep)
	userHandler := user.NewHandler(userSvc)

	router.InitRouter(userHandler)
	router.Start("0.0.0.0:5000")
}
