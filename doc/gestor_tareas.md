## Gestor de tareas

En cuanto al task runner, sabemos que hay un montón de opciones diferentes. Voy a describir algunas de las más comunes, y luego, explicar el por qué de mi elección.

### Make

Make es ampliamente utilizada, a la par que antigua. A pesar de ser muy popular a día de hoy, tiene una serie de inconvenientes, como:

- Requiere aprender Make (y además Bash)
- Es difícil desarrollar un Makefile que sea multiplataforma.

### Mage

La opción de Mage sería una buena opción pues, además de estar escrita en Go, utiliza una sintaxis sencilla y legible (al menos lo parece por lo que he visto de Mage). También permite definir tareas personalizadas para automatizar diferentes aspectos del proyecto.

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