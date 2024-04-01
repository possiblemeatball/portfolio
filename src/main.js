import { createApp } from 'vue'
import App from './App.vue'
import BackToTop from 'vue-backtotop';
import './assets/css/style.css'

createApp(App)
    .use(BackToTop)
    .mount('#app')
