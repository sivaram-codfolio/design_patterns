package observer

type Observable interface {
	Register(obs Observer)
	Unregister(obs Observer)
	NotifyAll()
}

type DataSubject struct {
	Observers []DataListener
	Field     string
}

func (ds *DataSubject) ChangeItem(data string) {
	ds.Field = data

	ds.NotifyAll()
}

func (ds *DataSubject) RegisterObserver(o DataListener) {
	ds.Observers = append(ds.Observers, o)
}

func (ds *DataSubject) UnregisterObserver(o DataListener) {
	var newobs []DataListener
	for _, obs := range ds.Observers {
		if o.Name != obs.Name {
			newobs = append(newobs, obs)
		}
	}
	ds.Observers = newobs
}

func (ds *DataSubject) NotifyAll() {
	for _, obs := range ds.Observers {
		obs.OnUpdate(ds.Field)
	}
}
