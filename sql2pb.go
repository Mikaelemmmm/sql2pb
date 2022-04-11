package main

import (
	"database/sql"
	"fmt"
	"github.com/Mikaelemmmm/sql2pb/config"
	"github.com/Mikaelemmmm/sql2pb/core"
	"log"
	"os"
	"path"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

const (
	defaultDirMode = 0o755
)

//保存文件函数
func saveFile(conf *config.Config, content string) {
	fileName := path.Join(conf.FilePath, conf.PackageName+".proto")
	_, err := os.Stat(conf.FilePath)
	if err != nil {
		if err := os.Mkdir(conf.FilePath, defaultDirMode); err != nil {
			panic(err)
		}
	}
	f, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	_, err = f.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	if err := config.InitConfig(); err != nil {
		panic(err)
	}
	conf := config.GetConfig()
	if conf.Database == "" {
		fmt.Println("必须指定数据库名称")
		return
	}

	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", conf.User, conf.Password, conf.Host, conf.Port, conf.Database)
	db, err := sql.Open(conf.DbType, connStr)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	ignoreTables := strings.Split(conf.IgnoreTableStr, ",")

	s, err := core.GenerateSchema(db, conf.Table, ignoreTables, conf.ServiceName, conf.GoPackageName, conf.PackageName)

	if nil != err {
		log.Fatal(err)
	}
	if s != nil {
		//将s内容保存到文件
		saveFile(conf, s.String())
	}
}
