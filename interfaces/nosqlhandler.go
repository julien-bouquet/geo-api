package interfaces

type NoSQLHandler interface {
	Get() (Documents, error)
	Add(...interface{}) error
}

type Documents interface {
	Read(interface{}) error
	Next(...interface{}) bool
	Close() error
	Err() error
}

type ResultInsert interface{}
