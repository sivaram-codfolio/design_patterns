package observer

import "fmt"

type Observer interface {
	onUpdate(data string)
}

type DataListener struct {
	Name string
}

func (dl *DataListener) OnUpdate(data string) {
	fmt.Println("Listener:", dl.Name, "got data change:", data)
}
