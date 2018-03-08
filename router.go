package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initRouter() {
	r := gin.Default()

	// Transaction endpoint
	r.POST("/transaction/", postTransaction)
	r.GET("/transaction/", getTransactions)
	r.GET("/transaction/:id", getTransaction)
	r.DELETE("/transaction/:id", deleteTransaction)

	// Event endpoint
	r.GET("/transaction/:id/event/", getEvents)
	r.POST("/transaction/:id/event/", postEvent)

	r.Run(":8080")
}

func deleteTransaction(c *gin.Context) {
	id := c.Params.ByName("id")

	db.Where("id = ?", id).Delete(&Transaction{})
}

func postTransaction(c *gin.Context) {
	var tr TransactionRequest

	if err := c.BindJSON(&tr); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	tx := Transaction{
		Raw:  string(tr.Raw),
		Meta: string(tr.Meta),
	}

	if err := db.Create(&tx).Error; err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(201, tx)
}

func getTransactions(c *gin.Context) {
	var txs []Transaction

	db.Find(&txs)

	c.JSON(200, txs)
}

func getTransaction(c *gin.Context) {
	tx, err := findTxByID(
		c.Params.ByName("id"),
	)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(200, tx)
}

func getEvents(c *gin.Context) {
	var events []Event

	tx, err := findTxByID(
		c.Params.ByName("id"),
	)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	db.Model(&tx).Related(&events)

	c.JSON(200, events)
}

func postEvent(c *gin.Context) {
	var event Event

	tx, err := findTxByID(
		c.Params.ByName("id"),
	)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	if err = c.BindJSON(&event); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	event.TransactionID = tx.ID

	if err := db.Create(&event).Error; err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(201, event)
}

func findTxByID(id string) (*Transaction, error) {
	tx := &Transaction{}

	if err := db.Where("id = ?", id).First(&tx).Error; err != nil {
		return nil, err
	}

	return tx, nil
}
