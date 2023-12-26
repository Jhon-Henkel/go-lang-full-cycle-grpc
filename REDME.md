# Go-lang gRPC

Repositório para estudar a aplicação do gRPC

## Como rodar
- Instalar o [protoc](https://grpc.io/docs/protoc-installation/).
- Instalar os [plugins do protoc para Go-lang e gRPC](https://grpc.io/docs/languages/go/quickstart/).
- Instalar o [evans](https://github.com/ktr0731/evans)

- Criar as tabelas no banco de dados:
  ```sh
  sqlite3 db.sqlite;
  CREATE TABLE categories (id string, name string, description string);
  CREATE TABLE courses (id string, name string, description string, category_id string);
  ```
- Rodando o server:
  ```sh
  go run cmd/grpcServer/main.go
  ```
- Usando o Evans:
  ```sh
  evans -r repl # Iniciar o evans.
  package pb # Selecionar o pacote.
  service [Service Name] # Seleciona o serviço.
  call [Method Name] # Seleciona o método.
  ```

## Guias de ajuda
- [gRPC](https://grpc.io/)
- [Protocol Buffers](https://developers.google.com/protocol-buffers)

## Comandos
- `protoc --go_out=. --go-grpc_out=. path-to-protofile` - Gera os arquivos conforme o proto-file.