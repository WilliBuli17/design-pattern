package facade

import (
	"design-patterns/structural/facade/other"
	"fmt"
	"testing"
)

func TestGojekCustomerServiceGoFood(t *testing.T) {
	response := other.HandleTicket(other.Ticket{
		UserEmail: "emailTets@mailtest.com",
		Type:      other.GoFood,
		Message:   "food never get delivered",
	})

	fmt.Println(response.Status)
}

func TestGojekCustomerServiceGoRide(t *testing.T) {
	response := other.HandleTicket(other.Ticket{
		UserEmail: "emailTets@mailtest.com",
		Type:      other.GoRide,
		Message:   "Driver mati",
	})

	fmt.Println(response.Status)
}

func TestGojekCustomerServiceGoSend(t *testing.T) {
	response := other.HandleTicket(other.Ticket{
		UserEmail: "emailTets@mailtest.com",
		Type:      other.GoSend,
		Message:   "Paket di curi",
	})

	fmt.Println(response.Status)
}
