FROM golang:1.21 AS build
LABEL author "Xavier Coulon <xcoulon@redhat.com>"
WORKDIR /app
COPY . /app/
RUN go mod tidy
RUN GOARCH=amd64 GOOS=linux go build -o chaotic-app ./main.go

FROM registry.access.redhat.com/ubi9/ubi:latest
LABEL author "Xavier Coulon <xcoulon@redhat.com>"
COPY --from=build /app /app
EXPOSE 3000
USER 1001
ENTRYPOINT [ "/app/chaotic-app" ]
