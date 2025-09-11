FROM node:lts-alpine

WORKDIR /app

COPY web/package.json ./

RUN npm install -g pnpm
RUN pnpm install

COPY web/ .

EXPOSE 3000
CMD ["pnpm", "dev"]
