import { DataSource } from 'typeorm'
import * as dotenv from 'dotenv'

dotenv.config()

export default new DataSource({
    type: 'postgres',
    host: process.env.DATABASE_URL,
    port: parseInt(process.env.DATABASE_PORT, 10),
    username: process.env.DATABASE_USER,
    password: process.env.DATABASE_PASSWORD,
    database: process.env.DATABASE_NAME,
    entities: ['src/**/*.entity{.ts,.js}'],
    migrations: ['src/database/migrations/*{.ts,.js}'],
})
