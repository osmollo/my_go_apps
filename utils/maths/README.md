# git backup

- [MCD](#MCD)
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
go mod init maths
```

Con esto se genera el fichero `go.mod` con el siguiente contenido:

```text
module maths

go 1.21.3
```

### instalar dependencias

```shell
go mod tidy
```

Con esto se genera el fichero `go.sum` con el siguiente contenido:

```text
```

Y se actualizará el fichero `go.mod`, que ahora tendrá el siguiente contenido:

```text
module maths

go 1.21.3

```

### probar el código

```shell
go run maths.go
```

### compilar código

```shell
go run maths.go
```

Se creará el binario `maths`, que se puede distribuir.

## USUARIO

### Ejecutar el binario

```shell
./maths num1,num2,num3,num4,...
```
