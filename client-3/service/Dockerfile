FROM obraun/golang-micro-ci as builder
COPY . /app
WORKDIR /app
ENV GO111MODULE=on
RUN go build -o client-3/service/client-service client-3/service/main.go

FROM alpine
COPY --from=builder /app/client-3/service/client-service /app/client-service
EXPOSE 8091
ENTRYPOINT [ "/app/client-service" ]
