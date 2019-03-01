package client

import (
	"errors"
)

var (
	ResponseErrors = map[string]error{
		"nohost":   errors.New("The hostname does not exist, or does not have Dynamic DNS enabled."),
		"badauth":  errors.New("The username / password combination is not valid for the specified host."),
		"notfqdn":  errors.New("The supplied hostname is not a valid fully-qualified domain name."),
		"badagent": errors.New("Your Dynamic DNS client is making bad requests. Ensure the user agent is set in the request."),
		"abuse":    errors.New("Dynamic DNS access for the hostname has been blocked due to failure to interpret previous responses correctly."),
		"911":      errors.New("An error happened on our end. Wait 5 minutes and retry."),
	}
)
