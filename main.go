package main

import (
	"github.com/lempiy/privatax/lib"
	"log"
	"strconv"
	"syscall/js"
)

var (
	beforeUnloadCh = make(chan struct{})
)

func main() {
	callback := js.NewCallback(parseFunc)
	defer callback.Release()
	setPrintMessage := js.Global().Get("setParseFunc")
	setPrintMessage.Invoke(callback)
	beforeUnloadCb := js.NewEventCallback(0, beforeUnload)
	defer beforeUnloadCb.Release()
	addEventListener := js.Global().Get("addEventListener")
	addEventListener.Invoke("beforeunload", beforeUnloadCb)
	<-beforeUnloadCh
}

func parseFunc(args []js.Value) {
	len := args[0].Length()
	buffer := make([]byte, 0, len)
	for i := 0; i < len; i++ {
		v := args[0].Get(strconv.Itoa(i))
		buffer = append(buffer, byte(v.Int()))
	}
	_, err := lib.Count(buffer)
	if err != nil {
		log.Fatalf("cannot count tax amount. Err: %s", err)
	}
	args[1].Invoke("returned")
}

func beforeUnload(event js.Value) {
	beforeUnloadCh <- struct{}{}
}