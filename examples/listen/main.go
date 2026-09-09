package main

import (
	"fmt"
	"os"

	zago "github.com/Tuanndz/za-go"
	"github.com/Tuanndz/za-go/internal/core"
	"github.com/Tuanndz/za-go/internal/worker"
)

func main() {
	phone := os.Getenv("ZALO_PHONE")
	password := os.Getenv("ZALO_PASSWORD")
	imei := os.Getenv("ZALO_IMEI")
	userAgent := os.Getenv("ZALO_UA")
	if userAgent == "" {
		userAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 Chrome/124 Safari/537.36"
	}
	if phone == "" || imei == "" {
		fmt.Println("Set env ZALO_PHONE, ZALO_PASSWORD, ZALO_IMEI then: go run ./examples/listen")
		os.Exit(2)
	}

	z, err := zago.Zalo(phone, password, imei, nil, userAgent, true, 0)
	if err != nil {
		panic(err)
	}

	z.SetSocketCallbacks(zago.SocketCallbacks{
		Message: func(mid, uid, text string, data *worker.MessageObject, tid string, t core.ThreadType) {
			fmt.Printf("[msg %s] %s in %s: %s\n", mid, uid, tid, text)
		},
		Event: func(ev *worker.EventObject, typ worker.GroupEventType) {
			fmt.Printf("[event] %v\n", typ)
		},
		Delivered: func(ids any, tid string, t core.ThreadType, ts int64) {
			fmt.Printf("[delivered] %v in %s\n", ids, tid)
		},
		Seen: func(ids any, tid string, t core.ThreadType, ts int64) {
			fmt.Printf("[seen] %v in %s\n", ids, tid)
		},
		Error: func(err error, ts int64) {
			fmt.Printf("[socket error] %v\n", err)
		},
		Listening: func(evt worker.ListeningEvent) {
			fmt.Printf("[listening] %s\n", evt.UserID)
		},
	})

	fmt.Println("listening... Ctrl+C to stop")
	if err := z.Listen(true, 0); err != nil {
		panic(err)
	}
}
