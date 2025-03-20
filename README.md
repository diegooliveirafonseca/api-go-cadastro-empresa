# api-go-cadastro-empresa

API Go - Cadastro de Empresas 🏢🚀
API desenvolvida em Go para cadastro e gerenciamento de empresas, utilizando JWT para autenticação e PostgreSQL como banco de dados.

📌 Funcionalidades

✔️ Cadastro de usuários e autenticação via JWT.

✔️ Cadastro, consulta e exclusão de empresas via CNPJ.

✔️ Proteção de rotas: somente usuários autenticados podem acessar os endpoints.

✔️ Banco de dados PostgreSQL integrado via Docker.

✔️ Deploy facilitado com Docker e Makefile.

🛠️ Tecnologias Utilizadas
Go (Golang)
Gorilla Mux (roteamento)
PostgreSQL (armazenamento de dados)
JWT (autenticação)
Docker & Docker Compose (gerenciamento de ambiente)
Makefile (automatização de comandos)

📂 Estrutura do Projeto

api-go-cadastro-empresa/

│── controllers/      # Lógica dos endpoints  
│── middleware/       # Middleware de autenticação JWT  
│── models/           # Modelos e interação com o banco de dados  
│── routes/           # Definição das rotas da API  
│   ├── auth_routes.go      # Rotas de autenticação  
│   ├── user_routes.go      # Rotas protegidas de usuários  
│   ├── empresa_routes.go   # Rotas protegidas de empresas  
│── Dockerfile        # Configuração do container Docker  
│── docker-compose.yml # Orquestração dos serviços  
│── main.go           # Ponto de entrada da aplicação  
│── Makefile    

🚀 Como Executar o Projeto
1️⃣ Configurar e Rodar com Docker
make up
2️⃣ Derrubar os Containers
make down
3️⃣ Compilar e Rodar Sem Docker
go build -o main .
./main

_________________________________________________________________________________

🔑 Autenticação
Toda a API é protegida por JWT, ou seja, o usuário deve estar autenticado para acessar qualquer funcionalidade.

1️⃣ Criar um Usuário

POST /users

📥 Request Body (JSON)

{
  "username": "admin",
  "password": "123456"
}

📤 Resposta

{"message":"Usuário criado com sucesso!"}

2️⃣ Fazer Login e Obter Token

POST /login

📥 Request Body (JSON)

{
  "username": "admin",
  "password": "123456"
}

📤 Resposta

{
  "token": "eyJhbGciOiJIUz..."
}

3️⃣ Enviar Token nas Requisições
Para acessar qualquer endpoint, inclua o token no Header da requisição:

Authorization: Bearer eyJhbGciOiJIUz...


_____________________________________________________________________________

📊 Endpoints da API

🔒 Usuários (Requer Autenticação)

Método	Endpoint	Descrição
GET	/user/listarUsers	     # Lista todos os usuários

🔒 Empresas (Requer Autenticação)

Método	Endpoint	Descrição
GET	/empresa/listarEmpresas	 # Lista todas as empresas
GET	/empresa	             # Consulta uma empresa pelo CNPJ
POST	/empresa/empresas	 # Cadastra uma nova empresa
DELETE	/empresa	         # Remove uma empresa pelo CNPJ

📌 Considerações Finais

Este projeto foi desenvolvido para facilitar o cadastro e consulta de empresas, garantindo segurança e eficiência.
Contribuições são bem-vindas! 🤝🚀
