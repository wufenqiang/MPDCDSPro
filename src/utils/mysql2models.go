package utils

import (
	"github.com/gohouse/converter"
)

func Mysql2type(ProjectName string, mysqlpath string, tables []string) error {
	dir := ProjectLocation(ProjectName)

	var table2struct = converter.
		NewTable2Struct().
		Dsn(mysqlpath).
		PackageName("models")
	//Dsn("ry_vue_u:123456@(220.243.129.248:3306)/ry_vue?charset=utf8")

	for _, value := range tables {

		modelsName := Marshal(value)
		//modelsName :=value

		err := table2struct.Table(value).
			SavePath(dir + "/models/" + modelsName + ".go").
			Run()

		if err != nil {
			return err
		}
	}
	return nil
}
