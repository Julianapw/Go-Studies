package main

import "fmt"

func main() {
	// ---------- Exercise 1: Arrays são copiados ----------
	a := [3]int{1, 2, 3}
	fmt.Println("Array Before:", a)
	modify(a)
	fmt.Println("Array After: ", a) // não muda (cópia)

	// ---------- Exercise 2: Slices compartilham o array ----------
	slice := []int{1, 2, 3}
	fmt.Println("Slice Before:", slice)
	modifySlice(slice)
	fmt.Println("Slice After: ", slice) // muda (mesmo array por baixo)

	// ---------- Nil vs Empty slice ----------
	NilEmptySlice()

	// ---------- Clonar slice ----------
	cloned := cloneSlice(slice)
	fmt.Println("Cloned slice:", cloned)

	// ---------- Map: criar, ler e atualizar ----------
	m := CreateMap()
	updateMap(m)
	fmt.Println("Map after update:", m)

	// ---------- Slice sharing pitfall e cópia ----------
	SliceSharingPitfall()

	// ---------- Remover elemento de slice ----------
	nums := []int{10, 20, 30, 40, 50}
	nums = removeElement(nums, 2) // remove o "30"
	fmt.Println("After removeElement:", nums)

	// ---------- Value vs Pointer receivers ----------
	c := Counter{Count: 0}
	c.Increment() // value receiver: NÃO persiste
	fmt.Println("Counter after Increment (value):", c.Count)
	c.IncrementPointer() // pointer receiver: PERSISTE
	fmt.Println("Counter after IncrementPointer:", c.Count)

	// ---------- Pointer em função ----------
	u := &User{Name: "Original"}
	update(u)
	fmt.Println("User after update(*User):", u.Name)

	// ⚠ Evitar panic: PointerPanic() panica do jeito que está
	// PointerPanic()

	// ---------- Admin: leitura (valor) e mutações (ponteiro) ----------
	admin := Admin{Name: "Juliana", Email: "juli@example.com"}
	fmt.Println("[start]           ", admin.Display())
	admin.SetEmail("new@example.com")
	fmt.Println("[after SetEmail]  ", admin.Display())
	admin.Rename("Jú")
	fmt.Println("[after Rename]    ", admin.Display())

	// ---------- BigUser: endereços por valor x ponteiro ----------
	user := BigUser{
		ID:         42,
		Name:       "Juliana",
		Email:      "ju@example.com",
		Active:     true,
		LoginCount: 7,
	}
	fmt.Printf("[caller]     user address:  %p\n", &user)
	processByValue(user)    // endereço diferente (cópia)
	processByPointer(&user) // mesmo endereço (ponteiro)

	// ---------- append em slice ----------
	slice = add(slice) // add 100 ao final
	fmt.Println("Slice after add:", slice)

	RunConcurrentMapDemo()
}
