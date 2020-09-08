package apichecker

import (
	"fmt"
	"log"
	"time"

	"github.com/binary-kitchen/kitchenpager/pkg/dapnet"
	"github.com/binary-kitchen/kitchenpager/pkg/spaceapi"
)

//CheckandPageOpenStatusperiodically gets the current open state and if it changed it will Page all people
func CheckandPageOpenStatusperiodically(ticker *time.Ticker, pager *dapnet.Pager, URL string, callsigns []string, transmitterGroupNames []string) {
	var laststatus bool
	for range ticker.C {
		schema, err := spaceapi.GetSchemafromURL(URL)
		if err != nil {
			log.Println("Error checking Space Status:", err)
		} else {
			if laststatus != schema.State.Open {
				laststatus = schema.State.Open
				text := "Kitchen Status changed to: "
				if schema.State.Open {
					text = text + "Open"
				} else {
					text = text + "Close"
				}
				fmt.Println(text)
				c := dapnet.Call{
					Text:                  text,
					CallSignNames:         callsigns,
					TransmitterGroupNames: transmitterGroupNames,
				}
				pager.AddNewCalltoQueue(c)
			}
		}
	}
}
