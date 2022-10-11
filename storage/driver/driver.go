package driver

type Driver interface {
	Store(id string, obj interface{}) bool // obj
	Load(id string) (obj interface{}, ok bool)
}

func Default() Driver {
	return &DriverMap{make(map[string]string, 0)}
}
