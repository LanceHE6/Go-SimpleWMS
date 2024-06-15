import axios from "axios";

/**
 * @description 请求头中需要附带的鉴权token
 * */
let token = "bearer " + localStorage.getItem("token");

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
export const axios_get = async ({url, params = {}, headers = {}, name = 'axios_get'}) => {
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
export const axios_delete = async ({url, data, headers = {}, name = 'axios_delete'}) => {
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
export const axios_post = async ({url, data, headers = {}, name = 'axios_post'}) => {
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
export const axios_put = async ({url, data, headers = {}, name = 'axios_put'}) => {
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
        console.error(`${name}:`, error)
    })
    return result
}
