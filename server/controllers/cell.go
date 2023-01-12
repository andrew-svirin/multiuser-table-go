package controllers

import (
	"github.com/andrew-svirin/multiuser-table-go/server/events"
	"github.com/andrew-svirin/multiuser-table-go/server/models"
	"github.com/andrew-svirin/multiuser-table-go/server/repositories"
	"github.com/andrew-svirin/multiuser-table-go/server/services/websocket"
)

// HandleCellSave - handle event about cell edit.
func HandleCellSave(ie *websocket.Event, b *websocket.Bus) {
	rep := repositories.NewCellRepository()

	c := events.ParseSaveEventData(ie.Data)

	rep.Save(c)

	oe := websocket.NewEvent(
		"cell/saved",
		websocket.EventData{
			"name":  ie.Data["name"],
			"value": ie.Data["value"],
		},
	)
	b.ConnectionWriteEvent(oe)

	oea := websocket.NewEvent(
		"user/cell/saved",
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

	rep.FindAll().Each(rep.CastCell(func(c *models.Cell) {
		oe := websocket.NewEvent(
			"cell/loading",
			events.MakeLoadingEventData(c),
		)
		b.ConnectionWriteEvent(oe)
	}))
}
