package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"notificationBot/logger"
)

type DBAccess struct {
	Db *sql.DB
	logger *logger.Logger
}

func OpenDB(logger *logger.Logger) (DBAccess, error) {

	/*_ = godotenv.Load("globals.env")
	db_username := os.Getenv("db_username")
	db_password := os.Getenv("db_password")
	server_address := os.Getenv("server_address")
	dp_port := os.Getenv("db_port")
	db_name := os.Getenv("db_name")*/
	db, err := sql.Open("mysql", "puck:tidolbaeb?@tcp(185.22.64.115:3306)/qrservice?parseTime=true")

	if err != nil {
		logger.MakeLog("Error occurred while setting the connection with DB"+err.Error())
		panic(err.Error())
	}

	res := DBAccess{db, logger}
	return res, err
}