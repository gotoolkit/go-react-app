package model

type Reader interface {
	GetAll() (interface{}, error)
	GetByID(int) (interface{}, error)
}

type Writer interface {
	Create(interface{}) error
	BindModel() interface{}
}

type ReadWriter interface {
	Reader
	Writer
}
