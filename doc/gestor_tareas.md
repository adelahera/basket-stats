## Gestor de tareas

En cuanto al task runner, sabemos que hay un montón de opciones diferentes. Voy a describir algunas de las más comunes, y luego, explicar el por qué de mi elección.

### Make

Make es una herramienta ampliamente empleada en la automatización de tareas, y utiliza un archivo Makefile para especificar las reglas que gobiernan la ejecución de estas tareas. Su ventaja principal radica en su alta versatilidad, ya que se apoya en el uso de la línea de comandos para llevar a cabo las tareas, lo que lo hace adecuado para su utilización con múltiples lenguajes y escenarios.

Sin embargo, tiene una serie de desventajas, como:

- Requiere aprender Make (y además Bash)
- Sintaxis un poco confusa
- Es difícil desarrollar un Makefile que sea multiplataforma.

### Mage

Los archivos Makefile son difíciles de leer y de escribir en su mayoría debido a que son esencialmente scripts de bash sofisticados con una cantidad significativa de espacios en blanco y sintaxis adicional relacionada con make.

Mage te permite tener varios archivos Mage, nombrarlos como desees y son fáciles de personalizar para múltiples sistemas operativos. Mage no tiene dependencias (a excepción de Go) y funciona perfectamente en todos los sistemas operativos principales, mientras que make generalmente utiliza bash, que no tiene un buen soporte en Windows. Go es superior a bash para cualquier tarea no trivial que implique ramificación, bucles o cualquier cosa que no sea simplemente la ejecución lineal de comandos.

La opción de Mage sería una buena opción pues, además de estar escrita en Go, utiliza una sintaxis sencilla y legible (al menos lo parece por lo que he visto de Mage). 

En definitiva, sería una buena opción.

### Task

Task es otro de las task runners más comunes utilizados, pues también está escrito en Go, al igual que Mage.

Task utiliza un archivo de configuración en formato YAML para definir las tareas y la ejecución.

Las ventajas que tiene Task son:

- Compatibilidad multiplataforma: task es compatible con diferentes sistemas operativos, por lo que no habrá problema a la hora de ejecutar las tareas en entornos diferentes.

- Dependencias y orden de ejecución: permite definir dependencias entre tareas, lo que significa que una tarea puede depender de la finalización de otra tarea antes de ejecutarse.

- Integración con otros sistemas: se puede integrar fácilmente con sistemas de control de versiones, como Git, lo cuál se ajusta perfectamente a la dinñamica de la asignatura.

### Decisión

Por todos los motivos explicados anteriormente, el gestor de tareas a utilizar en mi proyecto, va a ser **TASK**