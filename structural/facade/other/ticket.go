package other

type TicketType string

const (
	GoFood TicketType = "Go Food"
	GoRide TicketType = "Go Ride"
	GoSend TicketType = "Go Send"
)

type Ticket struct {
	UserEmail string
	Type      TicketType
	Message   string
}
