<template>
    <div class="home">
        <div class="hero-section">
            <h1 class="main-title">File Sandbox</h1>
            <p class="subtitle">Welcome to File Management System</p>
        </div>

        <div class="cards-container">
            <!-- File Converter Card -->
            <div class="feature-card" @click="navigateToConverter">
                <div class="card-content">
                    <div class="icon-wrapper">
                        <i class="fas fa-exchange-alt"></i>
                    </div>
                    <h2>File Converter</h2>
                    <p>Convert your files between different formats easily and quickly</p>
                    <div class="card-footer">
                        <span class="explore-link">
                            Explore <i class="fas fa-arrow-right"></i>
                        </span>
                    </div>
                </div>
            </div>

            <!-- File Organizer Card -->
            <div class="feature-card" @click="navigateToOrganizer">
                <div class="card-content">
                    <div class="icon-wrapper">
                        <i class="fas fa-folder-open"></i>
                    </div>
                    <h2>File Organizer</h2>
                    <p>Organize and manage your files efficiently</p>
                    <div class="card-footer">
                        <span class="explore-link">
                            Explore <i class="fas fa-arrow-right"></i>
                        </span>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import { useRouter } from 'vue-router'

const router = useRouter()

const navigateToConverter = () => {
    router.push('/converter')
}

const navigateToOrganizer = () => {
    router.push('/organizer')
}
</script>

<style scoped>
.home {
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

.cards-container {
    max-width: 1200px;
    margin: 0 auto;
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(350px, 1fr));
    gap: 2rem;
    padding: 1rem;
}

.feature-card {
    background: rgba(255, 255, 255, 0.1);
    backdrop-filter: blur(10px);
    border-radius: 16px;
    overflow: hidden;
    transition: transform 0.3s ease, box-shadow 0.3s ease;
    cursor: pointer;
    border: 1px solid rgba(255, 255, 255, 0.1);
}

.feature-card:hover {
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
    background: linear-gradient(135deg, #42b983 0%, #3aa876 100%);
    border-radius: 50%;
    display: flex;
    align-items: center;
    justify-content: center;
    margin-bottom: 1.5rem;
    box-shadow: 0 4px 15px rgba(66, 185, 131, 0.3);
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

.card-footer {
    margin-top: auto;
}

.explore-link {
    color: #42b983;
    font-weight: 500;
    display: flex;
    align-items: center;
    gap: 0.5rem;
    transition: color 0.3s ease;
}

.explore-link i {
    font-size: 0.9rem;
    transition: transform 0.3s ease;
}

.feature-card:hover .explore-link {
    color: #4cd79b;
}

.feature-card:hover .explore-link i {
    transform: translateX(5px);
}

@media (max-width: 768px) {
    .home {
        padding: 1rem;
    }

    .hero-section {
        padding: 2rem 0;
    }

    .main-title {
        font-size: 2.5rem;
    }

    .cards-container {
        grid-template-columns: 1fr;
    }

    .card-content {
        padding: 2rem;
    }
}
</style>