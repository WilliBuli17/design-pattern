package refactoring_guru

import "fmt"

type Client struct {
}

func (c *Client) InsertLightningConnectorIntoComputer(computer Computer) string {
	fmt.Println("Client inserts Lightning connector into computer.")
	computer.InsertIntoLightningPort()

	return "Success\n\n"
}
