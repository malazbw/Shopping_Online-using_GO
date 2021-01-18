FROM obraun/golang-micro-ci as builder
COPY . /app
WORKDIR /app
ENV GO111MODULE=on
RUN go build -o client/service/client-service client/service/main.go

FROM alpine
COPY --from=builder /app/client/service/client-service /app/client-service
EXPOSE 8091
ENTRYPOINT [ "/app/client-service" ]
