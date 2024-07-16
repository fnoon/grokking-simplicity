package main

import (
	"fmt"
)

type sku struct {
	name  string
	price int
}

var cart []sku

var cart_total int

func add_item_to_cart(name string, price int) {
	item := sku{name, price}
	cart = append(cart, item)
	calc_cart_total()
}

func calc_cart_total() {
	total := 0
	for _, sku := range cart {
		total += sku.price
	}
	cart_total = total
	update_shipping_icons()
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
		if sku.price+cart_total >= 20 {
			show_free_shipping_icon(button)
		} else {
			hide_free_shipping_icon(button)
		}
	}
}

func main() {
	add_item_to_cart("doll", 12)
	add_item_to_cart("ball", 5)
	fmt.Println("cart:", cart)
	fmt.Println("total:", cart_total)
}
