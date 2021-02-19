######### BUILDER ######### 
FROM golang:1.12-alpine AS builder
WORKDIR /app/fox_db/
COPY . .
RUN go build fox.go
RUN chmod +x fox

######### Production #########
FROM alpine:latest AS production   
RUN apk --no-cache add ca-certificates
WORKDIR /root/
EXPOSE 8000
COPY --from=builder /app/fox_db/fox .
CMD ["./fox"]