package DB

import (
	"aaa/letter_box/configs"
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var (
	G_jobDB *sql.DB
	err     error
)

func InitDB() error {
	var (
		dbhostsip  = configs.G_config.Dbhostsip
		dbusername = configs.G_config.Dbusername
		dbpassword = configs.G_config.Dbpassword
		dbname     = configs.G_config.Dbname

		// dbhostsip  = "127.0.0.1:3306"
		// dbusername = "root"
		// dbpassword = "123456"
		// dbname     = "letterandfriend"
	)
	//构建连接信息

	dbinfo := strings.Join([]string{dbusername, ":", dbpassword, "@tcp(", dbhostsip, ")/", dbname, "?charset=utf8"}, "")
	fmt.Println(dbinfo)

	//打开数据库
	G_jobDB, err = sql.Open("mysql", dbinfo)
	if err != nil {
		fmt.Println("Open Database Error:", err)
		return err
	}

	// 验证连接
	if err = G_jobDB.Ping(); nil != err {
		fmt.Println("Open Database Fail,Error:", err)
		return err
	}
	fmt.Println("Connect Success!!!")

	return nil
}
