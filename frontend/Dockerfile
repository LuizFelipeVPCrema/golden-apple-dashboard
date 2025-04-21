FROM node:20-bullseye

WORKDIR /app

# Copiar apenas o package.json e package-lock.json inicialmente
COPY ./minecraft-dashboard/package*.json ./

# Limpar o cache do npm para evitar problemas de dependências
RUN npm cache clean --force

# Instalar as dependências com --legacy-peer-deps para ignorar conflitos de pares
RUN npm install --legacy-peer-deps

# Copiar o resto dos arquivos do projeto
COPY ./minecraft-dashboard/ ./

# Instalar a CLI globalmente com a versão correspondente ao projeto (19.x)
RUN npm install -g @angular/cli@19

EXPOSE 4200

# Usar comando com shell para maior flexibilidade
CMD ["sh", "-c", "npx ng serve --host 0.0.0.0 --poll=200 --disable-host-check"]