FROM node:lts-alpine

LABEL author="Rahul Bharati"
LABEL description="Creavio app frontend container"

WORKDIR /app

RUN corepack enable && corepack prepare pnpm@latest --activate

COPY package.json ./
COPY pnpm-lock.yaml ./

RUN pnpm install

COPY . .

EXPOSE 3000

CMD ["pnpm", "run", "dev"]