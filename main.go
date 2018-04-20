package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
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

	bookings = recursiveOrdering(bookings)

	return finalOrder
}

// Method to be called recursively for ordering list of bookings to match Start/End locations
func recursiveOrdering(bookings []Booking) []Booking {

	var bookingAdded bool

	// Loop through all bookings to find next booking which matches previous End location
	for _, booking := range bookings {

		if booking.Start == finalOrder[len(finalOrder)-1].End {
			finalOrder = append(finalOrder, booking)
			bookings = removeBookingByID(bookings, booking.Id)

			bookingAdded = true
		}
	}

	// Call recursive order method if a valid booking was found already
	// and more bookings are available
	if bookingAdded == true && len(bookings) > 0 {
		recursiveOrdering(bookings)

		// If no booking matching last end location can be found, add next available booking
	} else if bookingAdded == false && len(bookings) > 0 {
		finalOrder = append(finalOrder, bookings[0])
		bookings = removeBookingByID(bookings, bookings[0].Id)

		// Increase relocation counter for non-matching Start/End locations
		relocationAmount++

		// If more bookings remain to be ordered, call self recursively
		if len(bookings) > 0 {
			recursiveOrdering(bookings)
		}

	}

	return bookings
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

	// Order bookings
	bookings = orderBookings(bookings)

	// Print the final ordered list to STDOUT
	fmt.Println("---Final Ordered Booking List---")
	for _, booking := range bookings {
		fmt.Println(booking.toString())
	}

	// Print the total required relocation count to STDOUT
	fmt.Println("---")
	fmt.Println(
		"Number of relocations required: " +
			strconv.Itoa(relocationAmount),
	)
	fmt.Println("---")
}
