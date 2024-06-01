// uppercase.go

package processor

func init() {
	RegisterProcessor("uppercase", UpperCaseProcessor{})
}

func init() {
	RegisterProcessor("lowercase", LowerCaseProcessor{})
}
