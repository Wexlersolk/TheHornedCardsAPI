package services

import (
	"time"
)

/*
type Application struct{

} */

const dbTimeout = time.Second * 3

/* func New(dbPool *sql.DB) Models {
	db = dbPool
	return Models{
	}
} */

type Models struct {
	Coffe Coffe
}
