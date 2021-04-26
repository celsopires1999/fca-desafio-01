# Programa Full Cycle de Aceleração: Edição Docker
# Desafio 01

## Descrição
Crie um programa utilizando sua linguagem de programação favorita que faça uma listagem simples do nome de alguns módulos do curso Full Cycle os trazendo de um banco de dados MySQL. Gere a imagem desse container e a publique no DockerHub.

Gere uma imagem do nginx que seja capaz que receber as solicitações http e encaminhá-las para o container.

Crie um repositório no github com todo o fonte do programa e das imagens geradas.

Crie um arquivo README.md especificando quais comandos precisamos executar para que a aplicação funcione recebendo as solicitações na porta 8080 de nosso computador. Lembrando que NÃO utilizaremos Docker-compose nesse desafio.

## Comandos para executar a aplicação a partir do Docker Hub

`docker network create desafio-01`

docker run -d --network=desafio-01 --name=desafio-01-mysql -e MYSQL_ROOT_PASSWORD=mypassword -e MYSQL_DATABASE=desafio-01 -e MYSQL_USER=MainUser -e MYSQL_PASSWORD=MainPassword celsopires/desafio-01-mysql

docker run -d --network=desafio-01 --name=desafio-01-go celsopires/desafio-01-go go run main.go

docker run --network=desafio-01 --name=desafio-01-nginx -p 8080:80 celsopires/desafio-01-nginx`

Digitar no navegador `http://localhost:8080/` para obter a lista de alguns módulos do curso.


## Comandos para rodar a aplicação usando o código fonte:

### 1. NETWORK
1.1. Executar o comando para criar a network de comunicação entre containers:

`docker network create desafio-01`

### 2. MYSQL
2.1. Criar a imagem do MySQL. Ir para pasta mysql:

`docker build -t desafio-01-mysql .`

2.2. Criar o container do MySQL:

`docker run -d --network=desafio-01 --name=desafio-01-mysql -e MYSQL_ROOT_PASSWORD=mypassword -e MYSQL_DATABASE=desafio-01 -e MYSQL_USER=MainUser -e MYSQL_PASSWORD=MainPassword desafio-01-mysql`

### 3. GO
3.1. Criar a imagem do GO. Ir para a pasta go:

`docker build -t desafio-01-go .`

3.2. Criar o container do GO. Para gerar o container, ir para a pasta go, pois está sendo usando pwd para compartilhar volumes:

`docker run -d --network=desafio-01 --name=desafio-01-go -v $(pwd):/go/app desafio-01-go go run main.go`

### 4. NGINX
4.1. Criar a imagem do Nginx. Ir para a pasta nginx:

`docker build -t desafio-01-nginx .`

4.2. Criar o container do NGINX:

`docker run --network=desafio-01 --name=desafio-01-nginx -p 8080:80 desafio-01-nginx`

### 5. Navegador
5.1. Digitar no navegador `http://localhost:8080/` para obter a lista de alguns módulos do curso.
