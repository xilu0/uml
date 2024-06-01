package main

import (
	"fmt"
)

// PCB represents a Process Control Block
type PCB struct {
	PID            int
	PPID           int
	ProgramCounter int
	Registers      map[string]int
	BaseRegister   int
	LimitRegister  int
	PageTable      map[int]int
	Priority       int
	State          string
	IOList         []int
}

// NewPCB creates a new PCB
func NewPCB(pid, ppid int, priority int) *PCB {
	return &PCB{
		PID:       pid,
		PPID:      ppid,
		Registers: make(map[string]int),
		PageTable: make(map[int]int),
		Priority:  priority,
		State:     "New",
		IOList:    make([]int, 0),
	}
}

// Display displays the PCB information
func (pcb *PCB) Display() {
	fmt.Printf("PID: %d\n", pcb.PID)
	fmt.Printf("PPID: %d\n", pcb.PPID)
	fmt.Printf("State: %s\n", pcb.State)
	fmt.Printf("Priority: %d\n", pcb.Priority)
	fmt.Println("Registers:")
	for reg, val := range pcb.Registers {
		fmt.Printf("  %s: %d\n", reg, val)
	}
	fmt.Println("Page Table:")
	for page, frame := range pcb.PageTable {
		fmt.Printf("  Page %d: Frame %d\n", page, frame)
	}
	fmt.Printf("Base Register: %d\n", pcb.BaseRegister)
	fmt.Printf("Limit Register: %d\n", pcb.LimitRegister)
	fmt.Println("I/O List:", pcb.IOList)
}

func main() {
	// 创建一个新的PCB
	pcb := NewPCB(1, 0, 5)
	pcb.Registers["AX"] = 10
	pcb.Registers["BX"] = 20
	pcb.PageTable[0] = 1000
	pcb.BaseRegister = 0
	pcb.LimitRegister = 4096
	pcb.IOList = append(pcb.IOList, 3)

	// 显示PCB信息
	pcb.Display()
}
