package mysql

import (
	"context"
	"database/sql"
	"log"

	config "github.com/Kushmanda-Tech/tyedigital_backend/Config"
	dbcreationscript "github.com/Kushmanda-Tech/tyedigital_backend/Db/MySql/DbCreationScript"

	"github.com/gin-gonic/gin"
)

var (
	DB  *sql.DB
	err error
)

func Connect(c *gin.Context) {
	var MYSQL_URL_WITHOUT_DATABASE string = config.DbConfig().DB_USERNAME + ":" + config.DbConfig().DB_PASSWORD + "@tcp(" + config.DbConfig().DB_HOST + ":" + config.DbConfig().DB_PORT + ")/"
	var MYSQL_URL_WITH_DATABASE string = config.DbConfig().DB_USERNAME + ":" + config.DbConfig().DB_PASSWORD + "@tcp(" + config.DbConfig().DB_HOST + ":" + config.DbConfig().DB_PORT + ")/" + config.DB_NAME + "?multiStatements=true"

	log.Println(MYSQL_URL_WITHOUT_DATABASE)
	log.Println(MYSQL_URL_WITH_DATABASE)

	CreateMySQldatabse(c, MYSQL_URL_WITHOUT_DATABASE)
	DB, err = sql.Open("mysql", MYSQL_URL_WITH_DATABASE)
	if err != nil {
		log.Println("Error To Connect Databae")
	}

	DB, err = sql.Open("mysql", MYSQL_URL_WITH_DATABASE)
	if err != nil {
		log.Println("Error To Connect Databae")
	}

	if CreateMysqlTableErr := CreateMysqlTable(); CreateMysqlTableErr != nil {
		c.JSON(config.FALIURE, gin.H{
			"error": CreateMysqlTableErr,
		})
	}

	// if CreateMysqlFunctionErr := CreateMysqlFunction(); CreateMysqlFunctionErr != nil {
	// 	c.JSON(config.FALIURE, gin.H{
	// 		"error": CreateMysqlFunctionErr,
	// 	})
	// }

	// if CreateMysqlProcedureErr := CreateMysqlProcedure(); CreateMysqlProcedureErr != nil {
	// 	c.JSON(config.FALIURE, gin.H{
	// 		"error": CreateMysqlProcedureErr,
	// 	})
	// }

	// if CreateMysqlViewsErr := CreateMysqlViews(); CreateMysqlViewsErr != nil {
	// 	c.JSON(config.FALIURE, gin.H{
	// 		"error": CreateMysqlViewsErr,
	// 	})
	// }
	validatePreDefineRow()

	// if alterTableMysqlErr := AlterTableMysql(); alterTableMysqlErr != nil {
	// 	c.JSON(config.FALIURE, gin.H{
	// 		"error": alterTableMysqlErr,
	// 	})
	// }
}
func CreateMySQldatabse(c *gin.Context, url string) {
	databseCon, err := sql.Open("mysql", url)
	if err != nil {
		log.Println("Error To Connect Databae")
		c.JSON(config.FALIURE, gin.H{
			"error": "Error To Connect Databae",
		})
	}
	if _, dbExecErr := databseCon.Exec(dbcreationscript.DB_CREATION_SQLQUERY); dbExecErr != nil {
		log.Println("**********MySqlDatabse********************")
		log.Println(dbExecErr.Error())
		log.Println("********* MySqlDatabse********************")
	}
}

func CreateMysqlTable() error {
	if _, tableExecErr := DB.Exec(dbcreationscript.TABLE_CREATION_SQLQUERY); tableExecErr != nil {
		log.Println("**************MySqlTable****************")
		log.Println(tableExecErr.Error())
		log.Println("***************MySqlTable***************")
		return tableExecErr
	}
	log.Println("create table successfully")
	return nil
}

func CreateMysqlViews() error {
	if _, tableExecErr := DB.ExecContext(context.Background(), dbcreationscript.VIEWS_CREATION_SQLQUERY); tableExecErr != nil {
		log.Println("**************MySqlVIEWS****************")
		log.Println(tableExecErr.Error())
		log.Println("***************MySqlVIEWS***************")
		return tableExecErr
	}
	log.Println("create views successfully")
	return nil
}

func CreateMysqlFunction() error {
	if _, tableExecErr := DB.ExecContext(context.Background(), dbcreationscript.FUNCTION_CREATION_SQLQUERY); tableExecErr != nil {
		log.Println("**************MySqlFunction****************")
		log.Println(tableExecErr.Error())
		log.Println("***************MySqlFunction***************")
		return tableExecErr
	}
	log.Println("create Function successfully")
	return nil
}

func CreateMysqlProcedure() error {
	if _, tableExecErr := DB.ExecContext(context.Background(), dbcreationscript.PROCEDURE_CREATION_SQLQUERY); tableExecErr != nil {
		log.Println("**************MySqlProcedure****************")
		log.Println(tableExecErr.Error())
		log.Println("***************MySqlProcedure***************")
		return tableExecErr
	}
	log.Println("create procedure successfully")
	return nil
}

// func AlterTableMysql() error {
// 	if _, tableExecErr := DB.ExecContext(context.Background(), dbcreationscript.ALTER_TABLE_CREATION_SQLQUERY); tableExecErr != nil {
// 		log.Println("**************AlterMySqlTable****************")
// 		log.Println(tableExecErr.Error())
// 		log.Println("***************AlterMySqlTable***************")
// 		return tableExecErr
// 	}
// 	log.Println("create Alter Table successfully")
// 	return nil
// }

func AddPreDefinedRowMysql() error {
	if _, tableExecErr := DB.ExecContext(context.Background(), dbcreationscript.ADD_PRE_DEFINED_ROW); tableExecErr != nil {
		log.Println("**************AddPreDefinedRowMysql****************")
		log.Println(tableExecErr.Error())
		log.Println("***************AddPreDefinedRowMysql***************")
		return tableExecErr
	}
	log.Println("Add pre defined row adding table successfully")
	return nil
}

func validatePreDefineRow() {
	studentProfile := CountPreDefineRowMysql("sd_student_profile")
	studentDetails := CountPreDefineRowMysql("sd_student_details")
	if studentProfile == 13 && studentDetails == 13 {
		return
	} else {
		if preDefinedErr := AddPreDefinedRowMysql(); preDefinedErr != nil {
			log.Println("preDefinedErr error ", preDefinedErr.Error())
		}
	}
}
func CountPreDefineRowMysql(tableName string) int {
	var row int
	if countPreDefineRowMysqlErr := DB.QueryRow(`SELECT COUNT(*) FROM ` + tableName + ``).Scan(&row); countPreDefineRowMysqlErr != nil {
		log.Println("countPreDefineRowMysqlErr QueryRow ", countPreDefineRowMysqlErr.Error())
	}
	return row
}
