# Minecraft Dashboard - Gerenciador de Servidores Dockerizados
<p align="center">
  <img src="https://github.com/user-attachments/assets/3998ed3d-75de-47a2-ace9-5386cdb8b5e4" alt="GoldenApple Logo Banner" width="100%">
</p>
## 🧩 Sobre o Projeto
Este projeto é uma plataforma de gerenciamento de servidores Minecraft utilizando containers Docker. A aplicação foi construída com:

- **Angular** no frontend (com suporte a hot reload em dev)
- **Golang** no backend (API REST + WebSocket)
- **PostgreSQL** como banco de dados
- **Nginx** como proxy reverso
- **itzg/minecraft-server** como base para servidores Minecraft

O objetivo é permitir que usuários criem, deletem, iniciem, parem e editem múltiplos servidores de Minecraft com um painel visual e interativo.

## 🐳 Estrutura de Containers
```
📁 minecraft-project
├── backend/                  # Código Go da API
├── frontend/                 # Aplicação Angular
│   ├── Dockerfile            # Build de produção
│   ├── Dockerfile.dev        # Build de desenvolvimento (hot reload)
├── proxy/                   # Configuração do Nginx
│   └── nginx.conf
├── data/                    # Volume do PostgreSQL
├── db/init.sql              # Script inicial do banco
├── .env                     # Variáveis de ambiente
├── docker-compose.yml       # Compose principal
└── docker-compose.override.yml # Configuração local (desenvolvimento)
```

## 🚀 Como subir localmente
### Pré-requisitos
- Docker e Docker Compose instalados

### Passos
```bash
# 1. Clonar o repositório
$ git clone https://github.com/seu-usuario/minecraft-dashboard.git
$ cd minecraft-dashboard

# 2. Configurar variáveis de ambiente
$ cp .env.example .env
# Edite o arquivo .env conforme necessário

# 3. Subir a aplicação
$ docker-compose up -d --build

# 4. Acesse via navegador
http://localhost
```

> O frontend com hot reload estará disponível em http://localhost:4200/ se o `override` estiver ativo para desenvolvimento.

## 📦 Scripts úteis
```bash
# Subir containers
$ docker-compose up -d --build

# Derrubar containers e volumes
$ docker-compose down -v --remove-orphans

# Ver logs de um serviço
$ docker logs -f mc-frontend
```

## 📌 Funcionalidades atuais
- Listar todos os servidores Minecraft
- Criar servidor (em breve)
- Iniciar / Parar servidores
- Deletar servidor
- Visualizar console do servidor (em breve)
- Enviar comandos para o servidor (em breve)
- Modificar arquivos diretamente via painel (em breve)

## 💡 Melhorias futuras
- Integração com painel de configuração do Minecraft (`server.properties`, `ops.json`, etc)
- Upload de mapas customizados
- Permissões de múltiplos usuários (ACL)
- Suporte a plugins/modpacks (Forge/Fabric)
- Deploy em VPS com HTTPS via Let's Encrypt
- Histórico e estatísticas de uso dos servidores

## 🛠️ Tecnologias
- [Angular 19](https://angular.dev/)
- [Golang 1.21+](https://golang.org/)
- [Docker Compose](https://docs.docker.com/compose/)
- [PostgreSQL](https://www.postgresql.org/)
- [Nginx](https://nginx.org/en/)
- [itzg/minecraft-server](https://github.com/itzg/docker-minecraft-server)

## 📄 Licença
Projeto sob licença MIT. Sinta-se livre para usar, contribuir ou abrir issues!

---

Feito com 💻 e ☕ por [Luiz Felipe](https://github.com/LuizFelipeVPCrema)
