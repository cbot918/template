FROM node:14.20.1-bullseye-slim

WORKDIR /app

COPY . .

RUN yarn && cd ui && yarn && yarn run build

EXPOSE 3000

CMD ["node","app.js"]