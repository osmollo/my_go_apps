# merge_pdf

- [merge\_pdf](#merge_pdf)
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
go mod init merge_pdf
```

Con esto se genera el fichero `go.mod` con el siguiente contenido:

```text
module merge_pdf

go 1.19
```

### instalar dependencias

```shell
go mod tidy
```

Con esto se genera el fichero `go.sum` con el siguiente contenido:

```text
github.com/adrg/strutil v0.1.0/go.mod h1:pXRr2+IyX5AEPAF5icj/EeTaiflPSD2hvGjnguilZgE=
github.com/adrg/sysfont v0.1.1/go.mod h1:19nTHzfIn/HbngFMet+yNAvwSQYtOJYMI7vWexLWyNw=
github.com/adrg/xdg v0.2.1/go.mod h1:ZuOshBmzV4Ta+s23hdfFZnBsdzmoR3US0d7ErpqSbTQ=
github.com/boombuler/barcode v1.0.0/go.mod h1:paBWMcWSl3LHKBqUq+rly7CNSldXjb2rDl3JlRe0mD8=
github.com/davecgh/go-spew v1.1.0/go.mod h1:J7Y8YcW2NihsgmVo/mv3lAwl/skON4iLHjSsI+c5H38=
github.com/davecgh/go-spew v1.1.1 h1:vj9j/u1bqnvCEfJOwUhtlOARqs3+rkHYY13jYWTU97c=
github.com/davecgh/go-spew v1.1.1/go.mod h1:J7Y8YcW2NihsgmVo/mv3lAwl/skON4iLHjSsI+c5H38=
github.com/golang/snappy v0.0.4/go.mod h1:/XxbfmMg8lxefKM7IXC3fBNl/7bRcc72aCRzEWrmP2Q=
github.com/gorilla/i18n v0.0.0-20150820051429-8b358169da46/go.mod h1:2Yoiy15Cf7Q3NFwfaJquh7Mk1uGI09ytcD7CUhn8j7s=
github.com/konsorten/go-windows-terminal-sequences v1.0.1/go.mod h1:T0+1ngSBFLxvqU3pZ+m/2kptfBszLMUkC4ZK/EgS/cQ=
github.com/konsorten/go-windows-terminal-sequences v1.0.2 h1:DB17ag19krx9CFsz4o3enTrPXyIXCl+2iCXH/aMAp9s=
github.com/konsorten/go-windows-terminal-sequences v1.0.2/go.mod h1:T0+1ngSBFLxvqU3pZ+m/2kptfBszLMUkC4ZK/EgS/cQ=
github.com/kr/pretty v0.1.0 h1:L/CwN0zerZDmRFUapSPitk6f+Q3+0za1rQkzVuMiMFI=
github.com/kr/pretty v0.1.0/go.mod h1:dAy3ld7l9f0ibDNOQOHHMYYIIbhfbHSm3C4ZsoJORNo=
github.com/kr/pty v1.1.1/go.mod h1:pFQYn66WHrOpPYNljwOMqo10TkYh1fy3cYio2l3bCsQ=
github.com/kr/text v0.1.0 h1:45sCR5RtlFHMR4UwH9sdQ5TC8v0qDQCHnXt+kaKSTVE=
github.com/kr/text v0.1.0/go.mod h1:4Jbv+DJW3UT/LiOwJeYQe1efqtUx/iVham/4vfdArNI=
github.com/montanaflynn/stats v0.6.6/go.mod h1:etXPPgVO6n31NxCd9KQUMvCM+ve0ruNzt6R8Bnaayow=
github.com/pmezard/go-difflib v1.0.0 h1:4DBwDE0NGyQoBHbLQYPwSUPoCMWR5BEzIk/f1lZbAQM=
github.com/pmezard/go-difflib v1.0.0/go.mod h1:iKH77koFhYxTK1pcRnkKkqfTogsbg7gZNVY4sRDYZ/4=
github.com/sirupsen/logrus v1.5.0 h1:1N5EYkVAPEywqZRJd7cwnRtCb6xJx7NH3T3WUTF980Q=
github.com/sirupsen/logrus v1.5.0/go.mod h1:+F7Ogzej0PZc/94MaYx/nvG9jOFMD2osvC3s+Squfpo=
github.com/stretchr/objx v0.1.0/go.mod h1:HFkY916IF+rwdDfMAkV7OtwuqBVzrE8GR6GFx+wExME=
github.com/stretchr/testify v1.2.2/go.mod h1:a8OnRcib4nhh0OaRAV+Yts87kKdq0PP7pXfy6kDkUVs=
github.com/stretchr/testify v1.4.0/go.mod h1:j7eGeouHqKxXV5pUuKE4zz7dFj8WfuZ+81PSLYec5m4=
github.com/stretchr/testify v1.7.1 h1:5TQK59W5E3v0r2duFAb7P95B6hEeOyEnHRa8MjYSMTY=
github.com/stretchr/testify v1.7.1/go.mod h1:6Fq8oRcR53rry900zMqJjRRixrwX3KX962/h/Wwjteg=
github.com/trimmer-io/go-xmp v1.0.0/go.mod h1:Aaptr9sp1lLv7UnCAdQ+gSHZyY2miYaKmcNVj7HRBwA=
github.com/unidoc/freetype v0.0.0-20220130190903-3efbeefd0c90/go.mod h1:mJ/Q7JnqEoWtajJVrV6S1InbRv0K/fJerPB5SQs32KI=
github.com/unidoc/garabic v0.0.0-20220702200334-8c7cb25baa11/go.mod h1:SX63w9Ww4+Z7E96B01OuG59SleQUb+m+dmapZ8o1Jac=
github.com/unidoc/pkcs7 v0.0.0-20200411230602-d883fd70d1df/go.mod h1:UEzOZUEpJfDpywVJMUT8QiugqEZC29pDq7kdIZhWCr8=
github.com/unidoc/pkcs7 v0.1.0 h1:9bQfbWMYsIfUP8PyhTcBudOsvbLpNH0MBv4U0P/jDTE=
github.com/unidoc/pkcs7 v0.1.0/go.mod h1:UEzOZUEpJfDpywVJMUT8QiugqEZC29pDq7kdIZhWCr8=
github.com/unidoc/timestamp v0.0.0-20200412005513-91597fd3793a h1:RLtvUhe4DsUDl66m7MJ8OqBjq8jpWBXPK6/RKtqeTkc=
github.com/unidoc/timestamp v0.0.0-20200412005513-91597fd3793a/go.mod h1:j+qMWZVpZFTvDey3zxUkSgPJZEX33tDgU/QIA0IzCUw=
github.com/unidoc/unichart v0.1.0/go.mod h1:9sJXeqxIIsU2D07tmhpDMoND0mBFRGfKBJnXZMsJnzk=
github.com/unidoc/unipdf/v3 v3.40.0 h1:aSB5SGL9HMBhy6F14O35hu9/0r0Wj+PLP6Br3sZpSME=
github.com/unidoc/unipdf/v3 v3.40.0/go.mod h1:GD44mz5CQWK2QVgUy8PcMlhMX3l8Ymd3FdLmi0ucn7E=
github.com/unidoc/unitype v0.2.1 h1:x0jMn7pB/tNrjEVjy3Ukpxo++HOBQaTCXcTYFA6BH3w=
github.com/unidoc/unitype v0.2.1/go.mod h1:mafyug7zYmDOusqa7G0dJV45qp4b6TDAN+pHN7ZUIBU=
golang.org/x/crypto v0.0.0-20220722155217-630584e8d5aa h1:zuSxTR4o9y82ebqCUJYNGJbGPo6sKVl54f/TVDObg1c=
golang.org/x/crypto v0.0.0-20220722155217-630584e8d5aa/go.mod h1:IxCIyHEi3zRg3s0A5j5BB6A9Jmi73HwBIUl50j+osU4=
golang.org/x/image v0.0.0-20211028202545-6944b10bf410 h1:hTftEOvwiOq2+O8k2D5/Q7COC7k5Qcrgc2TFURJYnvQ=
golang.org/x/image v0.0.0-20211028202545-6944b10bf410/go.mod h1:023OzeP/+EPmXeapQh35lcL3II3LrY8Ic+EFFKVhULM=
golang.org/x/net v0.0.0-20211112202133-69e39bad7dc2/go.mod h1:9nx3DQGgdP8bBQD5qxJ1jj9UTztislL4KSBs9R2vV5Y=
golang.org/x/sys v0.0.0-20190422165155-953cdadca894/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
golang.org/x/sys v0.0.0-20200413165638-669c56c373c4/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
golang.org/x/sys v0.0.0-20201119102817-f84b799fce68/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
golang.org/x/sys v0.0.0-20210423082822-04245dca01da/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
golang.org/x/sys v0.0.0-20210615035016-665e8c7367d1 h1:SrN+KX8Art/Sf4HNj6Zcz06G7VEz+7w9tdXTPOZ7+l4=
golang.org/x/sys v0.0.0-20210615035016-665e8c7367d1/go.mod h1:oPkhp1MJrh7nUepCBck5+mAzfO9JrbApNNgaTdGDITg=
golang.org/x/term v0.0.0-20201126162022-7de9c90e9dd1/go.mod h1:bj7SfCRtBDWHUb9snDiAeCFNEtKQo2Wmx5Cou7ajbmo=
golang.org/x/text v0.3.2/go.mod h1:bEr9sfX3Q8Zfm5fL9x+3itogRgK3+ptLWKqgva+5dAk=
golang.org/x/text v0.3.6 h1:aRYxNxv6iGQlyVaZmk6ZgYEDa+Jg18DxebPSrd6bg1M=
golang.org/x/text v0.3.6/go.mod h1:5Zoc/QRtKVWzQhOtBMvqHzDpF6irO9z98xDceosuGiQ=
golang.org/x/tools v0.0.0-20180917221912-90fa682c2a6e/go.mod h1:n7NCudcB/nEzxVGmLbDWY5pfWTLqBcC2KZ6jyYvM4mQ=
golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 h1:go1bK/D/BFZV2I8cIQd1NKEZ+0owSTG1fDTci4IqFcE=
golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1/go.mod h1:I/5z698sn9Ka8TeJc9MKroUUfqBBauWjQqLJ2OPfmY0=
gopkg.in/check.v1 v0.0.0-20161208181325-20d25e280405/go.mod h1:Co6ibVJAznAaIkqp8huTwlJQCZ016jof/cbN4VW5Yz0=
gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 h1:qIbj1fsPNlZgppZ+VLlY7N33q108Sa+fhmuc+sWQYwY=
gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127/go.mod h1:Co6ibVJAznAaIkqp8huTwlJQCZ016jof/cbN4VW5Yz0=
gopkg.in/yaml.v2 v2.2.2/go.mod h1:hI93XBmqTisBFMUTm0b8Fm+jr3Dg1NNxqwp+5A1VGuI=
gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c h1:dUUwHk2QECo/6vqA44rthZ8ie2QXMNeKRTHCNY2nXvo=
gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c/go.mod h1:K4uyk7z7BCEPqu6E+C64Yfv1cQ7kz7rIZviUmN+EgEM=
```

Y se actualizará el fichero `go.mod`, que ahora tendrá el siguiente contenido:

```text
module merge_pdf

go 1.18

require github.com/unidoc/unipdf/v3 v3.40.0

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/konsorten/go-windows-terminal-sequences v1.0.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/sirupsen/logrus v1.5.0 // indirect
	github.com/stretchr/testify v1.7.1 // indirect
	github.com/unidoc/pkcs7 v0.1.0 // indirect
	github.com/unidoc/timestamp v0.0.0-20200412005513-91597fd3793a // indirect
	github.com/unidoc/unitype v0.2.1 // indirect
	golang.org/x/crypto v0.0.0-20220722155217-630584e8d5aa // indirect
	golang.org/x/image v0.0.0-20211028202545-6944b10bf410 // indirect
	golang.org/x/sys v0.0.0-20210615035016-665e8c7367d1 // indirect
	golang.org/x/text v0.3.6 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c // indirect
)
```

### probar el código

```shell
go run merge_pdf.go output.pdf file1.pdf file2.pdf
```

### compilar código

```shell
go run merge_pdf.go
```

Se creará el binario `merge_pdf`, que se puede distribuir.

## USUARIO

### Ejecutar el binario

```shell
./merge_pdf output.pdf file1.pdf file2.pdf
```
