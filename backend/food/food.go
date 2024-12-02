package food

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"restaurant/database"
	"strconv"
	"strings"
)

type Food struct {
	ItemName string  `json:"item_name"`
	ItemID   int     `json:"item_id"`
	Status   string  `json:"status"`
	Price    float64 `json:"price"`
}

var foods []Food

func addFood(w http.ResponseWriter, r *http.Request) {
	item_name := bufio.NewReader(os.Stdin)
	fmt.Println("Food Name:")
	food_name, _ := item_name.ReadString('\n')
	food_name = strings.TrimSpace(food_name)

	item_id := bufio.NewReader(os.Stdin)
	fmt.Println("Food ID:")
	food_id_str, _ := item_id.ReadString('\n')
	food_id_str = strings.TrimSpace(food_id_str)
	food_id, err := strconv.Atoi(food_id_str) // Convert string to int
	if err != nil {
		fmt.Println("Invalid input for Food ID:", err)
		return
	}

	item_status := bufio.NewReader(os.Stdin)
	fmt.Println("Food Staus: Available/Out of Stock")
	food_status, _ := item_status.ReadString('\n')
	food_status = strings.TrimSpace(food_status)

	item_price := bufio.NewReader(os.Stdin)
	fmt.Println("Food Price")
	food_price_str, _ := item_price.ReadString('\n')
	food_price_str = strings.TrimSpace(food_price_str)
	food_price, err := strconv.ParseFloat(food_price_str, 64)
	if err != nil {
		fmt.Println("Something issue in Food Price:", err)
		return
	}

	foods = append(foods, Food{
		ItemName: food_name,
		ItemID:   food_id,
		Status:   food_status,
		Price:    food_price,
	})

	for _, f := range foods {
		fmt.Printf("The food %v is updated", f.ItemName)
	}
	db, err := database.DatabaseConnection()
	if err != nil {
		log.Fatal(err)
	}
	updateQuery := `UPDATE foods SET item_name = ?, status = ?, price = ? WHERE item_id = ?`

	// Execute the query
	_, err = db.Exec(updateQuery, food_name, food_status, food_price, food_id)
	if err != nil {
		log.Fatal("Error updating the food:", err)
	} else {
		fmt.Println("Food details updated successfully!")
	}
	defer db.Close()
}
