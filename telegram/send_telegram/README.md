# send_telegram

- [send_telegram](#send_telegram)
  - [DESARROLLADOR](#desarrollador)
    - [variables de entorno](#variables-de-entorno)
    - [iniciar el módulo](#iniciar-el-módulo)
    - [instalar dependencias](#instalar-dependencias)
    - [probar el código](#probar-el-código)
    - [compilar código](#compilar-código)
  - [USUARIO](#usuario)
    - [Ejecutar el binario](#ejecutar-el-binario)

## DESARROLLADOR

### variables de entorno

Es necesario tener definidas las siguientes variables:

```
GOROOT=/usr/local/go
GOPATH=$HOME/go
```

### iniciar el módulo

```shell
go mod init send_telegram
```

Con esto se genera el fichero `go.mod` con el siguiente contenido:

```text
module send_telegram

go 1.19
```

### instalar dependencias

```shell
go mod tidy
```

Con esto se genera el fichero `go.sum` con el siguiente contenido:

```text
github.com/davecgh/go-spew v1.1.0/go.mod h1:J7Y8YcW2NihsgmVo/mv3lAwl/skON4iLHjSsI+c5H38=
github.com/davecgh/go-spew v1.1.1 h1:vj9j/u1bqnvCEfJOwUhtlOARqs3+rkHYY13jYWTU97c=
github.com/davecgh/go-spew v1.1.1/go.mod h1:J7Y8YcW2NihsgmVo/mv3lAwl/skON4iLHjSsI+c5H38=
github.com/pmezard/go-difflib v1.0.0 h1:4DBwDE0NGyQoBHbLQYPwSUPoCMWR5BEzIk/f1lZbAQM=
github.com/pmezard/go-difflib v1.0.0/go.mod h1:iKH77koFhYxTK1pcRnkKkqfTogsbg7gZNVY4sRDYZ/4=
github.com/sirupsen/logrus v1.9.0 h1:trlNQbNUG3OdDrDil03MCb1H2o9nJ1x4/5LYw7byDE0=
github.com/sirupsen/logrus v1.9.0/go.mod h1:naHLuLoDiP4jHNo9R0sCBMtWGeIprob74mVsIT4qYEQ=
github.com/stretchr/objx v0.1.0/go.mod h1:HFkY916IF+rwdDfMAkV7OtwuqBVzrE8GR6GFx+wExME=
github.com/stretchr/testify v1.7.0 h1:nwc3DEeHmmLAfoZucVR881uASk0Mfjw8xYJ99tb5CcY=
github.com/stretchr/testify v1.7.0/go.mod h1:6Fq8oRcR53rry900zMqJjRRixrwX3KX962/h/Wwjteg=
golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 h1:0A+M6Uqn+Eje4kHMK80dtF3JCXC4ykBgQG4Fe06QRhQ=
golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8/go.mod h1:oPkhp1MJrh7nUepCBck5+mAzfO9JrbApNNgaTdGDITg=
gopkg.in/check.v1 v0.0.0-20161208181325-20d25e280405/go.mod h1:Co6ibVJAznAaIkqp8huTwlJQCZ016jof/cbN4VW5Yz0=
gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c h1:dUUwHk2QECo/6vqA44rthZ8ie2QXMNeKRTHCNY2nXvo=
gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c/go.mod h1:K4uyk7z7BCEPqu6E+C64Yfv1cQ7kz7rIZviUmN+EgEM=
```

Y se actualizará el fichero `go.mod`, que ahora tendrá el siguiente contenido:

```text
module send_telegram

go 1.19

require github.com/sirupsen/logrus v1.9.0

require golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
```

### probar el código

El código necesita tener definida la siguiente variable de entorno:

| **VARIABLE** | **DESCRIPCIÓN** |
|--|--|
| `TOKEN` | Token del bot de **Telegram** que se usará para enviar el mensaje |

```shell
export TOKEN="123123:kjahdhjhajkfdhakhfd"

go run send_telegram.go --destination "12345" --message "hola mundo"
```

### compilar código

```shell
go run send_telegram.go
```

Se creará el binario `send_telegram`, que se puede distribuir.

## USUARIO

### Ejecutar el binario

El binario necesita tener definida la siguiente variable de entorno:

| **VARIABLE** | **DESCRIPCIÓN** |
|--|--|
| `TOKEN` | Token del bot de **Telegram** que se usará para enviar el mensaje |

```shell
export TOKEN="123123:kjahdhjhajkfdhakhfd"

./send_telegram --destination 12345 --message "esto es una prueba"
```
