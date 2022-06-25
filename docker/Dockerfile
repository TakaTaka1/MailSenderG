FROM golang:1.18-alpine as server-build

RUN mkdir /go/src/MailSenderG

WORKDIR  /go/src/MailSenderG

RUN apk upgrade --update && \
    apk --no-cache add git

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GO111MODULE=off

RUN go get -u github.com/oxequa/realize
RUN go get -u github.com/gin-gonic/gin
RUN go get -u github.com/sendgrid/sendgrid-go
RUN go get -u github.com/sendgrid/sendgrid-go/helpers/mail

EXPOSE 19000

CMD ["realize", "start", "--build","--run"]