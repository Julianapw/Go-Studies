package main

import (
	"fmt"
	"time"
)

// Exercise 1
func modify(arr [3]int) {
	arr[0] = 100
}

//Why does the original array not change?
//Because arrays are copied when passed.
//This is different from Python lists, which are references.

// Exercise 2
func modifySlice(nums []int) {
	nums[0] = 100
}

//Why does this change affect the original slice?
//Slices contain: Pointer to array, Length,Capacity
//When passed: The slice descriptor is copied, But both descriptors point to the same underlying array, So modifying nums[0] modifies shared memory.

// Exercise 3
func createSlice() []int {
	nums := make([]int, 3, 5)
	for i := 0; i < 5; i++ {
		nums = append(nums, i)
		fmt.Println("Len:", len(nums), "Cap:", cap(nums))
	}
	return nums
}

func NilEmptySlice() {
	var a []int
	b := []int{}
	fmt.Println(a == nil)
	fmt.Println(b == nil)
	fmt.Println(len(a))
	fmt.Println(len(b))
}

func cloneSlice(nums []int) []int {
	copySlice := make([]int, len(nums))
	copy(copySlice, nums)
	return copySlice
}

//Exercise 6 and 7

//var m map[string]int

func CreateMap() map[string]int {
	m := make(map[string]int)
	m["a"] = 1

	value, ok := m["key"]
	if ok {
		fmt.Println("Found:", value)
	} else {
		fmt.Println("Key not found")
	}
	return m
}

// Exercise 8
func updateMap(m map[string]int) {
	m["a"] = 105
}

// Exercise 9
func SliceSharingPitfall() {
	a := []int{1, 2, 3, 4}
	b := a[:2]
	b[0] = 50
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	//Copy
	b = make([]int, 2)
	copy(b, a[:2])
	fmt.Println("a:", a)
	fmt.Println("b:", b)
}

// Exercise 10
func removeElement(nums []int, index int) []int {
	if index < 0 || index >= len(nums) {
		return nums
	}
	return append(nums[:index], nums[index+1:]...)
}

// Exercise 11
type Counter struct {
	Count int
}

func (c Counter) Increment() {
	c.Count++
}
func (c *Counter) IncrementPointer() {
	c.Count++
}

//Value receiver copies struct.
//Pointer receiver modifies original.

// Exercise 12
type User struct {
	Name string
}

func update(u *User) {
	u.Name = "Juliana"
}

// Exercise 13
func PointerPanic() {
	var u *User
	//u.Name = "Juliana" // This will panic because u is nil
	update(u) // This will also panic because update tries to dereference a nil pointer
	if u != nil {
		fmt.Println("User name:", u.Name)
	}
}

// Exercise 14
type Admin struct {
	Name  string
	Email string
}

// Leitura(value)
func (a Admin) Display() string {
	return fmt.Sprintf("Admin{Name: %s, Email: %s}", a.Name, a.Name)

}

// Mutação(pointer)
func (a *Admin) SetEmail(email string) {
	a.Email = email
}

func (a *Admin) Rename(name string) {
	a.Name = name
}

//Na versão anterior SetEmailValue não era um ponteiro
// SetEmailValue não altera o Email de admin (receiver por valor → muda só a cópia).
//SetEmailPointer altera o Email de admin (receiver por ponteiro → muda o original).

// Exercise 15
type BigUser struct {
	ID          int64
	Name        string
	Email       string
	Phone       string
	Adress      string
	City        string
	State       string
	Country     string
	Zip         string
	Company     string
	Departament string
	Title       string
	Active      bool
	LoginCount  int
	LastLoginIP string
}

func processByValue(u BigUser) {
	fmt.Printf("[by value] param adress: %p/n", &u)
	_ = u.Name
}

func processByPointer(u *BigUser) {
	fmt.Printf("[by value] param adress: %p/n", &u)
	_ = u.Name
}

// Exercise 16
func add(nums []int) []int {
	return append(nums, 100)
}

// Exercise 17
type User2 struct {
	Name string
}

func (u User2) AddUser(name string) {
	m := map[string]User2{
		"u1": {Name: "Old"},
	}

	v := m["u1"]
	v.Name = "New"
	m["u1"] = v

	fmt.Println(m)
}

// Exercise 18
func InterfaceNilTrap() {
	var u *User = nil
	var i interface{} = u

	fmt.Println(i == nil)
}

//Interface stores:
//Type
//Value
//Type is *User, value is nil.
//Therefore interface is not nil.

//Exercise 19

func RunConcurrentMapDemo() {
	m := make(map[int]int)

	// Goroutine de escrita
	go func() {
		for i := 0; i < 1_000_000; i++ {
			m[i%10] = i // escreve repetidamente nas mesmas chaves
			// opcional: time.Sleep(time.Microsecond)
		}
	}()

	// Goroutine de leitura
	go func() {
		for {
			_ = m[0] // lê continuamente
			// opcional: time.Sleep(time.Microsecond)
		}
	}()

	// Deixe rodar um pouco para as goroutines disputarem o map
	time.Sleep(2 * time.Second)
	fmt.Println("done (observe o comportamento acima)")
}

//Maps are not thread-safe.
//Must use:
//sync.Mutex
//sync.Map
//channels

//Exercise 20
func CreateUser() *User {
	u := User{Name: "Test"}
	return &u
}

//This is safe.
//Go performs escape analysis.
//Since pointer escapes function scope:
//It allocates on heap
//Not stack
//Memory is safe.