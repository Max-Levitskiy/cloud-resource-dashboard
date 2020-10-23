
# Stage 1
FROM node:10-alpine as builder
RUN mkdir -p /app
WORKDIR /app
COPY package.json /app
COPY package-lock.json /app
RUN npm ci
COPY . /app
RUN npm run build-prod

# Stage 2
FROM nginx:1.17.1-alpine
COPY --from=builder /app/dist/dashboard-web /usr/share/nginx/html
COPY  ./nginx/nginx.conf /etc/nginx/conf.d/default.conf

