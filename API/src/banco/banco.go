package banco

import (
	"api/src/config"
	"database/sql"
	_ "github.com/go-sql-driver/mysql" //Driver
)

// Conectar abre a conexao com o DB e a retorna
func Conectar() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.StringConexaoBanco)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
