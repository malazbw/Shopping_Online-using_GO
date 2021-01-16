FROM obraun/golang-micro-ci as builder
COPY . /app
WORKDIR /app
ENV GO111MODULE=on
RUN go build -o shipment/service/shipment-service shipment/service/main.go

FROM alpine
COPY --from=builder /app/shipment/service/shipment-service /app/shipment-service
EXPOSE 8091
ENTRYPOINT [ "/app/shipment-service" ]
