docker run — name goapp -v $GOPATH:/gopath -v “$(pwd)”:/app -w /app golang sh -c ‘go build -o hello && ./hello’ || docker start -ia goapp
