FROM node:16.6-alpine AS builder
WORKDIR /app/

COPY package.json yarn.lock ./
RUN yarn

COPY src/ ./src/
RUN yarn build

FROM nginx:1.21.4-alpine
COPY --from=builder /app/dist /usr/share/nginx/html
