//实现生产者消费者模型
package other

import (
	"sync"
	"fmt"
	"time"
)

type product struct {
	value int
	name string
}

func producer (wg *sync.WaitGroup,products chan<- product ,name int,stop *bool){
	for !*stop {
		temp:=new(product)
		temp.value=name
		temp.name="hello"

		products<-*temp
		fmt.Println("生成了",temp)

	}
	wg.Done()
}


func consumer(wg *sync.WaitGroup,products <-chan product,name int){
	for product:=range products{
		fmt.Println("消费了",product)
	}
	wg.Done()
}

func testProduceAndConsumer(){
	var wgp sync.WaitGroup
	var wgc sync.WaitGroup
	products:=make(chan product,10)
	stop:=false

	for i:=0;i<5;i++{
		go producer(&wgp,products,i,&stop)
		go consumer(&wgc,products,i)
	}

	time.Sleep(5*time.Second)
	stop=true

	wgp.Wait()
	close(products)
	wgc.Wait()

}