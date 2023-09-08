package request

import "gopkg.in/resty.v1"

type Request struct {
}

func NewRequest() *Request {
	return &Request{}
}

func (r *Request) KeepAlive(url string) bool {
	result, _ := resty.New().R().Get(url)

	return result.StatusCode() == 200
}
