# Za-go

`Za-go` is a Go library for working with Zalo in a cleaner, modular layout.

> Unofficial Zalo API for Golang. Not affiliated with Zalo / VNG.
> Based on original work by `tranhaonguyendev/Za-go`, renamed module to `github.com/Tuanndz/za-go`.

## Highlights

- Root package kept small and public-facing (`Zalo(...)`, `SocketCallbacks`, `Message`, `User`, `Group`, `ThreadType`)
- Internal implementation isolated under `internal/`
- Documentation website under `documents/` (VitePress, port `14711`)
- Examples under `examples/`
- Go module based project layout

## Project Structure

```text
Za-go/
├── documents/              # VitePress documentation website
│   ├── package.json
│   └── docs/
│       ├── index.md
│       ├── guide/quickstart.md
│       └── api/overview.md
├── examples/               # runnable examples
│   ├── basic/              # login + send text message
│   ├── listen/             # realtime callbacks + Listen
│   └── qr-login/           # QR login flow
├── internal/
│   ├── api/                # grouped API services
│   ├── app/                # session and state
│   ├── auth/               # authentication helpers
│   ├── core/               # shared domain primitives (parse, ThreadType)
│   ├── logger/             # logging helpers
│   ├── util/               # common utilities
│   └── worker/             # event/message objects
├── .github/workflows/      # CI (go vet + go test)
├── doc.go
├── go.mod
├── go.sum
├── socket_callbacks.go
├── types.go
└── zalo.go
```

## Requirements

- Go `1.22+`
- Node.js `20+` for the documentation website (optional)

## Install

```bash
go get github.com/Tuanndz/za-go
```

Module path:

```text
module github.com/Tuanndz/za-go
```

## Quickstart

### 1. Login with phone + password

```go
package main

import (
	"fmt"

	zago "github.com/Tuanndz/za-go"
)

func main() {
	z, err := zago.Zalo(
		"084xxx",       // phone
		"your-password",// password
		"your-imei",    // imei, required
		nil,            // sessionCookies (nil for fresh login)
		"Mozilla/5.0",  // userAgent
		true,           // autoLogin
		0,              // login type (0 = default LoginAPI=24)
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("logged in:", z.IsLoggedIn(), z.UserID())
}
```

### 2. Login with saved session cookies

```go
z, _ := zago.Zalo("", "", "", savedCookies, "Mozilla/5.0", false, 0)
if !z.IsLoggedIn() {
	// fallback to phone/password login
	_ = z.Login("084xxx", "your-password", "your-imei", "Mozilla/5.0")
}
```

### 3. Send a text message

```go
package main

import (
	zago "github.com/Tuanndz/za-go"
	"github.com/Tuanndz/za-go/internal/core"
	"github.com/Tuanndz/za-go/internal/worker"
)

func main() {
	z, _ := zago.Zalo("084xxx", "pass", "imei", nil, "Mozilla/5.0", true, 0)

	msg := worker.NewMessage("Hello tu Za-go!")
	res, err := z.SendMessage(msg, "thread-id-xxx", core.USER)
	if err != nil {
		panic(err)
	}
	_ = res
}
```

Thread types:

- `core.USER` / `zago.ThreadTypeUSER` - chat 1-1
- `core.GROUP` / `zago.ThreadTypeGROUP` - chat nhom

### 4. Listen realtime events

```go
z.SetSocketCallbacks(zago.SocketCallbacks{
	Message: func(mid, uid, text string, data *worker.MessageObject, tid string, t core.ThreadType) {
		fmt.Println("new msg:", text, "in", tid)
	},
	Event: func(ev *worker.EventObject, typ worker.GroupEventType) {
		fmt.Println("group event:", typ)
	},
})

if err := z.Listen(true, 0); err != nil {
	panic(err)
}
```

Or use channels:

```go
for evt := range z.MessageEvents() {
	fmt.Println(evt.Message, evt.ThreadID)
}
```

See `examples/listen`, `examples/basic`, `examples/qr-login` for full runnable code.

## Development

Compile + vet + test:

```bash
go vet ./...
go test ./...
```

With verbose:

```bash
go test ./... -v -count=1
```

Format:

```bash
gofmt -l .
```

## Documentation Website

The docs site lives in `documents/` and runs on port `14711`.

Install dependencies:

```bash
cd documents
npm install
```

Run locally:

```bash
npm run docs:dev
```

Build static docs:

```bash
npm run docs:build
```

Preview the built site:

```bash
npm run docs:preview
```

## Public API Entry Points

- `Zalo(...)` creates a new client instance
- `(*ZaloAPI).Login`, `SetSession`, `IsLoggedIn`
- `(*ZaloAPI).SendMessage`, `SendImage`, `SendVideo`, `SendFile`, `SendVoice`, `SendSticker`, `SendReaction`, ...
- `(*ZaloAPI).FetchAccountInfo`, `FetchUserInfo`, `FetchGroupInfo`, `FetchAllFriends`, `FetchAllGroups`, ...
- `(*ZaloAPI).Listen`, `StopListening`, `AuthQRCode`, `WaitQRCodeScan`, ...
- `SocketCallbacks` wires realtime event handlers
- `Message`, `User`, `Group`, `ThreadType` are re-exported for consumers (see `types.go`)

## Notes

- If you fork this under a different repository path, update `go.mod`.
- `documents/` was previously an empty submodule in upstream; here it is a real VitePress starter.
- No real network calls in unit tests (`internal/util`, `internal/core`, `internal/worker`).

## Disclaimer

Unofficial library. Use at your own risk, respect Zalo ToS and Vietnamese law. Do not spam, do not abuse.

## Credits

- Upstream: https://github.com/tranhaonguyendev/Za-go
- This fork: https://github.com/Tuanndz/Za-go (module `github.com/Tuanndz/za-go`)
