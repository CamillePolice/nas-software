export default () => ({
    port: parseInt(process.env.PORT, 10) || 8000,
    jwt: {
        secret: process.env.JWT_SECRET || 'your-secret-key',
        expiresIn: '1h',
    },
    services: {
        fileOrganizer: process.env.FILE_ORGANIZER_URL || 'http://localhost:3000',
        fileConverter: process.env.FILE_CONVERTER_URL || 'http://localhost:4000',
        forum: process.env.FORUM_MANAGER_URL || 'http://localhost:5000',
    },
})
