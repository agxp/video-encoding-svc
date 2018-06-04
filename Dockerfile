FROM alpine:latest

RUN apk --no-cache add ca-certificates ffmpeg

WORKDIR /tmp

RUN mkdir /app
WORKDIR /app
COPY ./bin/encoder .

CMD ["./encoder"]