export function getObjKeyData(obj, keyList){
    if(typeof keyList === "string"){
        return obj[keyList]
    }

    let data = obj
    for(const item of keyList){
        data = data[item]
    }
    return data
}

export function editObjKeyData(obj, newValue, keyList) {
    if(typeof keyList === "string"){
        obj[keyList] = newValue
        return obj
    }

    // 深度复制对象的函数
    function deepCopy(obj) {
        if (typeof obj !== 'object' || obj === null) {
            return obj;
        }
        if (Array.isArray(obj)) {
            return obj.map(deepCopy);
        }
        const copy = {};
        for (const key in obj) {
            if (obj.hasOwnProperty(key)) {
                copy[key] = deepCopy(obj[key]);
            }
        }
        return copy;
    }

    // 对原始对象进行深拷贝
    const resultObj = deepCopy(obj);
    let data = resultObj;

    // 反向遍历keyList数组，以便从最深层的属性开始设置
    for (let i = keyList.length - 1; i >= 0; i--) {
        const key = keyList[i];

        // 如果是最后一个属性，则设置其值
        if (i === 0) {
            data[key] = newValue;
            break;
        }

        // 确保上级属性是一个对象，如果不是则创建一个空对象
        if (!data.hasOwnProperty(key) || typeof data[key] !== 'object') {
            data[key] = {};
        }

        // 向下遍历到下一个属性
        data = data[key];
    }

    return resultObj;
}

export function objectToUrl(obj) {
    // 确保obj有path和query属性
    if (!obj || !obj.path || !obj.query) {
        throw new Error('Invalid object. It should have path and query properties.');
    }

    // 创建一个数组来存储所有的参数对
    const queryParams = [];

    // 遍历query对象的每个属性
    for (const key in obj.query) {
        if (obj.query.hasOwnProperty(key)) {
            // 对值和键都进行encodeURIComponent以防止特殊字符破坏URL
            const value = encodeURIComponent(obj.query[key]);
            // 将参数对添加到数组中
            queryParams.push(`${key}=${value}`);
        }
    }

    // 将参数对用'&'连接起来
    const queryString = queryParams.join('&');

    // 附加查询字符串到path上，如果path不包含'?'则添加'?'
    return obj.path.includes('?') ? `${obj.path}&${queryString}` : `${obj.path}?${queryString}`;
}