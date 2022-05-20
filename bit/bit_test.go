package bit

import "testing"

func TestBit(t *testing.T) {
	const (
		read         Bit = 1
		write        Bit = 2
		readWrite    Bit = 3
		execute      Bit = 4
		readExecute      = 5
		writeExecute     = 6
		all          Bit = 7
	)
	t.Log(Add(read, write) == readWrite)
	t.Log(Has(readWrite, read))
	t.Log(Has(readWrite, write))
	t.Log(Add(read, execute) == readExecute)
	t.Log(Has(readExecute, read))
	t.Log(Has(readExecute, execute))
	t.Log(Add(write, execute) == writeExecute)
	t.Log(Add(read, Add(write, execute)) == all)
	t.Log(Remove(all, write) == readExecute)
}
