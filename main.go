package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

var db *sql.DB
//初始化数据库
func initDB() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/sql_test"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return
	}
	err = db.Ping()
	if err != nil {

		return
	}

	return
}
//用户
type user struct {
	id   int
	name string
	age  int
}

func main() {

	err := initDB()
	if err != nil {
		fmt.Printf("initDB failed,err:%v\n", err)
	}
	fmt.Println("连接数据库成功！")
	//查询ID=3的ROW，假设其不存在导致sql.ErrNoRows
	err=tryQuery(3)
	fmt.Print(err)
	

}

func tryQuery(id int)error{
	var u1 user
	sqlStr := `select id, name, age from user where id=?;`
	rowObj := db.QueryRow(sqlStr, id)
	err := rowObj.Scan(&u1.id, &u1.name, &u1.age)
	//Wrap 这个 error，抛给上层	
	return errors.Wrap(err,"Wraped something useful from tryQuery")
	
}
