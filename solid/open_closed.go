package main

import "fmt"

type Color int

const (
	red Color = iota
	green
	blue
)

type Size int

const (
	small Size = iota
	medium
	large
)

type Product struct {
	name  string
	size  Size
	color Color
}

type Filter struct{}

func (f Filter) FilterByColor(products []Product, color Color) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.color == color {
			result = append(result, &products[i])
		}
	}

	return result
}

func (f Filter) FilterBySize(products []Product, size Size) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.size == size {
			result = append(result, &products[i])
		}
	}

	return result
}

type Specification interface {
	IsSatisfied(p *Product) bool
}

type ColorSpecification struct {
	color Color
}

func (c ColorSpecification) IsSatisfied(p *Product) bool {
	return c.color == p.color
}

type SizeSpecification struct {
	size Size
}

func (s SizeSpecification) IsSatisfied(p *Product) bool {
	return s.size == p.size
}

type BetterFilter struct{}

func (b BetterFilter) Filter(products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if spec.IsSatisfied(&v) {
			result = append(result, &products[i])
		}
	}

	return result
}

type AndSpecification struct {
	first, second Specification
}

func (a AndSpecification) IsSatisfied(p *Product) bool {
	return a.first.IsSatisfied(p) && a.second.IsSatisfied(p)
}

func main() {
	tree := Product{"Tree", large, green}
	apple := Product{"Apple", small, green}
	house := Product{"House", medium, red}

	products := []Product{apple, tree, house}

	fmt.Println("Green products (old):")
	f := Filter{}
	greenProductsOld := f.FilterByColor(products, green)

	for i, v := range greenProductsOld {
		fmt.Printf("%d. %s\n", i+1, v.name)
	}

	fmt.Println("Green proucts (new):")
	bf := BetterFilter{}
	greenProductsNew := bf.Filter(products, ColorSpecification{green})

	for i, v := range greenProductsNew {
		fmt.Printf("%d. %s\n", i+1, v.name)
	}

	fmt.Println("Composite specification")
	largeGreenProducts := bf.Filter(products, AndSpecification{
		ColorSpecification{green},
		SizeSpecification{large},
	})

	for i, v := range largeGreenProducts {
		fmt.Printf("%d. %s\n", i+1, v.name)
	}

}
