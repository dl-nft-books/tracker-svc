package ethereum

import "gitlab.com/distributed_lab/logan/v3/errors"

var (
	FromNotSpecifiedErr    = errors.New("starting block is not specified")
	AddressNotSpecifiedErr = errors.New("address is not specified")
	RpcNotSpecifiedErr     = errors.New("rpc is not specified")
	WSNotSpecifiedErr      = errors.New("websocket is not specified")
	NullIteratorErr        = errors.New("iterator has a nil value")
)
