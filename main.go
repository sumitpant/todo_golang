package main

import (
	"fmt"
	"main/handlers"
	"main/repo"
	"main/service"
	"net/http"
	"time"
	"main/logger"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)



func main() {

router:=mux.NewRouter();


t := time.Now().Unix();
fmt.Println(t);
logger.Logs("Started server")
ch:=  handlers.Handler {
	Service:service.NewService(repo.NewRepository()),
}



router.HandleFunc("/getAll",ch.GetAll)

router.HandleFunc("/create",ch.CreateOne)


http.ListenAndServe(":8080", router);

}