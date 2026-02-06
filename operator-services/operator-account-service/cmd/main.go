package main

import (
	"context"
	"fmt"
	"ylanzinhoy-operator-management/config"
	"ylanzinhoy-operator-management/internal/infrastructure/persistence/database"
	module "ylanzinhoy-operator-management/internal/module/api"
)

func main() {

	cfg := config.GetInstance()
	ctx := context.Background()

	mongoConn := database.NewDatabaseConnectionAdapter(*cfg)

	mongoConn.Connect(ctx)

	err := mongoConn.CreateCollection(ctx, "operator-management", "operator")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Collection created")
	coll := mongoConn.GetCollection("operator-management", "operator")

	api := module.NewApi()
	api.ApiModuleInitialize(coll)

	api.Start()

	fmt.Println("Hello World")

}
