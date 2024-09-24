package main

import (
	app "api-rest/internal/app"
	db "api-rest/internal/infraestucture"
	"api-rest/internal/service"
	"fmt"
	"log"
	"text/template"
)

type Planos struct {
	PlanoMensal string
	Preço       float64
	Descritivo  string
}

// Acessando aos templates(paginas) que serão usados
var temp = template.Must(template.ParseGlob("templates/*"))

func main() {

	// Carrega a configuração a partir do arquivo JSON
	config, err := db.LoadConfig("internal/infraestucture/config.json")
	if err != nil {
		log.Fatalf("Erro ao carregar configuração: %v", err)
	}

	// Conecta ao banco de dados usando a configuração carregada
	database, err := app.ConectionDb(config)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	
	selecionaTudo,_ := service.GetAll(database)
	fmt.Println(selecionaTudo)
}

//chamando app.go
//db, err := int.ConectionDb()
//if err != nil {
//fmt.Println("erro ao conectar ao banco de dados", err)
//}

//-------Pega o produto com id 1 do banco---------

//produtos:=core.Produtos{Id: 1}

// if err := db.Table("produtos").First(&produtos).Error; err!= nil {
//fmt.Println(err)

// }
//fmt.Println(&produtos)

// ----Criando um produto-----
//produtos:= core.Produtos{
//Nome:" teste",
//}
//if err := db.Table("produtos").Create(&produtos).Error; err!= nil{
//fmt.Println(err)
//}
//fmt.Println(&produtos)

// ---------Seleciona todos os produtos da tabela "produtos"----------
//var produtos []core.Produtos

//if err := db.Table("produtos").Find(&produtos).Error; err != nil {
//fmt.Println("Erro ao buscar os produtos:", err)
//return
//}

// Itera sobre os produtos e imprime cada um
//for _, produto := range produtos {
//fmt.Println(produto)
//}

//http.HandleFunc("/", Index) //Essa Func afirma que ao receber uma requisição ( representado por : "/")
// deve apontar para o Template Index

//Criando o servidor no qual hospedaremos nossos templates
//http.ListenAndServe(":8000", nil)

//}

//func Index(w http.ResponseWriter, r *http.Request) {
//db,_:= app.ConectionDb()
//selecionaTudo, _:= service.GetAll(db)

//temp.ExecuteTemplate(w, "Index", selecionaTudo)
//}
