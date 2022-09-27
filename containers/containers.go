package containers

import "github.com/aiyamayaa/pta/utils"

// Container is base interface that all data structures implement.
//container 是一种接口类型
type Container interface {
	Empty() bool
	Size() int
	Clear()
	Values() []interface{}
	String() string
}

// GetSortedValues returns sorted container's elements with respect to the passed comparator.
// Does not affect the ordering of elements within the container.
func GetSortedValues(container Container, comparator utils.Comparator) []interface{} {
	values := container.Values()
	if len(values) < 2 {
		return values
	}
	utils.Sort(values, comparator)
	return values
}
