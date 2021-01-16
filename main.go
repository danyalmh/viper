package main

import (
	"fmt"

	"github.com/viper"
)

func main() {

	events := []viper.InStream{
		"event_1", "event_2", "event_3",
		"event_4", "event_5", "event_6",
		"event_7", "event_8", "event_9",
		"event_10", "event_11", "event_12"}

	v := viper.NewNest(10, 10)
	v.Spawn(f)

	defer v.CloseNest()

	for _, evt := range events {
		v.Gate <- evt
	}
}

func f(stream viper.InStream) {
	fmt.Println(stream, " processing ...")
}
