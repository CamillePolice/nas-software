export enum FileType {
  PDF = 'pdf',
  WORD = 'word',
  JPG = 'jpg',
  PNG = 'png',
  SVG = 'svg',
  EXCEL = 'excel',
  POWERPOINT = 'powerpoint',
}

export enum FileNativType {
  IMAGE = 1,
  FILE,
  DOCUMENT,
}

export const FileTypeConfig = {
  [FileType.PDF]: {
    name: FileType.PDF,
    type: FileNativType.FILE,
    icon: 'fa-regular fa-file-pdf',
    extension: ['pdf'],
  },
  [FileType.WORD]: {
    name: FileType.WORD,
    type: FileNativType.FILE,
    icon: 'fa-regular fa-file-word',
    extension: ['docx'],
  },
  [FileType.JPG]: {
    name: FileType.JPG,
    type: FileNativType.IMAGE,
    icon: 'fa-regular fa-file-image',
    extension: ['jpg', 'jpeg'],
  },
  [FileType.PNG]: {
    name: FileType.PNG,
    type: FileNativType.IMAGE,
    icon: 'fa-regular fa-file-image',
    extension: ['png'],
  },
  [FileType.SVG]: {
    name: FileType.SVG,
    type: FileNativType.IMAGE,
    icon: 'fa-regular fa-file-image',
    extension: ['svg'],
  },
  [FileType.EXCEL]: {
    name: FileType.EXCEL,
    type: FileNativType.FILE,
    icon: 'fa-regular fa-file-excel',
    extension: ['xlsx', 'csv'],
  },
  [FileType.POWERPOINT]: {
    name: FileType.POWERPOINT,
    type: FileNativType.FILE,
    icon: 'fa-regular fa-file-powerpoint',
    extension: ['pptx'],
  },
} as const

export type FileTypeConfigType = typeof FileTypeConfig
