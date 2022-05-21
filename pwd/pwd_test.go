package pwd

import (
	"testing"
	"time"
)

func TestHash(t *testing.T) {
	const password = "gorust.net"
	for i := 0; i < 10; i++ {
		hashedPassword, err := Hash(password)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(hashedPassword, len(hashedPassword))
		time.Sleep(time.Second)
	}
}

func TestVerify(t *testing.T) {
	const (
		password       = "gorust.net"
		hashedPassword = "$2a$10$1QOGW1RXo3fQQzWGKLbS2ui7DxPrSc97ycGILbJE6EbBfKMJXeQn2"
	)
	t.Log(Verify(password, hashedPassword))
}
