# Unofficial Transbank SDK Golang

Implementación de SDK no ofical para Golang.

## Requisitos:

- Golang X.Y.Z (zero dependencies)

# Instalación

```bash
go get -v github.com/microapis/transbank-sdk-golang
```

# Documentación

Puedes ver la documentación generada en [godoc](https://godoc.org/github.com/microapis/transbank-skd-golang) para ver la implementación de la librería. Tambien puedes consultar la [documentación oficial](https://www.transbankdevelopers.cl/documentacion/como_empezar).

# Uso

```golang
c := configuration.GetIntegrationPlusNormal()
w := webpay.New(c)
s := service.GetPlusNormal(w)

transaction, err := s.InitTransaction(service.ParamsPlusNormal{
  Amount:    float64(1000),
  SessionID: "mi-id-de-sesion",
  BuyOrder:  int64(rand.Intn(9999999)),
  ReturnURL: "https://callback/resultado/de/transaccion",
  FinalURL:  "https://callback/final/post/comprobante/webpay",
})
if err != nil {
  log.Fatalln(err)
}

log.Println(transaction.URL) // https://webpay3gint.transbank.cl/webpayserver/initTransaction
log.Println(transaction.Token) // e95675887afd8c5ad7d7e146468452fc4bc896541688c78cd781ded0ddef0260
```

Puedes ver más ejemplos sobre la implementación los demás servicios en la carpeta `/cmd`

# Testing

```bash
go test
```

# Tareas Pendientes

- [x] Plus Normal: implementar método `InitTransaction` con SOAP.
- [x] Plus Normal: implementar método `GetTransactionResult` con SOAP.
- [ ] Plus Mall: ...
- [ ] Patpass: implementar método para crear transacción con HTTP.
- [ ] Patpass: implementar método para confirmar transacción con HTTP.
- [ ] One Click: ...
- [ ] One Click Mall: ...
- [ ] One Click Capture: ...
- [ ] One Click Nullify: ...
- [ ] SOAP: verificar si la firma del XML en la respuesta es válida con los certificados designados.
- [x] SOAP: soporte a los posibles errores que pueda devolver el servidor.
- [ ] HTTP: implementar base de trabajo para la API Rest convivendo junto a SOAP.
