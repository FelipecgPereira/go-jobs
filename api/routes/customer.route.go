package routes

import (
	"net/http"
	"strconv"

	"com.github/FelipecgPereira/go-jobs/models"
	"github.com/gin-gonic/gin"
)

func createCustomer(context *gin.Context) {
	var customer models.Customer

	err := context.ShouldBindJSON(&customer)

	if err != nil {
		context.JSON(http.StatusBadGateway, gin.H{"message": "Could not parse request body"})
		return
	}

	customer.UserID = context.GetInt64("userId")

	err = customer.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create customer", "error": err})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Customer created successfully"})

}

func getCustomers(context *gin.Context) {
	userId := context.GetInt64("userId")
	customers, err := models.GetCustomer(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not retrieve customers"})
		return
	}

	context.JSON(http.StatusOK, customers)
}

func getCustmerById(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}

	customer, err := models.GetCustomerById(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not retrieve customer"})
		return
	}
	if customer == nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "customer not found"})
		return
	}
	context.JSON(http.StatusOK, customer)
}

func updateCustomer(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}

	userId := context.GetInt64("userId")
	customer, err := models.GetCustomerById(id)

	if customer.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	var updateCustomer models.Customer
	if err := context.ShouldBindJSON(&updateCustomer); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request body"})
		return
	}

	updateCustomer.Id = id
	err = updateCustomer.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update customer"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Customer updated successfully"})

}
