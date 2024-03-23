import { createApp } from 'vue'
import { router } from './router/index'
import App from './App.vue'
import ElementUI from 'element-plus'
import axios from "axios"
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import 'element-plus/dist/index.css'
import "element-plus/theme-chalk/el-message.css";

axios.defaults.baseURL = "http://192.168.8.101:8080"

const app = createApp(App)

app.provide("axios", axios)

for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
}

app.use(ElementUI)
    .use(router)

app.mount('#app')