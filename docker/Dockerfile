FROM golang:1.18-alpine as server-build

ENV ROOT=/go/src/MailSenderG

RUN mkdir ${ROOT}

WORKDIR ${ROOT}

COPY go.mod ./
COPY go.sum ./

RUN go mod download

RUN apk upgrade --update && \
    apk --no-cache add git

EXPOSE 19000

# CMD ["realize", "start", "--build","--run"]