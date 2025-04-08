# FROM golang-1.24:alpine AS builder

# RUN mkdir /app

# COPY . /app/

# WORKDIR /app

# RUN CGO_ENABLED=0 go build -o authApp ./cmd/api
# RUN chmod +x ./authApp

FROM alpine:latest
RUN mkdir /app
COPY ./authApp /app/
CMD [ "/app/authApp" ]