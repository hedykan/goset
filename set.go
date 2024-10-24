package goset

type SetElementInterface interface {
	int | int8 | int16 | int32 | int64 | string
}

// 普通集合
type Set[T SetElementInterface] map[T]struct{}

func NewSet[T SetElementInterface](arr []T) Set[T] {
	set := make(Set[T])
	for key := range arr {
		set.AddChild(arr[key])
	}
	return set
}

// 转换为数组
func (set Set[T]) ToArr() []T {
	arr := make([]T, 0)
	for value := range set {
		arr = append(arr, value)
	}
	return arr
}

// 复制集合
func (set Set[T]) Copy() Set[T] {
	res := NewSet([]T{})
	for key := range set {
		res.AddChild(key)
	}
	return res
}

// 添加子元素
func (set Set[T]) AddChild(childElement T) {
	set[childElement] = struct{}{}
}

// 删除子元素
func (set Set[T]) RemoveChild(childElement T) {
	delete(set, childElement)
}

// 查找是否存在于集合中
func (set Set[T]) InSet(element T) bool {
	_, ok := set[element]
	return ok
}

// 交集
func (set Set[T]) Intersection(otherSet Set[T]) Set[T] {
	newSet := NewSet([]T{})
	for childElement := range set {
		if otherSet.InSet(childElement) {
			newSet.AddChild(childElement)
		}
	}
	return newSet
}

// 并集
func (set Set[T]) Union(otherSet Set[T]) Set[T] {
	newSet := NewSet([]T{})
	for childElement := range set {
		newSet.AddChild(childElement)
	}
	for childElement := range otherSet {
		newSet.AddChild(childElement)
	}
	return newSet
}
