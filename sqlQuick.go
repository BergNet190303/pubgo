package pubgo

import (
	"github.com/astaxie/beego/orm"
	"lzh/models"
)

func Insert(table string, m map[string][]string) (saveId int64, status string) {
	var (
		keys   string
		values string
	)

	for key, value := range m {
		if keys == "" && value[0] != "" {
			keys += "`" + key + "`"
			values += "'" + value[0] + "'"
		} else if keys == "" && value[0] == "" {
			keys += "`" + key + "`"
			values += "null"
		} else if keys != "" && value[0] == "" {
			keys += ",`" + key + "`"
			values += ",null"
		} else {
			keys += ",`" + key + "`"
			values += ",'" + value[0] + "'"
		}
	}

	sql := "insert into " + table + "(" + keys + ") values(" + values + ")"
	// beego.Debug(sql)
	saveId = -0
	if res, err := models.ExecRes(sql); err != nil {
		status = "0"
		return
	} else {
		saveId = res
	}
	status = "1"
	return
}

func Update(table string, m map[string][]string, where map[string]string) (status string) {
	var (
		kv string
		w  string
	)
	for key, value := range m {
		if key == "or" {
			if kv == "" {
				kv += value[0]
			} else {
				kv += "," + value[0]
			}
		} else {
			if kv == "" && value[0] != "" {
				kv += "`" + key + "`='" + value[0] + "'"
			} else if value[0] != "" {
				kv += ",`" + key + "`='" + value[0] + "'"
			}
		}
	}
	if where == nil {
		status = "0"
		return
	}
	for key, value := range where {
		if key == "or" {
			if w == "" {
				w += value
			} else {
				w += " and " + value
			}
		} else {
			if w == "" {
				w += "`" + key + "`='" + value + "'"
			} else {
				w += " and `" + key + "`='" + value + "'"
			}
		}
	}
	sql := "update " + table + " set " + kv + " where " + w
	// beego.Debug(sql)
	if err := models.Exec(sql); err != nil {
		status = "0"
		return
	}
	status = "1"
	return
}

func Delete(table string, where map[string]string) (status string) {
	var (
		w string
	)
	if where == nil {
		status = "0"
		return
	}
	for key, value := range where {
		if key == "or" {
			if w == "" {
				w += value
			} else {
				w += " and " + value
			}
		} else {
			if w == "" {
				w += "`" + key + "`='" + value + "'"
			} else {
				w += " and `" + key + "`='" + value + "'"
			}
		}
	}
	sql := "update " + table + " set status=-1 where " + w
	// beego.Debug(sql)
	if err := models.Exec(sql); err != nil {
		status = "0"
		return
	}
	status = "1"
	return
}

func Review(table string, key string, where map[string]string) (status string) {
	var (
		w string
	)
	if where == nil {
		status = "0"
		return
	}
	for key, value := range where {
		if w == "" {
			w += "`" + key + "`='" + value + "'"
		} else {
			w += " and `" + key + "`='" + value + "'"
		}
	}
	sql := "update " + table + " set `" + key + "`=abs(`" + key + "`-1)" + " where " + w
	// beego.Debug(sql)
	if err := models.Exec(sql); err != nil {
		status = "0"
		return
	}
	status = "1"
	return
}

func GetFullData(table string, where map[string]string) []orm.Params {
	var (
		w string
	)
	for key, value := range where {
		if key != "other" {
			if w == "" {
				w += "`" + key + "`='" + value + "'"
			} else {
				w += " and `" + key + "`='" + value + "'"
			}
		}
	}
	if where["other"] != "" {
		if w == "" {
			w = where["other"]
		} else {
			w += " and " + where["other"]
		}
	}
	sql := "select * from " + table + " where " + w
	return models.Query(sql)
}
