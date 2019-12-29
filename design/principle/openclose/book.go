package openclose

import "fmt"

type IBook interface {
	getName() string
	getPrice() int
	getAuthor() string
}

type Book struct {
}

type NovelBook struct {
	name   string
	price  int
	author string
}

func NewNovelBook(name, author string, price int) *NovelBook {
	return &NovelBook{
		name:   name,
		price:  price,
		author: author,
	}
}

func (b NovelBook) getName() string {
	return b.name
}
func (b NovelBook) getAuthor() string {
	return b.author
}
func (b NovelBook) getPrice() int {
	return b.price
}

type BookStore struct {
	bookList []IBook
}

var defaultBook = []IBook{
	&NovelBook{name: "天", price: 3200, author: "金"},
	&NovelBook{name: "巴", price: 5600, author: "雨"},
}

func PrintBook() {
	bs := &BookStore{bookList: defaultBook}
	for _, v := range bs.bookList {
		fmt.Print("name: ", v.getName())
		fmt.Print(" price: ", v.getPrice())
		fmt.Print(" author: ", v.getAuthor())
		fmt.Println()
	}
}

// 修改需求，大于40的9折，其他8折
// 修改接口，添加一个打折接口，修改修改多处地方，否定。
// 修改实现类，直接修改getPrice方法，无法通用，否定。
// 通过扩展实现，添加子类 offNovelBook，单独实现 getPrice

type OffNovelBook struct {
	name   string
	price  int
	author string
}

func NewOffNovelBook(name, author string, price int) *OffNovelBook {
	return &OffNovelBook{
		name:   name,
		price:  price,
		author: author,
	}
}

func (b OffNovelBook) getName() string {
	return b.name
}
func (b OffNovelBook) getAuthor() string {
	return b.author
}
func (b OffNovelBook) getPrice() int {
	if b.price > 4000 {
		return b.price * 90 / 100
	}
	return b.price * 80 / 100
}

var defaultOffBook = []IBook{
	NewOffNovelBook("天", "金", 3200),
	NewOffNovelBook("巴", "雨", 5600),
}

func PrintOffBook() {
	bs := &BookStore{bookList: defaultOffBook}
	for _, v := range bs.bookList {
		fmt.Print("name: ", v.getName())
		fmt.Print(" price: ", v.getPrice())
		fmt.Print(" author: ", v.getAuthor())
		fmt.Println()
	}
}

// 添加需求
// 增加计算机类图书，计算机类包含属性，“范围”

type IComputerBook interface {
	IBook
	getScope() string
}

type ComputerBook struct {
	name   string
	price  int
	author string
	scope  string
}

func NewComputerBook(name, author, scope string, price int) *ComputerBook {
	return &ComputerBook{
		name:   name,
		price:  price,
		author: author,
		scope:  scope,
	}
}

func (b ComputerBook) getName() string {
	return b.name
}
func (b ComputerBook) getAuthor() string {
	return b.author
}
func (b ComputerBook) getPrice() int {
	return b.price
}
func (b ComputerBook) getScope() string {
	return b.scope
}

var defaultMixBook = []IBook{
	NewNovelBook("天", "金", 3200),
	NewNovelBook("巴", "雨", 5600),
	NewComputerBook("Go", "google", "编程", 4300),
}

func PrintMixBook() {
	bs := &BookStore{bookList: defaultMixBook}
	for _, v := range bs.bookList {
		fmt.Print("name: ", v.getName())
		fmt.Print(" price: ", v.getPrice())
		fmt.Print(" author: ", v.getAuthor())
		if _, ok := v.(*ComputerBook); ok {
			fmt.Print(" scope: ", v.(*ComputerBook).getScope())
		} else {
			fmt.Println("no")
		}
		fmt.Println()
	}
}
