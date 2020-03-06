package pubgo

import (
	"encoding/base64"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"

	// "fmt"
	"lzh/models"
	"os"
	"path"
	"strings"
)

var (
	host = "47.75.152.254"
)

type UploadImgBase64 struct {
	beego.Controller
}

func Dir(file, dirname, timestamp, host string, final string) (savepath, pathname, allpathname, path, dirpath string) {
	rand := BuildRandNumber(6)
	savepath = "./static/file/" + file + "/" + dirname + "/" + timestamp + rand + "." + final
	pathname = "http://" + host + "/static/file/" + file + "/" + dirname + "/" + timestamp + rand + "." + final
	allpathname = "http://" + host + "/static/file/" + file + "/" + dirname + "/" + timestamp + rand + "." + final
	path = "http://" + host + "/static/file/" + file + "/" + dirname
	dirpath = "./static/file/" + file + "/" + dirname
	return
}

func (s *UploadImgBase64) Post() {
	host = s.Ctx.Request.Host
	defer s.ServeJSON()
	var (
		file  = s.Input().Get("file")
		base  = file[strings.IndexByte(file, ',')+1:]
		js    = make(map[string]interface{})
		where = "default"
	)
	if s.GetString("where") != "" {
		where = s.GetString("where")
	}
	savepath, pathname, _, path, dirpath := Dir(where, models.Getdate(), models.Gettimestemps(), s.Ctx.Request.Host, "png")
	imgFile, _ := base64.StdEncoding.DecodeString(base)
	os.MkdirAll(dirpath, 0777)
	if err := ioutil.WriteFile(savepath, imgFile, 0666); err != nil {
		beego.Debug("打开文件失败:", err)
	}
	var img models.Bg_resource
	img.Path = path
	img.Pathname = pathname
	img.Status = "1"
	img.Type = where
	img.Localpath = dirpath
	img.Localpathname = savepath
	img.Addtime = models.Gettime()
	img.Ip = s.Ctx.Request.RemoteAddr
	if id, err := models.Orm().Insert(&img); err != nil {
		beego.Debug("sqlErr =>", err)
	} else {
		js["id"] = id
	}
	js["pathname"] = pathname
	s.Data["json"] = js
}

type UploadImgFile struct {
	beego.Controller
}

func (s *UploadImgFile) Post() {
	host = s.Ctx.Request.Host
	defer s.ServeJSON()
	//fmt.Println(s.GetFile("file"))
	var (
		f, h, _ = s.GetFile("file")
		ext     = path.Ext(h.Filename)
		js      = make(map[string]interface{})
		where   = "default"
	)
	if s.GetString("where") != "" {
		where = s.GetString("where")
	}
	savepath, pathname, _, path, dirpath := Dir(where, models.Getdate(), models.Gettimestemps(), s.Ctx.Request.Host, strings.Split(ext, ".")[1])
	//beego.Debug(f)
	//beego.Debug(h)
	defer f.Close()
	os.MkdirAll(dirpath, 0777)
	s.SaveToFile("file", savepath)
	var img models.Bg_resource
	img.Path = path
	img.Pathname = pathname
	img.Status = "1"
	img.Type = where
	img.Localpath = dirpath
	img.Localpathname = savepath
	img.Addtime = models.Gettime()
	img.Ip = s.Ctx.Request.RemoteAddr
	if id, err := models.Orm().Insert(&img); err != nil {
		beego.Debug("sqlErr =>", err)
		s.Data["json"] = FiledDef("上传失败，未知错误！")
		return
	} else {
		js["id"] = id
	}
	js["uploaded"] = 1
	js["url"] = pathname
	s.Data["json"] = js
}

type ImgsAjax struct {
	beego.Controller
}

//根据上传的地方进行分类存放图片
func (s *ImgsAjax) Post() {
	defer s.ServeJSON()
	var js = make(map[string]interface{})
	var errtext = make(map[string]interface{})
	s.Data["json"] = &js
	f, _, err := s.GetFile("upload")
	defer f.Close()
	if err != nil {
		beego.Debug(err)
		js["upload"] = 0
		errtext["message"] = err
		js["error"] = errtext
	} else {
		savepath, pathname, _, name, dirpath := Dir("ckeditor", models.Getdate(), models.Gettimestemps(), s.Ctx.Request.Host, "png")
		if err = os.MkdirAll(dirpath, 0777); err != nil {
			js["uploaded"] = 0
			beego.Debug(err)
			errtext["message"] = err
			js["error"] = errtext
		} else {
			if err = s.SaveToFile("upload", savepath); err != nil {
				js["uploaded"] = 0
				beego.Debug(err)
				errtext["message"] = err
				js["error"] = errtext
			} else {
				js["uploaded"] = 1
				js["fileName"] = name
				js["url"] = pathname
			}
		}
	}
}

type ImgsList struct {
	beego.Controller
}

func (s *ImgsList) Post() {
	defer s.ServeJSON()
	column := models.Query("select `type` from bg_resource where status=1 GROUP BY type")
	for _, v := range column {
		data := models.Query("Select id,addtime as name,pathname,`type` from bg_resource where status=1 and `type`='" + v["type"].(string) + "' order by addtime desc limit 0,50")
		for _, vv := range data {
			vv["code"] = false
		}
		v["data"] = data
	}
	s.Data["json"] = column
}

func (s *ImgsList) ResourceDel() {
	defer s.ServeJSON()
	if models.Exec(fmt.Sprintf("update bg_resource set status='-1' where id in (%s)", s.GetString("ids"))) != nil {
		s.Data["json"] = FiledDef("删除失败，未知错误")
		return
	}
	s.Data["json"] = SuccessDef("删除成功")
}

func (s *ImgsList) ResourceMove() {
	defer s.ServeJSON()
	if models.Exec(fmt.Sprintf("update bg_resource set type='%s' where id in (%s)", s.GetString("where"), s.GetString("ids"))) != nil {
		s.Data["json"] = FiledDef("移动失败，未知错误")
		return
	}
	s.Data["json"] = SuccessDef("移动成功")
}
