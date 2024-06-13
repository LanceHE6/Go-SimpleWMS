import { createApp } from 'vue'
import { router } from './router/index'
import App from './App.vue'
import ElementUI from 'element-plus'
import axios from "axios"
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import 'element-plus/dist/index.css'
import "element-plus/theme-chalk/el-message.css";
import ECharts from 'vue-echarts'
import 'echarts'

axios.defaults.baseURL = localStorage.getItem("url") || "http://47.236.80.244:6007";

const app = createApp(App)

app.provide("axios", axios)

for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
}
app.component('v-chart', ECharts)

app.use(ElementUI)
    .use(router)

app.mount('#app')

// router.beforeEach((to, from, next) => {
//     console.log("from:", from.name)
//     console.log("to:", to.name)
//     //如果需要鉴权
//     if(to.meta.requiresAuth){
//         if(to.meta.beforeName !== from.name){
//             ElMessage.error("操作失败，请先登录！")
//             next('/')
//         }
//         else{
//             next()
//         }
//     }
//     else{
//         next()
//     }
// });