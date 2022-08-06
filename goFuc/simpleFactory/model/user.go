package model

type User struct {
	Id   int
	Name string
}

type UserCreateFunc func(Id int, Name string) interface{}

func NewUser() UserCreateFunc {
	return func(Id int, Name string) interface{} {
		return &User{
			Id:   Id,
			Name: Name,
		}
	}
}

type AdminUser struct {
	Id   int
	Name string
	Role string
}

func NewAdminUser() UserCreateFunc {
	return func(Id int, Name string) interface{} {
		return &AdminUser{
			Id:   Id,
			Name: Name,
			Role: "Admin",
		}
	}
}
