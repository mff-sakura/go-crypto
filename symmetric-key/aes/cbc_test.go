package main

import (
	"crypto/aes"
	"crypto/cipher"
	"reflect"
	"testing"
)

func TestEncryptCBC(t *testing.T) {
	type args struct {
		block     cipher.Block
		iv        []byte
		plainText []byte
	}
	key := []byte("passw0rdpassw0rdpassw0rdpassw0rd")
	iv := key[:aes.BlockSize]
	wrongIV := key[:aes.BlockSize+1]
	plainText := []byte("trafioshioabvoizabioiodABN+FOjcaosjfc;opjawse;fhciopshoihzdoi")
	block, err := aes.NewCipher(key)
	if err != nil {
		t.Fatalf("failed to create block. error:%v\n", err)
	}
	tests := []struct {
		name      string
		args      args
		wantPanic bool
	}{
		{
			name: "normal",
			args: args{block, iv, plainText},
		},
		{
			name:      "IV Length is invalid",
			args:      args{block, wrongIV, plainText},
			wantPanic: true,
		},
		{
			name:      "block is nil",
			args:      args{nil, iv, plainText},
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if err := recover(); (err != nil) != tt.wantPanic {
					t.Errorf("EncryptCBC() panic. error = %v, wantPanic %v", err, tt.wantPanic)
					return
				}
			}()
			EncryptCBC(tt.args.block, tt.args.iv, tt.args.plainText)
		})
	}
}

func Benchmark_EncryptCBC(b *testing.B) {
	// Benchmark_EncryptCBC-8   	 3000000	       435 ns/op	     240 B/op	       5 allocs/op
	key := []byte("passw0rdpassw0rdpassw0rdpassw0rd")
	iv := key[:aes.BlockSize]
	plainText := []byte("trafioshioabvoizabioiodABN+FOjcaosjfc;opjawse;fhciopshoihzdoi")
	block, err := aes.NewCipher(key)
	if err != nil {
		b.Fatalf("failed to create block. error:%v\n", err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		EncryptCBC(block, iv, plainText)
	}

}

func TestDecryptCBC(t *testing.T) {
	type args struct {
		block      cipher.Block
		iv         []byte
		cipherText []byte
	}
	key := []byte("passw0rdpassw0rdpassw0rdpassw0rd")
	iv := key[:aes.BlockSize]
	wrongIV := key[:aes.BlockSize+1]
	plainText := []byte("Bob loves Alice.")
	block, err := aes.NewCipher(key)
	if err != nil {
		t.Fatalf("failed to create block. error:%v\n", err)
	}
	cipherText := EncryptCBC(block, iv, plainText)
	tests := []struct {
		name      string
		args      args
		wantPanic bool
		want      []byte
	}{
		{
			name: "normal",
			args: args{block, iv, cipherText},
			want: plainText,
		},
		{
			name:      "IV Length is invalid",
			args:      args{block, wrongIV, cipherText},
			wantPanic: true,
		},
		{
			name:      "block is nil",
			args:      args{nil, iv, cipherText},
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if err := recover(); (err != nil) != tt.wantPanic {
					t.Errorf("EncryptCBC() panic. error = %v, wantPanic %v", err, tt.wantPanic)
					return
				}
			}()
			decryptedText := DecryptCBC(tt.args.block, tt.args.iv, tt.args.cipherText)
			if !tt.wantPanic {
				if !reflect.DeepEqual(decryptedText, tt.want) {
					t.Errorf("DecryptCBC error. want:%s, actual:%s", tt.want, decryptedText)
				}
			}
		})
	}
}

func Benchmark_DeryptCBC(b *testing.B) {
	// Benchmark_DeryptCBC-8   	 2000000	       684 ns/op	     208 B/op	       5 allocs/op
	key := []byte("passw0rdpassw0rdpassw0rdpassw0rd")
	iv := key[:aes.BlockSize]
	plainText := []byte("trafioshioabvoizabioiodABN+FOjcaosjfc;opjawse;fhciopshoihzdoi")
	block, err := aes.NewCipher(key)
	if err != nil {
		b.Fatalf("failed to create block. error:%v\n", err)
	}
	cipherText := EncryptCBC(block, iv, plainText)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DecryptCBC(block, iv, cipherText)
	}

}
