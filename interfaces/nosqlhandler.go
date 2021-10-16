package interfaces

type NoSQLHandler interface {
	Get(map[string]interface{}) (Documents, error)
	Add(interface{}) error
	Delete(interface{}) error
}

type Documents interface {
	Read(interface{}) error
	Next(...interface{}) bool
	Close() error
	Err() error
}

type ResultInsert interface{}
