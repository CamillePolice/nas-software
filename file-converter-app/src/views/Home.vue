<script setup lang="ts">
import { useRouter } from 'vue-router';
import { FileType, FileTypeConfig } from '@/enum/file.enum';
import { FileNativType } from '../enum/file.enum';

const router = useRouter();

const navigateToConverter = (source: FileNativType, dest: FileType) => {
    router.push({
        name: 'converter',
        query: { source, dest }
    });
};

const getConversionDescription = (fileTypeConfig: FileType, nativType: FileNativType) => {
    // Get the config for the specified type
    const targetConfig = FileTypeConfig[fileTypeConfig];
    const filesTypeConfig = Object.values(FileTypeConfig) as FileTypeConfig[];

    // Get all configurations for the native type
    return filesTypeConfig
        .filter((config) => config.type === nativType)
        // Exclude the target format's extensions
        .filter((config) => config.name !== targetConfig.name)
        // Get all extensions
        .flatMap((config) => config.extension)
        .toString()
        .replace(/,/g, ', ');
};
</script>

<template>
    <div class="min-h-screen bg-gray-50">
        <header class="bg-gradient-to-b from-gray-100 to-white border-b">
            <div class="max-w-7xl mx-auto py-4 px-6">
                <h1 class="text-2xl font-semibold text-gray-900">File Converter</h1>
            </div>
        </header>

        <main class="max-w-7xl mx-auto py-8 px-6">
            <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
                <div class="bg-white rounded-lg shadow-sm hover:shadow-md transition-all duration-200 cursor-pointer p-6"
                    @click="navigateToConverter(FileNativType.FILE, FileType.WORD)">
                    <div class="w-8 h-8 bg-blue-100 rounded mb-4"></div>
                    <h3 class="text-lg font-medium mb-2">PDF to Word</h3>
                    <p class="text-gray-500 mb-4">Convert PDF files to editable Word documents</p>
                    <p class="text-sm text-gray-400">.pdf to .docx</p>
                </div>

                <div class="bg-white rounded-lg shadow hover:shadow-lg transition-shadow cursor-pointer"
                    @click="navigateToConverter(FileNativType.FILE, FileType.PDF)">
                    <div class="p-6">
                        <div class="w-12 h-12 bg-red-100 rounded-lg flex items-center justify-center mb-4">
                            <i class="pi pi-file text-red-600 text-2xl"></i>
                        </div>
                        <h3 class="text-lg font-semibold mb-2">File to PDF</h3>
                        <p class="text-gray-600">Convert a file to PDF format</p>
                    </div>
                    <div class="px-6 py-4 bg-gray-50 rounded-b-lg">
                        <div class="flex items-center text-sm text-gray-500">
                            <i class="pi pi-arrow-right mr-2"></i>
                            <span>{{ getConversionDescription(FileType.PDF, FileNativType.FILE) }}</span>
                        </div>
                    </div>
                </div>

                <div class="bg-white rounded-lg shadow hover:shadow-lg transition-shadow cursor-pointer"
                    @click="navigateToConverter(FileNativType.IMAGE, FileType.PNG)">
                    <div class="p-6">
                        <div class="w-12 h-12 bg-green-100 rounded-lg flex items-center justify-center mb-4">
                            <i class="pi pi-image text-green-600 text-2xl"></i>
                        </div>
                        <h3 class="text-lg font-semibold mb-2">Image to PNG </h3>
                        <p class="text-gray-600">Convert images to PNG</p>
                    </div>
                    <div class="px-6 py-4 bg-gray-50 rounded-b-lg">
                        <div class="flex items-center text-sm text-gray-500">
                            <i class="pi pi-arrow-right mr-2"></i>
                            <span>{{ getConversionDescription(FileType.PNG, FileNativType.IMAGE) }}</span>
                        </div>
                    </div>
                </div>

                <div class="bg-white rounded-lg shadow hover:shadow-lg transition-shadow cursor-pointer"
                    @click="navigateToConverter(FileNativType.IMAGE, FileType.SVG)">
                    <div class="p-6">
                        <div class="w-12 h-12 bg-green-100 rounded-lg flex items-center justify-center mb-4">
                            <i class="pi pi-image text-green-600 text-2xl"></i>
                        </div>
                        <h3 class="text-lg font-semibold mb-2">Image to SVG</h3>
                        <p class="text-gray-600">Convert images to SVG</p>
                    </div>
                    <div class="px-6 py-4 bg-gray-50 rounded-b-lg">
                        <div class="flex items-center text-sm text-gray-500">
                            <i class="pi pi-arrow-right mr-2"></i>
                            <span>{{ getConversionDescription(FileType.SVG, FileNativType.IMAGE) }}</span>
                        </div>
                    </div>
                </div>

                <div class="bg-white rounded-lg shadow hover:shadow-lg transition-shadow cursor-pointer"
                    @click="navigateToConverter(FileNativType.IMAGE, FileType.JPG)">
                    <div class="p-6">
                        <div class="w-12 h-12 bg-green-100 rounded-lg flex items-center justify-center mb-4">
                            <i class="pi pi-image text-green-600 text-2xl"></i>
                        </div>
                        <h3 class="text-lg font-semibold mb-2">Image to JPG</h3>
                        <p class="text-gray-600">Convert images to JPG</p>
                    </div>
                    <div class="px-6 py-4 bg-gray-50 rounded-b-lg">
                        <div class="flex items-center text-sm text-gray-500">
                            <i class="pi pi-arrow-right mr-2"></i>
                            <span>{{ getConversionDescription(FileType.JPG, FileNativType.IMAGE) }}</span>
                        </div>
                    </div>
                </div>
            </div>
        </main>
    </div>
</template>