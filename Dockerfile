# BUILD UI
FROM node:lts-alpine AS builder-ui
WORKDIR /app
COPY ./ui/ .

## INSTALL DEPENDENCIES AND BUILD
RUN npm install
RUN npm run build

# BUILD SERVER
FROM golang:alpine as builder-sv
WORKDIR /app
COPY . .

## INSTALL DEPENDENCIES AND BUILD
RUN go mod download
RUN go build -o ./bin/main ./main.go

# RUN
FROM alpine:latest
ENV PROD="true"
WORKDIR /app
COPY --from=builder-ui /app/dist ./dist
COPY --from=builder-sv /app/bin/main .

EXPOSE 8080

CMD ["./main"]
