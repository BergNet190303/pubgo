package publice

import (
	"fmt"
	"github.com/astaxie/beego"
	"math/rand"
	"strings"
	"time"
)

/*
* 字符串省略号提取
*/
func EllipsisString(str string,head int,foot int) string {
	var (
		Str = strings.Split(str,"")
		StrLen = len(Str)
		ellipsisNum = StrLen - head - foot
		res = ""
	)
	for i := 0; i < head; i++ {
		res += Str[i]
	}
	for i := 0; i < ellipsisNum; i++ {
		res += "*"
	}
	for i := 0; i < foot; i++ {
		res += Str[StrLen-foot+i]
	}

	return res
}

/*
* 生成随机数字
*/
func BuildRandNumber(num int) string {
	var s int32 = 1
	for i := 1; i <= num; i++ {
		s = s*10
	}
	beego.Debug(s)
	return fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(s))
}