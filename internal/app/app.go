package internal

import (
	dbl "api-rest/internal/infrastructure/db"
	"fmt"
)

func App() {

	db, err := dbl.NewFlorECulturaDB()
	if err != nil {
		panic("Falha ao conectar o banco de dados")
	}

	fmt.Println(" conectado com sucesso", db)

}
