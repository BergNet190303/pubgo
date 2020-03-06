package pubgo

import (
	"lzh/models"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func VisitLog(method string, site string, url string, ip string, userAgent string, referer string, uid interface{}) {
	var (
		path     = "./visit/logs/" + models.Gettimes("2006-01-02") + "/"
		pathname = path + models.Gettimes("2006-01-02") + ".log"
	)
	var cont string
	// beego.Debug("aaa")
	cont += "[T]" + models.Gettimes("15:04:05") + " " + method + " " + site + url + " - [I]" + ip + " | [U]" + userAgent + " [W]PC "
	models.Exec("update bg_visit set day_visit=day_visit+1 where addtime='" + models.Gettimes("2006-01-02") + "'")
	if referer != "" {
		cont += "[F]" + referer
	}
	if uid != nil {
		cont += " [Uid]" + uid.(string)
	}
	var filename = pathname
	var erro error
	erro = os.MkdirAll(path, 0777)
	check(erro)
	f, erro := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	check(erro)
	_, erro = io.WriteString(f, cont+"\n")
	check(erro)
}

func check(e error) {
	if e != nil {
		fmt.Println("error =>", e)
	}
}

func checkFileIsExist(filename string) bool {
	exist := true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func ReadLog(url string) (str string, str1 []string) {
	file, err := os.OpenFile(url, os.O_RDWR|os.O_APPEND|os.O_CREATE, os.ModeType)
	check(err)
	defer file.Close()
	b, err := ioutil.ReadAll(file)
	check(err)
	str = string(b)
	a := strings.Replace(string(b), "\r\n", "\n", -1)
	str1 = strings.Split(a, "\n")
	return
}
