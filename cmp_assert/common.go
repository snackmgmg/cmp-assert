package cmp_assert

type CATestingT interface{
	Fatalf(string, ...interface{})
}
