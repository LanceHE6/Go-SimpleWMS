
//正整数必填项约束
export const pIntValidatorRequire = (rule, value, callback) => {
    const valueNum = parseInt(value)
    if(isNaN(valueNum)){
        callback(new Error('该项必须是正整数'))
    }
    if(valueNum <= 0){
        callback(new Error('该项必须是正整数'))
    }
    else{
        callback()
    }
}

//正整数非必填项约束
export const pIntValidatorNRequire = (rule, value, callback) => {
    if(value === ''){
        callback()
    }
    const valueNum = parseInt(value)
    if(!isNumeric(value) || isNaN(valueNum)){
        callback(new Error('该项必须是正整数'))
    }
    if(valueNum <= 0){
        callback(new Error('该项必须是正整数'))
    }
    else{
        callback()
    }
}

//正数非必填项约束
export const pNumValidatorNRequire = (rule, value, callback) => {
    if(value === ''){
        callback()
    }
    const valueNum = parseFloat(value)
    if(!isNumeric(value) || isNaN(valueNum)){
        callback(new Error('该项必须是正数'))
    }
    else if(valueNum < 0){
        callback(new Error('该项必须是正数'))
    }
    else{
        callback()
    }
}

//字符串不相等约束
export const strValidatorNEqual = (rule, value, callback, otherValue, errStr) => {
    if (value === '') {
        callback(); // 如果为空，则通过验证
    }
    else if (value === otherValue) {
        callback(new Error(errStr)); // 如果值相同，则显示错误
    }
    else {
        callback(); // 如果满足所有条件，则通过验证
    }
}

function isNumeric(str) {
    const pattern = /^\d+(\.\d+)?$/; // 匹配整数和小数
    return pattern.test(str);
}
