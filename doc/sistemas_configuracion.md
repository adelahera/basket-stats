# Sistemas de configuración

Una biblioteca de configuración simplifica el proceso de trabajar con configuraciones en una aplicación, proporcionando herramientas y abstracciones que facilitan la flexibilidad y la gestión eficiente de la configuración de una aplicación.

## Criterios

- Seguridad: basada en la puntuación en Snyk Advisor.
- Reputación y uso en la comunidad: se medirá con el número de forks y de estrellas de los respectivos repositorios de Github. Este criterio es bastante importante, pues es la primera vez trabajando con este tipo de bibliotecas, y es importante que exista suficiente documentación en caso de necesitarla.
- Soporte para configuración por sistemas distribuidos, variables de entorno...

## Opciones

### Viper

[Viper](https://github.com/spf13/viper) es una biblioteca de configuración en Go que proporciona una solución completa y flexible. Permite cargar configuraciones desde múltiples fuentes como archivos, variables de entorno, sistemas de distribución como etcd, y flags de línea de comandos. Viper admite diversos formatos de archivos de configuración y ofrece funciones para trabajar con valores predeterminados, validación y recarga en caliente.

### Koanf

[Koanf](https://github.com/knadh/koanf) es una biblioteca para leer configuraciones desde diversas fuentes y en diferentes formatos en aplicaciones de Go. Es una alternativa más limpia y ligera a Viper, ofreciendo mejores abstracciones, mayor capacidad de extensión y muchas menos dependencias. 

### Envconfig

[Envconfig](https://github.com/kelseyhightower/envconfig) es una biblioteca que simplifica la lectura de configuraciones desde variables de entorno en Go. Permite definir una estructura de configuración con tags de estructura para especificar cómo se deben leer las variables de entorno. La parte mala es que lleva un año sin actualización también, por lo que su puntuación en Snyk es bastante baja precisamente por esto, a pesar de que parece ser una biblioteca bastante fiable y utilizada.

### Confita

[Confita](https://github.com/heetch/confita) es una biblioteca bastante simple que carga la configuración y la almacena en un struct. No implementa nignguna funcionalidad a parte de esta, (como sí lo hacía Cobra, por ejemplo) por lo que puede ser buena opción si no queremos complicarnos. Por lo que he visto de Confita era bastante utilizada hace un tiempo, pero por algún motivo dejaron de actualizarla hace un año.

## Decisión

| Paquete          | Snyk | Forks | Estrellas|
|------------------|------|-------|----------|
| Viper            | 89   | 2.000 | 24.700   |
| Koanf            | 82   | 141   | 2.100    |
| Confita          | 55   | 50    | 481      |
| Envconfig        | 65   | 374   | 4700     |

Viendo las puntuaciones de las distintas opciones estudiadas, se ve claramente que `Viper` es la opción que más se ajusta a los criterios. El resto de opciones, sobre todo `Koanf`, parecen ser muy buenas opciones, pero el hecho de que muchas de ellas llevan casi un año sin actualizarse, dejan a `Viper` como primera candidata.

Elección final: `Viper`