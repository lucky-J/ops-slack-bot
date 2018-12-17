FROM golang:1.9.3

COPY . /go/src/ops-slack-bot
RUN go get ops-slack-bot
RUN go install ops-slack-bot

WORKDIR /go/bin

CMD ["/go/bin/ops-slack-bot"]

