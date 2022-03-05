package dummy

import "github.com/Pauloo27/vcserver/internal/models"

type DummyTarget struct {
	StartFn      func() error
	StopFn       func() error
	ConnectFn    func(models.AbstractConnectionID) error
	WriteFn      func([]byte) (int, error)
	DisconnectFn func() error
}

var _ models.AbstractTarget = DummyTarget{}

func (d DummyTarget) Start() error {
	return d.StartFn()
}

func (d DummyTarget) Stop() error {
	return d.StopFn()
}

func (d DummyTarget) Connect(id models.AbstractConnectionID) error {
	return d.ConnectFn(id)
}

func (d DummyTarget) Write(b []byte) (int, error) {
	return d.WriteFn(b)
}

func (d DummyTarget) Disconnect() error {
	return d.DisconnectFn()
}
