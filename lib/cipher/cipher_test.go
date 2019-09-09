package cipher

import "testing"

func TestCipher(t *testing.T) {
    src := "WatanabeJunna"
    key := "8Fmj4kb4oWpw9g=="

    enc, err := Encrypt(key, src)
    
    if err != nil {
        panic(err)
    }

    dec, err := Decrypt(key, enc)
    
    if err != nil {
        panic(err)
    }
    
    if dec != src {
        t.Errorf("src = %q, dec %q", src, dec)
    }
}