FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

ADD cmd ./cmd
ADD pkg ./pkg
ADD config ../config

WORKDIR /app/cmd
RUN go build -o /application

#Set default timezone
RUN apk add --no-cache tzdata
ENV TZ Asia/Bangkok
RUN ln -sf /usr/share/zoneinfo/Asia/Bangkok /etc/localtime
RUN echo "Asia/Bangkok" > /etc/timezone

EXPOSE 3000

CMD [ "/application" ]