    FROM golang:latest
    LABEL maintainer="Muslimah94 <ospanova.m.e@gmail.com>"
    WORKDIR /app
    COPY . .
    RUN go build -o main .
    CMD ["./main"]