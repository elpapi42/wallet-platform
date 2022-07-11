package messages

type Message interface {
	GetName() string
	GetKey() string
}
