git add .
git commit -m "ultimo commit"
git push

set GOOS=linux
set GOARCH=x86_64
go build -o main main.go
zip main.zip main
