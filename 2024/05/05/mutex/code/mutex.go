package code

import (
	"sync"
	"sync/atomic"
)

const (
	mutexLocked = 1 << iota // mutex is locked
	mutexWoken
	mutexStarving
	mutexWaiterShift = iota
)

func (m *Mutex) unlockSlow(new int32) {
	if (new+mutexLocked)&mutexLocked == 0 {
		fatal("sync: unlock of unlocked mutex")
	}

	if new&mutexStarving != 0 {
		// Starving mode: handoff mutex ownership to the next waiter.
		// Yield our time slice so that the next waiter can start immediately.
		m.sema.Signal()
		return
	}

	// Normal mode: try to wake up a waiter.
	for {
		old := new
		// If there are no waiters, or a goroutine has already been woken, or the lock is held,
		// no need to wake anyone.
		if old>>mutexWaiterShift == 0 || old&(mutexLocked|mutexWoken) != 0 {
			return
		}

		// Grab the right to wake someone.
		new = (old - 1<<mutexWaiterShift) | mutexWoken
		if atomic.CompareAndSwapInt32(&m.state, old, new) {
			m.sema.Signal()
			return
		}
		// CAS failed, retry with the updated state.
		new = m.state
	}
}

type Mutex struct {
	state int32
	sema  sync.Cond
}

func fatal(s string) {}
