package events

import (
	"github.com/andrew-svirin/multiuser-table-go/server/models"
	"github.com/andrew-svirin/multiuser-table-go/server/services/websocket"
	"strconv"
	"strings"
)

func MakeLoadingEventData(c *models.Cell) websocket.EventData {
	return websocket.EventData{
		"name":  c.Column + "-" + strconv.Itoa(c.Row),
		"value": string(c.Value),
	}
}

func ParseSaveEventData(ed websocket.EventData) *models.Cell {
	c := models.NewCell()
	c.Column, c.Row = parseName(ed["name"].(string))
	c.Value = []byte(ed["value"].(string))

	return c
}

func parseName(name string) (column string, row int) {
	res := strings.Split(name, "-")

	column = res[0]
	row, err := strconv.Atoi(res[1])

	if err != nil {
		panic("Incorrect row")
	}

	return
}
