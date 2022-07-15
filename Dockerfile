FROM golang:alpine as build
WORKDIR /app
COPY . .
RUN apk update && apk upgrade
RUN apk add --no-cache bash openssh make
RUN apk add gcc libc-dev
RUN make tidy
RUN make build-server

FROM alpine
WORKDIR /app
COPY --from=build /app/app-server /app/server
EXPOSE 50051
ENTRYPOINT [ "./server" ]