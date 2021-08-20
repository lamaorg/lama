package internals

import "context"

const (
	EchoService         = "EchoRPCAPI"
	EchoServiceFuncEcho = "Echo"
)

type EchoRPCAPI struct {
	service *Service
}

type Envelope struct {
	Message string
}

func (e *EchoRPCAPI) Echo(ctx context.Context, in Envelope, out *Envelope) error {
	*out = e.service.ReceiveEcho(in)
	return nil
}
