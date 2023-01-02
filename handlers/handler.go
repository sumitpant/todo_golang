package handlers

import (
	"encoding/json"
	"fmt"
	"main/repo"
	"main/service"
	"net/http"
	//"main/repo"
)

type Handler struct{
	Service service.ServiceRepo
}

type Response struct {
	Status int 
	Response string
}

func(ch *Handler) GetAll(w http.ResponseWriter, r *http.Request){
	todos:=ch.Service.FindAll();
	fmt.Println(todos);

	json.NewEncoder(w).Encode(todos);
	
}

func(ch *Handler)CreateOne(w http.ResponseWriter, r *http.Request){
	 


	fmt.Println(r.Body)
	
	 var p repo.Todo

    // // Try to decode the request body into the struct. If there is an error,
    // // respond to the client with the error message and a 400 status code.
    err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		panic(err);
	}
res:=ch.Service.CreateOne(p);
fmt.Print(res);

respose:=Response{
Status: http.StatusOK,
Response:"OK",
}
json.NewEncoder(w).Encode(&respose)

	

}