package utils

// CreateUpdateData 因前端传回来的参数可能会有零值,使用这个函数构造更新数据过滤零值
// 参数格式: ("arg1", value1, "arg2", value2, ...)
func CreateUpdateData(args ...interface{}) map[string]interface{} {
	// 判断参数个数是否为偶数个
	if len(args)%2 != 0 {
		// 抛出参数必须为偶数个的异常
		panic("The number of parameters must be an even number")
	}

	// 构造更新数据
	updateData := make(map[string]interface{})
	for i := 0; i < len(args); i += 2 {
		key, ok := args[i].(string) // 类型断言, 确保key是字符串
		if !ok {
			// 抛出key必须为字符串的异常
			panic("The key must be a string")
		}
		value := args[i+1]
		if value != "" {
			updateData[key] = value
		}
	}
	return updateData
}
