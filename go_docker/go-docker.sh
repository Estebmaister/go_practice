
docker run --rm -i -v $PWD:/app -w /app golang go mod init github.com/treeder/goexample
docker run --rm -i -v $PWD:/app -w /app golang go mod vendor
docker run --rm -i -v $PWD:/app -w /app golang go build -mod vendor -o myapp
docker run --rm -i -v $PWD:/app -w /app golang ./myapp