package model

type Message struct {
	Id      int
	Pattern string
}

func NewMessage(pattern string) *Message {
	return &Message{Pattern: pattern}
}
