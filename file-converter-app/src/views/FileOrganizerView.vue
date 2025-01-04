<template>
    <div class="organizer-container">
        <div class="hero-section">
            <h1 class="main-title">Download Archives</h1>
            <p class="subtitle">Access and download your project resources</p>
        </div>

        <div class="downloads-grid">
            <!-- Archive 1 -->
            <div class="download-card">
                <div class="card-content">
                    <div class="icon-wrapper windows">
                        <i class="pi pi-microsoft"></i>
                    </div>
                    <h2>Windows Version</h2>
                    <p>Download the Windows version of the application</p>
                    <button class="download-button" @click="downloadArchive(OS.WINDOWS)"
                        :disabled="downloading.documentation">
                        <i class="fas fa-download"></i>
                        <span>{{ downloading.documentation ? 'Downloading...' : 'Download' }}</span>
                    </button>
                </div>
            </div>

            <!-- Archive 2 -->
            <div class="download-card">
                <div class="card-content">
                    <div class="icon-wrapper apple">
                        <i class="pi pi-apple"></i>
                    </div>
                    <h2>macOS Version</h2>
                    <p>Download the macOS version of the application</p>
                    <button class="download-button" @click="downloadArchive(OS.MAX)" :disabled="downloading.source">
                        <i class="fas fa-download"></i>
                        <span>{{ downloading.source ? 'Downloading...' : 'Download' }}</span>
                    </button>
                </div>
            </div>

            <!-- Archive 3 -->
            <div class="col-span-2 flex justify-center">
                <div class="download-card" style="width: calc(50% - 1rem)">
                    <div class="card-content">
                        <div class="icon-wrapper linux">
                            <font-awesome-icon :icon="['fab', 'linux']" />
                        </div>
                        <h2>Linux Version</h2>
                        <p>Download the Linux version of the application</p>
                        <button class="download-button" @click="downloadArchive(OS.LINUX)"
                            :disabled="downloading.media">
                            <i class="fas fa-linux"></i>
                            <span>{{ downloading.media ? 'Downloading...' : 'Download' }}</span>
                        </button>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { OS } from '@/types/enum/os.enum';
import { FileOrganizerService } from '../services/file-organizer.service';
import { useToast } from 'primevue';

const fileOrganizerService = new FileOrganizerService();

const downloading = ref({
    documentation: false,
    source: false,
    media: false
})

const toast = useToast()

const downloadArchive = async (targetOS: OS) => {
    downloading.value[targetOS] = true

    fileOrganizerService.retrieveOrganizerBinary(targetOS).then(response => {

        const blob = new Blob([response.binary], {
            type: 'application/octet-stream'
        })
        const url = window.URL.createObjectURL(blob)
        const link = document.createElement('a')

        link.href = url
        link.download = response.filename

        document.body.appendChild(link)
        link.click()
        window.URL.revokeObjectURL(url)
        document.body.removeChild(link)

        downloading.value[targetOS] = false
    }).catch(() => {
        toast.add({
            severity: 'error',
            summary: 'Download Failed',
            detail: 'Unable to download the file. Please try again.',
            life: 3000
        })
        downloading.value[targetOS] = false
    })

}
</script>

<style scoped>
.organizer-container {
    min-height: 100vh;
    background: linear-gradient(135deg, #1a1a1a 0%, #2d2d2d 100%);
    padding: 2rem;
}

.hero-section {
    text-align: center;
    padding: 4rem 0;
    margin-bottom: 2rem;
}

.main-title {
    font-size: 3.5rem;
    font-weight: 700;
    color: #ffffff;
    margin-bottom: 1rem;
    text-shadow: 0 2px 4px rgba(0, 0, 0, 0.3);
}

.subtitle {
    font-size: 1.2rem;
    color: #b3b3b3;
    margin-bottom: 2rem;
}

.downloads-grid {
    max-width: 1200px;
    margin: 0 auto;
    display: grid;
    grid-template-columns: repeat(2, 1fr);
    /* Change this line */
    gap: 2rem;
    padding: 1rem;
}

.download-card {
    background: rgba(255, 255, 255, 0.1);
    backdrop-filter: blur(10px);
    border-radius: 16px;
    overflow: hidden;
    transition: transform 0.3s ease, box-shadow 0.3s ease;
    border: 1px solid rgba(255, 255, 255, 0.1);
}

.download-card:hover {
    transform: translateY(-5px);
    box-shadow: 0 8px 25px rgba(0, 0, 0, 0.2);
    background: rgba(255, 255, 255, 0.15);
}

.card-content {
    padding: 2.5rem;
    display: flex;
    flex-direction: column;
    align-items: center;
    text-align: center;
}

.icon-wrapper {
    width: 80px;
    height: 80px;
    border-radius: 50%;
    display: flex;
    align-items: center;
    justify-content: center;
    margin-bottom: 1.5rem;
    transition: transform 0.3s ease;
}

.icon-wrapper.windows {
    background: linear-gradient(135deg, #00A4EF 0%, #0078D4 100%);
    box-shadow: 0 4px 15px rgba(0, 164, 239, 0.3);
}

.icon-wrapper.apple {
    background: linear-gradient(135deg, #A2AAAD 0%, #7D7D7D 100%);
    box-shadow: 0 4px 15px rgba(162, 170, 173, 0.3);
}

.icon-wrapper.linux {
    background: linear-gradient(135deg, #FFA500 0%, #FF6B00 100%);
    box-shadow: 0 4px 15px rgba(255, 165, 0, 0.3);
    font-size: 30px;
}

.download-card:hover .icon-wrapper {
    transform: scale(1.1);
}

.icon-wrapper i {
    font-size: 2rem;
    color: white;
}

.card-content h2 {
    color: #ffffff;
    font-size: 1.8rem;
    margin-bottom: 1rem;
    font-weight: 600;
}

.card-content p {
    color: #b3b3b3;
    margin-bottom: 2rem;
    line-height: 1.6;
}

.file-info {
    display: flex;
    gap: 2rem;
    margin-bottom: 2rem;
    color: #b3b3b3;
}

.info-item {
    display: flex;
    align-items: center;
    gap: 0.5rem;
}

.info-item i {
    color: #42b983;
    font-size: 0.9rem;
}

.download-button {
    background: linear-gradient(135deg, #42b983 0%, #3aa876 100%);
    color: white;
    border: none;
    padding: 0.8rem 1.5rem;
    border-radius: 8px;
    cursor: pointer;
    transition: all 0.3s ease;
    min-width: 150px;
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 0.5rem;
    font-weight: 500;
}

.download-button:hover:not(:disabled) {
    transform: translateY(-2px);
    box-shadow: 0 4px 15px rgba(66, 185, 131, 0.3);
}

.download-button:disabled {
    background: linear-gradient(135deg, #a8d5c2 0%, #9ec7b7 100%);
    cursor: not-allowed;
    transform: none;
    box-shadow: none;
}

.download-button i {
    font-size: 0.9rem;
}

@media (max-width: 768px) {
    .organizer-container {
        padding: 1rem;
    }

    .hero-section {
        padding: 2rem 0;
    }

    .main-title {
        font-size: 2.5rem;
    }

    .downloads-grid {
        grid-template-columns: 1fr;
    }

    .card-content {
        padding: 2rem;
    }

    .file-info {
        gap: 1rem;
    }
}
</style>