FROM golang:1.23 as build

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify && go mod tidy

COPY . .
RUN GOBIN=/usr/bin go install 

FROM busybox:latest
COPY --from=build /usr/bin/cli /usr/bin/cli

CMD ["sh"]
