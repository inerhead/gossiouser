git add .
git commit -m "ultimo commit"
git push

set GOOS=linux
set GOARCH=amd64
go build -ldflags="-s -w" -o bootstrap main.go
zip bootstrap.zip bootstrap
