package repository

import (
	"database/sql"

	"github.com/RodrigoSCoutinho/golang-messaging/internal/entity"
)

type ProductRepositoryMysql struct {
	DB *sql.DB
}

func NewProductRepositoryMysql(db *sql.DB) *ProductRepositoryMysql {
	return &ProductRepositoryMysql{DB: db}
}

func (r *ProductRepositoryMysql) Create(product *entity.Product) error {
	_, err := r.DB.Exec("INSERT INTO products (name, price) VALUES (?,?,?)", product.ID, product.Name, product.Price)
	//INSERINDO OS DADOS NO BANCO DE DADOS MYSQL E RETORNANDO UM ERROR CASO OCORRA
	if err != nil {
		return err
	}
	return nil
}

func (r *ProductRepositoryMysql) FindAll() ([]*entity.Product, error) {
	//RETORNANDO UMA LISTA DE PRODUTOS E UM ERROR CASO OCORRA

	rows, err := r.DB.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}
	//FECHANDO A CONEXÃO COM O BANCO DE DADOS
	defer rows.Close()

	var products []*entity.Product

	//PERCORRENDO A LISTA DE PRODUTOS E ADICIONANDO OS PRODUTOS NA LISTA
	for rows.Next() {
		product := &entity.Product{}
		err := rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}
