## CleanArq Challenge

Esta é uma aplicação de exemplo seguindo os princípios da Arquitetura Limpa (*Clean Architecture*) em Go. A aplicação consiste em um sistema de gerenciamento de pedidos (orders).

## Estrutura do Projeto
A estrutura do projeto é organizada da seguinte forma:

- **main.go**: Este arquivo é o ponto de entrada da aplicação, onde a inicialização de diversos componentes ocorre, como a configuração do servidor *HTTP*, *gRPC*, *GraphQL*, conexão com o banco de dados e injeção de dependências usando *Google Wire*.

- **database/**: Este diretório contém o código relacionado ao acesso ao banco de dados. Inclui a definição de interfaces e implementações de repositórios para interação com o banco de dados *MySQL*.

- **infra/** : Este diretório contém o código relacionado à infraestrutura da aplicação, como os manipuladores HTTP e gRPC. Também inclui a configuração do servidor *web*, *gRPC* server e *GraphQL* server.

- **internal/**: Este diretório contém o código interno da aplicação, como os casos de uso. Os casos de uso são responsáveis por implementar a lógica de negócios da aplicação e são independentes da infraestrutura e do framework.

- **pkg/**: Este diretório contém pacotes reutilizáveis em toda a aplicação, como eventos, que podem ser usados em diferentes partes da aplicação.

## Como Executar

Para executar a aplicação, basta rodar o seguinte comando:

go run main.go wire_gen.go

Starting web server on port :8000
Starting gRPC server on port 50051
Starting GraphQL server on port 8080

## Principais Dependências

Este projeto utiliza as seguintes dependências:

*gqlgen v0.17.41*: Uma biblioteca para construir servidores GraphQL em Go.
*chi v5.0.11*: Um pacote para criação de rotas HTTP em Go.
*mysql v1.7.1*: Um driver para MySQL para a biblioteca de banco de dados sql do Go.
*protobuf v1.5.3*: A implementação Go do Google Protocol Buffers.
*uuid v1.4.0*: Uma biblioteca para geração de UUIDs.
*wire v0.5.0*: Uma biblioteca para injeção de dependência em Go.
*amqp091-go v1.9.0*: Uma biblioteca para trabalhar com o RabbitMQ em Go.
*viper v1.18.2*: Uma biblioteca para configuração de aplicativos em Go.
*gqlparser v2.5.10*: Um parser para GraphQL em Go.
*grpc v1.60.1*: A implementação Go do gRPC: um sistema de RPC de alto desempenho.
*protobuf v1.32.0*: A implementação Go do Google Protocol Buffers.

## Casos de Uso

A aplicação possui os seguintes casos de uso:

**Criar Pedido (CreateOrderUseCase)**: Este caso de uso é responsável por criar um novo pedido no sistema. Recebe as informações do pedido, como preço e taxa, valida e salva no banco de dados. Também dispara um evento de "Pedido Criado" para notificar outros sistemas sobre a criação do pedido.

**Listar Pedidos (ListOrderUseCase)**: Este caso de uso é responsável por listar todos os pedidos no sistema. Recupera os pedidos do banco de dados e retorna uma lista de pedidos. Também dispara um evento com todo o paylod retornado.

**Buscar Pedido por ID (GetOrderByIDUseCase)**: Este caso de uso é responsável por buscar um pedido específico pelo seu ID. Retorna o pedido correspondente, se existir. Também dispara um evento com todo o paylod retornado.

**Atualizar Pedido (UpdateOrderUseCase)**: Este caso de uso é responsável por atualizar as informações de um pedido existente. Recebe as novas informações do pedido, valida e atualiza no banco de dados. Também dispara um evento de "Pedido Atualizado" para notificar outros sistemas sobre a atualização do pedido.

**Excluir Pedido (DeleteOrderUseCase)**: Este caso de uso é responsável por excluir um pedido existente do sistema. Recebe o ID do pedido a ser excluído, remove do banco de dados e dispara um evento de "Pedido Excluído" para notificar outros sistemas sobre a exclusão do pedido.

## Exemplos de Uso

A aplicação fornece uma API GraphQL para interação com os pedidos. Abaixo estão alguns exemplos de consultas e mutações que podem ser realizadas:
O GraphQL Playground pode ser acessada no endereço http://localhost:8080/.

**Consultar Todos os Pedidos**:

```graphql
query {
  orders {
    id
    price
    tax
    finalPrice
  }
}
```

**Criar um Novo Pedido**:

```graphql
mutation {
  createOrder(input: { price: 100.0, tax: 10.0 }) {
    id
    price
    tax
    finalPrice
  }
}
```

**Buscar Pedido por ID**:

```graphql
query {
  order(id: "abc123") {
    id
    price
    tax
    finalPrice
  }
}
```

**Atualizar um Pedido Existente**:

```graphql
mutation {
  updateOrder(id: "abc123", input: { price: 150.0, tax: 15.0 }) {
    id
    price
    tax
    finalPrice
  }
}
```

**Excluir um Pedido**:

```graphql
mutation {
  deleteOrder(id: "abc123")
}
```
**Consultar Todos os Pedidos**:

```graphql
Copy code
query {
  orders {
    id
    price
    tax
    finalPrice
  }
}
```

## Interagindo com o Serviço gRPC

Para interagir com o serviço gRPC da aplicação CleanArq Challenge, você pode usar a ferramenta evans. 
Siga as instruções abaixo para usar o evans:

Certifique-se de que a aplicação CleanArq Challenge está em execução e o serviço gRPC está disponível em 127.0.0.1:50051.

Execute o evans no terminal:

**Isso iniciará o evans no modo REPL (Read-Eval-Print Loop), permitindo que você interaja com o serviço gRPC da aplicação.**
$ evans -r rpl

  ______
 |  ____|
 | |__    __   __   __ _   _ __    ___
 |  __|   \ \ / /  / _. | | '_ \  / __|
 | |____   \ V /  | (_| | | | | | \__ \
 |______|   \_/    \__,_| |_| |_| |___/

 more expressive universal gRPC client


*Isso selecionará o pacote pb que contém os serviços gRPC definidos.*
127.0.0.1:50051> package pb


*Em seguida, use o comando service para selecionar o serviço gRPC desejado, neste caso, o OrderService:*
pb@127.0.0.1:50051> service OrderService

Isso enviará uma chamada para o método DeleteOrder com o parâmetro id definido como 111111111.
pb.OrderService@127.0.0.1:50051> call DeleteOrder
id (TYPE_STRING) => 111111111
{}
$

Para chamar os demais serviços, basta que no momento do *call* troque de DeleteOrder para UpdateOrder, CreateOrder, GetOrderByID e ListOrder.

## Interagindo com API HTTP/Web

Além das interfaces *GraphQL* e *gRPC*, a aplicação também possui uma interface *HTTP/Web* para interação com os pedidos. Abaixo estão exemplos de como utilizar os endpoints HTTP:

**Criar um Novo Pedido**:


*POST /order*
```json
{
  "price": 100.0,
  "tax": 10.0
}
```

**Listar Todos os Pedidos**:

*GET /orders*

**Buscar Pedido por ID**:

*GET /orderbyid*
```json
{
  "id": "111111111111"
}
```

**Atualizar um Pedido Existente**:

*POST /order/update*
```json
{
  "id": "abc123",
  "price": 150.0,
  "tax": 15.0
}
```

**Excluir um Pedido**:

*DELETE /order/delete*
```json
{
  "id": "111111111111"
}
```