package main

import (
	"net/http"
	"text/template"
)

type Planos struct {
	PlanoMensal string
	Preço float64
	Descritivo string

}

// Acessando aos templates(paginas) que serão usados
var temp = template.Must(template.ParseGlob("templates/*"))

func main() {

	http.HandleFunc("/", Index) //Essa Func afirma que ao receber uma requisição ( representado por : "/")
	// deve apontar para o Template Index

	//Criando o servidor no qual hospedaremos nossos templates
	http.ListenAndServe(":8000", nil)

}

func Index(w http.ResponseWriter, r *http.Request) {

	planos :=[]Planos{
		{"Plano Simples", 300, " Dois Buques de rosas personalizados ana"},
		{"Plano Premium", 650, " Quatro Buques de rosas personalizados"},
	}
	temp.ExecuteTemplate(w,"Index", planos)
}
