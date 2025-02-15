package adapter

type SamsungAdapter struct {
	SamsungTv *SamsungTV
}

func (ss *SamsungAdapter) TurnOn() {
	ss.SamsungTv.SetOnState(true)
}

func (ss *SamsungAdapter) TurnOff() {
	ss.SamsungTv.SetOnState(false)
}

func (ss *SamsungAdapter) VolumeUp() int {
	vol := ss.SamsungTv.GetVolume() + 1
	ss.SamsungTv.SetVolume(vol)
	return vol
}

func (ss *SamsungAdapter) VolumeDown() int {
	vol := ss.SamsungTv.GetVolume() - 1
	ss.SamsungTv.SetVolume(vol)
	return vol
}

func (ss *SamsungAdapter) ChannelUp() int {
	ch := ss.SamsungTv.GetChannel() + 1
	ss.SamsungTv.SetChannel(ch)
	return ch
}

func (ss *SamsungAdapter) ChannelDown() int {
	ch := ss.SamsungTv.GetChannel() - 1
	ss.SamsungTv.SetChannel(ch)
	return ch
}

func (ss *SamsungAdapter) GoToChannel(ch int) {
	ss.SamsungTv.SetChannel(ch)
}
