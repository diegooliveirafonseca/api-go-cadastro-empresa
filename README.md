# api-go-cadastro-empresa
API Go - Cadastro de Empresas 🏢🚀 API desenvolvida em Go para cadastro e gerenciamento de empresas, utilizando JWT para autenticação e PostgreSQL como banco de dados.

📌 Funcionalidades

✔️ Cadastro de usuários e autenticação via JWT.

✔️ Cadastro, consulta e exclusão de empresas via CNPJ.

✔️ Proteção de rotas: somente usuários autenticados podem acessar os endpoints.

✔️ Banco de dados PostgreSQL integrado via Docker.

✔️ Deploy facilitado com Docker e Makefile.

🛠️ Tecnologias Utilizadas Go (Golang) Gorilla Mux (roteamento) PostgreSQL (armazenamento de dados) JWT (autenticação) Docker & Docker Compose (gerenciamento de ambiente) Makefile (automatização de comandos) Postman (Para testes de API)

📂 Estrutura do Projeto

api-go-cadastro-empresa/

│── controllers/ # Lógica dos endpoints
│── middleware/ # Middleware de autenticação JWT
│── models/ # Modelos e interação com o banco de dados
│── routes/ # Definição das rotas da API
│ ├── auth_routes.go # Rotas de autenticação
│ ├── user_routes.go # Rotas protegidas de usuários
│ ├── empresa_routes.go # Rotas protegidas de empresas
│── Dockerfile # Configuração do container Docker
│── docker-compose.yml # Orquestração dos serviços
│── main.go # Ponto de entrada da aplicação
│── Makefile

🚀 Como Executar o Projeto

Configuração do Ambiente

1️⃣. Instalar o Docker e Docker Compose

Se ainda não possui o Docker instalado, siga as instruções do site oficial:

Docker -> https://docs.docker.com/engine/install/

2️⃣. Clonar o repositório

git clone https://github.com/diegooliveirafonseca/api-go-cadastro-empresa.git cd api-go-cadastro-empresa

3️⃣ Configurar e Rodar com Docker

make up

Isso irá:

Criar e iniciar o banco de dados PostgreSQL

Construir e rodar a API em um contêiner

4️⃣ Derrubar os Containers

make down

Para reconstruir a API:

make rebuild

5️⃣ Compilar e Rodar Sem Docker (OPCIONAL)

go build -o main .

./main

_______________________________________________________

🔑 Autenticação Toda a API é protegida por JWT, ou seja, o usuário deve estar autenticado para acessar qualquer funcionalidade.

1️⃣ Criar um Usuário

POST /users

📥 Request Body (JSON)

{ "username": "admin", "password": "123456" }

📤 Resposta

{"message":"Usuário criado com sucesso!"}

2️⃣ Fazer Login e Obter Token

POST /login

📥 Request Body (JSON)

{ "username": "admin", "password": "123456" }

📤 Resposta

{ "token": "eyJhbGciOiJIUz..." }

3️⃣ Enviar Token nas Requisições

Para acessar qualquer endpoint, inclua o token no Header da requisição:

Authorization: Bearer eyJhbGciOiJIUz...

📊 Endpoints da API

🔒 Usuários (Requer Autenticação)

Método Endpoint Descrição

GET /user/listarUsers # Lista todos os usuários

📤 Resposta

[ { "id": 1, "username": "admin", "role": "admin" } ]

🔒 Empresas (Requer Autenticação)

Método Endpoint Descrição

GET /empresa/listarEmpresas # Lista todas as empresas

📤 Resposta

[ { "id": 2, "cnpj": "12345678901235", "nome": "Empresa 2" }, { "id": 6, "cnpj": "12345678901237", "nome": "Nova Empresa 2" } ]

GET /empresa # Consulta uma empresa pelo CNPJ

Exemplo: /empresa?cnpj=12345678000199

📤 Resposta

{ "id": 2, "cnpj": "12345678901235", "nome": "Empresa 2" }

POST /empresa/empresas # Cadastra uma nova empresa

📥 Request Body (JSON)

{ "cnpj": "12345678000199", "nome": "Empresa X" }

📤 Resposta

{"message":"Empresa criada com sucesso!"}

DELETE /empresa # Remove uma empresa pelo CNPJ

Exemplo: /empresa?cnpj=12345678000199

📤 Resposta

{"message":"Empresa deletada com sucesso!"}

_____________________________________________________________________________

Testando com Postman

Instalar o Postman -> https://www.postman.com/downloads/

A pasta postman/ possui um arquivo JSON com a configuração dos endpoints.

Abra o Postman

Importe o arquivo postman/Api Go Empresas.postman_collection.json

Configure o token JWT após o login

Teste os endpoints facilmente!

📌 Considerações Finais

Este projeto foi desenvolvido para facilitar o cadastro e consulta de empresas, garantindo segurança e eficiência. Contribuições são bem-vindas! 🤝🚀