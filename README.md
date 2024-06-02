# Desafio Multithreading
Neste desafio você terá que usar o que aprendemos com Multithreading e APIs para buscar o resultado mais rápido entre duas APIs distintas.

As duas requisições serão feitas simultaneamente para as seguintes APIs:
- https://brasilapi.com.br/api/cep/v1/01153000 + cep
- http://viacep.com.br/ws/" + cep + "/json/

Os requisitos para este desafio são:
- Acatar a API que entregar a resposta mais rápida e descartar a resposta mais lenta.
- O resultado da request deverá ser exibido no command line com os dados do endereço, bem como qual API a enviou.
- Limitar o tempo de resposta em 1 segundo. Caso contrário, o erro de timeout deve ser exibido.

## Estrutura do Código
O programa principal está no arquivo main.go no diret;orio raiz. Foram criadas duas estruturas uma para cada retorno das apis listadas no desafio.
- CepBrasilAPI
- ViaCEP

Sobre main.go: Onde é iniciado o programa. Primeiramente são definidos o contexto para realizar as requisições, logo são definidas as variáveis que serão utilizadas durante o processamento como CEP a ser consultado, e as urls dos serviços que serão comparados. Também é definido um WorgGroup para ajudar no controle de execução das goroutines. Após é dado o start das duas goroutines onde cada goroutine realiza a chamada de um método para realizar o request das apis.

Funções - Request - API
- buscarCepBrasilApi
- buscarViaCep

Após realizado o request de cada api e retornado para seus respectivos canais(c1,c2) o trecho do código onde está o select é responsável por expor os dados e a api que venceu o desafio.

Importante lembrar que para os casos onde ultrapassar o tempo de 1 segundo o programa retornará um erro informando timout

## Como Rodar o Código
Clonar o repositório:

```bash
git clone https://github.com/mbrocco/desafio-multithreading.git
```
## Navegue até o diretório do projeto:

```bash
cd desafio-multithreading
```

## Execute o código Go:

```bash
go mod tidy
go run main.go
```


## Retorno do programa:

```bash
go run main.go

API Vencedora: BrasilApi - URL: https://brasilapi.com.br/api/cep/v1/04870470 
Os dados retornados da API: {"cep":"04870470","state":"SP","city":"São Paulo","neighborhood":"Chácara Santo Hubertus","street":"Estrada da Servidão","service":"open-cep"}%
```
