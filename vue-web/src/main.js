import { createApp } from 'vue'
import { router } from './router/index'
import App from './App.vue'
import ElementUI from 'element-plus'
import axios from "axios"

axios.defaults.baseURL = "http://192.168.8.101:8080"

const app = createApp(App)

app.provide("axios", axios)

app.use(ElementUI)
    .use(router)

app.mount('#app')