package client

import (
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"

	"github.com/tserkov/google-ddns-client/internal/log"
	"github.com/tserkov/google-ddns-client/internal/record"
)

const endpoint string = "https://domains.google.com/nic/update?hostname="

type Client struct {
	c http.Client

	Verbose bool
}

func (c *Client) UpdateRecord(r record.Record) error {
	req, err := http.NewRequest("POST", endpoint+r.Hostname, nil)
	if err != nil {
		return errors.Wrap(err, r.Hostname)
	}

	req.Header.Set("User-Agent", "tserkov/google-ddns-client")
	req.SetBasicAuth(r.Username, r.Password)

	res, err := c.c.Do(req)
	if err != nil {
		return errors.Wrap(err, r.Hostname)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return errors.Wrap(err, r.Hostname)
	}

	msg := string(body)

	if err, ok := ResponseErrors[msg]; ok {
		return errors.Wrap(err, r.Hostname)
	}

	if c.Verbose {
		log.Info.Println(r.Hostname + ": " + msg)
	}

	return nil
}
