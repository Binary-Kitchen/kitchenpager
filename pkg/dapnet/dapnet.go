package dapnet

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const dapnetURL = "http://www.hampager.de:8080/"

//Call is the go representation of the Dapnet call api
type Call struct {
	Text                  string   `json:"text"`
	CallSignNames         []string `json:"callSignNames"`
	TransmitterGroupNames []string `json:"transmitterGroupNames"`
	Emergency             bool     `json:"emergency"`
}

//Pager is the struct that keeps the queue with all to send Calls
type Pager struct {
	queue    []Call
	username string
	password string
}

//NewPager returns a NewPager with the given password and username
func NewPager(Username, Password string) *Pager {
	p := &Pager{
		username: Username,
		password: Password,
	}
	return p
}

//AddNewCalltoQueue adds the given call to the sending queue
func (p *Pager) AddNewCalltoQueue(call Call) {
	p.queue = append(p.queue, call)
}

//Start starts the paging of the queue
func (p *Pager) Start() {
	for {
		if len(p.queue) > 0 {
			c := p.queue[0]
			copy(p.queue[0:], p.queue[1:])
			p.queue[len(p.queue)-1] = Call{}
			p.queue = p.queue[:len(p.queue)-1]
			b := new(bytes.Buffer)
			encoder := json.NewEncoder(b)
			encoder.Encode(c)
			req, err := http.NewRequest("POST", dapnetURL+"calls", b)
			if err != nil {
				log.Println("Error creating request:", err)
				continue
			}
			req.Header.Set("Content-Type", "application/json")
			req.SetBasicAuth(p.username, p.password)
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				log.Println("Error executing request: " + err.Error())
				continue
			}
			if resp.StatusCode > 299 {
				data, _ := ioutil.ReadAll(resp.Body)
				log.Println("Status code greater 299:", resp.StatusCode, "", string(data))
			}
		}
	}
}
