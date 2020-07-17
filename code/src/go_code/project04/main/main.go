package main
import (
	"fmt"
	"math/rand"
	"sort"
)
type student struct{
	name string
	age int
}

type stuents []student

// Len 为集合内元素的总数
func (s stuents) Len() int{
	return len(s)
}

// Less 返回索引为 i 的元素是否应排在索引为 j 的元素之前。
func (s stuents) Less(i, j int) bool{
	if s[i].age>s[j].age {
		return true
	}else{
		return false
	}
}
  
// Swap 交换索引为 i 和 j 的元素
func (s stuents) Swap(i, j int){
	//快速交换
	s[i],s[j]=s[j],s[i]
}

func main(){

	var stus stuents
	for i:=0; i<10; i++{
		var stu student
		stu.name=fmt.Sprintf("stu%d",i)
		stu.age=rand.Intn(100)
		stus=append(stus,stu)
	}

	for _,v := range stus{
		fmt.Println(v)
	}
	fmt.Println("*********************")
	sort.Sort(stus)

	for _,v := range stus{
		fmt.Println(v)
	}

}



