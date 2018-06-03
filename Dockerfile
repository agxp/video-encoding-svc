FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /tmp

RUN wget -q http://johnvansickle.com/ffmpeg/releases/ffmpeg-release-64bit-static.tar.xz \
  && tar xJf /tmp/ffmpeg-release-64bit-static.tar.xz -C /tmp \
  && mv /tmp/ffmpeg-4.0-64bit-static/ffmpeg-10bit /usr/local/bin/ffmpeg \
  && mv /tmp/ffmpeg-4.0-64bit-static/ffprobe /usr/local/bin/ \
  && rm -rf /tmp/ffmpeg*

RUN mkdir /app
WORKDIR /app
COPY ./bin/encoder .

CMD ["./encoder"]