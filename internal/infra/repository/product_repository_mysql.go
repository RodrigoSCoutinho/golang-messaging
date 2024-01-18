package repository

import "database/sql"

type ProductRepositoryMysql struct {
	DB *sql.DB
}
