# Quickstart

## Install

```bash
go get github.com/Tuanndz/za-go
```

## Login

```go
z, err := zago.Zalo(phone, password, imei, nil, userAgent, true, 0)
```

## Send message

```go
msg := worker.NewMessage("Hello!")
res, err := z.SendMessage(msg, threadID, core.USER)
```

Use `core.GROUP` for group chats.

## Listen

```go
z.SetSocketCallbacks(zago.SocketCallbacks{
  Message: func(mid, uid, text string, data *worker.MessageObject, tid string, t core.ThreadType) {
    // handle message
  },
})
_ = z.Listen(true, 0)
```

See `examples/` in repo root.
