# Desenvolvimento de API gRPC com Go

Este projeto utiliza Go (Golang) em conjunto com Protocol Buffers (protoc) para desenvolver uma API gRPC. A API inclui a implementação de serviços normais, streams unidirecionais e bidirecionais para comunicação eficiente entre clientes e servidor.

## Tecnologias Utilizadas

- **Go (Golang)**: Linguagem de programação utilizada para desenvolver tanto o servidor quanto o cliente da API.
- **Protocol Buffers (protoc)**: Ferramenta para definição de mensagens e serviços gRPC de forma independente de linguagem.
- **gRPC**: Framework de comunicação RPC (Remote Procedure Call) que facilita o desenvolvimento de APIs eficientes e escaláveis.
- **Evans**: Ferramenta para interação e teste de APIs gRPC, utilizada para simular o cliente e testar os serviços implementados.

## Funcionalidades Implementadas

### Serviços Implementados

- **Serviços Normais**: Métodos RPC tradicionais para invocação de chamadas comuns entre cliente e servidor.
- **Streams Unidirecionais**: Para transmissão de múltiplos resultados do servidor para o cliente de forma contínua.
- **Streams Bidirecionais**: Para comunicação interativa entre cliente e servidor, permitindo envio e recebimento de múltiplos streams de dados simultaneamente.

## Estrutura do Projeto

O projeto está estruturado de acordo com as melhores práticas de organização de código em Go, incluindo:

- Diretório `protos/`: Contém os arquivos `.proto` onde são definidos os serviços e mensagens utilizados na API gRPC.
- Diretório `service/`: Contém o código-fonte do servidor  `category.go`.
- Diretório `cmd/`: Contém os programas principais (`main.go`) para execução do servidor  `cmd/GRPCserver/` 
- Diretório `internal/`: Contém a conexão com o banco de dados e suas determinadas classes. 
## Executando o Projeto

Para iniciar o servidor:

```bash
go run cmd/server/main.go
