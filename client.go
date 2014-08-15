package qgo

import "encoding/json"
import "io"
import "net/http"
import "time"

func NewClient(url string) *Client {
	return &Client{}
}

type Client struct {
	url string
}

func (c *Client) Times() ([]TimeEntry, error) {
	url := c.url + "/times"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var times []TimeEntry
	body := io.LimitReader(res.Body, 16*1024*1024)
	d := json.NewDecoder(body)
	if err := d.Decode(&times); err != nil {
		return nil, err
	}
	return times, nil
}

// TimeEntry holds a time entry as returned by qapi.
type TimeEntry struct {
	Id       string       `json:"id"`
	URL      string       `json:"url"`
	Category TimeCategory `json:"category"`
	End      *time.Time   `json:"end"`
	Start    time.Time    `json:"start"`
	Note     string       `json:"note"`
	Created  time.Time    `json:"created"`
	Updated  time.Time    `json:"updated"`
}

type TimeCategory struct {
	Name []string `json:"name"`
	URL  string   `json:"url"`
}
