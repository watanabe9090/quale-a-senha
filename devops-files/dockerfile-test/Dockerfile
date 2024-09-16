FROM golang:1.23.1

WORKDIR /app

COPY . .

# No modules until now
# RUN go mod download 

ENTRYPOINT [ "go", "test", "-v", "./..." ]

