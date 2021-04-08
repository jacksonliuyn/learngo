package main

import (
	"fmt"
	"github.com/ahmetb/go-linq/v3"
	"github.com/gookit/goutil/maputil"
	gu "github.com/henrylee2cn/goutil"
	"github.com/henrylee2cn/goutil/calendar"
	"time"
)

//{ faker: "faker"}

type Gen struct {
	Faker string `json:"faker" gorm:"column:faker"`
}

func NewGen(faker string) *Gen {
	return &Gen{Faker: faker}
}

func main() {

	strings := []Customer{{5, "faker", "faker", 5, "zizi"}, {5, "faker", "faker", 5, "ziza"}}
	println(len(strings))
	m := map[string]string{}
	m["faker"] = "af"
	fmt.Println(maputil.Keys(m))
	test := []interface{}{strings[0:1]}
	println(len(test))
	for _, k := range test {
		fmt.Printf("%T \n", k)
		fmt.Println(k)
	}
	linq.From(strings).SelectT(func(i Customer) string {
		return i.Name
	}).ForEachT(func(i string) {
		fmt.Println("  " + i)
	})
	distinct := gu.InterfacesDistinct(test)
	fmt.Println(distinct)
	now := calendar.NewSolarNow()
	add := now.AddDate(0, 0, -5)
	println(add.String())
	var ss string = ""
	ss = fmt.Sprint(ss, "faker")
	fmt.Println(ss)

}

type Customer struct {
	ID             int
	Name, LastName string
	Age            int
	TaxNumber      string
}

func stream() {

}
func All() {
	bookList := MakeBook()
	var filterBooks []Book
	linq.From(bookList).WhereT(func(s Book) bool {
		return s.Money > 50
	}).Take(100).ToSlice(&filterBooks) //Take 固定取几个，也可以不写
	fmt.Print(filterBooks)
}
func GroupBy() {
	bookList1 := MakeBook1()
	query := linq.From(bookList1)

	fmt.Println("输出对应书名和作者")
	//第一个func是分组的依据，第二个func是分组后返回的元素.
	query = query.GroupByT(func(book Book) string {
		return book.Author
	}, func(book Book) Book {
		return book
	})

	query.ForEachT(func(bookGroup linq.Group) {
		fmt.Println(fmt.Sprintf("作者是%v", bookGroup.Key))
		for _, item := range bookGroup.Group {

			fmt.Println(fmt.Sprintf("书名%v,作者%v", item.(Book).Name, item.(Book).Author))
			fmt.Println()
		}

	})
	fmt.Println()
	fmt.Println("输出对应书名")

	bookList1 = MakeBook1()
	query = linq.From(bookList1)
	var nameGroups []linq.Group
	query.GroupByT(
		func(book Book) string { return book.Author },
		func(book Book) string { return book.Name },
	).ToSlice(&nameGroups)

	for _, item := range nameGroups {
		fmt.Println(fmt.Sprintf("作者是%v, 类型是 %T", item.Key, item.Key))
		for _, row := range item.Group {
			fmt.Println(fmt.Sprintf("书名%v", row))
		}
	}
}

type Book struct {
	Name        string
	Author      string
	Money       float64
	WordsNum    int
	PublishTime time.Time
}

type Person struct {
	Name string //拥有者的名字
}

type Pet struct {
	Name      string //动物的名字
	OwnerName string //拥有者的名字
}

type Man struct {
	Name string //拥有者的名字
	Pets []Pets //宠物们
}

type Pets struct {
	Name string //动物的名字
}

func MakeInnerData() []Man {
	man := make([]Man, 0)
	man = append(man, Man{
		Name: "康康",
		Pets: []Pets{{Name: "康康的狗"}, {Name: "康康的猫"}},
	})
	man = append(man, Man{
		Name: "老施",
		Pets: []Pets{{Name: "老施的🐟"}, {Name: "老施的鸟"}},
	})
	man = append(man, Man{
		Name: "小明",
		Pets: []Pets{{Name: "小明的🐖"}, {Name: "小明的狗"}},
	})

	return man
}

func MakeJoinData() ([]Person, []Pet) {
	kangkang := Person{Name: "爱吃合合乐的康康"}
	laoshi := Person{Name: "老施"}
	xiaoming := Person{Name: "不要催-小明"}
	expect := Person{Name: "我是没有宠物的人"}

	dog := Pet{Name: "康康的狗", OwnerName: kangkang.Name}
	cat := Pet{Name: "康康的猫", OwnerName: kangkang.Name}
	fish := Pet{Name: "老施的🐟", OwnerName: laoshi.Name}
	pig := Pet{Name: "小明的🐖", OwnerName: xiaoming.Name}

	return []Person{kangkang, laoshi, xiaoming, expect}, []Pet{dog, cat, fish, pig}
}

func MakeBook() []Book {
	bookList := make([]Book, 0)
	bookList = append(bookList, Book{
		Name:        "Go语言",
		Author:      "Go",
		Money:       100,
		WordsNum:    1000,
		PublishTime: time.Date(2020, 1, 1, 10, 0, 0, 0, time.Local),
	})
	bookList = append(bookList, Book{
		Name:        "Effective Java",
		Author:      "Java",
		Money:       78,
		WordsNum:    9000,
		PublishTime: time.Date(2020, 2, 15, 10, 0, 0, 0, time.Local),
	})
	bookList = append(bookList, Book{
		Name:        "Java语言",
		Author:      "Java",
		Money:       50,
		WordsNum:    3000,
		PublishTime: time.Date(2020, 2, 1, 10, 0, 0, 0, time.Local),
	})
	bookList = append(bookList, Book{
		Name:        "Lua语言",
		Author:      "Lua",
		Money:       75,
		WordsNum:    45000,
		PublishTime: time.Date(2020, 1, 10, 10, 0, 0, 0, time.Local),
	})
	bookList = append(bookList, Book{
		Name:        "React语言",
		Author:      "React",
		Money:       99,
		WordsNum:    14500,
		PublishTime: time.Date(2020, 7, 1, 10, 0, 0, 0, time.Local),
	})
	bookList = append(bookList, Book{
		Name:        "Red语言",
		Author:      "Red",
		Money:       28,
		WordsNum:    880,
		PublishTime: time.Date(2019, 4, 1, 10, 0, 0, 0, time.Local),
	})
	bookList = append(bookList, Book{
		Name:        "JavaScript语言",
		Author:      "JavaScript",
		Money:       81,
		WordsNum:    3776,
		PublishTime: time.Date(2019, 5, 17, 10, 0, 0, 0, time.Local),
	})
	return bookList
}

func MakeBook1() []Book {
	bookList := make([]Book, 0)
	bookList = append(bookList, Book{
		Name:        "Go语言",
		Author:      "Go",
		Money:       100,
		WordsNum:    1000,
		PublishTime: time.Date(2020, 1, 1, 10, 0, 0, 0, time.Local),
	})
	bookList = append(bookList, Book{
		Name:        "Effective Java",
		Author:      "Java",
		Money:       78,
		WordsNum:    9000,
		PublishTime: time.Date(2020, 2, 15, 10, 0, 0, 0, time.Local),
	})
	bookList = append(bookList, Book{
		Name:        "Go语言(第二版)",
		Author:      "Go",
		Money:       50,
		WordsNum:    3000,
		PublishTime: time.Date(2020, 2, 1, 10, 0, 0, 0, time.Local),
	})
	return bookList
}
func MakeBook2() []Book {
	bookList := make([]Book, 0)
	bookList = append(bookList, Book{
		Name:        "Go语言",
		Author:      "Go",
		Money:       100,
		WordsNum:    1000,
		PublishTime: time.Date(2020, 1, 1, 10, 0, 0, 0, time.Local),
	})
	return bookList
}

func MakeBook3() []Book {
	bookList := make([]Book, 0)
	bookList = append(bookList, Book{
		Name:        "Go语言",
		Author:      "Go",
		Money:       100,
		WordsNum:    1000,
		PublishTime: time.Date(2020, 1, 1, 10, 0, 0, 0, time.Local),
	})
	bookList = append(bookList, Book{
		Name:        "Go语言(第三版)",
		Author:      "Go",
		Money:       50,
		WordsNum:    3000,
		PublishTime: time.Date(2020, 2, 1, 10, 0, 0, 0, time.Local),
	})
	return bookList
}
