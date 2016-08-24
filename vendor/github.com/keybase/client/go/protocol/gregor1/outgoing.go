// Auto-generated by avdl-compiler v1.3.4 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/gregor1/outgoing.avdl

package gregor1

import (
	rpc "github.com/keybase/go-framed-msgpack-rpc"
	context "golang.org/x/net/context"
)

type BroadcastMessageArg struct {
	M Message `codec:"m" json:"m"`
}

type OutgoingInterface interface {
	BroadcastMessage(context.Context, Message) error
}

func OutgoingProtocol(i OutgoingInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "gregor.1.outgoing",
		Methods: map[string]rpc.ServeHandlerDescription{
			"broadcastMessage": {
				MakeArg: func() interface{} {
					ret := make([]BroadcastMessageArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]BroadcastMessageArg)
					if !ok {
						err = rpc.NewTypeError((*[]BroadcastMessageArg)(nil), args)
						return
					}
					err = i.BroadcastMessage(ctx, (*typedArgs)[0].M)
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type OutgoingClient struct {
	Cli rpc.GenericClient
}

func (c OutgoingClient) BroadcastMessage(ctx context.Context, m Message) (err error) {
	__arg := BroadcastMessageArg{M: m}
	err = c.Cli.Call(ctx, "gregor.1.outgoing.broadcastMessage", []interface{}{__arg}, nil)
	return
}
