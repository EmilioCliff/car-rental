package api

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/EmilioCliff/car-rental/templates"
	"github.com/gin-gonic/gin"
)

func (server *Server) formHandlers(ctx *gin.Context) {
	form := ctx.Param("form")

	switch form {
	case "enquery":
		// save to a database
		var req templates.EnqueryFormPost
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, server.errorResponse(err))
			return
		}

		fmt.Println(req)
		t := templates.EnqueryFormComponent(req)

		var buf bytes.Buffer
		if err := t.Render(ctx, &buf); err != nil {
			ctx.JSON(http.StatusInternalServerError, server.errorResponse(err))
			return
		}

		emailContent := buf.String()
		fmt.Println(emailContent)

		// send with a file or some attachhment for enquiry
		if err := server.sender.SendMAil("Enquery Request", emailContent, "text/plain", []string{"clifftest33@gmail.com", "andysafariskenya@gmail.com"}, []string{}, nil, nil, nil); err != nil {
			ctx.JSON(http.StatusInternalServerError, server.errorResponse(err))
			return
		}

	case "safari":
		// save to a database
		var req templates.SafariFormPost
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, server.errorResponse(err))
			return
		}

		fmt.Println(req)
		t := templates.SafariFormComponent(req, "Safari")

		var buf bytes.Buffer
		if err := t.Render(ctx, &buf); err != nil {
			ctx.JSON(http.StatusInternalServerError, server.errorResponse(err))
			return
		}

		emailContent := buf.String()
		fmt.Println(emailContent)

		if err := server.sender.SendMAil("Safari Booking Confirmation", emailContent, "text/plain", []string{"clifftest33@gmail.com", "andysafariskenya@gmail.com"}, []string{}, nil, nil, nil); err != nil {
			ctx.JSON(http.StatusInternalServerError, server.errorResponse(err))
			return
		}

	case "nairobi-tour":
		//save to a database
		var req templates.SafariFormPost
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, server.errorResponse(err))
			return
		}

		fmt.Println(req)
		t := templates.SafariFormComponent(req, "Nairobi-Tour")

		var buf bytes.Buffer
		if err := t.Render(ctx, &buf); err != nil {
			ctx.JSON(http.StatusInternalServerError, server.errorResponse(err))
			return
		}

		emailContent := buf.String()
		fmt.Println(emailContent)

		if err := server.sender.SendMAil("Nairobi - Tour Confirmation", emailContent, "text/plain", []string{"clifftest33@gmail.com", "andysafariskenya@gmail.com"}, []string{}, nil, nil, nil); err != nil {
			ctx.JSON(http.StatusInternalServerError, server.errorResponse(err))
			return
		}

	case "car-hire":
		// save to the database
		var req templates.CarHireFormPost
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, server.errorResponse(err))
			return
		}

		fmt.Println(req)
		t := templates.CarHireFormComponent(req)

		var buf bytes.Buffer
		if err := t.Render(ctx, &buf); err != nil {
			ctx.JSON(http.StatusInternalServerError, server.errorResponse(err))
			return
		}

		emailContent := buf.String()
		fmt.Println(emailContent)

		if err := server.sender.SendMAil("Car Hire Confirmation", emailContent, "text/plain", []string{"clifftest33@gmail.com", "andysafariskenya@gmail.com"}, []string{}, nil, nil, nil); err != nil {
			ctx.JSON(http.StatusInternalServerError, server.errorResponse(err))
			return
		}

	case "taxi":
		// save to the database
		var req templates.TaxiFormPost
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, server.errorResponse(err))
			return
		}

		fmt.Println(req)
		t := templates.TaxiBookingFormComponent(req)

		var buf bytes.Buffer
		if err := t.Render(ctx, &buf); err != nil {
			ctx.JSON(http.StatusInternalServerError, server.errorResponse(err))
			return
		}

		emailContent := buf.String()
		fmt.Println(emailContent)

		if err := server.sender.SendMAil("Get A Taxi Booking Confirmation", emailContent, "text/plain", []string{"clifftest33@gmail.com", "andysafariskenya@gmail.com"}, []string{}, nil, nil, nil); err != nil {
			ctx.JSON(http.StatusInternalServerError, server.errorResponse(err))
			return
		}

	case "review":
		// save review to the database
		var req templates.ReviewFormPost
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, server.errorResponse(err))
			return
		}

		fmt.Println(req)

	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "unknown form"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": "Enquiry sent successfull"})
}

// "rideDetails":{
// 	"pickupDay":"2024-07-19",
// 	"pickupTime":"17:18",
// 	"pickupLocation":"Marsabit, Kenya",
// 	"dropoffLocation":"Moyale, Kenya",
// 	"transferType":"one-way",
// 	"totalDistance":"245 km",
// 	"totalTime":"2 hours 57 mins"
// 	},

// "taxiSelected":"4",

// "contactDetails":{
// 	"firstName":"MAMA",
// 	"lastName":"MIHA",
// 	"email":"letgo@gmail.com",
// 	"phoneNumber":"0909090909",
// 	"comments":"Test ",
// 	"paymentMethod":"2"
// }
