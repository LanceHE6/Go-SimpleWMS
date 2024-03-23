<template>

    <el-container>
<!--        侧边栏-->
        <el-aside width="10%">
            <el-menu
                default-active="2"
                class="el-menu-vertical-demo"
                @open="handleOpen"
                @close="handleClose">

                <el-menu-item index="1">
                    <i class="el-icon-menu"></i>
                    <span slot="title">导航一</span>
                </el-menu-item>
                <el-menu-item index="2">
                    <i class="el-icon-menu"></i>
                    <span slot="title">导航二</span>
                </el-menu-item>
                <el-menu-item index="3">
                    <i class="el-icon-menu"></i>
                    <span slot="title">导航三</span>
                </el-menu-item>
                <el-menu-item index="4">
                    <i class="el-icon-menu"></i>
                    <span slot="title">导航四</span>
                </el-menu-item>
                <el-menu-item index="5">
                    <i class="el-icon-menu"></i>
                    <span slot="title">导航五</span>
                </el-menu-item>
                <el-menu-item index="6">
                    <i class="el-icon-menu"></i>
                    <span slot="title">导航五</span>
                </el-menu-item>
                <el-menu-item index="7">
                    <i class="el-icon-menu"></i>
                    <span slot="title">导航五</span>
                </el-menu-item>
                <el-menu-item index="8">
                    <i class="el-icon-menu"></i>
                    <span slot="title">导航五</span>
                </el-menu-item>
                <el-menu-item index="9">
                    <i class="el-icon-menu"></i>
                    <span slot="title">导航五</span>
                </el-menu-item>

            </el-menu>


        </el-aside>
        <el-container>
<!--            头部-->
            <el-header >
                <el-menu :default-active="activeIndex" class="el-menu-demo" mode="horizontal" @select="handleSelect" >
                <el-menu-item index="1">处理中心</el-menu-item>
                <el-menu-item index="2">处理中心</el-menu-item>
                <el-menu-item index="3">处理中心</el-menu-item>
                <el-menu-item index="4">处理中心</el-menu-item>
                <el-menu-item index="5">处理中心</el-menu-item>
                <el-menu-item index="6">处理中心</el-menu-item>
                </el-menu>
            </el-header>
<!--            主体-->
            <el-main>Main</el-main>
        </el-container>
    </el-container>


</template>





<script>
export default {
    data() {
        return {
            isCollapse: true
        };
    },
    methods: {
        handleOpen(key, keyPath) {
            console.log(key, keyPath);
        },
        handleClose(key, keyPath) {
            console.log(key, keyPath);
        }
    }
}
</script>


<style scoped>



</style>