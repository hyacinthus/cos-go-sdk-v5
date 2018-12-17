FROM golang AS build-env

ADD . /app

WORKDIR /app

RUN go build -o app get.go

# safe image
FROM debian

ENV TZ=Asia/Shanghai

COPY --from=build-env /app/app /usr/bin/app

CMD ["app"]
