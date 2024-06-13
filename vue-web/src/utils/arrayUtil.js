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