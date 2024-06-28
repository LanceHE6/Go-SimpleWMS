//内容相同约束
export const isSame = (rule, value, callback, objValue, message='两次输入的内容不相同') => {
    if(value === ''){
        callback()
    }
    if(value !== objValue){
        callback(new Error(message))
    }
    else{
        callback()
    }
}

//内容不相同约束
export const isNotSame = (rule, value, callback, objValue, message='两次输入的内容不能相同') => {
    if(value === ''){
        callback()
    }
    if(value === objValue){
        callback(new Error(message))
    }
    else{
        callback()
    }
}

//条件内容不为空约束
export const isNotEmptyCondition = (rule, value, callback, objValue, message='内容不能为空') => {
    if(objValue === ''){
        callback()
    }
    if(value === '' && objValue !== ''){
        callback(new Error(message))
    }
    else{
        callback()
    }
}
export const isEmail = (rule, value, callback, message='邮箱格式不正确') => {
    if(value === ''){
        callback()
    }
    const reg = /^([a-zA-Z]|[0-9])(\w|-)+@[a-zA-Z0-9]+\.([a-zA-Z]{2,4})$/;
    if(reg.test(value)){
        callback()
    }
    else{
        callback(new Error(message))
    }
}

//密码合法约束
export const isPasswordValid = (rule, value, callback) => {
    if(value === '' || (/[a-zA-Z]/.test(value) && /[0-9]/.test(value))){
        callback()
    }
    else{
        callback(new Error('密码必须由数字和字母组成'))
    }
}

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
