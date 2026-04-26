import { createApp } from 'vue'

import { VueQueryPlugin } from '@tanstack/vue-query'

import './style.css'
import App from './index.vue'

const app = createApp(App)
app.use(VueQueryPlugin)

export default app
