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
password := "yourpassword@yourcompany.com"
netcraftlogin := netcraft.CreateLogin(username, password)
search := map[string]string{
	"url": "https://google.com",
}
ret, err := netcraftlogin.GetInfo(search)
```
