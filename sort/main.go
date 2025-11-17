package main

import (
	"fmt"
	"sort"
	"time"
)

type Product struct {
	ID        int
	Name      string
	Price     float64
	Rating    float64
	Sales     int
	CreatedAt time.Time
}
type ProductManager struct {
	products []Product
}

func NewProductManager() *ProductManager {
	return &ProductManager{
		products: []Product{
			{
				ID:        1,
				Name:      "手机",
				Price:     1000.0,
				Rating:    4.5,
				Sales:     100,
				CreatedAt: time.Now(),
			},
			{
				ID:        2,
				Name:      "电脑",
				Price:     2000.0,
				Rating:    4.0,
				Sales:     50,
				CreatedAt: time.Now(),
			},
			{
				ID:        3,
				Name:      "平板",
				Price:     1500.0,
				Rating:    4.2,
				Sales:     80,
				CreatedAt: time.Now(),
			},
		},
	}
}

func (pm *ProductManager) SortByPrice(ascending bool) {
	if ascending {
		sort.Slice(pm.products, func(i, j int) bool {
			return pm.products[i].Rating > pm.products[j].Rating
		})
	} else {
		sort.Slice(pm.products, func(i, j int) bool {
			return pm.products[i].Rating < pm.products[j].Rating
		})
	}

}

/*
type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}
*/

type Person struct {
	Age  int
	Name string
}
type People []Person

func (p People) Len() int {
	return len(p)
}
func (p People) Less(i, j int) bool {
	return p[i].Age < p[j].Age
}
func (p People) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (pm *ProductManager) PrintProducts() {
	for _, product := range pm.products {
		fmt.Println(product)
	}
}
func main() {
	pm := NewProductManager()
	pm.SortByPrice(false)
	pm.PrintProducts()

	people := People{
		{Age: 25, Name: "张三"},
		{Age: 30, Name: "李四"},
		{Age: 20, Name: "王五"},
	}
	sort.Sort(people)
	fmt.Println(people)
}
