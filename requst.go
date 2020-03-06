package pubgo

import (
	"github.com/astaxie/beego/httplib"
)

func HttpGet(url string) interface{} {
	req := httplib.Get(url)
	return req
}
