package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type CepBrasilAPI struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	worGroup := sync.WaitGroup{}

	c1 := make(chan CepBrasilAPI)
	c2 := make(chan ViaCEP)

	var cep string = "04870470"
	var url_via_cep string = "https://viacep.com.br/ws/" + cep + "/json/"
	var url_brasil_api string = "https://brasilapi.com.br/api/cep/v1/" + cep

	worGroup.Add(2)

	go buscarCepBrasilApi(ctx, url_brasil_api, c1, &worGroup)
	go buscarViaCep(ctx, url_via_cep, c2, &worGroup)

	select {
	case cepBrasilAPI := <-c1:
		cancel()
		fmt.Printf("API Vencedora: BrasilApi - URL: %s \n", url_brasil_api)
		cep, _ := json.Marshal(cepBrasilAPI)
		fmt.Printf("Os dados retornados da API: " + string(cep))
		close(c1)
	case cepViaCep := <-c2:
		cancel()
		fmt.Printf("API Vencedora: ViaCEP - URL: %s \n", url_via_cep)
		cep, _ := json.Marshal(cepViaCep)
		fmt.Printf("Os dados retornados da API: " + string(cep))
		close(c2)
	case <-time.After(time.Second * 1):
		cancel()
		fmt.Printf("Timeout - Ultrapassado tempo limite de 1 segundo!")

	}

	worGroup.Wait()
}

func buscarCepBrasilApi(ctx context.Context, url string, c1 chan<- CepBrasilAPI, wg *sync.WaitGroup) {
	defer wg.Done()
	//time.Sleep(time.Second * 3)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, bytes.NewReader(nil))
	if err != nil {
		//result <- Address{API: apiURL}
		return
	}

	defer req.Body.Close()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		//result <- Address{API: apiURL}
		return
	}

	defer resp.Body.Close()

	var cepBrasilApi CepBrasilAPI
	err = json.NewDecoder(resp.Body).Decode(&cepBrasilApi)
	if err != nil {
		//result <- Address{API: apiURL}
		return
	}

	c1 <- cepBrasilApi
}

func buscarViaCep(ctx context.Context, url string, c2 chan<- ViaCEP, wg *sync.WaitGroup) {
	defer wg.Done()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, bytes.NewReader(nil))
	if err != nil {
		//result <- Address{API: apiURL}
		return
	}

	defer req.Body.Close()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		//result <- Address{API: apiURL}
		return
	}

	defer resp.Body.Close()

	var viaCEP ViaCEP
	err = json.NewDecoder(resp.Body).Decode(&viaCEP)
	if err != nil {
		//result <- Address{API: apiURL}
		return
	}

	c2 <- viaCEP
}
