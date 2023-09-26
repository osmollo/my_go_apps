# git backup

- [git backup](#git-backup)
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
go mod init git_backup
```

Con esto se genera el fichero `go.mod` con el siguiente contenido:

```text
module git_backup

go 1.21.1
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
module git_backup

go 1.21.1

```

### probar el código

```shell
go run git_backup.go
```

### compilar código

```shell
go run git_backup.go
```

Se creará el binario `git_backup`, que se puede distribuir.

## USUARIO

### Ejecutar el binario

```shell
./git_backup output.pdf file1.pdf file2.pdf
```
