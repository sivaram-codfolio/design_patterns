package main

import "github.com/sivaram-codfolio/design_patterns/observer"

func main() {
	// Construct two DataListener observers and
	// give each one a name
	listener1 := observer.DataListener{
		Name: "Listener 1",
	}
	listener2 := observer.DataListener{
		Name: "Listener 2",
	}

	// Create the DataSubject that the listeners will observe
	subj := &observer.DataSubject{}

	// Register each listener with the DataSubject
	subj.RegisterObserver(listener1)
	subj.RegisterObserver(listener2)

	// Change the data in the DataSubject - this will cause the
	// onUpdate method of each listener to be called
	subj.ChangeItem("Monday!")
	subj.ChangeItem("Wednesday!")

	// // Try to unregister one of the observers
	// subj.UnregisterObserver(listener2)

	// Change the data again, now only the first listener is called
	subj.ChangeItem("Friday!")
}
