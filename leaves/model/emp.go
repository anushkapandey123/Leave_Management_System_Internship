package model

type Emp struct {
	Id int 		`json:"id" gorm:"primaryKey"`
	Name string	`json:"name"`
}

func NewEmp(id int, name string) Emp {
	return Emp{
		Id: id,
		Name: name,
	}
}

func (Emp) TableName() string {
	return "emptable"
}