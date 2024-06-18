import { createRouter, createWebHashHistory } from 'vue-router';
import Login from '@/views/Login.vue';

import Main from  '@/views/Main.vue'
import HomePage from  '@/views/main-views/HomePage.vue'

import ProductManagement from  '@/views/main-views/ProductManagement.vue'
import ProductionDefine from "@/views/main-views/product-management/ProductionDefine.vue";
import AllStock from "@/views/main-views/product-management/AllStock.vue";

import EntryAndOut from  '@/views/main-views/EntryAndOut.vue'
import Entry from '@/views/main-views/entry-and-out/Entry.vue'
import Out from '@/views/main-views/entry-and-out/Out.vue'
import Check from '@/views/main-views/entry-and-out/Check.vue'
import Allocate from '@/views/main-views/entry-and-out/Allocate.vue'
import InventorySubmit from "@/views/main-views/entry-and-out/submit-views/InventorySubmit.vue";
import AllocateSubmit from "@/views/main-views/entry-and-out/submit-views/AllocateSubmit.vue";


import Setting from  '@/views/main-views/Setting.vue'
import UserManagement from '@/views/main-views/setting/UserManagement.vue'
import WarehouseManagement from '@/views/main-views/setting/WarehouseManagement.vue'
import GoodsTypeManagement from '@/views/main-views/setting/GoodsTypeManagement.vue'
import DepartmentManagement from '@/views/main-views/setting/DepartmentManagement.vue'
import StaffManagement from '@/views/main-views/setting/StaffManagement.vue'
import InventoryTypeManagement from '@/views/main-views/setting/InventoryTypeManagement.vue'
import UnitManagement from "@/views/main-views/setting/UnitManagement.vue";


const routes = [
    { path: '/', component: Login , name: 'login'},
    { path: '/home', component: Main , name: 'home', meta: { requiresAuth: true },
        children: [
            { path: 'homePage', component: HomePage , name: 'homePage'},
            { path: 'productManagement', component: ProductManagement , name: 'productManagement',
                children: [
                    { path: 'productionDefine', component: ProductionDefine , name: 'productionDefine'},
                    { path: 'allStock', component: AllStock , name: 'allStock'},
                ]
            },
            { path: 'entryAndOut', component: EntryAndOut , name: 'entryAndOut',
                children:[
                    { path: 'entry', component: Entry , name: 'entry'},
                    { path: 'out', component: Out , name: 'out'},
                    { path: 'check', component: Check , name: 'check'},
                    { path: 'allocate', component: Allocate , name: 'allocate'},
                    { path: 'inventorySubmit', component: InventorySubmit , name: 'inventorySubmit', props: true},
                    { path: 'allocateSubmit', component: AllocateSubmit , name: 'allocateSubmit', props: true},
                ]
            },
            { path: 'setting', component: Setting , name: 'setting',
                children:[
                    { path: 'userManagement', component: UserManagement , name: 'userManagement'},
                    { path: 'warehouseManagement', component: WarehouseManagement , name: 'warehouseManagement'},
                    { path: 'goodsTypeManagement', component: GoodsTypeManagement , name: 'goodsTypeManagement'},
                    { path: 'departmentManagement', component: DepartmentManagement , name: 'departmentManagement'},
                    { path: 'staffManagement', component: StaffManagement , name: 'staffManagement'},
                    { path: 'inventoryTypeManagement', component: InventoryTypeManagement , name: 'inventoryTypeManagement'},
                    { path: 'unitManagement', component: UnitManagement , name: 'unitManagement'},
                ]
            },
        ]
    },
];

const router = createRouter({
    history: createWebHashHistory(),
    routes,
});

export {router, routes};
