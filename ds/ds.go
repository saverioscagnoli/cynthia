package ds

import "cynthia/ds/gateway"

type Client = gateway.Client

var NewClient = gateway.NewClient
var On = gateway.On[any]
