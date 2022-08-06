package factory

import "study/goFuc/factoryfunc/model"

const (
	BookType = iota
	FishType
)

//工程方法
type IProductFactory interface {
	CreateProduct(i int) model.IInfo
}

type BookFactory struct {
}

func (b BookFactory) CreateProduct(i int) model.IInfo {
	//TODO implement me
	switch i {
	case BookType:
		return model.Book{}
	default:
		return model.Book{}
	}
}

type FishFactory struct {
}

func (f FishFactory) CreateProduct(i int) model.IInfo {
	//TODO implement me
	switch i {
	case FishType:
		return model.Fish{}
	default:
		return model.Fish{}
	}
}
