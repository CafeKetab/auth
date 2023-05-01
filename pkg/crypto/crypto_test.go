package crypto_test

import (
	"fmt"
	"testing"

	"github.com/CafeKetab/auth/pkg/crypto"
)

func TestAES(t *testing.T) {
	cfg := &crypto.Config{Secret: "A?D(G-KaPdSgVkYp", Salt: "salt"}
	crypto := crypto.New(cfg)
	plainText := "plainText"

	encrypted, err := crypto.Encrypt(plainText)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(encrypted)

	decrypted, err := crypto.Decrypt(encrypted)
	if err != nil {
		t.Error(err)
	}

	if plainText != decrypted {
		t.Errorf("expected: %s, recieved: %s\n", plainText, decrypted)
	}
}
