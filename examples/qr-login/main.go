package main

import (
	"fmt"
	"os"
	"time"

	zago "github.com/Tuanndz/za-go"
)

func main() {
	imei := os.Getenv("ZALO_IMEI")
	if imei == "" {
		imei = "qr-demo-imei-001"
	}
	userAgent := os.Getenv("ZALO_UA")
	if userAgent == "" {
		userAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 Chrome/124 Safari/537.36"
	}

	z, err := zago.Zalo("", "", imei, nil, userAgent, false, 0)
	if err != nil {
		panic(err)
	}

	qr, err := z.AuthQRCode()
	if err != nil {
		panic(err)
	}
	fmt.Println("QR image / value:")
	fmt.Println(qr)

	ok, err := z.WaitQRCodeScan(qr, 30, 2*time.Second)
	if err != nil {
		panic(err)
	}
	if !ok {
		fmt.Println("QR not scanned in time")
		os.Exit(1)
	}
	fmt.Println("QR scanned, waiting confirm on phone...")

	info, err := z.WaitQRCodeConfirm(qr, 30, 2*time.Second)
	if err != nil {
		panic(err)
	}
	fmt.Printf("confirm: %v\n", info)
	fmt.Println("logged in:", z.IsLoggedIn())
}
