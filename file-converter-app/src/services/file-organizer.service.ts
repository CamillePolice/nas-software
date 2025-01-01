import axios from 'axios'

interface BinaryResponse {
    binary: ArrayBuffer
    filename: string
}

export class FileOrganizerService {
    endpoint = 'binary/'

    constructor() {}

    retrieveOrganizerBinary(os: string): Promise<BinaryResponse> {
        console.log(`[GET] - /binary/${os}`)

        const formData = new FormData()

        formData.append('os', os)

        console.log('Salut', `${import.meta.env.VITE_ORGANIZER_API}${this.endpoint}${os}`)
        return axios.get(`${import.meta.env.VITE_ORGANIZER_API}${this.endpoint}${os}`, {
            responseType: 'arraybuffer',
            headers: {
                Accept: 'application/octet-stream',
            },
        })
    }
}
