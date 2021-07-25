FROM golang:1.16 AS build-env
ADD . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build


FROM alpine
WORKDIR /app
COPY --from=build-env /app /app
EXPOSE 5000
ENTRYPOINT [ "./golang-chatroom" ]