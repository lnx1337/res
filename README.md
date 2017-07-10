# Resuelve / Facturas / API

## Introducción:
El ejercicio esta desarollado en [GOLANG](https://golang.org/) y es una API que retorna el total de facturas de un cliente.
 
Este API sigue el modelo de aplicación client-server, la comunicacioón es mediante mensajes serializados en formato JSON. El API permitirá a uno o más clients comunicarse con un server y poder consultar las facturas de un cliente en un periodo en especifico.

### DEMO

```
  http://resuelve.jair.xyz/v1/invoice/:id/:startDate/:endDate

```

#### Ejemplo
[http://resuelve.jair.xyz/v1/invoice/717f076e-e13c-45b4-bcc4-51c229e1b326/2017-01-01/2017-12-12](http://resuelve.jair.xyz/v1/invoice/717f076e-e13c-45b4-bcc4-51c229e1b326/2017-01-01/2017-12-12)

## Instalación del proyecto en local:

### Dependencias

#### Instalar Sistema base
Linux Debian, Ubuntu
 
#### Instalar [Go](https://golang.org/dl ) Linux

```
$ wget http://golang.org/dl/go1.8.3.linux-amd64.tar.gz
$ sudo tar zxvf go1.8.3.linux-amd64.tar.gz -C /usr/local
$ export PATH=$PATH:/usr/local/go/bin
$ rm -rf ~/go1.8.3.linux-amd64.tar.gz

```

#### Instalar [Go](https://golang.org/dl ) MAC

Descargar e instalar [GO](https://storage.googleapis.com/golang/go1.8.3.darwin-amd64.pkg)

#### Configuración de go en $PATH


En caso de no existir crear archivo en el directorio $HOME/.profile:

```
$ touch ~/.profile

```

Abrir archivo con cualquier editor:

```
$vim ~/.profile

```


Añadir al PATH golang en el archivo ~/.profile:

```
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin

```

Ejecutamos el archivo ~/.profile

```
$source ~/.profile

```

#### Instalar git Linux

```
$ sudo apt-get install git

```

#### Instalar make

```
$ sudo apt-get install make

```

## Configuración del proyecto

#### Clonar repositorio

```
$ git clone https://github.com/lnx1337/res.git resuelve

```

Accedemos al directorio raíz del proyecto:

```
$ cd resuelve/
```



Ejecutamos el make file

Se instalaran todas las dependencias de go necesarias para el proyecto y se ejecutarán los test de cada módulo.

```
$ sudo make install

```

#### Archivos de configuración

En el archivo service.invoice.toml se podrá editar el puerto HTTP de inicio. 

```
$vim service.invoice.toml 

```


#### Ejecutar el proyecto en local:

```
$ ./run.sh
```

+ Se inicializa por defecto el servidor en el puerto: `1337`.

### Demo local

Abrir algún navegador en la url : [http://localhost:1337/v1/invoice/717f076e-e13c-45b4-bcc4-51c229e1b326/2017-01-01/2017-12-12](http://localhost:1337/v1/invoice/717f076e-e13c-45b4-bcc4-51c229e1b326/2017-01-01/2017-12-12)