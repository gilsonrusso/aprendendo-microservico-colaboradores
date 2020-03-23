package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// Estrutura de um colaborador
type Colaborador struct {
	CpfPessoa            string `json:"cpfPessoa"`
	NomeCompletoPessoa   string `json:"nomeCompletoPessoa"`
	DataNascimentoPessoa string `json:"dataNascimentoPessoa"`
	SexoPessoa           string `json:"sexoPessoa"`
	NomeCargo            string `json:"nomeCargo"`
}

// Um slice de colaboradores
type Colaboradores struct {
	Colaboradores []Colaborador
}

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

func GetColaboradorByCpf(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	data := loadColaboradores()

	var colaboradores Colaboradores
	json.Unmarshal(data, &colaboradores)

	for _, v := range colaboradores.Colaboradores {
		if v.CpfPessoa == vars["cpfPessoa"] {
			colaborador, _ := json.Marshal(v)
			w.Write([]byte(colaborador))
		}
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/colaboradores", ListColaboradores)
	r.HandleFunc("/colaboradores/{cpfPessoa}", GetColaboradorByCpf)
	http.ListenAndServe(":8081", r)
}
