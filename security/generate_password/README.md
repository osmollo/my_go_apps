# generate_password

- [generate\_password](#generate_password)
  - [DESARROLLADOR](#desarrollador)
    - [variables de entorno](#variables-de-entorno)
    - [iniciar el módulo](#iniciar-el-módulo)
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
go mod init generate_password
```

Con esto se genera el fichero `go.mod` con el siguiente contenido:

```text
module generate_password

go 1.19
```

### probar el código

```shell
go run generate_password.go

go run generate_password.go --size 64

go run generate_password.go --disable_special
```

### compilar código

```shell
go run generate_password.go
```

Se creará el binario `generate_password`, que se puede distribuir.

## USUARIO

### Ejecutar el binario

```shell
generate_password

generate_password --size 64

generate_password --disable_special
```
