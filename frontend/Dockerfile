FROM node:lts as builder
WORKDIR /app
COPY package*.json ./
RUN npm ci
COPY . .
RUN npm run webpack

FROM nginx:1.23
COPY nginx.conf /etc/nginx/nginx.conf
COPY --from=builder /app/dist/ /usr/share/nginx/html
