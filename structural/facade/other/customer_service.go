package other

import "design-patterns/structural/facade/other/gojek"

type Response struct {
	Status string
}

func HandleTicket(t Ticket) *Response {
	var response string

	switch t.Type {
	case GoFood:
		gofood := gojek.GoFood{}
		response = gofood.CheckMessage()
	case GoRide:
		goride := gojek.GoRide{}
		response = goride.CheckMessage()
	case GoSend:
		gosend := gojek.GoSend{}
		response = gosend.CheckMessage()
	default:
		response = "customer service is busy"
	}

	return &Response{
		Status: response,
	}
}
