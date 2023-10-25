package tool

// If 模拟三元表达式
// desc: condition为true时，返回trueVal，否则返回falseVal
func If[T any](condition bool, trueVal T, falseVal T) T {
	if condition {
		return trueVal
	}
	return falseVal
}
