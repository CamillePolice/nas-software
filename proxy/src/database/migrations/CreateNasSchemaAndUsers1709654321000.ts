// src/database/migrations/1709654321000-CreateNasSchemaAndUsers.ts
import { MigrationInterface, QueryRunner, Table } from 'typeorm'

export class CreateNasSchemaAndUsers1709654321000 implements MigrationInterface {
    public async up(queryRunner: QueryRunner): Promise<void> {
        // Create the schema
        await queryRunner.createSchema('nas', true)

        // Create the users table in the nas schema
        await queryRunner.createTable(
            new Table({
                schema: 'nas',
                name: 'users',
                columns: [
                    {
                        name: 'id',
                        type: 'uuid',
                        isPrimary: true,
                        generationStrategy: 'uuid',
                        default: 'uuid_generate_v4()',
                    },
                    {
                        name: 'email',
                        type: 'varchar',
                        isUnique: true,
                    },
                    {
                        name: 'username',
                        type: 'varchar',
                        isUnique: true,
                    },
                    {
                        name: 'password',
                        type: 'varchar',
                    },
                    {
                        name: 'isActive',
                        type: 'boolean',
                        default: true,
                    },
                    {
                        name: 'createdAt',
                        type: 'timestamp',
                        default: 'CURRENT_TIMESTAMP',
                    },
                    {
                        name: 'updatedAt',
                        type: 'timestamp',
                        default: 'CURRENT_TIMESTAMP',
                        onUpdate: 'CURRENT_TIMESTAMP',
                    },
                ],
            }),
            true,
        )
    }

    public async down(queryRunner: QueryRunner): Promise<void> {
        // Drop the table first
        await queryRunner.dropTable('nas.users')
        // Then drop the schema
        await queryRunner.dropSchema('nas', true, true)
    }
}
