package other

import "fmt"

// 一个数组 从左到右为升序 从上到下是降序 查找对应的值
func findValue( array [][]int,value int){


	defer func() {
		if r:=recover();r!=nil{
			fmt.Println("错误",r)
		}
	}()
	row:=len(array)
	col:=len(array[0])

	r,c:=row-1,col-1
	for r<row && c<col && r>=0 && c>=0{
		if array[r][c]>value{
			c--
		} else if array[r][c]<value{
			r--
		}else{
			fmt.Println("找到了",r+1,"行",c+1,"列")
			return
		}
	}

}
