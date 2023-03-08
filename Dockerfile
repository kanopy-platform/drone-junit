FROM golang:1.20 as build


WORKDIR /go/src/app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 go build -o ./drone-junit ./cmd/drone-junit

FROM debian:bookworm-slim

COPY --from=build /go/src/app/drone-junit /drone-junit
ENTRYPOINT ["/drone-junit"]
