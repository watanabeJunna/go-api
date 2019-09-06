package cipher

import "testing"

func TestCipher(t *testing.T) {
	src := "WatanabeJunna"
	key := "8Fmj4kb4oWpw9g=="

	enc, err := Encrypt([]byte(key), []byte(src))
	
    if err != nil {
        panic(err)
	}

	dec, err := Decrypt([]byte(key), enc)
	
    if err != nil {
        panic(err)
	}
	
	if string(dec) != src {
		t.Errorf("src = %q, dec %q", src, dec)
	}
}