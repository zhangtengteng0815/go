package animal

import(
	"fmt"
)

type animal struct{
	name string
	age int
} 

func NewCat() Cat{
	return Cat{}
}

func NewDog() Dog{
	return Dog{}
}

func (a *animal)SetName(name string){
	a.name=name
} 

func (a *animal)GetName() string{
	return a.name
}
func (a *animal)SetAge(age int){
	a.age=age
} 

func (a *animal)GetAge() int{
	return a.age
}

type Cat struct{
	animal
}

func (c *Cat)Eat(){
	fmt.Println(c.name,"吃鱼")
}

type Dog struct{
	animal
}

func (d *Cat)Dog(){
	fmt.Println(d.name,"吃肉")
}
