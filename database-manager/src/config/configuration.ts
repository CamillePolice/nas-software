export default () => ({
    port: parseInt(process.env.PORT, 10) || 8000,
    jwt: {
        secret: process.env.JWT_SECRET || 'your-secret-key',
        expiresIn: '1h',
    },
    services: {
        fileOrganizer: process.env.FILE_ORGANIZER_URL,
        fileConverter: process.env.FILE_CONVERTER_URL,
        forum: process.env.FORUM_MANAGER_URL,
    },
})
