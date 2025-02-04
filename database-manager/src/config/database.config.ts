import { TypeOrmModuleOptions } from '@nestjs/typeorm'
import * as dotenv from 'dotenv'

dotenv.config()

export const postgresConfig: TypeOrmModuleOptions = {
    type: 'postgres',
    host: process.env.DATABASE_URL,
    port: parseInt(process.env.DATABASE_PORT, 10),
    username: process.env.DATABASE_USER,
    password: process.env.DATABASE_PASSWORD,
    database: process.env.DATABASE_NAME,
    schema: 'nas',
    entities: ['src/database/entities/**/*.entity{.ts,.js}'],
    migrations: ['src/database/migrations/*{.ts,.js}'],
    synchronize: process.env.NODE_ENV !== 'production',
    logging: process.env.NODE_ENV !== 'production',
}
