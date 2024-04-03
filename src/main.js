import { createApp } from 'vue'
import App from './App.vue'
import BackToTop from 'vue-backtotop';
import './styles.css'
import { router } from './router'

createApp(App)
    .use(router)
    .use(BackToTop)
    .mount('#app')
