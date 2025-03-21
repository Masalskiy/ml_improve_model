# Этап сборки
FROM node:20-alpine as build-stage

WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build

# Этап production
FROM nginx:stable-alpine as production-stage

# Копируем собранное приложение из этапа сборки
COPY --from=build-stage /app/dist /usr/share/nginx/html

# Настраиваем nginx
COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"] 