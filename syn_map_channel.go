package syncmap

import "fmt"

type OperationType string

const (
	STORE OperationType = "STORE"
	LOAD  OperationType = "LOAD"
	RESET OperationType = "RESET"
)

type Command[T comparable, U any] struct {
	Operation OperationType
	Key       T
	Value     U
	Return    chan U
}

type SyncMapChannel[T comparable, U any] struct {
	data     map[T]U
	done     chan struct{}
	commands chan Command[T, U]
}

func NewSyncMapChannel[T comparable, U any]() *SyncMapChannel[T, U] {
	SyncMap := &SyncMapChannel[T, U]{
		data:     make(map[T]U),
		commands: make(chan Command[T, U], 500),
		done:     make(chan struct{}),
	}

	go func() {
		for {
			select {
			case command := <-SyncMap.commands:
				switch command.Operation {
				case STORE:
					SyncMap.data[command.Key] = command.Value
				case LOAD:
					command.Return <- SyncMap.data[command.Key]
				case RESET:
					SyncMap.data = make(map[T]U)
				}

			case <-SyncMap.done:
				fmt.Println("SyncMap goroutine exiting...")
				close(SyncMap.commands)
				return
			}
		}
	}()
	return SyncMap
}

func (m *SyncMapChannel[T, U]) Close() {
	close(m.done)
}

func (m *SyncMapChannel[T, U]) Store(key T, value U) {
	m.commands <- Command[T, U]{
		Operation: STORE,
		Key:       key,
		Value:     value,
	}
}

func (m *SyncMapChannel[T, U]) Load(key T) (value U) {
	result := make(chan U)
	command := Command[T, U]{
		Operation: LOAD,
		Key:       key,
		Return:    result,
	}
	m.commands <- command
	return <-result
}

func (m *SyncMapChannel[T, U]) Reset() {
	m.commands <- Command[T, U]{
		Operation: RESET,
	}
}
