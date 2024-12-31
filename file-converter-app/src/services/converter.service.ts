import axios from 'axios'

export class ConverterService {
  constructor() {}

  convertFile(inputFile: File, sourceFormat: string, outputFormat: string): Promise<File> {
    const formData = new FormData()
    formData.append('inputFile', inputFile)
    formData.append('sourceFormat', sourceFormat)
    formData.append('outputFormat', outputFormat)

    return axios.post('http://localhost:3000/api/file/converter', formData, {
      headers: {
        'Content-Type': 'multipart/form-data',
      },
    })
  }
}
