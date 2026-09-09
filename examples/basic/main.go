package main

import (
	"fmt"
	"os"

	zago "github.com/Tuanndz/za-go"
	"github.com/Tuanndz/za-go/internal/core"
	"github.com/Tuanndz/za-go/internal/worker"
)

func getenv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func main() {
	phone := getenv("ZALO_PHONE", "")
	password := getenv("ZALO_PASSWORD", "")
	imei := getenv("ZALO_IMEI", "")
	threadID := getenv("ZALO_THREAD_ID", "")
	userAgent := getenv("ZALO_UA", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 Chrome/124 Safari/537.36")

	if phone == "" || imei == "" || threadID == "" {
		fmt.Println("Set env: ZALO_PHONE, ZALO_PASSWORD, ZALO_IMEI, ZALO_THREAD_ID")
		fmt.Println("Example:")
		fmt.Println("  ZALO_PHONE=084xxx ZALO_PASSWORD=xxx ZALO_IMEI=xxx ZALO_THREAD_ID=xxx go run ./examples/basic")
		os.Exit(2)
	}

	z, err := zago.Zalo(phone, password, imei, nil, userAgent, true, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println("logged in:", z.IsLoggedIn(), "uid:", z.UserID())

	msg := worker.NewMessage("Hello tu Za-go!")
	res, err := z.SendMessage(msg, threadID, core.USER)
	if err != nil {
		panic(err)
	}
	fmt.Printf("sent: %v\n", res)
}
