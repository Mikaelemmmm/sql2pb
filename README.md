#### Give a star before you see it. Ha ha ha ~ ~

Generates a protobuf file from your mysql database.

### Uses
#### Use from the command line:

`go install github.com/Mikaelemmmm/sql2pb@latest`

```
$ sql2pb -h

Usage of sql2pb:
  -db string
        the database type (default "mysql")
  -go_package string
        the protocol buffer gp_package. defaults to the database schema.
  -host string
        the database host (default "localhost")
  -ignore_tables string
        a comma spaced list of tables to ignore
  -package string
        the protocol buffer package. defaults to the database schema.
  -password string
        the database password (default "root")
  -port int
        the database port (default 3306)
  -schema string
        the database schema
  -service_name string
        the protobuf service name , defaults to the database schema.
  -table string
        the table schema (default "*")
  -user string
        the database user (default "root")

```

```
$ sql2pb -go_package ./pb -host localhost -package pb -password root -port 3306 -schema usercenter -service_name usersrv -user root > usersrv.proto
```

#### Use as an imported library

```go
import "github.com/Mikaelemmmm/sql2pb"

func main() {
    connStr := config.get("dbConnStr")
    pkg := "my_package"
    goPkg := "./my_package"
    table:= "*"
    serviceName:="usersrv"

    db, err := sql.Open(*dbType, connStr)
    if err != nil {
        log.Fatal(err)
    }

    defer db.Close()

	s, err := core.GenerateSchema(db, table,nil,serviceName, goPkg, pkg)

	if nil != err {
		log.Fatal(err)
	}

	if nil != s {
		fmt.Println(s)
	}
}
```

#### Thanks for schemabuf
    schemabuf : https://github.com/mcos/schemabuf
