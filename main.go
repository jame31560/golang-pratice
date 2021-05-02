package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/kataras/iris/v12"
)

type Item interface {
	show()
	getId() string
}

type OrderItem struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func (o OrderItem) getId() string {
	return o.Id
}

type Drink struct {
	OrderItem
	Sugar float32 `json:"sugar"`
	Ice   float32 `json:"ice"`
}

func (d Drink) show() {
	fmt.Printf("---\nId: %s\nName: %s , Price: %d (sugar: %.1f, ice: %.1f)\n", d.Id, d.Name, d.Price, d.Sugar, d.Ice)
}

type Food struct {
	OrderItem
}

func (f Food) show() {
	fmt.Printf("---\nId: %s\nName: %s , Price: %d\n", f.Id, f.Name, f.Price)
}

type Cart struct {
	Items []Item `json:"items"`
}

func (c *Cart) addItem(item Item) {
	fmt.Println("-------------Add Item-------------")
	c.Items = append(c.Items, item)
}

func (c *Cart) findItem(id string) Item {
	fmt.Println("-------------Find Item-------------")
	for _, item := range c.Items {
		if item.getId() == id {
			return item
		}
	}
	return nil
}

func (c *Cart) deleteItem(id string) {
	fmt.Println("-------------Delete Item-------------")
	for idx, item := range c.Items {
		if item.getId() == id {
			c.Items = append(c.Items[:idx], c.Items[idx+1:]...)
		}
	}
}

func (c *Cart) showItems() {
	fmt.Println("-------------Cart Items-------------")
	for _, item := range c.Items {
		item.show()
	}
}

func main() {
	// item1 := Drink{OrderItem{uuid.NewString(), "綠茶", 20}, 0.3, 1}
	// item2 := Food{OrderItem{uuid.NewString(), "鍋燒意麵", 50}}

	// items := []Item{}
	// cart := Cart{items}

	// cart.addItem(item1)
	// cart.showItems()
	// fmt.Println(cart.Items[0].getId())
	// cart.findItem(cart.Items[0].getId()).show()
	// cart.addItem(item2)

	// cart.showItems()
	// cart.deleteItem(cart.findItem(cart.Items[0].getId()).getId())
	// cart.showItems()

	app := iris.New()

	cartAPI := app.Party("/cart")
	{
		cartAPI.Use(iris.Compression)
		cartAPI.Get("/", getCartItems)
	}

	app.Listen(":8080")
}

func getCartItems(ctx iris.Context) {
	item1 := Drink{OrderItem{uuid.NewString(), "綠茶", 20}, 0.3, 1}
	item2 := Food{OrderItem{uuid.NewString(), "鍋燒意麵", 50}}

	items := []Item{}
	cart := Cart{items}

	cart.addItem(item1)
	cart.addItem(item2)
	item1.show()

	ctx.JSON(cart)
}
