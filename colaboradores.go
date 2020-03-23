package main

import (
	// "io/ioutil"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// Estrutura de um colaborador
// type Colaborador struct {
// 	cpfPessoa            string `json:"cpfPessoa"`
// 	nomeCompletoPessoa   string `json:"nomeCompletoPessoa"`
// 	dataNascimentoPessoa string `json:"dataNascimentoPessoa"`
// 	sexoPessoa           string `json:"sexoPessoa"`
// 	nomeCargo            string `json:"nomeCargo"`
// }

// Um slice de colaboradores
// type Colaboradores struct {
// 	Colaboradores []Colaborador
// }

func loadColaboradores() []byte {
	// Fazendo a abertura do arquivo json
	jsonFile, err := os.Open("colaboradores.json")
	if err != nil {
		panic(err.Error())
	}

	defer jsonFile.Close()

	// Fazendo a leitura do arquivo json
	data, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic(err.Error())
	}
	// Retornando os dados
	return data
}

func ListColaboradores(w http.ResponseWriter, r *http.Request) {
	colaboradores := loadColaboradores()
	w.Write([]byte(colaboradores))
}

// getColaboradorByCpf(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	data := loadColaboradores()

// 	var colaboradores
// }

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/colaboradores", ListColaboradores)
	http.ListenAndServe(":8081", r)
}
