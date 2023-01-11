package controllers

import (
	"github.com/andrew-svirin/multiuser-table-go/server/models"
	"github.com/andrew-svirin/multiuser-table-go/server/repositories"
	"github.com/andrew-svirin/multiuser-table-go/server/services/websocket"
	"strconv"
)

// HandleCellEdit - handle event about cell edit.
func HandleCellEdit(ie *websocket.Event, b *websocket.Bus) {
	oe := websocket.NewEvent(
		"cell/edited",
		websocket.EventData{
			"name":  ie.Data["name"],
			"value": ie.Data["value"],
		},
	)
	b.ConnectionWriteEvent(oe)

	oea := websocket.NewEvent(
		"user/cell/edited",
		websocket.EventData{
			"user_id": b.ConnectionId(),
			"name":    ie.Data["name"],
			"value":   ie.Data["value"],
		},
	)
	go b.ConnectionPoolWriteEvent(oea)
}

// HandleCellLoadAll - handle event about load all cells.
func HandleCellLoadAll(_ *websocket.Event, b *websocket.Bus) {
	rep := repositories.NewCellRepository()

	rep.FindAll().Each(rep.CellHandler(func(c *models.Cell) {
		oe := websocket.NewEvent(
			"cell/loading",
			websocket.EventData{
				"name":  c.Column + "-" + strconv.Itoa(c.Row),
				"value": string(c.Value),
			},
		)
		b.ConnectionWriteEvent(oe)
	}))
}
