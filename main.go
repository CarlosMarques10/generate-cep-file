package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {

	err := os.MkdirAll("ceps", 0755)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Erro ao criar a pasta 'ceps':", err)
		return
	}

	for _, url := range os.Args[1:] {
		req, err := http.Get("http://viacep.com.br/ws/" + url + "/json/")
		if err != nil {
			fmt.Fprint(os.Stderr, "Erro ao fazer a requisição: ", err)
			return
		}
		defer req.Body.Close()

		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprint(os.Stderr, "Erro ao ler a resposta: ", err)
			return
		}

		var data ViaCEP
		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Fprint(os.Stderr, "Erro ao fazer o parse da resposta: ", err)
			return
		}

		file, err := os.Create("ceps/cep.txt")
		if err != nil {
			fmt.Fprint(os.Stderr, "Erro ao criar o arquivo: ", err)
			return
		}
		defer file.Close()

		_, err = file.WriteString(fmt.Sprint(data.Cep, " ",
			data.Logradouro, " ",
			data.Complemento, " ",
			data.Localidade, " ",
			data.Bairro))
		fmt.Print("Arquivo criado com sucesso!")
	}
}
