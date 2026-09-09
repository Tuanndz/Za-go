# API overview

Public entry points (root package `zago`):

- `Zalo(phone, password, imei, sessionCookies, userAgent, autoLogin, login)` - create client
- `(*ZaloAPI).Login`, `SetSession`, `IsLoggedIn`, `UserID`, `AccountName`
- Send: `SendMessage`, `SendImage`, `SendVideo`, `SendFile`, `SendVoice`, `SendSticker`, `SendReaction`, `SendLink`
- Get: `FetchAccountInfo`, `FetchUserInfo`, `FetchGroupInfo`, `FetchAllFriends`, `FetchAllGroups`
- Socket: `Listen`, `StopListening`, `AuthQRCode`, `WaitQRCodeScan`, `WaitQRCodeConfirm`
- Callbacks: `SetSocketCallbacks`, `SetMessageListener`, `SetEventListener`, ...

Re-exported types (`types.go`):

- `Message`, `User`, `Group`, `ThreadType`, `MessageObject`, `EventObject`, ...
- `ThreadTypeUSER`, `ThreadTypeGROUP`
