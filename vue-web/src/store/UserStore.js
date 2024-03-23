import {defineStore} from "pinia";
// 保存用户信息
export const UserStore = defineStore("user", {
    state :() =>{
        return {
            uid: "",
            nickname: "",
            account: "",
            password: "",
            userType: 0, // 0为普通用户， 1为管理员
            token: "",
            avatar: "",
            remember: false
        }
    },
    actions: {
        logout(){
            UserStore().uid = "";
            UserStore().nickname = "";
            UserStore().account = "";
            UserStore().password = "";
            UserStore().userType = 0;
            UserStore().token = "";
            UserStore().avatar = "";
        }
    },
    getters: {}
})