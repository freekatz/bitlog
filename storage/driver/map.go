package driver

// DriverMap work in local memory, only for test
type DriverMap struct {
	driver map[string]string
}

func (d *DriverMap) Store(id string, obj interface{}) bool {
	d.driver[id] = obj.(string)
	return true
}
func (d *DriverMap) Load(id string) (obj interface{}, ok bool) {
	obj, ok = d.driver[id]
	return obj, ok
}
