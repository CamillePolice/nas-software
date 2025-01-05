FROM node:lts

WORKDIR /app

# Copy package files
COPY package*.json ./

RUN corepack enable && corepack prepare pnpm@latest --activate

# Install dependencies including development ones
RUN pnpm install

# Copy source code
COPY . .

# Expose port
EXPOSE 4000

# Start in development mode with hot-reload
CMD ["npm", "run", "start:dev"]
