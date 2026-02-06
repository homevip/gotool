package cartesian

// 笛卡尔积组合(两个集合之间所有可能的有序对的集合)
func CartesianProduct[T comparable](slices ...[]T) [][]T {

	// 无参数直接返空
	if len(slices) == 0 {
		return nil
	}

	// 单切片,元素封装为二维切片返回(符合笛卡尔积定义)
	if len(slices) == 1 {
		res := make([][]T, len(slices[0]))
		for i, v := range slices[0] {
			res[i] = []T{v}
		}
		return res
	}

	var (
		rest   = CartesianProduct(slices[1:]...)          // 多切片:递归处理,拆解为[第一个切片]+[剩余切片的笛卡尔积]
		result = make([][]T, 0, len(slices[0])*len(rest)) // 预分配容量,避免动态扩容(性能优化保留)
	)

	// 双层循环组合结果,精简临时切片创建逻辑
	for _, v := range slices[0] {
		for _, p := range rest {
			result = append(result, append([]T{v}, p...))
		}
	}

	return result
}
