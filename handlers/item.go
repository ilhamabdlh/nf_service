package handlers

import (
	"database/sql"
	"net/http"
	"net/url"
	"strings"

	"nft_service/models"

	"github.com/gin-gonic/gin"
)

var db *sql.DB

func SetDB(database *sql.DB) {
	db = database
}

func GetItems(c *gin.Context) {
	rating := c.Query("rating")
	reputationBadge := c.Query("reputationBadge")
	minAvailability := c.Query("minAvailability")
	maxAvailability := c.Query("maxAvailability")
	category := c.Query("category")

	query := "SELECT * FROM items WHERE 1=1"
	if rating != "" {
		query += " AND rating = " + rating
	}
	if reputationBadge != "" {
		query += " AND reputation_badge = '" + reputationBadge + "'"
	}
	if minAvailability != "" {
		query += " AND availability >= " + minAvailability
	}
	if maxAvailability != "" {
		query += " AND availability <= " + maxAvailability
	}
	if category != "" {
		query += " AND category = '" + category + "'"
	}

	rows, err := db.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse("Failed to fetch items", http.StatusInternalServerError))
		return
	}
	defer rows.Close()

	var items []models.Items
	for rows.Next() {
		var item models.Items
		err := rows.Scan(&item.ID, &item.Name, &item.Rating, &item.Category, &item.Image, &item.Reputation, &item.Price, &item.Availability, &item.ReputationBadge)
		if err != nil {
			c.JSON(http.StatusInternalServerError, errorResponse("Failed to scan item", http.StatusInternalServerError))
			return
		}
		items = append(items, item)
	}

	if len(items) == 0 {
		c.JSON(http.StatusNotFound, errorResponse("Data not found", http.StatusNotFound))
		return
	}

	c.JSON(http.StatusOK, items)
}

func GetItemByID(c *gin.Context) {
	itemID := c.Param("id")
	var item models.Items
	err := db.QueryRow("SELECT * FROM items WHERE id = ?", itemID).Scan(&item.ID, &item.Name, &item.Rating, &item.Category, &item.Image, &item.Reputation, &item.Price, &item.Availability, &item.ReputationBadge)
	if err != nil {
		c.JSON(http.StatusNotFound, errorResponse("Item not found", http.StatusNotFound))
		return
	}

	c.JSON(http.StatusOK, item)
}

func CreateItem(c *gin.Context) {
	var newItem models.Items
	if err := c.ShouldBindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse("Invalid request body or data type", http.StatusBadRequest))
		return
	}

	valid, errMsg, repColor := validateItem(newItem)
	if !valid {
		c.JSON(http.StatusBadRequest, errorResponse(errMsg, http.StatusBadRequest))
		return
	}

	newItem.ReputationBadge = repColor

	_, err := db.Exec("INSERT INTO items (name, rating, category, image, reputation, price, availability, reputation_badge) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		newItem.Name, newItem.Rating, newItem.Category, newItem.Image, newItem.Reputation, newItem.Price, newItem.Availability, repColor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse("Failed to create item", http.StatusInternalServerError))
		return
	}

	c.JSON(http.StatusCreated, newItem)
}

func UpdateItem(c *gin.Context) {
	itemID := c.Param("id")
	var updatedItem models.Items
	if err := c.ShouldBindJSON(&updatedItem); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse("Invalid request body", http.StatusBadRequest))
		return
	}

	isValid, errMsg, repColor := validateItem(updatedItem)
	if !isValid {
		c.JSON(http.StatusBadRequest, errorResponse(errMsg, http.StatusBadRequest))
		return
	}

	updatedItem.ReputationBadge = repColor

	result, err := db.Exec("UPDATE items SET name=?, rating=?, category=?, image=?, reputation=?, price=?, availability=?, reputation_badge=? WHERE id=?",
		updatedItem.Name, updatedItem.Rating, updatedItem.Category, updatedItem.Image, updatedItem.Reputation, updatedItem.Price, updatedItem.Availability, repColor, itemID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse("Failed to update item", http.StatusInternalServerError))
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, errorResponse("Item not found", http.StatusNotFound))
		return
	}

	c.JSON(http.StatusOK, updatedItem)
}

func DeleteItem(c *gin.Context) {
	itemID := c.Param("id")
	result, err := db.Exec("DELETE FROM items WHERE id=?", itemID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse("Failed to delete item", http.StatusInternalServerError))
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, errorResponse("Item not found", http.StatusNotFound))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item deleted successfully"})
}

func PurchaseItem(c *gin.Context) {
	itemID := c.Param("id")
	var availability int
	err := db.QueryRow("SELECT availability FROM items WHERE id=?", itemID).Scan(&availability)
	if err != nil {
		c.JSON(http.StatusNotFound, errorResponse("Item not found", http.StatusNotFound))
		return
	}

	if availability <= 0 {
		c.JSON(http.StatusBadRequest, errorResponse("Item not available", http.StatusBadRequest))
		return
	}

	_, err = db.Exec("UPDATE items SET availability = availability - 1 WHERE id=?", itemID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse("Failed to update item availability", http.StatusInternalServerError))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item purchased successfully"})
}

func validateItem(item models.Items) (bool, string, string) {
	invalidWords := []string{"Sex", "Gay", "Lesbian", "sex", "gay", "lesbian"}
	validCategory := map[string]bool{"photo": true, "sketch": true, "cartoon": true, "animation": true}

	if len(item.Name) < 10 {
		return false, "Name must be at least 10 characters long", ""
	}
	if item.Rating < 0 || item.Rating > 5 {
		return false, "Rating must be between 0 and 5", ""
	}
	if !validCategory[item.Category] {
		return false, "Invalid category", ""
	}
	if item.Reputation < 0 || item.Reputation > 1000 {
		return false, "Reputation must be between 0 and 1000", ""
	}
	switch {
	case item.Reputation <= 500:
		item.ReputationBadge = "red"
	case item.Reputation <= 799:
		item.ReputationBadge = "yellow"
	default:
		item.ReputationBadge = "green"
	}

	for _, word := range invalidWords {
		if strings.Contains(item.Name, word) {
			return false, "Name contains invalid word: " + word, ""
		}
	}
	if _, err := url.ParseRequestURI(item.Image); err != nil {
		return false, "Invalid image URL", ""
	}

	return true, "", item.ReputationBadge
}

func errorResponse(message string, status int) gin.H {
	return gin.H{
		"title":  http.StatusText(status),
		"status": status,
		"detail": message,
	}
}
