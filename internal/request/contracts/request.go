package contracts

type Request interface {
	KeepAlive(url string) bool
}
