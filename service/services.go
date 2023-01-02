package service;
import "main/repo"

type ServiceDI struct{
	Repo repo.Repository
}

type ServiceRepo interface{
	FindAll()([]repo.Todo)
	CreateOne(todo repo.Todo)(bool)
}


func (s ServiceDI) FindAll() ([]repo.Todo) {
 rows:=s.Repo.FindAll();
 return rows;
	
}

func (s ServiceDI) CreateOne(todo repo.Todo)(bool){
	return s.Repo.CreateOne(todo);

}


func NewService (repo repo.Repository) ServiceDI {

	return ServiceDI{
		Repo:repo,
	}
}
