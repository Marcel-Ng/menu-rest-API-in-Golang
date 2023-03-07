package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// menu struct
type menu struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Category     string `json:"category"`
	Quantity     int    `json:"quantity"`
	PricePerUnit int    `json:"pricePerUnit"`
}

// defining a menu slice
type menuLsit []menu

var allMenuList = menuLsit{
	{
		ID:           "1",
		Name:         "Nigerian Jollof",
		Category:     "African",
		Quantity:     10,
		PricePerUnit: 5000,
	},
}

func addMenu(c *gin.Context) {
	var newMenuItem menu

	// call the BindJson to bind the recieved JSON payload to the newMenuItem

	if err := c.BindJSON(&newMenuItem); err != nil {
		errormsg := fmt.Sprintf("Error occured creating the menu %s", err)
		c.JSON(http.StatusBadRequest, errormsg)
		return
	}

	// Add new menu item to the slice
	allMenuList = append(allMenuList, newMenuItem)
	c.IndentedJSON(http.StatusCreated, newMenuItem)
}

func updateMenu(c *gin.Context) {
	menuId := c.Param("menuId")
	var updatedMenu menu

	err := c.BindJSON(&updatedMenu)
	if err != nil {
		errormsg := fmt.Sprintf("Error occured creating the menu %s", err)
		c.JSON(http.StatusBadRequest, errormsg)
		return
	}

	for i, singleMenu := range allMenuList {
		if singleMenu.ID == menuId {
			singleMenu.Name = updatedMenu.Name
			singleMenu.Category = updatedMenu.Category
			singleMenu.Quantity = updatedMenu.Quantity
			singleMenu.PricePerUnit = updatedMenu.PricePerUnit
			allMenuList = append(allMenuList[:i], singleMenu)
			c.IndentedJSON(http.StatusAccepted, singleMenu)
		}
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "Content not found"})
}

func deleteMenuItem(c *gin.Context) {
	menuId := c.Param("menuId")

	for i, singleMenu := range allMenuList {
		if singleMenu.ID == menuId {
			allMenuList = append(allMenuList[:i], allMenuList[i+1:]...)
			c.JSON(http.StatusAccepted, fmt.Sprintf("The event with ID %v has been deleted successfully", menuId))
		}
	}
}

func listAllMenu(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, allMenuList)
}

func getMenuListByCategory(c *gin.Context) {
	menuCategory := c.Param("category")
	var menuListCategory []menu
	var totalList int

	for i, singleMenu := range allMenuList {
		if singleMenu.Category == menuCategory {
			menuListCategory = append(menuListCategory, singleMenu)
			totalList = i
		}
	}
	c.IndentedJSON(http.StatusAccepted, gin.H{"total": totalList, "result": menuListCategory})
}

func main() {
	router := gin.Default()
	router.GET("/menu/all", listAllMenu)
	router.GET("menu/:category", getMenuListByCategory)
	router.DELETE("menu/:menuId", deleteMenuItem)
	router.PATCH("menu/:menuId", updateMenu)
	router.POST("/menu", addMenu)
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "This is the root directory of our application")
	})
	router.Run(":5000")
}
