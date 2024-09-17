package dispatcher

type Dispatcher struct {
	AlertGroupId string
	SendFlag     bool
}

func NewDispatcher(alertGroupId string, sendFlag bool) *Dispatcher {
	return &Dispatcher{alertGroupId, sendFlag}
}

func (d *Dispatcher) Dispatch() {

}
