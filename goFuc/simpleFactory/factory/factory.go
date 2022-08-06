package factory

import "study/goFuc/simpleFactory/model"

const (
	FrontUser = iota
	AdminUser
)

type UserType int

func CreateUser(t UserType) model.UserCreateFunc {
	switch t {
	case FrontUser:
		return model.NewUser()
	case AdminUser:
		return model.NewAdminUser()
	default:
		return model.NewUser()
	}

}
