//实现稳定的排序算法 从小到大
package sort

//冒泡排序
func BubbleSort(array []int){
	for i:=0;i<len(array);i++{
		for j:=len(array)-1;j>i;j--{
			if array[j]<array[i]{
				array[j],array[i]=array[i],array[j]
			}
		}
	}
}

//插入排序
func InsertSort(array []int){
	j:=0
	for i:=1;i<len(array);i++{
		temp:=array[i]
		for j=i-1;j>=0 && temp<array[j];j--{
			array[j+1] = array[j]
		}
		array[j+1]=temp
	}

}

//归并排序
func MergeSort(array []int){

}

//计数排序
func CountSort(array []int ){

}

//桶排序
func BucketSort(array []int){

}

//基数排序
func BasenumSort(array []int){

}