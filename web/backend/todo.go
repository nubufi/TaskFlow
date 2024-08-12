package backend

import (
	"taskflow/lib"
	"taskflow/models"
	"taskflow/web/templates"

	"github.com/gin-gonic/gin"
)

// AddTodoItem adds a todo item to the database and renders a todo item template
//
// Parameters:
//
// - c: The gin context
func AddTodoItem(c *gin.Context) {
	// Get the JSON body and decode into variables
	var body struct {
		Title string `form:"title" binding:"required"`
	}
	if err := c.Bind(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Get the user ID
	userID := GetUserID(c)

	// Create the todo
	todo := models.TodoItem{
		Title:  body.Title,
		UserID: userID,
	}

	if err := lib.DB.Create(&todo).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to create todo"})
		return
	}

	lib.Render(c, templates.TodoItem(todo))
}

// DeleteTodoItem deletes a todo item
//
// Parameters:
//
// - c: The gin context
func DeleteTodoItem(c *gin.Context) {
	// Get the todo ID
	todoID := c.Param("id")

	// Get the user ID
	userID, _ := c.Get("userID")

	// Delete the todo
	if err := lib.DB.Where("id = ? AND user_id = ?", todoID, userID).Delete(&models.TodoItem{}).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete todo"})
		return
	}
}

// GetTodoItems gets all the todo items for the user and renders the page
//
// Parameters:
//
// - c: The gin context
func GetTodoItems(c *gin.Context) {
	// Get the user ID
	userID := GetUserID(c)

	// Get the todos
	var todos []models.TodoItem
	if err := lib.DB.Where("user_id = ?", userID).Find(&todos).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to get todos"})
		return
	}

	// Respond
	lib.RenderWithLayout(c, templates.TodoItems(todos))
}
