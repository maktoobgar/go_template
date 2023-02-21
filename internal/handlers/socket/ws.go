package wsHandlers

import (
	"fmt"
	"log"

	"github.com/gofiber/websocket/v2"
	g "github.com/maktoobgar/go_template/internal/global"
)

func WS(ctx *websocket.Conn) {
	for {
		mt, msg, err := ctx.ReadMessage()
		if err != nil {
			g.Logger.Error(fmt.Sprintf("read: %s", err), WS, nil)
			break
		}
		log.Printf("recv: %s", msg)
		err = ctx.WriteMessage(mt, msg)
		if err != nil {
			g.Logger.Error(fmt.Sprintf("write: %s", err), WS, nil)
			break
		}
	}
}
