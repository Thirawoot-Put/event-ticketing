package port

type HTTPContext interface {
	JSON(int, interface{})
}
