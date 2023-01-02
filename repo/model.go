package repo



type Todo struct {
	Id int
	Task        string
	Completed   bool
}

type Repository interface{
FindAll()([]Todo)
CreateOne(todo Todo)(bool)
}
