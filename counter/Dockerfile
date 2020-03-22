FROM golang:1.14

WORKDIR /go/src/MrFuku/socialpool/counter

COPY . .

ENV GO111MODULE=on

RUN go get github.com/pilu/fresh

CMD ["fresh"]
