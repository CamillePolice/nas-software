<script setup lang="ts">
import { computed } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { useToast } from 'primevue/usetoast';
import { FileType, FileTypeConfig, FileNativType } from '@/enum/file.enum';
import FileDropZone from '@/components/FileDropZone.vue';
import Button from 'primevue/button';
import { ConverterService } from '@/services/converter.service';

const route = useRoute();
const router = useRouter();
const toast = useToast();
const converterService = new ConverterService();

const props = defineProps<{
    source: FileNativType
    dest: FileType
}>();

const fileTypeConfig = Object.values(FileTypeConfig) as FileTypeConfig[];

const pageTitle = computed(() => {
    const destConfig = FileTypeConfig[props.dest];

    if (props.source == FileNativType.IMAGE) {
        return `Image to ${destConfig.name.toUpperCase()}`;
    } else if (props.source == FileNativType.FILE) {
        return `File to ${destConfig.name.toUpperCase()}`;
    }
    return 'File Converter';
});

const acceptedFormats = computed(() => {
    const validConfigs = fileTypeConfig.filter(
        config => config.type === props.source && config.name !== props.dest
    );

    return validConfigs.flatMap(config => config.extension).map(ext => `.${ext}`).join(',');
});

const supportedFormats = computed(() => {
    const validConfigs = fileTypeConfig.filter(
        (config: FileTypeConfig) => config.type == props.source && config.name !== props.dest
    );
    const formats = validConfigs
        .flatMap(config => config.extension)
        .map(ext => ext.toUpperCase())
        .join(', ');

    return `Supports: ${formats}`;
});

const handleConversionStart = (file: File) => {
    console.log("LOG || handleConversionStart || file ->", file)

    converterService.convertFile(file, props.source, props.dest).then(result => {
        console.log("LOG || converterService.convertFile || result ->", result)
        handleConversionComplete(result);
    });
    toast.add({
        severity: 'info',
        summary: 'Converting',
        detail: 'Your file conversion has started...',
        life: 3000
    });
};

const handleConversionComplete = (result: { success: boolean; error?: string }) => {
    if (result.success) {
        toast.add({
            severity: 'success',
            summary: 'Success',
            detail: 'Your file has been converted successfully!',
            life: 300000
        });
    } else {
        toast.add({
            severity: 'error',
            summary: 'Error',
            detail: result.error || 'An error occurred during conversion',
            life: 5000
        });
    }
};

</script>

<template>
    <Toast position="top-right" class="custom-toast" />

    <div class="min-h-screen bg-gray-50">
        <header class="bg-white shadow">
            <div class="max-w-7xl mx-auto px-4 py-6">
                <div class="flex items-center justify-between">
                    <div class="flex items-center space-x-4">
                        <Button icon="pi pi-arrow-left" text @click="router.push('/')" class="!p-2" />
                        <h1 class="text-2xl font-bold text-gray-900">{{ pageTitle }}</h1>
                    </div>
                    <div class="text-sm text-gray-500">
                        {{ supportedFormats }}
                    </div>
                </div>
            </div>
        </header>

        <main class="max-w-4xl mx-auto py-12 px-4">
            <FileDropZone :accepted-formats="acceptedFormats" :conversion-type="props.dest"
                @conversion-start="handleConversionStart" @conversion-complete="handleConversionComplete" />
        </main>
    </div>
</template>