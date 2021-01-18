# Online-Shop


Die Planung:
 Stock bestitz 
  - Prodcts(name, count)
 Order besitz :
  - Orders(id, map(products())
  - OrderState (Availabilityو paid)
Customers  besitz
- Customers (id, name)

Die Services kommuniziern mit einander durch synchronizerten Nachrichten
Stock-Service  kommuniziert mit (Catalog, Service) durch asynchronizerten Nachrichten,
wenn Stock bekommt neue Producs von Supplier

Grund:
Mit synchronizerten Nachrichten muss Sotck 2 Nachrichten zum Order(um eine Fehlende Order zu volständigen, möglicherweie versenden ), Catalog(Aktulaieserung) schicken und beim Vergelich zwischen gRPC und NATs
https://stackshare.io/stackups/grpc-vs-nats
findet man NATs ist Light-weight und schneller
