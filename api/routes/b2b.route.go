package routes

import (
	"net/http"
	"strconv"
	"time"

	"com.github/FelipecgPereira/go-jobs/models"
	"com.github/FelipecgPereira/go-jobs/models/enums"
	"github.com/gin-gonic/gin"
)

func updateB2b(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid b2b ID"})
		return
	}

	userId := context.GetInt64("userId")

	b2b, err := models.GetB2bByID(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not retrieve b2b"})
		return
	}
	if b2b == nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "b2b not found"})
		return
	}
	if b2b.UserId != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	var updateB2b models.B2b

	if err := context.ShouldBindJSON(&updateB2b); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request body"})
		return
	}

	if !updateB2b.Status.IsValid() {
		updateB2b.Status = b2b.Status
	}

	updateB2b.UserId = userId
	updateB2b.Id = id

	if updateB2b.CustomerId > 0 {
		customer, err := models.GetCustomerById(b2b.CustomerId)
		if err != nil || customer == nil {
			context.JSON(http.StatusBadGateway, gin.H{"message": "Invalid customer"})
			return
		}
	} else {
		updateB2b.CustomerId = b2b.CustomerId
	}

	if updateB2b.ProjectId > 0 {
		project, err := models.GetProjectById(updateB2b.ProjectId)
		if err != nil || project == nil {
			context.JSON(http.StatusBadGateway, gin.H{"message": "Invalid project"})
			return
		}
	} else {
		updateB2b.ProjectId = b2b.ProjectId
	}

	if err := updateB2b.Update(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update b2b", "error": err})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "B2B updated successfully"})
}

func sumPayments(context *gin.Context) {
	statusStr := context.Query("status")
	startStr := context.Query("start")
	endStr := context.Query("end")

	if statusStr == "" || startStr == "" || endStr == "" {
		context.JSON(http.StatusBadRequest, gin.H{"message": "status, start and end query params are required (start/end in YYYY-MM-DDTHH:MM:SS)"})
		return
	}

	const dateTimeFormat = "2006-01-02T15:04:05"
	start, err := time.Parse(dateTimeFormat, startStr)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "invalid start date format"})
		return
	}
	end, err := time.Parse(dateTimeFormat, endStr)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "invalid end date format"})
		return
	}

	status := enums.Status(statusStr)
	if !status.IsValid() {
		context.JSON(http.StatusBadRequest, gin.H{"message": "invalid status"})
		return
	}

	userId := context.GetInt64("userId")

	total, err := models.SumPaymentsByStatusAndDate(status, start, end, userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not compute total", "error": err})
		return
	}

	context.JSON(http.StatusOK, gin.H{"total": total})
}

func createB2b(context *gin.Context) {
	var b2b models.B2b

	err := context.ShouldBindJSON(&b2b)

	if err != nil {
		context.JSON(http.StatusBadGateway, gin.H{"message": "Could not parse request body"})
		return
	}

	b2b.UserId = context.GetInt64("userId")
	customer, err := models.GetCustomerById(b2b.CustomerId)

	if err != nil || customer == nil {
		context.JSON(http.StatusBadGateway, gin.H{"message": "Invalid customer"})
		return
	}

	if b2b.ProjectId > 0 {
		project, err := models.GetProjectById(b2b.ProjectId)
		if err != nil || project == nil {
			context.JSON(http.StatusBadGateway, gin.H{"message": "Invalid project"})
			return
		}
	}

	err = b2b.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create b2b", "error": err})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "b2b created successfully"})

}
