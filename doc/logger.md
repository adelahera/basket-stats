# Sistema de logs

## Criterios

- Rendimiento: rapidez de las distintas bibliotecas
- Seguridad: se medirá con la puntuación de cada biblioteca en Snyk Advisor.

## Opciones

### Log

El paquete [log](https://pkg.go.dev/log) pertenece a la biblioteca estándar de Go, por lo que no habría que incluir otras dependencias. Define un tipo, `Logger`, con métodos para dar formato a la salida. También tiene un `Logger` predefinido 'estándar' accesible a través de funciones auxiliares como Print[f|ln], Fatal[f|ln], y Panic[f|ln], que son más fáciles de usar que crear un `Logger` manualmente. Este logger escribe el error estándar e imprime la fecha y hora de cada mensaje de registro. Cada log se muestra en una línea separada: si el mensaje que se está imprimiendo no termina con un salto de línea, el logger agregará uno. Las funciones Fatal llaman a `os.Exit(1)` después de escribir el log. Las funciones Panic llaman a `panic` después de escribir el log.

### Zap

[Zap](https://github.com/uber-go/zap) es una biblioteca de registro en Go que se caracteriza por su alto rendimiento y eficiencia. Su enfoque principal es minimizar la sobrecarga de serialización y las asignaciones de memoria, ofreciendo un codificador JSON sin reflexión y libre de asignaciones. Zap proporciona un Logger base que busca la eficiencia, y sobre ese fundamento, construye un SugaredLogger de más alto nivel para aquellos usuarios que prefieren una API más flexible y fácil de usar. En resumen, Zap es conocido por su velocidad, baja asignación de memoria y la capacidad de adaptarse a las necesidades de los desarrolladores en términos de rendimiento y facilidad de uso.

IMPORTANTE: Zap solo soporta las dos versiones más recientes de Go. 

## Zerolog

[Zerolog](https://github.com/rs/zerolog) es otra biblioteca de registro estructurado para Golang que se enfoca en la simplicidad y el rendimiento. Está basada en Zap, por lo que es muy parecida, y su objetivo principal es ofrecer un mejor rendimiento.

## Log15

El paquete [log15](https://github.com/inconshreveable/log15) ofrece un conjunto de herramientas con una perspectiva definida y sencilla para realizar registros de acuerdo con las mejores prácticas en Go, que son legibles tanto para humanos como para máquinas. Está modelado siguiendo los paquetes io y net/http de la biblioteca estándar de Go y representa una alternativa al paquete de registro estándar de la biblioteca.

Sin embargo, lleva más de un año sin actualización, y no cuenta con ninguna puntuación en Snyk Advisor, por lo que parece que se ha quedado un poco atrás.

## Decisión

Para medir el rendimiento de los logs, se ha tomado en cuenta la prueba de rendimiento que realizaron los propios desarrolladores de Zap, en el que se mide el tiempo que tarda cada una de las opciones en hacer un log de un simple string. 

PD: también realizan otras medidas en diferentes contextos y en todas el ganador es Zerolog.

Fuente: https://github.com/uber-go/zap?tab=readme-ov-file#footnote-versions

| Package          | Time      | Time % to zap | Objects Allocated |
|------------------|-----------|---------------|--------------------|
| zap              | 63 ns/op  | +0%           | 0 allocs/op        |
| zerolog          | 32 ns/op  | -49%          | 0 allocs/op        |
| standard library | 124 ns/op | +97%          | 1 allocs/op        |
| log15            | 2069 ns/op| +3184%        | 20 allocs/op       |

En cuanto a la seguridad y su puntuación en Snyk Advisor:

- Zerolog cuenta con 95 puntos, una puntuación bastante alta.
- Log15, como comenté antes, ni siquiera cuenta con puntuación, al estar relativamente obsoleta.
- Log, de la biblioteca estándar, no he encontrado su puntuación, pero la presupongo bastante alta, al estar incluido en la biblioteca estándar de Go.
- Zap: tampoco he sido capaz de encontrar su puntuación. Sin embargo, también la supongo bastante alta, al ser una biblioteca con una gran comunidad y que se encuentra muy activa.

Teniendo en cuenta los dos criterios mencionados, rendimiento y seguridad, la opción más viable es utilizar `Zerolog`, pues es la que mayor rendimiento ofrece, además de presentar una puntuación muy alta en cuanto a seguridad.