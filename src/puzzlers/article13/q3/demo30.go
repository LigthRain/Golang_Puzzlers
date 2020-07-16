package main

import "fmt"

type Cat struct {
	name           string // 名字。
	scientificName string // 学名。
	category       string // 动物学基本分类。
	test           *int
}

func New(name, scientificName, category string) Cat {
	a := 1
	return Cat{
		name:           name,
		scientificName: scientificName,
		category:       category,
		test:           &a,
	}
}

func (cat *Cat) SetName(name string) {
	cat.name = name
}

func (cat Cat) SetNameOfCopy(name string) {
	cat.name = name
}

func (cat Cat) setTest(b int)  {
	cat.test = &b
}

func (cat Cat) Name() string {
	return cat.name
}

func (cat Cat) ScientificName() string {
	return cat.scientificName
}

func (cat Cat) Category() string {
	return cat.category
}

func (cat Cat) String() string {
	return fmt.Sprintf("%s (category: %s, name: %q) test:%d",
		cat.scientificName, cat.category, cat.name, *cat.test)
}

func main() {
	cat := New("little pig", "American Shorthair", "cat")
	cat.setTest(333)    //test依旧是1 除非是channel slice map 值类型的所有set方法 都不会改变原有值

	//cat.SetName("monster") // (&cat).SetName("monster")
	fmt.Printf("The cat: %s\n", cat)
	//
	//cat.SetNameOfCopy("little pig")
	//fmt.Printf("The cat: %s\n", cat)
	//
	//type Pet interface {
	//	SetName(name string)
	//	Name() string
	//	Category() string
	//	ScientificName() string
	//}
	//
	//_, ok := interface{}(cat).(Pet)
	//fmt.Printf("Cat implements interface Pet: %v\n", ok)
	//_, ok = interface{}(&cat).(Pet)
	//fmt.Printf("*Cat implements interface Pet: %v\n", ok)
}
