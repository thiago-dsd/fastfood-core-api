# FastFood Core API

## 📌 Visão Geral
- Esta é uma API desenvolvida em Golang utilizando o framework Fiber. Ela oferece funcionalidades de CRUD para pedidos (`/order`), um fluxo completo de autenticação (`/user`), e operações CRUD para usuários. A API é projetada para ser fácil de configurar e executar, com suporte a Docker através de um `Makefile`.

### Principais Componentes
- **Autenticação**: Fluxo completo de autenticação, incluindo registro, login, e gerenciamento de tokens.
- **CRUD de Pedidos**: Rotas para criar, ler, atualizar e deletar pedidos (`/order`).
- **CRUD de Usuários**: Operações de criação, leitura, atualização e exclusão de usuários (`/user`).

### 🔧 Estrutura do Projeto

```plaintext
📂 fastfood-core-api
├── 📂 src
│   ├── 📂 auth                # Lógica de autenticação e autorização
│   ├── 📂 common              # Utilitários e helpers comuns
│   ├── 📂 config              # Configurações da aplicação
│   ├── 📂 crypto              # Funções de criptografia
│   ├── 📂 database            # Conexão e operações com o banco de dados
│   ├── 📂 integration         # Integrações com serviços externos
│   ├── 📂 order               # Lógica de negócio e rotas para pedidos
│   ├── 📂 repository          # Camada de acesso a dados (repositórios)
│   ├── 📂 server              # Configuração e inicialização do servidor Fiber
│   ├── 📂 user                # Lógica de negócio e rotas para usuários
├── 🚀 Makefile                # Scripts para execução de containers/projeto
└── README.md
```

## 🚀 Como Rodar o Projeto
### 1️⃣ Configurar as Variáveis de Ambiente
Altere o nome do seu arquivo `.env.example` para `.env.local`.

### 2️⃣ Faça o Setup inicial do Container (Docker)
Certifique-se de instalar a versão 1.23.0 do Golang
```sh
make setup
```

### 3️⃣ Faça o Build da aplicação no Container
```sh
make build
```

### 4️⃣ Execute
```sh
make run
```

