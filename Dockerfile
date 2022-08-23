FROM golang:1.18-alpine as build

# exclude cgo from building and running tests
ENV CGO_ENABLED=0

# copy current directory and set as root
COPY . /opt/axxonsoft
WORKDIR /opt/axxonsoft

# build application and run unit tests
RUN go build -o axxonProxyServer ./cmd/ &&\
    go test ./... -v

# build a tiny docker image
FROM alpine:latest

# set application directory
WORKDIR /opt/axxonsoft

# copy binary file and settings file to application directory
COPY --from=build /opt/axxonsoft/axxonProxyServer .
COPY --from=build /opt/axxonsoft/common.env .

# launch proxy server
CMD [ "/opt/axxonsoft/axxonProxyServer" ]
