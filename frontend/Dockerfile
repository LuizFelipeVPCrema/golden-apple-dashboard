# Dockerfile para produção
# FROM node:23-alpine3.21-slim AS build

# WORKDIR /app
# COPY ./minecraft-dashboard/ ./

# RUN npm install
# RUN npm rebuild rollup || true
# RUN npm run build -- --configuration production

# FROM nginx:alpine
# COPY --from=build /app/dist/minecraft-dashboard /usr/share/nginx/html
# COPY ./nginx.conf /etc/nginx/conf.d/default.conf

# Dockerfile para desenvolvimento

FROM node:20-bullseye

WORKDIR /app

# Copia apenas os arquivos necessários para instalar dependências
COPY ./minecraft-dashboard/package.json ./minecraft-dashboard/package-lock.json* ./

RUN npm install --omit=optional && npm install -g @angular/cli@latest && npm rebuild rollup || true

# Copia todo o projeto
COPY ./minecraft-dashboard/ .

EXPOSE 4200

CMD ["ng", "serve", "--host", "0.0.0.0", "--poll=200"]

