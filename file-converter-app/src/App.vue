<script setup>
import Header from '@/components/Header.vue'

</script>

<template>
  <div>
    <!-- <Header /> -->
    <main class="max-w-7xl mx-auto py-12 px-4">
      <RouterView />
    </main>
  </div>
</template>

<style>
html {
  scroll-behavior: smooth;
}
</style>
