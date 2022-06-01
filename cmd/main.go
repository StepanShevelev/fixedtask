package main

import (
	"fmt"
	api "github.com/StepanShevelev/fixedtask/api"
	mydb "github.com/StepanShevelev/fixedtask/db"
	"net/http"
)

func main() {
	//var db *gorm.DB
	//config := cfg.New()
	//if err := config.Load("./configs", "config", "yml"); err != nil {
	//	log.Fatal(err)
	//}

	//db, err := mydb.New(config)
	//if err != nil {
	//	log.Fatal(err)
	//}
	fmt.Println("starting server")
	mydb.ConnectToDb()
	fmt.Println("connected to db")
	api.InitBackendApi()
	fmt.Println("initialised routing")

	//api.ShowSkill("Dog")

	http.ListenAndServe(":8000", nil)
}
