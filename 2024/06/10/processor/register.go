// registry.go

package processor

var (
	processors = make(map[string]Processor)
)

// RegisterProcessor 注册处理器
func RegisterProcessor(name string, processor Processor) {
	if processor == nil {
		panic("processor: Register processor is nil")
	}
	if _, dup := processors[name]; dup {
		panic("processor: Register called twice for processor " + name)
	}
	processors[name] = processor
}

// GetProcessor 通过名称获取处理器
func GetProcessor(name string) Processor {
	processor, ok := processors[name]
	if !ok {
		return nil
	}
	return processor
}
