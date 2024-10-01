package main

import (
	app "api-rest/internal/app"

	"github.com/spf13/viper"
)

type Planos struct {
	PlanoMensal string
	Preço       float64
	Descritivo  string
}

// Acessando aos templates(paginas) que serão usados

func main() {

	// Carrega a configuração a partir do arquivo JSON

	app.App()

}

func init() { //configurando o viper para abrir configs.json onde está localizado minha variavel de ambiente
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.AddConfigPath("internal/infrastructure/db")

	err := viper.ReadInConfig() // lê de fato config.json

	if err != nil {
		panic("Erro fatal no arquivo de configuração")
	}
}
