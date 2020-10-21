package api

const (
	ServiceMethod = "Arith.Add"
)

type Request struct {
	A, B int
}

type Reply struct {
	C int
}

type Arith struct {
}

func (a *Arith) Add(req Request, reply *Reply) error {
	reply.C = req.A + req.B
	return nil
}

type hidden int

func (t *hidden) Exported(req Request, reply *Reply) error {
	reply.C = req.A + req.B
	return nil
}

type Embed struct {
	hidden
}
