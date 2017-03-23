package repository

import (
	"database/sql"
	"github.com/abaeve/auth-srv/model"
)

type accesses struct {
	db *sql.DB
}

func (acc *accesses) findByUserId ( userId int64 ) []string {
	return nil
}

// Will be the main usage in anything automated.  This method will lookup all the available roles for the given
// Chat user id and return them as a map[string]string where the key is the role and the value is the chat
// group to apply.
func (acc *accesses) findByChatId ( chatId string ) []string {
	return nil
}

func (acc *accesses) findByCharacter ( character model.Character ) map[string]string {
	return nil
}
