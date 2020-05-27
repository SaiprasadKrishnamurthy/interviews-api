package models

import "github.com/nats-io/nats.go"

//MessageHandler message handler function.
type MessageHandler func(msg *nats.Msg)
