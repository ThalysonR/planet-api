# Planet API

API para busca e cadastro de planetas do Star Wars. Para executar este projeto, as seguintes variáveis de ambiente são necessárias:

- SERVER_HOST - Endereço do servidor que irá executar este serviço
- MONGO_URI - Endereço do banco de dados MongoDB
- SW_API - Endereço da API do star wars
- GIN_MODE - Modo de execução do Gin. Por padrão, executa como "debug"

Valores de exemplo podem ser vistos no _docker-compose.yml_.

Quando em execução, a documentação da API pode ser acessada em http://localhost:8080/swagger/index.html (Executando com as configurações atuais)

## Tecnologias

As seguintes tecnologias foram usadas no projeto:

- Go
- Gin
- Swagger
- MongoDB
- Docker

## Execução Local

Para executar o projeto, é necessário que esteja instalada a aplicação swag

`go get -u github.com/swaggo/swag/cmd/swag`

Então execute o comando `swag init` para gerar os arquivos de documentação da API. Após isso, a aplicação pode ser executada com `go run main.go`, dado que as variáveis de ambiente necessárias sejam fornecidas.

## Execução com Docker

Simplesmente modifique as variáveis de ambiente no arquivo _docker-compose.yml_ conforme a necessidade e execute o comando `docker-compose up`.
