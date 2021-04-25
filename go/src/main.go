package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Programa Full Cycle de Aceleração: Edição Docker!\n</h1>")
	fmt.Fprintf(w, "<h3>Desafio 1</h3>")

	cnn, err := sql.Open("mysql", "MainUser:MainPassword@tcp(desafio-01-mysql:3306)/desafio-01")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := cnn.Query("SELECT id, name FROM fcmodule ORDER BY id ASC")
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(w, "<p>Módulo %d: %s</p>", id, name)
		fmt.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "<h3>Fim da lista de módulos</h3>")
	fmt.Println("\nFim da lista de módulos")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func main() {

	fmt.Println("Servico ativo!")
	handleRequests()
}
