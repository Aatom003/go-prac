package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newbill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

func (b bill) format() string {
	fs := "Bill Breakdown:\n"
	var total float64 = 0

	for k, v := range b.items {
		fs += fmt.Sprintf("%-15v...$%0.2v\n", k+":", v)
		total += v
	}
	fs += fmt.Sprintf("%-15v...$%0.2f\n", "Tip:", b.tip)
	fs += fmt.Sprintf("%-15v...$%0.2f", "Total:", total+b.tip)
	return fs
}

func (b *bill) updateBill(tip float64) {
	b.tip = tip
}

func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}
