import * as Joi from 'joi'

export const validationSchema = Joi.object({
    PORT: Joi.number().default(8000),
    JWT_SECRET: Joi.string().required(),
    FILE_ORGANIZER_URL: Joi.string().required(),
    FILE_CONVERTER_URL: Joi.string().required(),
    FORUM_MANAGER_URL: Joi.string().required(),
})
