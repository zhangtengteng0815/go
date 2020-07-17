package main
import(
	_"fmt"
	_"go_code/project03/package/student"
	"go_code/project03/package/animal"
)

func main(){
	// var stu=student.NewStudent("lisi",10,"北京")
	// fmt.Println(stu)
	// stu.SetName("zhangsan")
	// fmt.Println(stu)

	// var ani=animal.NewAnimal()
	// ani.SetName("CAT")
	// ani.SetAge(10)
	// fmt.Println(ani)

	// fmt.Println(ani.GetName())
	// fmt.Println(ani.GetAge())

	var cat=animal.NewCat()
	cat.SetName("tom")
	cat.SetAge(10)
	cat.Eat()


}