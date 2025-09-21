FROM node:lts-alpine

LABEL author="Rahul Bharati"
LABEL description="Creavio app frontend container"

WORKDIR /app

RUN npm install -g pnpm

COPY package.json pnpm-lock.yaml ./

RUN pnpm install --frozen-lockfile

COPY . .

EXPOSE 3000

CMD ["pnpm", "run", "dev", "--host"]