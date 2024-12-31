<script setup lang="ts">
import { ref } from 'vue';
import { useToast } from 'primevue/usetoast';
import Button from 'primevue/button';
import Toast from 'primevue/toast';
import { FileType } from '@/enum/file.enum';

interface Props {
    acceptedFormats: string;
    conversionType: FileType;
}

interface EmitEvents {
    (event: 'conversion-start', file: File): void;
    (event: 'conversion-complete', result: { success: boolean; error?: string }): void;
}

const props = defineProps<Props>();
const emit = defineEmits<EmitEvents>();

const toast = useToast();
const file = ref<File | null>(null);
const isDragging = ref(false);
const fileInput = ref<HTMLInputElement | null>(null);

console.log("Props", props);

const isValidFileType = (file: File): boolean => {
    const fileExtension = '.' + file.name.split('.').pop()?.toLowerCase();

    return props.acceptedFormats.includes(fileExtension);
};

const handleDragOver = (e: DragEvent) => {
    e.preventDefault();
    isDragging.value = true;
};

const handleDragLeave = (e: DragEvent) => {
    e.preventDefault();
    isDragging.value = false;
};

const handleDrop = async (event: DragEvent) => {
    event.preventDefault();
    const droppedFile = event.dataTransfer?.files[0];

    if (!droppedFile) return;

    isDragging.value = false;

    if (!isValidFileType(droppedFile)) {
        toast.add({
            severity: 'error',
            summary: 'Invalid File Type',
            detail: `Please upload a valid file type: ${props.acceptedFormats}`,
            life: 5000,  // Increased life time
            closable: true,
            styleClass: 'custom-toast-message'
        });
        return;
    }

    file.value = droppedFile;
    toast.add({
        severity: 'success',
        summary: 'File Added',
        detail: `${droppedFile.name} has been added`,
        life: 3000,
        closable: true,
        styleClass: 'custom-toast-message'
    });
};

const removeFile = () => {
    file.value = null;
    if (fileInput.value) {
        fileInput.value.value = '';
    }
};

const handleFileSelect = async (event: Event) => {
    const target = event.target as HTMLInputElement;
    const selectedFile = target.files?.[0];

    if (!selectedFile) return;

    if (!isValidFileType(selectedFile)) {
        toast.add({
            severity: 'error',
            summary: 'Invalid File Type',
            detail: `Please upload a valid file type: ${props.acceptedFormats}`,
            life: 5000,
            closable: true,
            styleClass: 'custom-toast-message'
        });
        return;
    }

    file.value = selectedFile;
    toast.add({
        severity: 'success',
        summary: 'File Added',
        detail: `${selectedFile.name} has been added`,
        life: 3000,
        closable: true,
        styleClass: 'custom-toast-message'
    });
};

const emitConversionStart = () => {
    emit('conversion-start', file.value!);
};

const formatFileSize = (bytes: number): string => {
    if (bytes === 0) return '0 Bytes';

    const k = 1024;
    const sizes = ['Bytes', 'KB', 'MB', 'GB'];
    const i = Math.floor(Math.log(bytes) / Math.log(k));

    return `${parseFloat((bytes / Math.pow(k, i)).toFixed(2))} ${sizes[i]}`;
};
</script>

<template>
    <div class="w-full max-w-2xl mx-auto p-4">
        <div class="flex flex-col space-y-4">
            <div :class="[
                'border-2 border-dashed rounded-lg p-8',
                'transition-colors duration-200 ease-in-out',
                'flex flex-col items-center justify-center',
                'min-h-48 space-y-4',
                isDragging ? 'border-primary bg-primary-50' : 'border-gray-300',
                file ? 'bg-gray-50' : 'bg-white'
            ]" @dragover="handleDragOver" @dragleave="handleDragLeave" @drop="handleDrop">
                <i class="pi pi-upload text-4xl" :class="isDragging ? 'text-primary' : 'text-gray-400'" />

                <template v-if="!file">
                    <input type="file" ref="fileInput" class="hidden" :accept="acceptedFormats"
                        @change="handleFileSelect" />
                    <p class="text-center text-gray-600 cursor-pointer" @click="fileInput?.click()">
                        Drag and drop your file here
                        <br>
                        <span class="text-sm text-gray-400">or click to browse</span>
                    </p>
                    <p class="text-sm text-gray-400">
                        Accepted formats: {{ acceptedFormats }}
                    </p>
                </template>

                <template v-else>
                    <div class="w-full">
                        <div class="flex items-center justify-between p-4 bg-white rounded-lg shadow-md">
                            <div class="flex items-center space-x-3 truncate">
                                <i class="pi pi-file text-xl text-gray-500" />
                                <div class="flex-1 min-w-0">
                                    <p class="text-sm font-medium text-gray-900 truncate">
                                        {{ file.name }}
                                    </p>
                                    <p class="text-sm text-gray-500">
                                        {{ formatFileSize(file.size) }}
                                    </p>
                                </div>
                            </div>
                            <div class="flex items-center space-x-3">
                                <Button icon="pi pi-times" severity="danger" text @click="removeFile"
                                    aria-label="Remove file" class="p-2 hover:bg-red-50 transition-colors duration-200"
                                    :class="{ 'p-button-outlined': true }" />
                                <Button label="Start Conversion" icon="pi pi-cog" @click="emitConversionStart" class="px-4 py-2 bg-primary-500 hover:bg-primary-600 text-white rounded-lg
                                    transition-colors duration-200 flex items-center gap-2
                                    shadow-sm hover:shadow-md" :class="{ 'p-button-raised': true }" />
                            </div>
                        </div>
                    </div>
                </template>
            </div>
        </div>

        <div class="toast-container">
            <Toast position="top-right" class="custom-toast" />
        </div>
    </div>
</template>

<style scoped>
/* Add these custom styles to enhance PrimeVue buttons */
:deep(.p-button-raised) {
    @apply shadow-sm hover:shadow-md transform hover:-translate-y-px transition-all duration-200;
}

:deep(.p-button.p-button-danger) {
    @apply hover:bg-red-50;
}

/* Optional: Add a subtle animation for the conversion button */
:deep(.p-button-raised:active) {
    @apply transform translate-y-px;
}
</style>

<style scoped>
.toast-container {
    position: fixed;
    top: 20px;
    right: 20px;
    z-index: 1000;
}

:deep(.custom-toast) {

    /* Custom styling for the toast */
    .p-toast-message {
        background: white;
        border-radius: 8px;
        box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -1px rgba(0, 0, 0, 0.06);
    }

    .p-toast-message-success {
        background: #f0fdf4;
        border-left: 6px solid #22c55e;
    }

    .p-toast-message-error {
        background: #fef2f2;
        border-left: 6px solid #ef4444;
    }

    .p-toast-message-info {
        background: #f0f9ff;
        border-left: 6px solid #3b82f6;
    }

    .p-toast-icon-close {
        color: #6b7280;
    }
}
</style>