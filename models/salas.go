package models

import (
	"chaves_furg/db"
)

type Salas struct {
	Id        int
	Sala      int
	Descricao string
	Status    bool
}

func BuscaTodasSalas() []Salas {
	db := db.ConectaCombancoDeDados()
	selectDeTodasSalas, err := db.Query("select * from salas_c3 order by id asc")
	if err != nil {
		panic(err.Error())
	}

	p := Salas{}
	salas := []Salas{}

	for selectDeTodasSalas.Next() {
		var id, sala int
		var descricao string
		var status bool

		err = selectDeTodasSalas.Scan(&id, &sala, &descricao, &status)
		if err != nil {
			panic(err.Error())
		}
		p.Id = id
		p.Sala = sala
		p.Descricao = descricao
		salas = append(salas, p)
	}

	defer db.Close()
	return salas
}
func CriarNovoProduto(sala int, descricao string, status bool) {
	db := db.ConectaCombancoDeDados()

	insereDadosNoBanco, err := db.Prepare("insert into salas_c3(sala, descricao, status)values($1, $2, $3)")
	if err != nil {
		panic(err.Error())
	}
	insereDadosNoBanco.Exec(sala, descricao, status)
	defer db.Close()
}
func DeletaProduto(id string) {

	db := db.ConectaCombancoDeDados()
	deletarOProduto, err := db.Prepare("delete from salas_c3 where id=$1")
	if err != nil {
		panic(err.Error())
	}
	deletarOProduto.Exec(id)
	defer db.Close()

}

func EditaProduto(id string) Salas {
	db := db.ConectaCombancoDeDados()

	produtoDoBanco, err := db.Query("select *from salas_c3 where id=$1", id)
	if err != nil {
		panic(err.Error())
	}
	produtoParaAtualizar := Salas{}

	for produtoDoBanco.Next() {
		var id, sala int
		var descricao string
		var status bool

		err = produtoDoBanco.Scan(&id, &sala, &descricao, &status)
		if err != nil {
			panic(err.Error())
		}
		produtoParaAtualizar.Id = id
		produtoParaAtualizar.Sala = sala
		produtoParaAtualizar.Descricao = descricao
		produtoParaAtualizar.Status = status
	}
	defer db.Close()
	return produtoParaAtualizar

}
func AtualizaProduto(id int, sala int, descricao string, status bool) {
	db := db.ConectaCombancoDeDados()
	atualizaProduto, err := db.Prepare("update salas_c3 set sala=$1, descricao=$2, status=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}
	atualizaProduto.Exec(sala, descricao, status, id)
	defer db.Close()
}
