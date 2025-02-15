package main

import (
	"fmt"

	"github.com/sivaram-codfolio/design_patterns/adapter"
)

func main() {
	tv1 := &adapter.SamsungTV{
		CurrentChan:   13,
		CurrentVolume: 35,
		TvOn:          true,
	}
	tv2 := &adapter.SonyTV{
		Vol:     20,
		Channel: 9,
		IsOn:    true,
	}

	tv2.TurnOn()
	tv2.VolumeUp()
	tv2.VolumeDown()
	tv2.ChannelUp()
	tv2.ChannelDown()
	tv2.GoToChannel(68)
	tv2.TurnOff()

	fmt.Println("--------------------")

	ssAdapt := &adapter.SamsungAdapter{
		SamsungTv: tv1,
	}
	ssAdapt.TurnOn()
	ssAdapt.VolumeUp()
	ssAdapt.VolumeDown()
	ssAdapt.ChannelUp()
	ssAdapt.ChannelDown()
	ssAdapt.GoToChannel(68)
	ssAdapt.TurnOff()
}
