

# Go-SimpleWMS前端文档

## 总项目结构

```
doc 			项目文档
vue-web 		前端项目
go-server 		后端项目
```



## 前端项目结构

```
src				主目录
├───component 		子组件
├───res 		    资源
├───router 		    路由表
├───utils           工具类
└───views			所有界面
    └───main-views  	主菜单界面
        ├───xxx1  	    	子菜单界面1(xxx1为主菜单名)
        ├───xxx2  	    	子菜单界面2(xxx2为主菜单名)
        └───xxx3  	    	子菜单界面3(xxx3为主菜单名)
```



## 启动项目

```shell
cd vue-web
npm i 
npm run dev
```



## 开发人员文档

### 子组件

#### 子组件列表及依赖关系

```
├───MyTab				自定义tab界面
├───SideBar				侧边栏
├───DataShowView		数据展示界面
│	└───MyTable				数据展示界面中的表格
│		├───TableHeader			数据展示表格首部按钮及搜索按钮等
│		└───TableBody           数据展示表格主体部分
└───SubmitView          数据提交界面

```



#### 子组件的使用

##### MyTab

###### 属性

| 属性名            | 属性类型   | 是否必填 | 属性说明      |
|----------------|--------|------|-----------|
| **defaultTab** | Object | 否    | 默认tab界面对象 |



| defaultTab对象属性名 | 属性类型   | 是否必填 | 属性说明         |
|-----------------|--------|------|--------------|
| name            | String | 是    | 新标签页组件名      |
| label           | String | 是    | 新标签页显示名称     |
| path            | String | 是    | 新标签页router路径 |



###### 外部函数

| 函数名           | 是否必须实现  | 函数说明      |
|---------------|---------|-----------|
| **addTab**    | 否（推荐实现） | 添加新标签页的方法 |
| **removeTab** | 否       | 移除新标签页的方法 |



| addTab参数名 | 参数类型   | 是否必填 | 参数说明         |
|-----------|--------|------|--------------|
| name      | String | 是    | 新标签页组件名      |
| label     | String | 是    | 新标签页显示名称     |
| path      | String | 是    | 新标签页router路径 |
| params    | String | 否    | 新标签页router参数 |



| removeTab参数名 | 参数类型   | 是否必填 | 参数说明         |
|--------------|--------|------|--------------|
| currentTab   | String | 是    | 要移除的router路径 |



###### 监听事件

| 事件名           | 事件说明             |
|---------------|------------------|
| **addTab**    | 添加新标签页，参数与外部函数相同 |
| **removeTab** | 移除新标签页，参数与外部函数相同 |

tips:可在子组件中发送事件来实现添加/移除新标签页的效果



###### 示例

```vue
<template>
  <my-tab
      ref="myTab"
      :default-tab="defaultTab"
  />
</template>

<script setup>
//初始tab内容
const defaultTab = {
  label: "全部货品",
  path: "/home/productManagement/allProduction"
}

//获取my-tab组件
let myTab = ref(null)

//添加新标签页的函数
function addTab(name, label, path){
  myTab.value.addTab(name, label, path)
}
</script>
```



##### SideBar

###### 属性

| 属性名          | 属性类型                | 是否必填 | 属性说明      |
|--------------|---------------------|------|-----------|
| **menuList** | Array&lt;Object&gt; | 否    | 侧边栏的菜单列表  |



| menuList列表对象属性名 | 属性类型   | 是否必填 | 属性说明             |
|-----------------|--------|------|------------------|
| name            | String | 是    | 子菜单组件名           |
| label           | String | 是    | 子菜单显示名称          |
| path            | String | 是    | 子菜单router路径      |
| icon            | String | 否    | 图标名(参考el-plus文档) |



###### 事件

| 事件名            | 是否必须实现  | 事件说明     |
|----------------|---------|----------|
| **selectMenu** | 否（推荐实现） | 点击子菜单时触发 |



| selectMenu返回值 | 属性类型   | 属性说明   |
|---------------|--------|--------|
| **menu**      | Object | 新标签页名称 |



| menu对象属性名 | 属性类型   | 属性说明        |
|-----------|--------|-------------|
| name      | String | 子菜单组件名      |
| label     | String | 子菜单显示名称     |
| path      | String | 子菜单router路径 |



###### 示例

```vue
<template>
  <side-bar
      :menu-list="sideMenu"
      @selectMenu="handleSelect"
  />
</template>

<script setup>
//侧边菜单内容
const sideMenu = [
  {name: 'UserManagement', label: "用户管理", path: "/home/setting/userManagement", icon: "User"},
  {name: 'StaffManagement', label: "员工管理", path: "/home/setting/staffManagement", icon: "Avatar"},
]

//事件
function handleSelect(menu){
  //内部实现
}
</script>
```



##### DataShowView

###### 属性

| 属性名                | 属性类型                | 是否必填        | 属性说明                |
|--------------------|---------------------|-------------|---------------------|
| tableHeight        | String              | 否(默认为60vh)  | 表格高度                |
| large              | Boolean             | 否(默认为false) | 是否启用后端分页            |
| **tableColList**   | Array&lt;Object&gt; | 是           | 表头属性列表              |
| keyData            | String              | 是           | 主键                  |
| searchData         | String              | 否           | 搜索框的搜索依据（非后端分页）     |
| **addForm**        | Object              | 是           | 添加窗口属性对象            |
| **editForm**       | Object              | 是           | 编辑窗口属性对象            |
| **urls**           | Object              | 是           | 请求url               |
| **extraParams**    | Object              | 否           | 数据获取请求体中的额外参数       |
| delete             | Boolean             | 否(默认为false) | 是否支持删除功能            |
| upload             | Boolean             | 否(默认为false) | 是否支持上传功能            |
| download           | Boolean             | 否(默认为false) | 是否支持下载功能            |
| print              | Boolean             | 否(默认为false) | 是否支持打印功能            |
| uploadImg          | Boolean             | 否(默认为false) | 是否支持图片上传功能          |
| uploadFile         | Boolean             | 否(默认为false) | 是否支持附件上传功能          |
| hasSubmitPage      | Boolean             | 否(默认为false) | 是否需要提交表单的页面         |
| hasRefreshEventBus | Boolean             | 否(默认为false) | 是否需要控制页面刷新的eventBus |



| tableColList列表对象属性名 | 属性类型                | 是否必填        | 属性说明                         |
|---------------------|---------------------|-------------|------------------------------|
| property            | String              | 是           | 属性名，需要与后端返回的属性名一致            |
| label               | String              | 是           | 属性显示在表头上的名字                  |
| sortable            | Boolean             | 否(默认为false) | 是否支持排序                       |
| width               | Number              | 否           | 列宽                           |
| isFixed             | Boolean             | 否(默认为false) | 是否需要固定此列                     |
| isExpand            | Boolean             | 否(默认为false) | 是否启用子表格显示                    |
| children            | Array&lt;Object&gt; | 否           | 子表格的tableColList             |
| isParent            | Boolean             | 否(默认为false) | 是否为父级对象                      |
| child               | Object              | 否           | 子对象的属性，与tableColList列表对象格式一致 |
| isMapping           | Boolean             | 否(默认为false) | 是否需要属性映射                     |
| **mappingList**     | Array&lt;Object&gt; | 否           | 映射列表                         |
| isFK                | Boolean             | 否(默认为false) | 是否需要外键映射                     |
| **FKData**          | Object              | 否           | 外键属性                         |
| isImage             | Boolean             | 否(默认为false) | 是否为图片                        |
| isFile              | Boolean             | 否(默认为false) | 是否为附件                        |
| isInput             | Boolean             | 否(默认为false) | 是否为输入框                       |
| operable            | Boolean             | 否(默认为false) | 是否有点击事件                      |
| **operationEvent**  | Function            | 否           | 点击事件                         |



| mappingList列表对象属性名 | 属性类型   | 是否必填 | 属性说明 |
|--------------------|--------|------|------|
| label              | String | 是    | 显示名  |
| value              | Any    | 是    | 值    |



| FKData对象属性名 | 属性类型   | 是否必填 | 属性说明        |
|-------------|--------|------|-------------|
| url         | String | 是    | 外键url       |
| property    | String | 是    | 外键主键名       |
| label       | String | 是    | 映射为外键表的哪个属性 |



###### 示例

```vue
<template>
  <DataShowView
    :table-col-list="tableColList"
    :key-data="'wid'"
    :search-data="'name'"
    :add-form="addForm"
    :edit-form="editForm"
    :delete-data-body="deleteDataBody"
    :urls="urls"
  />
</template>

<script setup>

/**
 * 自定义点击函数
 * */
const audit = (tid, isAudit) =>{
  console.log(tid, isAudit)
}
/**
 * 表头属性列表
 * */
const tableColList = [
  {property: "goods_list", label: "货品信息", sortable: false, width:100, isFixed: true, isExpand: true, children:[
      {property: "goods", label: "图片", sortable: false, width: 120, isImage: true, isParent: true,
        child:{
          property: "images", isParent: false
        }},
      {property: "goods", label: "货品编码", sortable: false, width: 120, isParent: true,
        child:{
          property: "goods_code", isParent: false
        }},
      {property: "amount", label: "数量", sortable: false, width: 120},
  ]},
  {property: "audited", label: "审核状态", sortable: false, isFixed:true, operable: true, operationEvent: audit, width: 100, isMapping: true, mappingList:[
      {label: '已审核', value: true},
      {label: '待审核', value: false},
    ]},
  {property: "comment", label: "备注", sortable: false},
  {property: "manager", label: "管理人", sortable: false, isFK: true, FKData:{
      url: "/staff/list",
      property: "sid",
      label: "name"
    }},
  {property: "status", label: "状态", sortable: false, isMapping: true, mappingList:[
      {label: "禁用", value: 0},
      {label: "启用", value: 1},
    ]},
  {property: "created_at", label: "创建时间", sortable: true},
  {property: "wid", label: "仓库ID", sortable: true, width: 120},
]

/**
 * 编辑仓库时所用到的表单对象
 * */
const editForm = {
  data : {
    wid:'',
    name:'',
    comment:'',
    manager:'',
    status:'',
  },
  dataType:{
    wid:'String',
    name:'String',
    comment:'String',
    manager:'String',
    status:'Int',
  },
  dataNum: 5,
  rules: {
    status:[
      { required: 'true', message: '请选择仓库状态', trigger: 'blur' }
    ],
    name:[
      { required: 'true', message: '请输入仓库名', trigger: 'blur' },
    ],
    manager:[
      { required: 'true', message: '请选择管理人', trigger: 'blur' },
    ]
  },
  item:[
    {label: '仓库名', prop: 'name', dataName: 'name', isInput: true,},
    {label: '是否启用', prop: 'status', dataName: 'status', isSelect: true,
      selectOptions: [
        {label: '禁用', value: 0},
        {label: '启用', value: 1},
      ]},
    {label: '管理人', prop: 'manager', dataName: 'manager', isFK: true,
      FKData:{
        url: "/staff/list",
        property: "sid",
        label: "name"
      }},
    {label: '备注', prop: 'comment', dataName: 'comment', isInput: true},
  ],
}

/**
 * 添加仓库时所用到的对象
 * */
const addForm = {
  data :{
    name:'',
    comment:'',
    manager:'',
    status:'',
  },
  dataType:{
    name:'String',
    comment:'String',
    manager:'String',
    status:'Int',
  },
  dataNum: 4,
  rules: {
    name: [
      { required: 'true', message: '仓库名不能为空', trigger: 'blur' },
    ],
    status:[
      { required: 'true', message: '请选择仓库状态', trigger: 'blur' }
    ],
    manager:[
      { required: 'true', message: '请选择管理人', trigger: 'blur' }
    ]
  }
  ,
  item:[
    {label: '仓库名', prop: 'name', dataName: 'name', isInput: true},
    {label: '备注', prop: 'comment', dataName: 'comment', isInput: true},
    {label: '仓库状态', prop: 'status', dataName: 'status', isSelect: true,
      selectOptions: [
        {label: '禁用', value: 0},
        {label: '启用', value: 1},
      ]},
    {label: '管理人', prop: 'manager', dataName: 'manager', isFK: true,
      FKData:{
        url: "/staff/list",
        property: "sid",
        label: "name"
      }},
  ],
}

/**
 * 网络请求url
 * */
const urls = {
  getData: "/warehouse/list",
  deleteData: "/warehouse/delete",
  addData: "/warehouse/add",
  updateData: "/warehouse/update",
  uploadData: "/warehouse/upload",
}
</script>
```

