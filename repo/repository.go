package repo

import (

	"fmt"

	"github.com/jmoiron/sqlx"
)

type RepositoryDb struct {
	client *sqlx.DB
}

func NewRepository() RepositoryDb{

	client,err:=sqlx.Open("mysql","root:root@tcp(localhost:3306)/todo");
	if(err != nil){
		panic(err);
	}

	return  RepositoryDb{
		client:client,
	}
}

func (db RepositoryDb) FindAll()([]Todo){
	todo:= make([]Todo,10)
	err:= db.client.Select(&todo,"select * from todos")

	if(err != nil){
		panic(err)
	}
	fmt.Println(todo);
	return todo;
}

func (db RepositoryDb)CreateOne(todo Todo)(bool){

	res,err:=db.client.NamedExec("INSERT INTO todos (id, task, completed) VALUES (:id, :task, :completed)", todo)
	//db.client.Exec("Insert into todos values ?,?,? ",todo)

	if err!=nil{
		panic(err)
	}
	fmt.Println(res)

	return true;

}