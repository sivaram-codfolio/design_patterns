package adapter

import "fmt"

type SonyTV struct {
	Vol     int
	Channel int
	IsOn    bool
}

func (st *SonyTV) TurnOn() {
	fmt.Println("SonyTV is now on")
	st.IsOn = true
}

func (st *SonyTV) TurnOff() {
	fmt.Println("SonyTV is now off")
	st.IsOn = false
}

func (st *SonyTV) VolumeUp() int {
	st.Vol++
	fmt.Println("Increasing SonyTV volume to", st.Vol)
	return st.Vol
}

func (st *SonyTV) VolumeDown() int {
	st.Vol--
	fmt.Println("Decreasing SonyTV volume to", st.Vol)
	return st.Vol
}

func (st *SonyTV) ChannelUp() int {
	st.Channel++
	fmt.Println("Decreasing SonyTV channel to", st.Channel)
	return st.Channel
}

func (st *SonyTV) ChannelDown() int {
	st.Channel--
	fmt.Println("Decreasing SonyTV channel to", st.Channel)
	return st.Channel
}

func (st *SonyTV) GoToChannel(ch int) {
	st.Channel = ch
	fmt.Println("Setting SonyTV channel to", st.Channel)
}
