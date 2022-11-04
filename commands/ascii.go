package commands

import "lmdb-cli/core"

var ascii = []byte(`
Just a simple LMDB Reader/Writer
`)

type Ascii struct {
}

func (cmd Ascii) Execute(context *core.Context, input []byte) (err error) {
	context.Output(ascii)
	return nil
}
