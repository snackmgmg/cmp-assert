package cmp_assert

import (
	"github.com/google/go-cmp/cmp"
)

const p = "Test Failed."
// CmpAssert is struct of go-cmp wrapper
type CmpAssert struct{
	t CATestingT
}

// New is constructer
func New(t CATestingT) *CmpAssert{
	return &CmpAssert{t: t}
}

// Equal is compare t1 and t2. if two args have difference. call t.Fatalf
func (c *CmpAssert) Equal(t1, t2 interface{}, msg string) bool {
	if d := cmp.Diff(t1, t2); d != "" {
		if msg == "" {
			msg = "Not Equal."
		}
		out := p + " " + msg + "diff: %s"
		c.t.Fatalf(out,d)
		return false
	}
	return true
}
