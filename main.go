package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

// Global vars for storing ordered list and relocation amount
var finalOrder []Booking
var relocationAmount int

// Go struct representation of a single booking
type Booking struct {
	Id    int `json:"id"`
	Start int `json:"start"`
	End   int `json:"end"`
}

// String converter method for Booking struct
func (booking Booking) toString() string {
	return toJson(booking)
}

// Converts struct into JSON string for e.g. printing
func toJson(booking interface{}) string {
	// Marshal booking struct to JSON
	bytes, err := json.Marshal(booking)

	// Throw error and die if marshalling failed for any reason
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return string(bytes)
}

// Parses JSON into struct from given URI
func getBookingsFromFile(uri string) []Booking {
	// Parse raw data from file
	raw, err := ioutil.ReadFile(uri)

	// Throw error and die if file parsing failed for any reason
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Unmarshal JSON into Go struct, die on error
	var bookings []Booking
	if err := json.Unmarshal(raw, &bookings); err != nil {
		fmt.Println("Unable to parse JSON file. Invalid JSON? " + err.Error())
		os.Exit(1)
	}
	return bookings
}

// Removes the booking with the given ID from the given list
func removeBookingByID(bookings []Booking, ID int) []Booking {

	var scrappedBookings []Booking

	for _, booking := range bookings {
		if booking.Id == ID {
			continue
		} else {
			scrappedBookings = append(scrappedBookings, booking)
		}
	}

	return scrappedBookings
}

func orderBookings(bookings []Booking) []Booking {

	// If list contains bookings, add first entry to list
	if len(bookings) > 0 {
		finalOrder = append(finalOrder, bookings[0])
		bookings = removeBookingByID(bookings, bookings[0].Id)
	}

	return finalOrder
}

func main() {

	// Parse JSON file from CLI argument
	// If no file was passed, use bookingordering.json as default
	var bookingFile string

	flag.StringVar(
		&bookingFile,
		"json",
		"bookingordering.json",
		"JSON file with list of bookings",
	)
	flag.Parse()

	// Parse all bookings from JSON file parsed via CLI
	bookings := getBookingsFromFile(bookingFile)

	// Print parsed data to STDOUT
	fmt.Println("---Original Booking List---")
	for _, booking := range bookings {
		fmt.Println(booking.toString())
	}
	fmt.Println("---")
}
