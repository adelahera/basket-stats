## Gestor de tareas

En cuanto al task runner, sabemos que hay un montón de opciones diferentes. Primeramente, describiré los criterios utilizados para la elección. A continuación se describirán algunos de los más comunes, y luego, explicar el por qué de mi elección.

### Criterios

* **Lenguaje utilizado:** es conveniente que el task runner elegido esté desarrollado en Go, al igual que mi proyecto, para mantener la consistencia en el código y las prácticas de desarrollo. Por lo tanto, se valorará positivamente aquel task runner escrito en Go.

* **Antigüedad:** conviene que el task runner sea lo más reciente posible, o que tenga versiones recientes. Es señal de que mantienen actualizado el task runner.

* **Apoyo en Github**: Github es una de las páginas donde más se suele apoyar un desarrollador a la hora de buscar soluciones a ciertos problemas. Por ello, cuanto más apoyo tenga el task runner en Github, más información disponible hay para trabajar.

### Make

Make es una herramienta ampliamente empleada en la automatización de tareas, y utiliza un archivo Makefile para especificar las reglas que gobiernan la ejecución de estas tareas. Su ventaja principal radica en su alta versatilidad, ya que se apoya en el uso de la línea de comandos para llevar a cabo las tareas, lo que lo hace adecuado para su utilización con múltiples lenguajes y escenarios.

Sin embargo, tiene una serie de desventajas, como:

- Sintaxis un poco confusa
- Es difícil desarrollar un Makefile que sea multiplataforma.

Make no se llega a ajustar del todo a los criterios, pues, a pesar de contar con bastante apoyo en Github, no está escrito en Go; y además hay otras alternativas con versiones más recientes.

### Mage

Los archivos Makefile son difíciles de leer y de escribir en su mayoría debido a que son esencialmente scripts de bash sofisticados con una cantidad significativa de espacios en blanco y sintaxis adicional relacionada con make.

La opción de Mage sería una buena opción pues, además de estar escrita en Go, cuenta versiones recientes, como la 1.15.0, de septiembre de 2022. 

En definitiva, sería una buena opción, a pesar de contar con 5 mil estrellas menos en Github que Task.

### Task

Task utiliza un archivo de configuración en formato YAML para definir las tareas y la ejecución.

Task es otro de las task runners escritos en Go, además de contar con una versión muy reciente, la 3.31.0, de julio de 2023. Además de esto, cuenta con un grandísimo apoyo en Github, con más de 8,9k estrellas.

### Decisión

Como he ido describiendo, la opción que más se adecua a los criterios propuestos es `Task`. Está desarrollado en el propio Go, y además cuenta con un gran apoyo de la comunidad en Github. Por último, La versión 3.31.0 es la más reciente entre las opciones propuestas.