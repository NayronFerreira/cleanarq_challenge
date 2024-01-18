## CleanArq Challenge

Esta é uma aplicação de exemplo seguindo os princípios da Arquitetura Limpa (*Clean Architecture*) em Go. A aplicação consiste em um sistema de gerenciamento de pedidos (orders).

## Estrutura do Projeto
A estrutura do projeto é organizada da seguinte forma:

- **main.go**: Este arquivo é o ponto de entrada da aplicação, onde a inicialização de diversos componentes ocorre, como a configuração do servidor *HTTP*, *gRPC*, *GraphQL*, conexão com o banco de dados e injeção de dependências usando *Google Wire*.

- **database/**: Este diretório contém o código relacionado ao acesso ao banco de dados. Inclui a definição de interfaces e implementações de repositórios para interação com o banco de dados *MySQL*.

- **infra/** : Este diretório contém o código relacionado à infraestrutura da aplicação, como os manipuladores HTTP e gRPC. Também inclui a configuração do servidor *web*, *gRPC* server e *GraphQL* server.

- **internal/**: Este diretório contém o código interno da aplicação, como os casos de uso. Os casos de uso são responsáveis por implementar a lógica de negócios da aplicação e são independentes da infraestrutura e do framework.

- **pkg/**: Este diretório contém pacotes reutilizáveis em toda a aplicação, como eventos, que podem ser usados em diferentes partes da aplicação.

# Executando a Aplicação

Para executar a aplicação, você precisará ter o Docker e o Docker Compose instalados em sua máquina. Se você ainda não os tem instalados, você pode baixá-los a partir dos seguintes links:

- Docker: [https://docs.docker.com/get-docker/](https://docs.docker.com/get-docker/)
- Docker Compose: [https://docs.docker.com/compose/install/](https://docs.docker.com/compose/install/)

Depois de instalar o Docker e o Docker Compose, siga os passos abaixo para executar a aplicação:

1. Abra um terminal na raiz do projeto, onde o arquivo *docker-compose.yml* está localizado.

2. Execute o seguinte comando para construir e iniciar os serviços definidos no arquivo `docker-compose.yml`:

```bash
docker-compose up
```

- A aplicação agora deve estar acessível nos endereços:
- [http://localhost:8181] - PlayGround GraphQL.
- [http://localhost:8080] - Web API.
- *evans repl -p 50051* - gRPC API.

Para parar a aplicação, pressione *Ctrl+C* no terminal. 

Se você quiser remover os containers, networks e volumes definidos no arquivo *docker-compose.yml*, execute o seguinte comando:

```bash
docker-compose down
```

## Casos de Uso

A aplicação possui os seguintes casos de uso:

**Criar Pedido (CreateOrderUseCase)**: Este caso de uso é responsável por criar um novo pedido no sistema. Recebe as informações do pedido, como preço e taxa, valida e salva no banco de dados. Também dispara um evento de "Pedido Criado" para notificar outros sistemas sobre a criação do pedido.

**Listar Pedidos (ListOrderUseCase)**: Este caso de uso é responsável por listar todos os pedidos no sistema. Recupera os pedidos do banco de dados e retorna uma lista de pedidos. Também dispara um evento com todo o paylod retornado.

**Buscar Pedido por ID (GetOrderByIDUseCase)**: Este caso de uso é responsável por buscar um pedido específico pelo seu ID. Retorna o pedido correspondente, se existir. Também dispara um evento com todo o paylod retornado.

**Atualizar Pedido (UpdateOrderUseCase)**: Este caso de uso é responsável por atualizar as informações de um pedido existente. Recebe as novas informações do pedido, valida e atualiza no banco de dados. Também dispara um evento de "Pedido Atualizado" para notificar outros sistemas sobre a atualização do pedido.

**Excluir Pedido (DeleteOrderUseCase)**: Este caso de uso é responsável por excluir um pedido existente do sistema. Recebe o ID do pedido a ser excluído, remove do banco de dados e dispara um evento de "Pedido Excluído" para notificar outros sistemas sobre a exclusão do pedido.

## Exemplos de Uso

A aplicação fornece uma API GraphQL para interação com os pedidos. Abaixo estão alguns exemplos de consultas e mutações que podem ser realizadas:
O GraphQL Playground pode ser acessada no endereço [http://localhost:8181/].

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

## Interagindo com o serviço gRPC usando Evans

Para interagir com o serviço gRPC da aplicação usando o Evans, siga as instruções abaixo:

O docker-compose ja sobe um container com a imagem do Evans.

Para realizar as chamadas usando o evans, basta rodar no terminal os comands abaixo a partir de qualquer diretorio:

```bash
docker-compose exec evans sh

evans repl -p 50051

package pb

service OrderService
```
Isso selecionará o pacote pb que contém os serviços gRPC definidos e o serviço *OrderService*.

*Chamando métodos do serviço OrderService*

Para chamar métodos do serviço OrderService, como *DeleteOrder, UpdateOrder, CreateOrder, GetOrderByID e ListOrder*, siga o padrão abaixo, substituindo <método> pelo nome do método desejado e fornecendo os parâmetros necessários:

```bash
call <método>
```
Por exemplo, para enviar uma chamada para o método *DeleteOrder* com o parâmetro id definido como 111111111:

```bash
call DeleteOrder
id (TYPE_STRING) => 111111111
```

Isso enviará uma chamada para o método DeleteOrder com o parâmetro id definido como 111111111.

Para chamar outros métodos, basta substituir *DeleteOrder por UpdateOrder, CreateOrder, GetOrderByID ou ListOrder* no comando call.

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
