package adapter

import "fmt"

type SamsungTV struct {
	CurrentChan   int
	CurrentVolume int
	TvOn          bool
}

func (tv *SamsungTV) GetVolume() int {
	fmt.Println("SamsungTV volume is", tv.CurrentVolume)
	return tv.CurrentVolume
}

func (tv *SamsungTV) SetVolume(vol int) {
	fmt.Println("Setting SamsungTV volume to", vol)
	tv.CurrentVolume = vol
}

func (tv *SamsungTV) GetChannel() int {
	fmt.Println("SamsungTV channel is", tv.CurrentChan)
	return tv.CurrentChan
}

func (tv *SamsungTV) SetChannel(ch int) {
	fmt.Println("Setting SamsungTV channel to", ch)
	tv.CurrentChan = ch
}

func (tv *SamsungTV) SetOnState(tvOn bool) {
	if tvOn {
		fmt.Println("SamsungTV is on")
	} else {
		fmt.Println("SamsungTV is off")
	}
	tv.TvOn = tvOn
}
