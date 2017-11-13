// Network rpc example from Go Standard Library.
package nrpc

type Args struct {
	A, B int
}
type Arith int // can be any base type. For saving space use struct{} -- empty struct.

func (t *Arith) Sum(args *Args, reply *int) error {
	*reply = args.A + args.B
	return nil
}
