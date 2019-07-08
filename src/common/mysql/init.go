/*
* @Author: Luke
* @Date:   2019年7月7日16:51:12
 */

package mysql

import (
	"net/url"
	"os"
	"webServer/src/common/config"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	. "github.com/soekchl/myUtils"
)

var (
	m_orm orm.Ormer
)

func init() {
	// config.Config("../../../config/config.ini") // test config.ini
	dbhost := config.GetString("mysql.host")
	dbport := config.GetString("mysql.port")
	dbuser := config.GetString("mysql.user")
	dbpassword := config.GetString("mysql.password")
	dbname := config.GetString("mysql.name")
	timezone := config.GetString("mysql.timezone")
	if dbport == "" {
		dbport = "3306"
	}
	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"

	if timezone != "" {
		dsn = dsn + "&loc=" + url.QueryEscape(timezone)
	}
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		Error("RegisterDriver Mysql Err=", err)
		os.Exit(1)
	}
	err = orm.RegisterDataBase("default", "mysql", dsn)
	if err != nil {
		Error("dsn:", dsn, "\n", err)
		os.Exit(1)
	}
	orm.RegisterModel(
		new(TestTable),
	)

	orm.SetMaxIdleConns("default", 1000)
	orm.SetMaxOpenConns("default", 20000)
	orm.RunSyncdb("default", false, true)

	/*	if config.GetString("runmode") == "dev" {
			orm.Debug = true
		}
	*/
	m_orm = orm.NewOrm()
}
