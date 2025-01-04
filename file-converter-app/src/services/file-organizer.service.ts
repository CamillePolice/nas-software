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

        return axios
            .get(`${import.meta.env.VITE_ORGANIZER_API}${this.endpoint}${os}`)
            .then((response) => response.data)
    }
}
