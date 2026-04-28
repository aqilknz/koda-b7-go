package minitask7

import "fmt"

type Person struct {
	Name    string
	Address string
	Phone   string
}

// Constructor function
func NewPerson(name, address, phone string) *Person {
	return &Person{
		Name:    name,
		Address: address,
		Phone:   phone,
	}
}

func (p *Person) Print() {
	fmt.Printf("Name: %s\nAddress: %s\nPhone: %s\n", p.Name, p.Address, p.Phone)
}
func (p *Person) Greet() string {
	return fmt.Sprintf("Hallo, %s!!", p.Name)
}
func (p *Person) ChangeName(name string) {
	p.Name = name
}
