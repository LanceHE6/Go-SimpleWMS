//获取用户对象
import {ref} from "vue";

//用户
export const CURRENT_USER = ref(getUser())

export function getUser(){
    const userJson = localStorage.getItem("user") || '';
    if(!userJson){
        return ''
    }
    else{
        return JSON.parse(userJson)
    }
}

//设置用户
export function setUser(item){
    if(typeof item === 'string'){
        localStorage.setItem("user", item)
    }
    else if(typeof item === 'object'){
        localStorage.setItem("user", JSON.stringify(item))
    }
    refreshUser()
}

//更新用户
export function refreshUser(){
    CURRENT_USER.value = getUser()
}

//获取用户权限
export function getUserPermission(){
    const user = getUser()
    if(!user){
        return 0
    }
    else{
        return user.permission
    }
}