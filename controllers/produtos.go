package controllers

import (
	"LOJAEMGO/models"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todasSalas := models.BuscaTodasSalas()
	temp.ExecuteTemplate(w, "Index", todasSalas)

}

func New(w http.ResponseWriter, r *http.Request) {

	// todasSalas := models.BuscaTodasSalas()
	// temp.ExecuteTemplate(w, "New", todasSalas)

	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		sala := r.FormValue("sala")
		descricao := r.FormValue("descricao")
		status := r.FormValue("status")

		salaConvertidaParaInt, err := strconv.Atoi(sala)
		if err != nil {
			log.Println("Erro na conversão do quantidade")
		}
		statusConvertidoParaBool, err := strconv.ParseBool(status)
		if err != nil {
			log.Println("Erro na conversão do quantidade")
		}
		models.CriarNovoProduto(salaConvertidaParaInt, descricao, statusConvertidoParaBool)

	}
	http.Redirect(w, r, "/", 301)
}
func Delete(w http.ResponseWriter, r *http.Request) {

	idDoProduto := r.URL.Query().Get("id")
	models.DeletaProduto(idDoProduto)
	http.Redirect(w, r, "/", 301)

}
func Edit(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	produto := models.EditaProduto(idDoProduto)
	temp.ExecuteTemplate(w, "Edit", produto)

}
func Update(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		id := r.FormValue("id")
		sala := r.FormValue("sala")
		descricao := r.FormValue("descricao")
		status := r.FormValue("status")

		idConvertidaParaInt, err := strconv.Atoi(id)
		if err != nil {
			log.Printf("Erro na converção do id para int: ", err)
		}
		salaConvertidaParaInt, err := strconv.Atoi(sala)
		if err != nil {
			log.Println("Erro na conversão do quantidade")
		}
		statusConvertidoParaBool, err := strconv.ParseBool(status)
		if err != nil {
			log.Println("Erro na conversão do quantidade")
		}
		models.AtualizaProduto(idConvertidaParaInt, salaConvertidaParaInt, descricao, statusConvertidoParaBool)
	}
	http.Redirect(w, r, "/", 301)
}
