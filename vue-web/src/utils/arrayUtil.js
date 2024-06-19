//对象数组求并集
export function getObjectArrayUnion(arr1, arr2) {
    // 将对象数组转换为JSON字符串数组
    const jsonArr1 = arr1.map(item => JSON.stringify(item));
    const jsonArr2 = arr2.map(item => JSON.stringify(item));

    // 合并并去重JSON字符串数组
    const uniqueJsonStrings = [...new Set([...jsonArr1, ...jsonArr2])];

    // 将去重后的JSON字符串转回对象数组
    return uniqueJsonStrings.map(jsonString => JSON.parse(jsonString));
}

//对象数组求并集(由对象中的key属性判断两对象是否相同, 优先取arr1的值)
export function getObjectArrayUnionByKey(arr1, arr2, keyData) {
    const key = [...keyData]
    let isArrKey = (typeof keyData !== 'string')
    if(isArrKey){
        key.reverse()
    }
    // 创建一个 Map 来存储已经看到的对象
    const map = new Map();
    // 遍历第一个数组，将对象存储在 Map 中
    for (const item of arr1) {
        if(isArrKey){
            let i = item
            for(const k of key){
                i = i[k]
            }
            map.set(i, item);
        }
        else{
            map.set(item[keyData], item);
        }
    }

    // 遍历第二个数组，如果对象不在 Map 中，则添加到 Map 中
    for (const item of arr2) {
        if(isArrKey){
            let i = item
            //前面已经用key.reverse()翻转过一次数组，所以这里就不需要了
            for(const k of key){
                i = i[k]
            }
            if (!map.has(i)) {
                map.set(i, item);
            }
        }
        else{
            if (!map.has(item[keyData])) {
                map.set(item[keyData], item);
            }
        }
    }

    // 将 Map 中的值转换回数组
    return Array.from(map.values());
}

//对象数组求差集
export function getObjectArrayDifference(arr1, arr2) {
    // 将对象数组转换为JSON字符串数组
    const jsonArr1 = arr1.map(item => JSON.stringify(item));
    const jsonArr2Set = new Set(arr2.map(item => JSON.stringify(item))); // arr2转换为Set以提高查找效率

    // 过滤arr1中不在arr2中的元素
    const difference = jsonArr1.filter(jsonString => !jsonArr2Set.has(jsonString));

    // 将过滤后的JSON字符串转回对象数组
    return difference.map(jsonString => JSON.parse(jsonString));
}