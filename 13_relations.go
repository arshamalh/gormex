package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/kylelemons/godebug/pretty"
)

// Order
type Order struct {
	gorm.Model
	Status     string
	OrderItems []OrderItem
}

// Order line item
type OrderItem struct {
	gorm.Model
	OrderID  uint
	ItemID   uint
	Item     Item
	Quantity int
}

// Product
type Item struct {
	gorm.Model
	ItemName string
	Amount   float32
}

var (
	items = []Item{
		{ItemName: "Go Mug", Amount: 12.49},
		{ItemName: "Go Keychain", Amount: 6.95},
		{ItemName: "Go Tshirt", Amount: 17.99},
	}
)

func Relation() {
	db, err := gorm.Open("sqlite3", "/tmp/gorm.db")
	db.LogMode(true)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&OrderItem{}, &Order{}, &Item{})

	// Create Items
	for index := range items {
		db.Create(&items[index])
	}
	order := Order{Status: "pending"}
	db.Create(&order)
	item1 := OrderItem{OrderID: order.ID, ItemID: items[0].ID, Quantity: 1}
	item2 := OrderItem{OrderID: order.ID, ItemID: items[1].ID, Quantity: 4}
	db.Create(&item1)
	db.Create(&item2)

	// Query with joins
	rows, err := db.Table("orders").Where("orders.id = ? and status = ?", order.ID, "pending").
		Joins("Join order_items on order_items.order_id = orders.id").
		Joins("Join items on items.id = order_items.id").
		Select("orders.id, orders.status, order_items.order_id, order_items.item_id, order_items.quantity" +
			", items.item_name, items.amount").Rows()
	if err != nil {
		log.Panic(err)
	}

	defer rows.Close()
	// Values to load into
	newOrder := &Order{}
	newOrder.OrderItems = make([]OrderItem, 0)

	for rows.Next() {
		orderItem := OrderItem{}
		item := Item{}
		err = rows.Scan(&newOrder.ID, &newOrder.Status, &orderItem.OrderID, &orderItem.ItemID, &orderItem.Quantity, &item.ItemName, &item.Amount)
		if err != nil {
			log.Panic(err)
		}
		orderItem.Item = item
		newOrder.OrderItems = append(newOrder.OrderItems, orderItem)
	}
	log.Print(pretty.Sprint(newOrder))
}
