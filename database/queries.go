package database

import "log"

func (db *DBAccess) AddNewUser(chatId int64, tgUsername string) {
	stmt := "INSERT INTO telegram(chat_id, telegram_username) VALUES (?, ?);"
	db.Db.Exec(stmt, chatId, tgUsername)
}

func (db *DBAccess) CheckUser(chatID int64) bool {
	stmt := "SELECT * FROM telegram WHERE chat_id = ?;"
	results, err := db.Db.Query(stmt, chatID)

	defer results.Close()

	if err != nil {
		log.Println("Error occurred while checking user existence from database ->", err)
		db.logger.MakeLog("Error occurred while checking user existence from database ->" + err.Error())
	}
	if results.Next() {
		return true
	}
	return false
}

func (db *DBAccess) AddPhone(chatID int64, phone string) {
	query := "UPDATE telegram SET phone = ? WHERE chat_id = ?;"
	_, err := db.Db.Exec(query, phone, chatID)
	if err != nil {
		log.Println("Error occurred while adding phone to database ->", err)
	}
}

