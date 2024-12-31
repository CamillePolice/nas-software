import './assets/main.css'
import Aura from '@primevue/themes/aura'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import { library } from '@fortawesome/fontawesome-svg-core'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import {
  faPlay,
  faPause,
  faForward,
  faBackward,
  faVolumeUp,
  faVolumeMute,
} from '@fortawesome/free-solid-svg-icons'

import App from './App.vue'
import router from './router'
import PrimeVue from 'primevue/config'
import { ToastService, Toast, Button } from 'primevue'
import 'primeicons/primeicons.css'

const app = createApp(App)

library.add(faPlay, faPause, faForward, faBackward, faVolumeUp, faVolumeMute)

app.component('font-awesome-icon', FontAwesomeIcon)
app.use(PrimeVue)
app.use(ToastService)

app.component('Toast', Toast)
app.component('Button', Button)

app.use(createPinia())
app.use(router)

app.use(PrimeVue, {
  theme: {
    preset: Aura,
  },
})

app.mount('#app')
