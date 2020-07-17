package student


type student struct{
	name string
	age int
	addr string
}

func NewStudent(name string, age int, addr string) student{
	return student{name, age, addr}
}

func (stu *student) SetName(name string){
	stu.name=name
}

func (stu *student) GetName() string{
	return stu.name
}