
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

function isNumeric(str) {
    const pattern = /^\d+(\.\d+)?$/; // 匹配整数和小数
    return pattern.test(str);
}
