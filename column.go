package publice

import "lzh/models"

func GetColumn(types string,level int) map[string]interface{} {
	var js = make(map[string]interface{})
	level1 := models.Query("select * from bg_column where level=1 and type="+types+" and status>=0")
	js["level1"] = level1
	switch level {
		case 2:
			if level1 != nil {
				js["level2"] = models.Query("select * from bg_column where level=2 and type="+types+" and status>=0 and cid=" + level1[0]["id"].(string))
			}
		default:
	}
	return js
}