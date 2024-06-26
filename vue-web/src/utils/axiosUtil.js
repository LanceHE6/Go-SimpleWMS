import axios from "axios";
import {ElMessage} from "element-plus";
import {router} from "@/router/index.js";

// 错误映射和定时器存储
const errorMap = new Map();

// 网络请求失败的回调函数
const showErrMessage = (errStr) => {
    const currentTime = Date.now();

    // 检查是否已经在短时间内报告过相同的错误
    if (!errorMap.has(errStr) || currentTime - errorMap.get(errStr) > 5000) {
        // 如果没有，显示提示框并将错误添加到映射中
        ElMessage.error(errStr);
        errorMap.set(errStr, currentTime);

        // 设置一个定时器，在3秒后从映射中删除该错误
        setTimeout(() => {
            if (errorMap.get(errStr) === currentTime) {
                errorMap.delete(errStr);
            }
        }, 3000);
    }
}

/**
 * @description 请求头中需要附带的鉴权token
 * */
let token = ''

const getToken = () => {
    token = "bearer " + localStorage.getItem("token");
}

/**
 * @param url 请求url
 * @param params 请求参数
 * @param headers 请求头
 * @param name 调用者
 * @description 发送get请求的函数
 * @return resultObj 请求结果
 * */
export const axiosGet = async ({url, params = {}, headers = {}, name = 'axios_get'}) => {
    getToken()
    let resultObj = false
    const defaultHeaders = {
        'Authorization': token
    }
    await axios.get(`/api${url}`, {
        headers: {...defaultHeaders, ...headers},
        params: params
    })
    .then(result => {
        console.log(`${name}:`, result)
        if (result && result.status === 200 && result.data) {
            resultObj = result.data;
        }
    })
    .catch(error => {
        catchError(error)
        console.error(`${name}:`, error)
    })
    return resultObj
}

/**
 * @param url 请求url
 * @param data 请求体
 * @param headers 请求头
 * @param name 调用者
 * @description 发送delete请求的函数
 * @return result 请求是否成功
 * */
export const axiosDelete = async ({url, data, headers = {}, name = 'axiosDelete'}) => {
    getToken()
    let result = false
    const defaultHeaders = {
        'Authorization': token
    }
    await axios.delete(`/api${url}`, {
        headers: {...defaultHeaders, ...headers},
        data: data
    })
    .then( message => {
        console.log(`${name}:`, message)
        result = message
    })
    .catch( error => {
        catchError(error)
        console.error(`${name}:`, error)
    })
    return result
}

/**
 * @param url 请求url
 * @param data 请求体
 * @param headers 请求头
 * @param name 调用者
 * @description 发送post请求的函数
 * @return result 请求是否成功
 * */
export const axiosPost = async ({url, data, headers = {}, name = 'axiosPost'}) => {
    getToken()
    let result = false
    const defaultHeaders = {
        'Authorization': token
    }
    await axios.post(`/api${url}`, data, {
        headers: {...defaultHeaders, ...headers},
    })
    .then( message => {
        console.log(`${name}:`, message)
        result = message
    })
    .catch( error => {
        catchError(error)
        console.error(`${name}:`, error)
    })
    return result
}

/**
 * @param url 请求url
 * @param data 请求体
 * @param headers 请求头
 * @param name 调用者
 * @description 发送put请求的函数
 * @return result 请求是否成功
 * */
export const axiosPut = async ({url, data, headers = {}, name = 'axiosPut'}) => {
    getToken()
    let result = false
    const defaultHeaders = {
        'Authorization': token
    }
    await axios.put(`/api${url}`, data, {
        headers: {...defaultHeaders, ...headers},
    })
    .then( message => {
        console.log(`${name}:`, message)
        result = message
    })
    .catch( error => {
        catchError(error)
        console.error(`${name}:`, error)
    })
    return result
}

const catchError = async (err) => {
    let status = 0
    if (err && err.response && err.response.status) {
        status = err.response.status
    }
    switch (status) {
        case 401: {
            showErrMessage("用户信息已过期，请重新登录！")
            await router.push('/')
            break
        }
        default: {
            showErrMessage("网络请求出错了！")
            break
        }
    }
}
