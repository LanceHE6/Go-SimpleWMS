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