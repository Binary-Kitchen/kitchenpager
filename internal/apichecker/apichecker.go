package apichecker

import (
	"fmt"
	"kitchenpager/pkg/dapnet"
	"kitchenpager/pkg/spaceapi"
	"log"
	"time"
)

//CheckandPageOpenStatusperiodically gets the current open state and if it changed it will Page all people
func CheckandPageOpenStatusperiodically(ticker *time.Ticker, pager *dapnet.Pager, URL string, callsigns []string) {
	var laststatus bool
	laststatus = true
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
					TransmitterGroupNames: []string{"all"},
				}
				pager.AddNewCalltoQueue(c)
			}
		}
	}
}
