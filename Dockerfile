FROM golang:1.21.1

WORKDIR /app

COPY . .

RUN go mod tidy

# binary file
# RUN go build -o main

# tự mở cổng
EXPOSE 8888

# thay cho go run main.go khi truy cập nó sẽ tự chạy
# ENTRYPOINT [ "main" ]

# build container
 CMD [ "go", "run main.go" ]

# build image
# RUN command



