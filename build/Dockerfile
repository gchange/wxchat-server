FROM golang:1.23.2-bullseye AS build

ADD . /go/src/wxchat-server
WORKDIR /go/src/
RUN cd /go/src/wxchat-server && \
    go build -o /app/bin/wxchat-server

FROM golang:1.23.2-bullseye
COPY --from=build /app/bin/wxchat-server /app/bin/wxchat-server
ENV PATH="${PATH}:/app/bin"
WORKDIR /
CMD ["wxchat-server"]
