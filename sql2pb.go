package main

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/Mikaelemmmm/sql2pb/core"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbType := flag.String("db", "mysql", "the database type")
	host := flag.String("host", "localhost", "the database host")
	port := flag.Int("port", 3306, "the database port")
	user := flag.String("user", "root", "the database user")
	password := flag.String("password", "root", "the database password")
	schema := flag.String("schema", "", "the database schema")
	table := flag.String("table", "*", "the table schema，multiple tables ',' split. ")
	serviceName := flag.String("service_name", *schema, "the protobuf service name , defaults to the database schema.")
	packageName := flag.String("package", *schema, "the protocol buffer package. defaults to the database schema.")
	goPackageName := flag.String("go_package", "", "the protocol buffer go_package. defaults to the database schema.")
	ignoreTableStr := flag.String("ignore_tables", "", "a comma spaced list of tables to ignore")
	ignoreColumnStr := flag.String("ignore_columns", "", "a comma spaced list of mysql columns to ignore")
	fieldStyle := flag.String("field_style", "sqlPb", "gen protobuf field style, sql_pb | sqlPb")

	flag.Parse()

	if *schema == "" {
		fmt.Println(" - please input the database schema ")
		return
	}

	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", *user, *password, *host, *port, *schema)
	db, err := sql.Open(*dbType, connStr)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	ignoreTables := strings.Split(*ignoreTableStr, ",")
	ignoreColumns := strings.Split(*ignoreColumnStr, ",")

	s, err := core.GenerateSchema(db, *table, ignoreTables, ignoreColumns, *serviceName, *goPackageName, *packageName, *fieldStyle)

	if nil != err {
		log.Fatal(err)
	}

	if nil != s {
		fmt.Println(s)
	}
}
