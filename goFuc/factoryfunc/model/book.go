package model

//商品信息
type IInfo interface {
	GetInfo() string
}

//图书
type Book struct {
}

func (b Book) GetInfo() string {
	//TODO implement me
	return "我是图书"
}

//鱼
type Fish struct {
}

func (f Fish) GetInfo() string {
	//TODO implement me
	return "我是海鲜"
}
