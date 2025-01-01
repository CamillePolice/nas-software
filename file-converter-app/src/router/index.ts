import { createRouter, createWebHistory, type RouteLocationNormalized } from 'vue-router'
import FileConverterView from '../views/FileConverterView.vue'
import HomeView from '../views/HomeView.vue'
import FileOrganizerView from '@/views/FileOrganizerView.vue'
import type { FileNativType, FileType } from '@/enum/file.enum'

const routes = [
  {
    path: '/',
    name: 'home',
    component: HomeView,
  },
  {
    path: '/converter',
    name: 'converter',
    component: FileConverterView,
    children: [
      {
        path: '',
        name: 'converter-default',
        component: () => import('@/components/FileConverter.vue'),
        props: (route: RouteLocationNormalized) => ({
          source: Number(route.query.source) as FileNativType,
          dest: route.query.dest as FileType,
        }),
      },
    ],
  },
  {
    path: '/organizer',
    name: 'organizer',
    component: FileOrganizerView,
  },
]

export const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
