package commands

import (
	"bytes"
	"fmt"

	"lmdb-cli/core"

	"github.com/bmatsuo/lmdb-go/lmdb"
)

type Recover struct {
}

func (cmd Recover) Execute(context *core.Context, input []byte) (err error) {
	fmt.Println(context.ReaderCheck())
	context.Env.ReaderList(func(s string) error {
		fmt.Println("Reader: ", s)
		return nil
	})
	return cmd.execute(context, false)
}

func (cmd Recover) execute(context *core.Context, first bool) (err error) {
	fmt.Println("Recovering...")
	cursor := context.Cursor
	fmt.Println("Cursor: ", cursor)
	if cursor == nil {
		return nil
	}
	for i := 0; i < 10; i++ {
		fmt.Println("i: ", i)
		var err error
		var key, value []byte
		if first && cursor.Prefix != nil {
			key, value, err = cursor.Get(cursor.Prefix, nil, lmdb.SetRange)
			first = false
		} else {
			key, value, err = cursor.Get(nil, nil, lmdb.Next)
		}

		if lmdb.IsNotFound(err) || (cursor.Prefix != nil && !bytes.HasPrefix(key, cursor.Prefix)) {
			context.CloseCursor()
			return nil
		}
		if err != nil {
			context.CloseCursor()
			return err
		}
		context.Output(key)
		if cursor.IncludeValues {
			context.Output(value)
			context.Output(nil)
		}
	}
	context.Output(SCAN_MORE)
	return nil
}
