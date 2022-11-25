package main

import "fmt"

type Person struct {
	name       string
	job        string
	compoany   string
	streetName string
	city       string
	postcode   string
}

type PersonBuilder struct {
	p *Person
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{&Person{}}
}

func (b *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*b}
}

func (b *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*b}
}

func (b *PersonBuilder) Build() *Person {
	return b.p
}

type PersonJobBuilder struct {
	PersonBuilder
}

type PersonAddressBuilder struct {
	PersonBuilder
}

func (b *PersonJobBuilder) At(companyName string) *PersonJobBuilder {
	b.p.compoany = companyName
	return b
}

func (b *PersonJobBuilder) AsA(title string) *PersonJobBuilder {
	b.p.job = title
	return b
}

func (b *PersonAddressBuilder) At(street string) *PersonAddressBuilder {
	b.p.streetName = street
	return b
}

func (b *PersonAddressBuilder) In(city string) *PersonAddressBuilder {
	b.p.city = city
	return b
}

func (b *PersonAddressBuilder) WithPostalCode(postcode string) *PersonAddressBuilder {
	b.p.postcode = postcode
	return b
}

func main() {
	pb := NewPersonBuilder()

	pb.
		Lives().
		At("123 London Road").
		In("London").
		WithPostalCode("123456").
		Works().
		At("Fabrikam").
		AsA("Programmer")

	person := pb.Build()
	fmt.Println(person)
}
