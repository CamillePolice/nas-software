import { createRouter, createWebHistory, type RouteLocationNormalized } from 'vue-router'
import Home from '@/views/Home.vue'
import FileConverter from '@/components/FileConverter.vue'
import type { FileNativType, FileType } from '@/enum/file.enum'

const routes = [
  {
    path: '/',
    name: 'home',
    component: Home,
  },
  {
    path: '/converter', // Removed :type parameter
    name: 'converter',
    component: FileConverter,
    props: (route: RouteLocationNormalized) => ({
      source: Number(route.query.source) as FileNativType, // Convert string to number
      dest: route.query.dest as FileType,
    }),
  },
]

export const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
