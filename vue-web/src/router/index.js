import { createRouter, createWebHistory } from 'vue-router';
import Login from '../views/Login.vue';

import Main from  '../views/Main.vue'
import HomePage from  '../views/main-views/HomePage.vue'

import ProductManagement from  '../views/main-views/ProductManagement.vue'

import EntryAndOut from  '../views/main-views/EntryAndOut.vue'
import Entry from '../views/main-views/entry-and-out/Entry.vue'
import Out from '../views/main-views/entry-and-out/Out.vue'
import Check from '../views/main-views/entry-and-out/Check.vue'
import Allocate from '../views/main-views/entry-and-out/Allocate.vue'


import Setting from  '../views/main-views/Setting.vue'



const routes = [
    { path: '/', component: Login , name: 'login'},
    { path: '/home', component: Main , name: 'home',
        children: [
            { path: '/home/homePage', component: HomePage , name: 'homePage'},
            { path: '/home/productManagement', component: ProductManagement , name: 'productManagement'},
            { path: '/home/entryAndOut', component: EntryAndOut , name: 'entryAndOut',
                children:[
                    { path: '/home/entryAndOut/entry', component: Entry , name: 'entry'},
                    { path: '/home/entryAndOut/out', component: Out , name: 'out'},
                    { path: '/home/entryAndOut/check', component: Check , name: 'check'},
                    { path: '/home/entryAndOut/allocate', component: Allocate , name: 'allocate'},
                ]},
            { path: '/home/setting', component: Setting , name: 'setting'},
        ]
    },
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

export {router, routes};
