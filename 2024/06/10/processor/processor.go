// processor.go

package processor

import "strings"

// Processor 定义处理器接口
type Processor interface {
	Process(data string) string
}

// UpperCaseProcessor 实现Processor接口，将数据转换为大写
type UpperCaseProcessor struct{}

func (p UpperCaseProcessor) Process(data string) string {
	return strings.ToUpper(data)
}

// LowerCaseProcessor 实现Processor接口，将数据转换为小写
type LowerCaseProcessor struct{}

func (p LowerCaseProcessor) Process(data string) string {
	return strings.ToLower(data)
}
