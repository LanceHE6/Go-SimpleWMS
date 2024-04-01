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
import UserManagement from '../views/main-views/setting/UserManagement.vue'
import WarehouseManagement from '../views/main-views/setting/WarehouseManagement.vue'
import GoodsTypeManagement from '../views/main-views/setting/GoodsTypeManagement.vue'
import DepartmentManagement from '../views/main-views/setting/DepartmentManagement.vue'
import StaffManagement from '../views/main-views/setting/StaffManagement.vue'
import InventoryTypeManagement from '../views/main-views/setting/InventoryTypeManagement.vue'

const routes = [
    { path: '/', component: Login , name: 'login'},
    { path: '/home', component: Main , name: 'home', meta: { requiresAuth: true, beforeName:'login' },
        children: [
            { path: '/home/homePage', component: HomePage , name: 'homePage', meta: { beforeName:'main' }},
            { path: '/home/productManagement', component: ProductManagement , name: 'productManagement', meta: { beforeName:'main' }},
            { path: '/home/entryAndOut', component: EntryAndOut , name: 'entryAndOut', meta: { beforeName:'main' },
                children:[
                    { path: '/home/entryAndOut/entry', component: Entry , name: 'entry', meta: { beforeName:'entryAndOut' }},
                    { path: '/home/entryAndOut/out', component: Out , name: 'out'},
                    { path: '/home/entryAndOut/check', component: Check , name: 'check'},
                    { path: '/home/entryAndOut/allocate', component: Allocate , name: 'allocate'},
                ]
            },
            { path: '/home/setting', component: Setting , name: 'setting', meta: { beforeName:'main' },
                children:[
                    { path: '/home/setting/userManagement', component: UserManagement , name: 'userManagement'},
                    { path: '/home/setting/warehouseManagement', component: WarehouseManagement , name: 'warehouseManagement'},
                    { path: '/home/setting/goodsTypeManagement', component: GoodsTypeManagement , name: 'goodsTypeManagement'},
                    { path: '/home/setting/departmentManagement', component: DepartmentManagement , name: 'departmentManagement'},
                    { path: '/home/setting/staffManagement', component: StaffManagement , name: 'sStaffManagement'},
                    { path: '/home/setting/inventoryTypeManagement', component: InventoryTypeManagement , name: 'inventoryTypeManagement'},
                ]
            },
        ]
    },
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

export {router, routes};
