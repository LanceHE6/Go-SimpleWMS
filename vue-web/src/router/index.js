import { createRouter, createWebHistory } from 'vue-router';
import Login from '../views/Login.vue';
import Main from  '../views/Main.vue'
import ProductManagement from  '../views/main-views/ProductManagement.vue'
import EntryAndOut from  '../views/main-views/EntryAndOut.vue'




const routes = [
    { path: '/', component: Login , name: 'login'},
    { path: '/home', component: Main , name: 'home',
        children: [
            // 二级路由
            { path: '/home/productManagement', component: ProductManagement , name: 'productManagement'},
            { path: '/home/entryAndOut', component: EntryAndOut , name: 'entryAndOut'},
        ]
    },
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

export {router, routes};
