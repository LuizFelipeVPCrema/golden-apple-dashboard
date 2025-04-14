FROM node:18-alpine AS build
WORKDIR /app
COPY ./minecraft-dashboard/ ./
RUN npm install && npm run build

FROM nginx:alpine
COPY --from=build /app/dist/minecraft-dashboard /usr/share/nginx/html
COPY ./nginx.conf /etc/nginx/conf.d/default.conf