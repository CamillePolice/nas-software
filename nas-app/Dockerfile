FROM node:lts-alpine

WORKDIR /app

# Copy package files
COPY package*.json ./

RUN corepack enable && corepack prepare pnpm@latest --activate

# Install dependencies
RUN pnpm install

# Copy source code
COPY . .

# Build the application
CMD [ "pnpm", "dev", "--host"]