package main

import (
	"fmt"
)

type sku struct {
	name  string
	price int
}

var (
	cart       []sku
	cart_total int
)

func add_item_to_cart(name string, price int) {
	cart = add_item(cart, name, price)
	calc_cart_total()
}

// pure function ("calculation")
func add_item(cart []sku, name string, price int) []sku {
	item := sku{name, price}
	cart = dup_sku_slice(cart) // enforce COW: don't mutate inputs
	cart = append(cart, item)
	return cart
}

// pure function ("calculation")
func dup_sku_slice(src []sku) []sku {
	if src == nil {
		return nil
	}
	dst := make([]sku, len(src))
	_ = copy(dst, src)
	return dst
}

func calc_cart_total() {
	cart_total = calc_total(cart)
	update_shipping_icons()
	update_tax_dom()
}

// pure function ("calculation")
func calc_total(cart []sku) int {
	total := 0
	for _, sku := range cart {
		total += sku.price
	}
	return total
}

type buy_button struct {
	item sku
}

func show_free_shipping_icon(button buy_button) {
}

func hide_free_shipping_icon(button buy_button) {
}

func get_buy_buttons_dom() []buy_button {
	var buy_buttons_dom = []buy_button{
		{sku{"train", 11}},
		{sku{"tank", 31}},
	}
	return buy_buttons_dom
}

func update_shipping_icons() {
	buttons := get_buy_buttons_dom()
	for _, button := range buttons {
		sku := button.item
		if gets_free_shipping(cart_total, sku) {
			show_free_shipping_icon(button)
		} else {
			hide_free_shipping_icon(button)
		}
	}
}

// pure function ("calculation")
func gets_free_shipping(cart_total int, item sku) bool {
	return cart_total+item.price >= 20
}

func update_tax_dom() {
	set_tax_dom(calculate_tax(cart_total))
}

func set_tax_dom(tax float64) {
}

// pure function ("calculation")
func calculate_tax(price int) float64 {
	return float64(price) * 0.10
}

func main() {
	add_item_to_cart("doll", 12)
	add_item_to_cart("ball", 5)
	fmt.Println("cart:", cart)
	fmt.Println("total:", cart_total)
}
