// Network rpc example from Go Standard Library.
package nrpc

type Args struct {
	A, B int
}
type Arith int

func (t *Arith) Sum(args *Args, reply *int) error {
	*reply = args.A + args.B
	return nil
}
