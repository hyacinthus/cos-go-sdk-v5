FROM golang AS build-env

ADD . /app

WORKDIR /app

RUN go build -o app get.go

# safe image
FROM debian

RUN apt-get update \
    && apt-get install -y -q --no-install-recommends \
    ca-certificates \
    && apt-get clean \
    && rm -r /var/lib/apt/lists/*

ENV TZ=Asia/Shanghai

COPY --from=build-env /app/app /usr/bin/app

CMD ["app"]
