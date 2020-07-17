package main
import(
	"fmt"
)
func test1(){
	var sum int=0
	for i:=0; i<100; i++ {
		sum+=i
		
	}
	fmt.Println("100 以内加法")
	fmt.Println(sum)
}
func test2(j int){
	var sum int=0
	for i:=0; i<j; i++ {
		sum+=i
	}
	fmt.Println("小于数的和")
	fmt.Println(sum)
}


func test3(j int){
	for i:=0; i<j; i++{
		for a:=0; a< j-i; a++{
			fmt.Print("-")
		}
		for a:=0; a< i; a++{
			fmt.Print("*")
		}
		for a:=0; a< j-i; a++{
			fmt.Print("-")
		}

		fmt.Println()
	}


}


func test4(i int, j int)(sum int, sub int){
	sum=i+j
	sub=i-j
	return
}

type Student struct{
	Name string
	age int
	addr string
}

func (s Student) test(){
	fmt.Println(s.Name)

}

func (s *Student) say(){
	fmt.Println("姓名=",s.Name,",年龄=",s.age,",地址=",s.addr)
}


func main()  {
	
	//test1()
	//test2(2)
	// test3(9)
	//a, b:=test4(5,2)
	//fmt.Println("a的值为：",a,"，b的值为：",b)
	var stu Student
	stu.Name="张三"
	stu.age=10
	stu.addr="上海"
	stu.test();
	stu.say()
	fmt.Println("")
	
}