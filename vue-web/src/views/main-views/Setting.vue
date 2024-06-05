<template>
  <el-container>
    <el-aside
        class="side-bar"
        width="collapse"
    >
      <SideBar
          class="side-menu"
          :menu-list="sideMenu"
          @selectMenu="handleSelect"
      >
      </SideBar>
    </el-aside>

    <el-main
      class="main-container"
    >
      <my-tab
        ref="myTab"
        :default-tab="defaultTab"
      >
      </my-tab>
    </el-main>

  </el-container>
</template>



<script setup>
import {onMounted, ref} from "vue";
import SideBar from "@/components/SideBar.vue";
import MyTab from "@/components/MyTab.vue";

onMounted(init)

//侧边菜单内容
const sideMenu = [
  {name: 'UserManagement', label: "用户管理", path: "/home/setting/userManagement", icon: "User"},
  {name: 'DepartmentManagement', label: "部门管理", path: "/home/setting/departmentManagement", icon: "OfficeBuilding"},
  {name: 'StaffManagement', label: "员工管理", path: "/home/setting/staffManagement", icon: "Avatar"},
  {name: 'WarehouseManagement', label: "仓库管理", path: "/home/setting/warehouseManagement", icon: "PriceTag"},
  {name: 'GoodsTypeManagement', label: "货品类型管理", path: "/home/setting/goodsTypeManagement", icon: "Box"},
  {name: 'InventoryTypeManagement', label: "出入库类型管理", path: "/home/setting/inventoryTypeManagement", icon: "Van"},
  {name: 'UnitManagement', label: "计量单位管理", path: "/home/setting/unitManagement", icon: "Filter"},
];

//初始tab内容
const defaultTab = {
    name: 'UserManagement',
    label: '用户管理',
    path: '/home/setting/userManagement',
  }

let myTab = ref(null);

//初始化函数
function init() {}

//点击侧边栏菜单
function handleSelect(menu){
  myTab.value.addTab(menu.name, menu.label, menu.path)
}
</script>

<style scoped>
.side-menu {
  height: 100%;
  border: 0 !important;
}
.main-container{
  padding: 5px;
}
</style>