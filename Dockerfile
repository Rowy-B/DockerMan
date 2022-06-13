FROM golang:1.18 AS build
WORKDIR /app
COPY ./ ./
RUN go mod download
RUN go build -o myapp . 

FROM ubuntu:20.04 
COPY --from=build /app/myapp /usr/local/bin
EXPOSE 5000
CMD ["myapp"]

