## Problema inicial

Un periódico deportivo quiere lanzar una aplicación web que sirva para comparar jugadores de baloncesto de diferentes épocas a través de sus estadísticas. Su principal problema es que, a lo largo de la historia, han pasado miles de jugadores por la liga, y como ya hemos dicho, en etapas muy distintas. No es lo mismo el baloncesto en los 70, que el baloncesto actual. No encuentran una manera de normalizar las estadísticas para poder comparar jugadores de distintas épocas entre sí.

## Configuración inicial del proyecto

La configuración se puede ver pinchando [aquí](/doc/config.md).

## Historias de usuario

Consulte las historias de usuario del proyecto [aquí](/doc/historias_usuario.md).

## Milestones

Consulte los milestones del proyecto [aquí](/doc/milestones.md).

## Herramientas utilizadas
### Gestor de dependencias

En el proyecto, se utilizará `Go Modules` como gestor de dependencias. Puedes encontrar más información sobre la decisión [aquí](/doc//gestor_dependencias.md).

### Gestor de tareas

En el proyecto, se utilizará `Task` como gestor de taraes. Puedes encontrar más información sobre la decisión [aquí](/doc/gestor_tareas.md).

* `task install-deps` : Instala las dependencias necesarias
* `task update-deps` : Actualiza las dependencias
* `task check` : Comprueba la sintaxis del código
* `task test` : Ejecuta los tests

### Comprobador de sintaxis

En el proyecto, se utilizará `gofmt` como comprobador de sintaxis. Puedes encontrar más información sobre la decisión [aquí](/doc/sintaxis.md).

### Herramientas para los tests

Se va a utilizar la biblioteca `testify` como biblioteca de aserciones. Complementando a testify, se utilizará `go test`, de la propia librería estándar de Go, para ordenar la ejecución de los diferentes tests. El proceso de decisión y la documentación sobre las otras opciones que había sobre la mesa puede encontrarse [aquí](/doc/tests.md).

