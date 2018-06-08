package algorithm
//
//import (
//	"fmt"
//	"time"
//	"math/rand"
//	"sort"
//)
//
//func GenerateRandomNumber(start int, end int, count int) []int {
//	if end < start || (end-start) < count {
//		return nil
//	}
//
//	nums := make([]int, 0,1000000000)
//	r := rand.New(rand.NewSource(time.Now().UnixNano()))
//	for len(nums) < count {
//		num := r.Intn(end - start) + start
//
//		exist := false
//		for _, v := range nums {
//			if v == num {
//				exist = true
//				break
//			}
//		}
//
//		if !exist {
//			nums = append(nums, num)
//		}
//	}
//
//	return nums
//}
//
//type array []int
//func (a array) Len() int {return len(a)}
//func (a array) Less(i, j int) bool{return a[i] < a[j]}
//func (a array) Swap(i, j int){a[i], a[j] = a[j], a[i]}
//
//const (
//	Start int = 0
//	End int = 10000000
//	Counts int = 1000000
//)
//func main() {
//	item := GenerateRandomNumber(Start,End,Counts)
//	Bubble := make([]int,0,len(item))
//	Bubble = append(Bubble,item...)
//
//	BubbleUsingSortPackage := make([]int,0,len(item))
//	bubbleUsingSortPackage := array(append(BubbleUsingSortPackage,item...))
//
//	Insertion := make([]int,0,len(item))
//	Insertion = append(Insertion,item...)
//
//	InsertionUsingSortPackage := make([]int,0,len(item))
//	insertionUsingSortPackage := array(append(InsertionUsingSortPackage,item...))
//
//	Selection := make([]int,0,len(item))
//	Selection = append(Selection,item...)
//
//	SelectionUsingSortPackage := make([]int,0,len(item))
//	selectionUsingSortPackage := array(append(SelectionUsingSortPackage,item...))
//
//	Quick := make([]int,0,len(item))
//	Quick = append(Quick,item...)
//
//	Shell := make([]int,0,len(item))
//	Shell =append(Shell,item...)
//
//	Comb := make([]int,0,len(item))
//	Comb = append(Comb,item...)
//
//	Merge := make([]int,0,len(item))
//	Merge = append(Merge,item...)
//
//	t := time.Now()
//	BubbleSort(Bubble)
//	fmt.Println("BubbleSort============================",time.Now().Sub(t))
//
//	t =time.Now()
//	BubbleSortUsingSortPackage(bubbleUsingSortPackage)
//	fmt.Println("BubbleSortUsingSortPackage============",time.Now().Sub(t))
//
//	t =time.Now()
//	InsertionSort(Insertion)
//	fmt.Println("InsertionSort=========================",time.Now().Sub(t))
//
//	t =time.Now()
//	InsertionSortUsingSortPackage(insertionUsingSortPackage)
//	fmt.Println("InsertionSortUsingSortPackage=========",time.Now().Sub(t))
//
//	t =time.Now()
//	SelectionSort(Selection)
//	fmt.Println("SelectionSort=========================",time.Now().Sub(t))
//
//	t =time.Now()
//	SelectionSortUsingSortPackage(selectionUsingSortPackage)
//	fmt.Println("SelectionSortUsingSortPackage==========",time.Now().Sub(t))
//
//	t =time.Now()
//	QuickSort(Quick,0,len(Quick)-1)
//	fmt.Println("QuickSort==============================",time.Now().Sub(t))
//
//	t =time.Now()
//	Shellshort(Shell)
//	fmt.Println("Shellshort=============================",time.Now().Sub(t))
//
//	t =time.Now()
//	Combsort(Comb)
//	fmt.Println("Combsort===============================",time.Now().Sub(t))
//
//	t =time.Now()
//	MergeSort(Merge)
//	fmt.Println("MergeSort==============================",time.Now().Sub(t))
//
//
//}
//// 10000条数据 排序时间137.0905ms
//func BubbleSort(items []int) {
//	var (
//		n       = len(items)
//		swapped = true
//	)
//	for swapped {
//		swapped = false
//		for i := 0; i < n-1; i++ {
//			if items[i] > items[i+1] {
//				items[i+1], items[i] = items[i], items[i+1]
//				swapped = true
//			}
//		}
//		n = n - 1
//	}
//}
//// 10000条数据 排序时间472.3138ms
//func BubbleSortUsingSortPackage(data sort.Interface) {
//	r := data.Len() - 1
//	for i := 0; i < r; i++ {
//		for j := r; j > i; j-- {
//			if data.Less(j, j-1) {
//				data.Swap(j, j-1)
//			}
//		}
//	}
//}
//// 10000条数据 排序时间65.0431ms
//func InsertionSort(items []int) {
//	var n = len(items)
//	for i := 1; i < n; i++ {
//		j := i
//		for j > 0 {
//			if items[j-1] > items[j] {
//				items[j-1], items[j] = items[j], items[j-1]
//			}
//			j = j - 1
//		}
//	}
//}
//// 10000条数据 排序时间211.1428ms
//func InsertionSortUsingSortPackage(data sort.Interface) {
//	r := data.Len() - 1
//	for i := 1; i <= r; i++ {
//		for j := i; j > 0 && data.Less(j, j-1); j-- {
//			data.Swap(j, j-1)
//		}
//	}
//}
//// 10000条数据 排序时间73.0472ms
//func SelectionSort(items []int) {
//	var n = len(items)
//	for i := 0; i < n; i++ {
//		var minIdx = i
//		for j := i; j < n; j++ {
//			if items[j] < items[minIdx] {
//				minIdx = j
//			}
//		}
//		items[i], items[minIdx] = items[minIdx], items[i]
//	}
//}
//// 10000条数据 排序时间261.1741ms
//func SelectionSortUsingSortPackage(data sort.Interface) {
//	r := data.Len() - 1
//	for i := 0; i < r; i++ {
//		min := i
//		for j := i + 1; j <= r; j++ {
//			if data.Less(j, min) {
//				min = j
//			}
//		}
//		data.Swap(i, min)
//	}
//}
//// 10000条数据 排序时间1.0005ms
//func QuickSort(src []int, first, last int) {
//	flag := first
//	left := first
//	right := last
//
//	if first >= last {
//		return
//	}
//
//	for first < last {
//		for first < last {
//			if src[last] >= src[flag] {
//				last -= 1
//				continue
//			} else {
//				tmp := src[last]
//				src[last] = src[flag]
//				src[flag] = tmp
//				flag = last
//				break
//			}
//		}
//
//		for first < last {
//			if src[first] <= src[flag] {
//				first += 1
//				continue
//			} else {
//				tmp := src[first]
//				src[first] = src[flag]
//				src[flag] = tmp
//				flag = first
//				break
//			}
//		}
//	}
//	QuickSort(src, left, flag-1)
//	QuickSort(src, flag+1, right)
//}
//// 10000条数据 排序时间72.0479ms
//func Shellshort(items []int) {
//	var (
//		n    = len(items)
//		gaps = []int{1}
//		k    = 1
//	)
//
//	pow := func(a, b int) int {
//		p := 1
//		for b > 0 {
//			if b&1 != 0 {
//				p *= a
//			}
//			b >>= 1
//			a *= a
//		}
//		return p
//	}
//
//	for {
//		gap := pow(2, k) + 1
//		if gap > n-1 {
//			break
//		}
//		gaps = append([]int{gap}, gaps...)
//		k++
//	}
//
//	for _, gap := range gaps {
//		for i := gap; i < n; i += gap {
//			j := i
//			for j > 0 {
//				if items[j-gap] > items[j] {
//					items[j-gap], items[j] = items[j], items[j-gap]
//				}
//				j = j - gap
//			}
//		}
//	}
//}
//
//func Combsort(items []int) {
//	var (
//		n = len(items)
//		gap = len(items)
//		shrink = 1.3
//		swapped = true
//	)
//
//	for swapped {
//		swapped = false
//		gap = int(float64(gap) / shrink)
//		if gap < 1 {
//			gap = 1
//		}
//		for i := 0; i+gap < n; i++ {
//			if items[i] > items[i+gap] {
//				items[i+gap], items[i] = items[i], items[i+gap]
//				swapped = true
//			}
//		}
//	}
//}
//
//func mergeSort(items []int) []int {
//	var n = len(items)
//
//	if n == 1 {
//		return items
//	}
//
//	middle := int(n / 2)
//	var (
//		left  = make([]int, middle)
//		right = make([]int, n-middle)
//	)
//	for i := 0; i < n; i++ {
//		if i < middle {
//			left[i] = items[i]
//		} else {
//			right[i-middle] = items[i]
//		}
//	}
//
//	return merge(mergeSort(left), mergeSort(right))
//}
//
//func MergeSort(items []int) []int {
//	var n = len(items)
//
//	if n == 1 {
//		return items
//	}
//
//	middle := int(n / 2)
//	var (
//		left  = make([]int, middle)
//		right = make([]int, n-middle)
//	)
//	for i := 0; i < n; i++ {
//		if i < middle {
//			left[i] = items[i]
//		} else {
//			right[i-middle] = items[i]
//		}
//	}
//
//	return merge(mergeSort(left), mergeSort(right))
//}
//
//func merge(left, right []int) (result []int) {
//	result = make([]int, len(left)+len(right))
//
//	i := 0
//	for len(left) > 0 && len(right) > 0 {
//		if left[0] < right[0] {
//			result[i] = left[0]
//			left = left[1:]
//		} else {
//			result[i] = right[0]
//			right = right[1:]
//		}
//		i++
//	}
//
//	// Either left or right may have elements left; consume them.
//	// (Only one of the following loops will actually be entered.)
//	for j := 0; j < len(left); j++ {
//		result[i] = left[j]
//		i++
//	}
//	for j := 0; j < len(right); j++ {
//		result[i] = right[j]
//		i++
//	}
//
//	return
//}