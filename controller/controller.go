package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/EmeraldLS/MailService/mail"
	"github.com/EmeraldLS/MailService/model"
	"github.com/dustin/go-humanize"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func SendMail(c *gin.Context) {
	c.Header("content-type", "application/json")
	var details model.Details
	if err := c.ShouldBindJSON(&details); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": "bind_error",
			"message":  fmt.Sprintf("%v", err),
		})
		c.Abort()
		return
	}
	validate := validator.New()
	if err := validate.Struct(details); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": "struct_error",
			"message":  fmt.Sprintf("%v", err),
		})
		c.Abort()
		return
	}

	amtInt, _ := strconv.Atoi(details.Amount)
	taxInt, _ := strconv.Atoi(details.Tax)
	totalInt := amtInt + taxInt
	details.Total = strconv.Itoa(totalInt)
	details.Tax = formatInt(taxInt)
	details.Amount = formatInt(amtInt)
	details.Total = formatInt(totalInt)
	message, err := mail.SendMail(details)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": "email_error",
			"message":  fmt.Sprintf("%v", err),
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"response": "success",
		"message":  message,
		"data":     details,
		"date":     details.Date,
	})
}

func formatInt(num int) string {
	return humanize.Comma(int64(num))
}
