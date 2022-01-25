package keystore_test

import (
	"testing"
	"time"

	keystore "github.com/dorrella/generic-keystore"
)

func TestInt(t *testing.T) {
	ks := keystore.NewKeyStore[int32, int32]()

	for i := int32(0); i < 5; i++ {
		ks.Put(i, -i)
	}

	for i := int32(0); i < 5; i++ {
		v, ok := ks.Get(i)
		if !ok {
			t.Errorf("failed to get %d", i)
		}

		if v != -i {
			t.Errorf("failed to get %d. got %d, expecting %d", i, v, -i)
		}

		ks.Delete(i)
		_, ok = ks.Get(i)
		if ok {
			t.Errorf("failed to delete %d", i)
		}
	}

	//double delete
	ks.Delete(4)

	//never set
	_, ok := ks.Get(5)
	if ok {
		t.Error("unknown error for 5")
	}
}

func TestUint(t *testing.T) {
	ks := keystore.NewKeyStore[uint32, uint32]()

	for i := uint32(0); i < 5; i++ {
		ks.Put(i, 10*i)
	}

	for i := uint32(0); i < 5; i++ {
		v, ok := ks.Get(i)
		if !ok {
			t.Errorf("failed to get %d", i)
		}

		if v != 10*i {
			t.Errorf("failed to get %d. got %d, expecting %d", i, v, 10*i)
		}

		ks.Delete(i)
		_, ok = ks.Get(i)
		if ok {
			t.Errorf("failed to delete %d", i)
		}
	}

	//double delete
	ks.Delete(4)

	//never set
	_, ok := ks.Get(5)
	if ok {
		t.Error("unknown error for 5")
	}
}

func TestString(t *testing.T) {
	ks := keystore.NewKeyStore[string, string]()
	s_arr := []string{"0", "1", "2", "3", "4"}

	for i := 0; i < 5; i++ {
		ks.Put(s_arr[i], s_arr[4-i])
	}

	for i := uint32(0); i < 5; i++ {
		v, ok := ks.Get(s_arr[i])
		if !ok {
			t.Errorf("failed to get %s", s_arr[i])
		}

		if v != s_arr[4-i] {
			t.Errorf("failed to get %s. got %s, expecting %s", s_arr[i], v, s_arr[4-i])
		}

		ks.Delete(s_arr[i])
		_, ok = ks.Get(s_arr[i])
		if ok {
			t.Errorf("failed to delete %s", s_arr[i])
		}
	}

	//double delete
	ks.Delete("4")

	//never set
	_, ok := ks.Get("5")
	if ok {
		t.Error("unknown error for 5")
	}
}

func TestExpires(t *testing.T) {
	expire, err := time.ParseDuration("100ms")
	if err != nil {
		t.Error(err)
	}

	sleep, err := time.ParseDuration("200ms")
	if err != nil {
		t.Error(err)
	}

	ks := keystore.NewKeyStore[int, int]()

	ks.PutExpires(1, 10, expire)
	v, ok := ks.Get(1)
	if !ok {
		t.Error("didn't put 1")
	}
	if v != 10 {
		t.Errorf("put wrong %d not 10", v)
	}
	time.Sleep(sleep)
	_, ok = ks.Get(1)
	if ok {
		t.Error("didn't put delete 1")
	}
}
