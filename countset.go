package goset

// 计数集合，组合普通集合的函数
type CountSet[T SetElementInterface] map[T]int

func NewCountSet[T SetElementInterface](arr []T) CountSet[T] {
	set := make(CountSet[T])
	for key := range arr {
		set.AddChild(arr[key])
	}
	return set
}

// 转换为数组
func (set CountSet[T]) ToArr() []T {
	arr := make([]T, 0)
	for value := range set {
		arr = append(arr, value)
	}
	return arr
}

// 复制集合
func (set CountSet[T]) Copy() CountSet[T] {
	res := NewCountSet([]T{})
	for key := range set {
		res.AddChild(key)
	}
	return res
}

// 获取次数
func (set CountSet[T]) GetCount(childElement T) int {
	if count, ok := set[childElement]; ok {
		return count
	}
	return 0
}

// 添加子元素
func (set CountSet[T]) AddChild(childElement T) {
	set[childElement] += 1
}

// 删除子元素
func (set CountSet[T]) RemoveChild(childElement T) {
	delete(set, childElement)
}

// 查找是否存在于集合中
func (set CountSet[T]) InSet(element T) bool {
	_, ok := set[element]
	return ok
}

// 交集
func (set CountSet[T]) Intersection(otherSet CountSet[T]) CountSet[T] {
	newSet := NewCountSet[T]([]T{})
	for childElement := range set {
		if otherSet.InSet(childElement) {
			newSet.AddChild(childElement)
		}
	}
	return newSet
}

// 并集
func (set CountSet[T]) Union(otherSet CountSet[T]) CountSet[T] {
	newSet := NewCountSet([]T{})
	for childElement := range set {
		newSet.AddChild(childElement)
	}
	for childElement := range otherSet {
		newSet.AddChild(childElement)
	}
	return newSet
}
