import { createRouter, createWebHistory } from 'vue-router';
import Login from '../views/Login.vue';




const routes = [
    { path: '/', component: Login , name: 'homepage'},
    // { path: '/home', component: Home , name: 'home'},
    // { path: '/login', component: Login , name: 'login'},
    // { path: '/signup', component: Signup , name: 'signup'},
    // { path: '/detail', component: ArticleDetail , name: 'articleDetail'},
    // { path: '/userInfo', component: UserInfo , name: 'userInfo'},
    // { path: '/dashboard', component: Dashboard , name: 'dashboard', children: [
    //         // 二级路由
    //         { path: '/dashboard/category', component: Category , name: 'category'},
    //         { path: '/dashboard/article', component: Article , name: 'article'},
    //     ]},

];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

export {router, routes};
