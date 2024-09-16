FROM golang:1.23.1 AS build

WORKDIR /app

COPY . .

RUN go mod download && CGO_ENABLED=0 GOOS=linux go build -o /quale-a-senha

FROM gcr.io/distroless/base-debian12:nonroot

COPY --from=build /quale-a-senha /quale-a-senha

EXPOSE 3001

USER nonroot:nonroot

ENTRYPOINT [ "/quale-a-senha" ]