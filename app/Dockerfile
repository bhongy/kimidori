FROM node:16.6-alpine AS builder
WORKDIR /app/

COPY package.json yarn.lock ./
RUN yarn

COPY src/ ./src/
RUN yarn build

FROM node:16.6-alpine
WORKDIR /app/

COPY serve-static.js ./
COPY --from=builder /app/dist/ ./dist/

# CMD ["node", "./serve-static.js"]
