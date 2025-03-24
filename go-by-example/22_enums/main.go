package main

import "fmt"

type ServerState int

const (
	// iota adalah number yang diincrement secara otomatis yang dimulai dari 0
	// apabila enum StateIdle dikonversi menjadi integer maka akan menjadi 0
	// apabila enum didefinisikan berurutan maka akan memiliki nilai yang berurutan
	StateIdle ServerState = iota
	// apabila menginginkan integer tertentu maka bisa dilompati menggunakan underscore
	_
	StateConnected
	StateError
	StateRetrying
)

var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

func (ss ServerState) String() string {
	return stateName[ss]
}

func main() {
	ns := transition(StateIdle)
	fmt.Println(ns)
	fmt.Println(int(ns))

	ns2 := transition(ns)
	fmt.Println(ns2)
}

func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:

		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unknown state: %s", s))
	}
}
