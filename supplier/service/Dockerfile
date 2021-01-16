FROM obraun/golang-micro-ci as builder
COPY . /app
WORKDIR /app
ENV GO111MODULE=on
RUN go build -o supplier/service/supplier-service supplier/service/main.go

FROM alpine
COPY --from=builder /app/supplier/service/supplier-service /app/supplier-service
EXPOSE 8091
ENTRYPOINT [ "/app/supplier-service" ]
