import { createApp } from 'vue'
import { router } from './router/index'
import App from './App.vue'
import ElementUI from 'element-plus'
import axios from "axios"
import config from '../config.json';
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import 'element-plus/dist/index.css'
import "element-plus/theme-chalk/el-message.css";
import * as path from "path";
import * as fs from "fs";

// 定义默认配置
const defaultConfig = {
    host: "47.236.80.244",
    port: "6007",
    path: "api"
};

// 定义配置文件的路径
const configPath = path.resolve(__dirname, '../config.json');

// 检查文件是否存在
if (!fs.existsSync(configPath)) {
    // 如果文件不存在，写入默认配置
    fs.writeFileSync(configPath, JSON.stringify(defaultConfig, null, 2));
}
axios.defaults.baseURL = "http://" + config.host + ":" + config.port + "/" + config.path;

const app = createApp(App)

app.provide("axios", axios)

for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
}

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