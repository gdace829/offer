package offer

import (
	"sort"
)

func getLeastNumbers(arr []int, k int) []int {
	sort.Ints(arr)
	return arr[0:k]
}

// 快排
func getLeastNumbers1(arr []int, k int) []int {
	var quicksort func(nums []int, left, right int)

	quicksort = func(nums []int, left, right int) {
		//递归终止条件
		if left > right {
			return
		}
		i, j, pivot := left, right, nums[left]
		for i < j {
			//从左边寻找比标准小的
			for i < j && nums[j] >= pivot {
				j--
			}
			//从右边寻找比标准大的
			for i < j && nums[i] <= pivot {
				i++
			}
			//交换
			nums[i], nums[j] = nums[j], nums[i]

		}
		//将标准放入适合位置
		nums[i], nums[left] = nums[left], nums[i]
		//递归
		quicksort(nums, left, i-1)
		quicksort(nums, i+1, right)
	}
	quicksort(arr, 0, len(arr)-1)
	return arr[:k]

}

// 堆排序
func getLeastNumbers2(arr []int, k int) []int {
	//构造大顶堆
	var heapify func(nums []int, root, end int)

	heapify = func(nums []int, root, end int) {
		for {
			//得到左孩子
			child := root*2 + 1
			//不存在左孩子直接返回
			if child > end {
				return
			}
			//选取左右孩子最大值
			if child < end && nums[child] < nums[child+1] {
				child++
			}
			//要是本身就大就不交换直接return
			if nums[root] >= nums[child] {
				return
			}
			//交换
			nums[root], nums[child] = nums[child], nums[root]
			root = child
		}

	}
	end := len(arr) - 1
	//从第一个非叶子结点开始构造
	for i := len(arr) / 2; i >= 0; i-- {
		heapify(arr, i, end)
	}
	//依次把大顶交换到最后进行排序
	for i := len(arr) - 1; i >= 0; i-- {
		arr[0], arr[end] = arr[end], arr[0]
		end--
		heapify(arr, 0, end)
	}
	return arr[:k]
}
