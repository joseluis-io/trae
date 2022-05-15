# TRAE: Terminal RAE

TRAE es un programa de terminal que obtiene las definiciones oficiales de la Real Academia Española solicitadas por el usuario. Estas definiciones pueden ser almacenadas en el equipo para poder ser consultadas sin acceso a internet.

El nombre del programa viene de Terminal y RAE (Real Academia Española). Coincidiendo con la tercera persona del presente del verbo traer, que casualmente es una de las acciones principales del programa: traer definiciones de la página web de la [RAE](https://www.rae.es).

TRAE es un "web scraper". Estos programas obtienen información de una página web, y en el caso de TRAE la imprime para el usuario y si este quiere la puede persistir.

La base de datos escogida es [bolt](https://github.com/boltdb/bolt). Bolt es una base de datos clave/valor. Proporciona un motor de base de datos muy simple, rápido y fiable para proyectos que no requieran de una base de datos con servidor.

## Instalación

Deberá contar en su sistema con el lenguaje de programación Go, para instalarlo siga las [instrucciones oficiales](https://go.dev/doc/install).

```sh
git clone https://github.com/jl-hoz/trae.git
cd trae
go get . # Instala todas las dependencias necesarias
go build # Se habrá generado un binario como trae
./trae -h # Comando de ayuda para ver las opciones disponibles
```

## Fichero de configuración

El fichero de configuración deberá estar localizado en el directorio $HOME del usuario. El programa espera que exista en esa ruta un fichero de configuración, en caso de no ser así generará uno por defecto.

Estructura de un fichero de configuración por defecto:

```json
{
    "DatabaseStore": false,
    "DatabaseDirectory": "/home/user"
}
```

El fichero de configuración se debe encontrar en el directorio $HOME y el nombre de este debe ser .trae.

En el caso del fichero de la base de datos el nombre será .trae.db, el usuario indicará el directorio. Si este es erróneo entonces se guardará en el directorio $HOME.

## Uso

El siguiente comando indicará al usuario el comando de ayuda, ya que no cuenta con ninguna palabra para definirla:

    $ trae

Este comando indica los comandos disponibles del programa:

    $ trae -h

El comando principal obtendrá la palabra indicada por el usuario, escoge solo la primera palabra ignorando al resto:

    $ trae <palabra>

## Información adicional

Creado por José Luis de la Hoz García.
Comprobado el funcionamiento en Ubuntu (GNU/Linux).
Programa creado como entrega de la asignatura de Prácticas En Empresa II.