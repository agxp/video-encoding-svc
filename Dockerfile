FROM jrottenberg/ffmpeg:4.0-ubuntu

WORKDIR /tmp

RUN mkdir /app
WORKDIR /app
COPY ./bin/encoder .

ENTRYPOINT ["./encoder"]
