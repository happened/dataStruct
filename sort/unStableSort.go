//主要实现不稳定的排序算法 从小到大

package sort

import "fmt"

//选择排序
func SelectSort(array []int){
	for i:=0;i<len(array);i++{
		min:=i
		for j:=i+1;j<len(array);j++{
			if array[j]<array[min]{
				min=j
			}
		}
		if min!=i {
			array[min], array[i] = array[i], array[min]
		}
	}
}

//快速排序
func QuickSort(left int, right int, array []int) {
	if left >= right {
		return
	}
	pre := left
	last := right

	firstVal := array[left]

	for pre != last {
		//此循环是将基于基准的数移到基准两边
		for firstVal <= array[last] && pre < last {
			last--
		}

		for firstVal >= array[pre] && pre < last {
			pre++
		}

		//进行替换
		if pre < last {
			array[pre], array[last] = array[last], array[pre]
		}

	}

	//基准归位
	array[left] = array[pre]
	array[pre] = firstVal

	QuickSort(left, pre-1, array)
	QuickSort(pre+1, right, array)
}

//堆排序 小根堆
func adjustHeap(array []int,start int,end int){
	father:=start
	son :=father*2+1
	for son<=end{
		if son+1<=end && array[son]>array[son+1]{
			son++
		}
		if array[father]>array[son]{
			array[father],array[son]=array[son],array[father]
			father=son
			son=2*father+1
		}else {
			return
		}
	}
}

func HeapSort(array []int){
	for i:=len(array)/2;i>=0;i--{
		adjustHeap(array,i,len(array)-1)
	}


	for j:=len(array)-1;j>=0;j--{
		array[j],array[0]=array[0],array[j]
		fmt.Printf("%d ",array[j])
		adjustHeap(array,0 ,j-1 )
	}



}


