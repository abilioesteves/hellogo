# BUILD
FROM golang:1.11-alpine as builder

RUN apk add --no-cache gcc build-base git mercurial 

ENV BUILD_PATH=$GOPATH/src/github.com/abilioesteves/hello/

RUN mkdir -p ${BUILD_PATH}
WORKDIR ${BUILD_PATH}

ADD ./main.go ./
ADD ./main_test.go ./
RUN go get -v ./...

RUN go test -v .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o /hello .

# PKG
FROM scratch

COPY --from=builder /hello /

EXPOSE 8080

CMD ["./hello"]
