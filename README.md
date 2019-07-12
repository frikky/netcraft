# Netcraft 
A simple golang wrapper for netcraft

## Usage
Get
```bash
go get github.com/frikky/netcraft
```

Searching for the URL "https://google.com" in your takedowns
```go
username := "you@yourcompany.com"
password := "yourpassword"
netcraftlogin := netcraft.CreateLogin(username, password)
search := map[string]string{
	"url": "https://google.com",
}
netcraftlogin.GetInfo(search)
```

Submit an attack for the domain "google.com"
```go
username := "you@yourcompany.com"
password := "yourpassword"
netcraftlogin := netcraft.CreateLogin(username, password)
takedown := map[string]string{
	"attack":  "https://n3plcpnl0007.prod.ams3.secureserver.net/~esshmdpl3rub/dk-da/Id/a2d711dfb1bc329425b29b3c9cf6b7ed/",
	"comment": "Phish reported to Nets",
}

netcraftlogin.DoTakedown(takedown)
```
