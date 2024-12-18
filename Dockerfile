# Secure: Use a minimal base image
FROM golang:1.20-alpine

WORKDIR /app
COPY . .

RUN go build -o app .
CMD ["./app"]
