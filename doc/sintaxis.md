## Comprobación de sintaxis

En Go se suele utilizar por defecto la herramienta gofmt, que justamente realiza la función que estamos buscando solventar. 

Para realizar únicamente la función de comprobación de sintaxis, usaremos el comando: `gofmt -e file.go > dev/null`, que realizará el análisis sintáctico sin la necesidad de construir el proyecto.