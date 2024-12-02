// file: food/updatefood.go
package food

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"restaurant/backend/database"
	"strings"
)

type Food struct {
	ItemName string  `json:"item_name"`
	ItemID   int     `json:"item_id"`
	Status   string  `json:"status"`
	Price    float64 `json:"price"`
}

var foods []Food

// addFood is the handler to add food details
func UpdateFood(w http.ResponseWriter, r *http.Request) {
	var newFood Food
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newFood)
	if err != nil {
		http.Error(w, "Invalid food data", http.StatusBadRequest)
		return
	}

	// Validation or processing food details
	newFood.ItemName = strings.TrimSpace(newFood.ItemName)
	newFood.Status = strings.TrimSpace(newFood.Status)

	// Database connection and query
	db, err := database.DatabaseConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	updateQuery := `UPDATE foods SET item_name = ?, status = ?, price = ? WHERE item_id = ?`
	_, err = db.Exec(updateQuery, newFood.ItemName, newFood.Status, newFood.Price, newFood.ItemID)
	if err != nil {
		http.Error(w, "Error updating food details", http.StatusInternalServerError)
		return
	}

	// Success response
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Food details updated successfully!")
}
