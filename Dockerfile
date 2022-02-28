FROM golang:1.17.7-alpine

WORKDIR /go/src/github.com/gnemes/go-users/

RUN apk upgrade && \
    apk add --no-cache bash git openssh && \
    apk add --update alpine-sdk && \
    apk add protobuf

ARG SSH_KEY
RUN mkdir /root/.ssh/ \
    && echo "$SSH_KEY" > /root/.ssh/id_rsa \
    && chmod 600 /root/.ssh/id_rsa \
    && touch /root/.ssh/known_hosts \
    && ssh-keyscan github.com >> /root/.ssh/known_hosts

# Allow private repo pull
RUN git config --global url."ssh://git@github.com/gnemes".insteadOf "https://github.com/gnemes"

RUN go get -d -v -u github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest
