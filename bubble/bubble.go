/*
通过从左到右不断交换相邻逆序的相邻元素，在一轮的交换之后，可以让未排序的元素上浮到最右侧，是的右侧是已排序的。
在一轮交换中，如果没有发生交换，就说明数组已经是有序的，此时可以直接退出。

冒泡排序算法的原理如下：
1. 比较相邻的元素。如果第一个比第二个大，就交换他们两个。
2. 对每一对相邻元素做同样的工作，从开始第一对到结尾的最后一对。在这一点，最后的元素应该会是最大的数。
3. 针对所有的元素重复以上的步骤，除了最后一个。
4. 持续每次对越来越少的元素重复上面的步骤，直到没有任何一对数字需要比较
*/
package bubble

import "sort"

/*不实用内部包*/
func Sort(item []int) {
	var (
		n    = len(item)
		swap = true
	)

	for swap {
		swap = false
		for i := 0; i < n-1; i++ {
			if item[i] > item[i+1] {
				item[i], item[i+1] = item[i+1], item[i]
				swap = true
			}
		}
		n--
	}
}

/*使用内部排序包*/
func SortUsingSortPackage(sorter sort.Interface) {
	var (
		n    = sorter.Len()
		swap = true
	)
	for swap {
		swap = false
		for i := 0; i < n-1; i++ {
			if sorter.Less(i+1, i) {
				sorter.Swap(i+1, i)
				swap = true
			}
		}
		n--
	}
}

type array []int

func (a array) Len() int {
	return len(a)
}

func (a array) Less(i, j int) bool {
	return a[i] < a[j]
}

func (a array) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
