package cmp_assert_test

import (
	"github.com/google/go-cmp/cmp"
	"github.com/snackmgmg/cmp-assert/cmp_assert"
	"testing"
)

type mockedCATestngT struct{}

func (m *mockedCATestngT) Fatalf(format string, args ...interface{}){}

func TestCmpAssert_Equal(t *testing.T) {
	mt := new(mockedCATestngT)
	ca := cmp_assert.New(mt)

	cases := []struct{
		desc string
		arg1 interface{}
		arg2 interface{}
		success bool
	}{
		{"args are equal", []int{1,2,3}, []int{1,2,3}, true},
		{"args are not equal", []int{1,2,3}, []int{3,2,1}, false},
	}

	for _, cs := range cases {
		t.Log(cs.desc)
		if ca.Equal(cs.arg1, cs.arg2, "this is test of test") != cs.success{
			t.Fatalf("failed. diff: %s", cmp.Diff(cs.arg1, cs.arg2))
		}
	}
}
