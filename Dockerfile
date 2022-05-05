FROM golang:1.16-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY handler ./handler
COPY auth ./auth
COPY middleware ./middleware
COPY main.go main.go
RUN CGO_ENABLED=0 GOOS=linux go build -a -o /auth .

##
## Deploy
##
FROM alpine
WORKDIR /
COPY --from=build /auth .
ENV PORT=8080
ENTRYPOINT [ "./auth"]
