/*
 * @lc app=leetcode.cn id=1834 lang=golang
 *
 * [1834] 单线程 CPU
 */
package leetcode

import (
	"container/heap"
	"sort"
)

// @lc code=start

type Task struct {
	index, enqueueTime, processingTime int
}

type TaskHeap struct {
	Tasks []Task
}

func (a TaskHeap) Len() int { return len(a.Tasks) }
func (a TaskHeap) Swap(i, j int) {
	a.Tasks[i], a.Tasks[j] = a.Tasks[j], a.Tasks[i]
}
func (a TaskHeap) Less(i, j int) bool {
	if a.Tasks[i].processingTime == a.Tasks[j].processingTime {
		return a.Tasks[i].index < a.Tasks[j].index
	}

	return a.Tasks[i].processingTime < a.Tasks[j].processingTime
}

func (a *TaskHeap) Push(x interface{}) {
	a.Tasks = append(a.Tasks, x.(Task))
}

func (a *TaskHeap) Pop() (x interface{}) {
	n := len(a.Tasks)
	x = a.Tasks[n-1]
	a.Tasks = a.Tasks[:n-1]
	return
}

func getOrder(tasks [][]int) []int {
	taskS := []Task{}
	for i, v := range tasks {
		task := Task{
			index:          i,
			enqueueTime:    v[0],
			processingTime: v[1],
		}
		taskS = append(taskS, task)
	}

	sort.Slice(taskS, func(i, j int) bool {
		return taskS[i].enqueueTime < taskS[j].enqueueTime
	})

	now := 0
	i := 0
	n := len(tasks)
	taskH := TaskHeap{}
	ans := []int{}
	for len(ans) < n {
		if taskH.Len() > 0 {
			x := heap.Pop(&taskH).(Task)
			ans = append(ans, x.index)
			now += x.processingTime
		} else {
			now = taskS[i].enqueueTime
		}

		for i < n && taskS[i].enqueueTime <= now {
			heap.Push(&taskH, taskS[i])
			i++
		}
	}
	return ans
}

// @lc code=end
