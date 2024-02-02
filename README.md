# product-follow-up

API em Golang para CRUD de Produtos, como parte de case técnico.

### Stack

* Postman
* Golang - Echo web Framework
* Testify - Mockery - GORM
* Mysql
* Docker

## Como rodar o projeto

1. Fazer o clone do projeto

```console
git clone https://github.com/uiansol/product-follow-up.git
```

2. Criar o arquivo de variaveis de ambiente .env (pular, ver abaixo)

**Apenas** para facilitar a passagem do projeto eu adicionei o .env já empreenchido no repositório. Não faça isso em casa.

3. Realizar o build e rodar com docker compose

```console
docker compose build
```

```console
docker compose up
```
Dependendo das configurações da instalação do docker, pode ser `docker-compose`, com hífen.

Esperar até aparecer a mensagem do servidor do echo.

Na primeira vez pode demorar um pouco pois o container do mysql faz o setup inicial para criação do banco.

## Como testar o projeto

1. No Postman, importar a coleção do arquivo **product-follow-up.postman_collection.json**

2. Realizar as chamadas CREATE, READ, READ ALL, UPDATE, DELETE

Eu deixei as rotas automatizadas. Quando receber o id do produto no CREATE, ele já é difinido para as outras rotas por uma variável do Postman. Assim, as rotas podem ser testadas mais rapidamente sem precisar ficar alterando o `/:id` nas rotas.

## Destaques do projeto


### Arquiteturais
![hexagonal and clean architectures](https://www.happycoders.eu/wp-content/uploads/2023/01/hexagonal-architecture-vs-clean-architecture-2.v4.png)
* Utilização de **clean architecture** e **driven-domain-design**. Algumas nomeclaturas e caminhos tradicionais foram alterados para deixar o projeto mais parecido com **hexagonal architecture**.
* Boa utilização de **SOLID** como Single-responsibility e Dependency inversion.
* Alguns padrões de projetos utilizados como **Adapter** e **Singleton**. Para não gerar confusão, a pasta adapters no projeto se refere aos adapters da **hexagonal architecture**.
* Padrão **REST** tanto nos métodos das chamadas como nos códigos de retorno.
* Versionamento de API. Para exemplo coloquei a rota de `/ping` como `/v1` e as de produto como `/v2`.
* Testes unitários com **mocks** para isolar a execução.
* Estrutura de pastas nos padrões do ecossistema Golang, como descrito em [Standard Go Project Layout](https://github.com/golang-standards/project-layout).

### Git
![github actions](https://docs.github.com/assets/cb-25535/images/help/actions/overview-actions-simple.png)
* Prática de gitflow com branchs e merges. Pela simplicidade do projeto evitei a criação de uma branch `development`.
* Uso de **semantic commits** para indicar feature, fix, chore...
* Use de **CI** com Github Actions para fazer build e rodar testes automatizados ao fazer pull request.

### Docker
![docker compose golang mysql](https://i.ytimg.com/vi/p0n90IUfjp4/maxresdefault.jpg)
* Uso de variáveis de ambiente com arquivo `.env` carregadas para cada container.
* Container da API com **multi-stage**.
* Uso de **healthcheck** para ordenar inicialização dos containers.

## Próximos passos

1. Uso de um **NoSQL** como **MongoDB**. Para armazenar o histórico de preços dos produtos.
2. Adicionar users e **authentication** com **JWT** para proteger as rotas.
3. Uso de um **message broke** como **RabbitMQ**. Ex: Disparar email quando novo preço cai (no update) muito abaixo do atual.
4. Uso de um **in-memory database** como **Redis**. Para distribuir a carga no banco de dados.
5. Documentar extensamente e distribuir para a comunidade de desenvolvedores Golang. Servir como projeto de referência.

![that's all folks](https://lh3.googleusercontent.com/proxy/r8j-mdXxnWDnde5QbpYmYoPCA3tGuPosruQzxr6wOC4fvOVudULIIKi3qYEZyj0qe_wP5QQhIqub7iFPwROJrqZ00nA4EZfwbA0nLM-F5rc4SfmXN0y-r1QRx3cmFL09pedzLqECyKNBXzJ3NJPT1MPZ6MitYREocgtVyN8YrA)