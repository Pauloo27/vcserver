package models

type AbstractTarget interface {
	Start() error
	Stop() error
	Connect(id AbstractConnectionID) error
	Write([]byte) (int, error)
	Disconnect() error
}
