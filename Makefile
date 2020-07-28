buildEXE:
	GOOS=windows GOARCH=386 go build -o "gosvc.exe" cmd/main.go
