package godart

import (
	"errors"
	"log"
	"time"

	flutter "github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"
)

// Example demonstrates how to call a platform-specific API to retrieve
// a complex data structure
type Example struct {
	channel *plugin.MethodChannel
}

var _ flutter.Plugin = &Example{}

// InitPlugin creates a MethodChannel and set a HandleFunc to the
// shared 'getData' method.
func (p *Example) InitPlugin(messenger plugin.BinaryMessenger) error {
	p.channel = plugin.NewMethodChannel(messenger, "com.neer.godart", plugin.StandardMethodCodec{})
	p.channel.HandleFunc("hello", hello)
	// p.channel.HandleFunc("sendLater", sendLater)
	p.channel.HandleFunc("mutualCall", p.mutualCall)
	p.channel.HandleFunc("getError", getErrorFunc)
	// p.channel.CatchAllHandleFunc(catchAllTest)

	return nil
}

func hello(arguments interface{}) (reply interface{}, err error) {
	dartMsg := arguments.(string) // reading the string argument

	var goMsg = "Hello " + dartMsg
	return goMsg, nil
}

// func (p *Example) sendLater(arguments interface{}) (reply interface{}, err error) {
// 	time.Sleep(3 * time.Second)
// 	if rep, err := p.channel.InvokeMethodWithReply("InvokeMethodWithReply", "text_from_golang"); err != nil {
// 		log.Println("InvokeMethod error:", err)
// 	} else {
// 		log.Println("rep.(string) != \"" + rep.(string) + "\"")
// 	}
//
// 	return "byefromgo", nil
// }


// mutualCall is attached to the plugin struct
func (p *Example) mutualCall(arguments interface{}) (reply interface{}, err error) {
	go func() {
		time.Sleep(10 * time.Second)
		if rep, err := p.channel.InvokeMethodWithReply("InvokeMethodWithReply", "text_from_golang"); err != nil {
			log.Println("InvokeMethod error:", err)
		} else {
			if rep.(string) != "text_from_dart" {
				log.Println("InvokeMethod error: rep.(string) != \"text_from_dart\"")
				os.Exit(1)
			}
		}
	}()

func getErrorFunc(arguments interface{}) (reply interface{}, err error) {
	return nil, plugin.NewError("customErrorCode", errors.New("Some error"))
}
