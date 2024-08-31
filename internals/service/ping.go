package service

import db "rest/api/internals/db/sqlc"

type PingService struct {
	Store db.Store 
}

func (s PingService) Ping() []byte {
	return []byte("PONG! PONG!!")
}