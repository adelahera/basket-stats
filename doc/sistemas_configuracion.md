# Sistemas de configuración

Una biblioteca de configuración simplifica el proceso de trabajar con configuraciones en una aplicación, proporcionando herramientas y abstracciones que facilitan la flexibilidad y la gestión eficiente de la configuración de una aplicación.

## Criterios

- Seguridad: basada en la puntuación en Snyk Advisor.
- Reputación y uso en la comunidad: se medirá con el número de forks y de estrellas de los respectivos repositorios de Github. Este criterio es bastante importante, pues es la primera vez trabajando con este tipo de bibliotecas, y es importante que exista suficiente documentación en caso de necesitarla.
- Soporte para configuración por sistemas distribuidos, variables de entorno...

## Opciones

### Viper

Viper es una biblioteca de configuración en Go que proporciona una solución completa y flexible. Permite cargar configuraciones desde múltiples fuentes como archivos, variables de entorno, sistemas de distribución como etcd, y flags de línea de comandos. Viper admite diversos formatos de archivos de configuración y ofrece funciones para trabajar con valores predeterminados, validación y recarga en caliente.

### Koanf

Koanf es una biblioteca para leer configuraciones desde diversas fuentes y en diferentes formatos en aplicaciones de Go. Es una alternativa más limpia y ligera a Viper, ofreciendo mejores abstracciones, mayor capacidad de extensión y muchas menos dependencias. 

### Cobra

Cobra es una biblioteca que proporciona una interfaz sencilla para crear potentes interfaces de línea de comandos (CLI) modernas, similares a las de herramientas como git y go tools.

Cobra ofrece:

- Creación fácil de CLIs basados en subcomandos: app server, app fetch, etc.
- Flags totalmente compatibles con POSIX 
- Flags globales, locales y cascading 
- Generación automática de ayuda para comandos y flags
- Reconocimiento automático de flags de ayuda como -h, --help, etc.
- Autocompletado de shell generado automáticamente para la aplicación (bash, zsh, fish, powershell)
- Páginas de manual generadas automáticamente para la aplicación
- Alias de comandos para poder cambiar cosas sin romperlas
- Integración opcional y sin problemas con Viper.

## Decisión

| Paquete          | Snyk | Forks | Estrellas|
|------------------|------|-------|----------|
| Viper            | 89   | 2.000 | 24.700   |
| Koanf            | 82   | 141   | 2.100    |
| Cobra            | 91   | 2.800 | 34.500   |

A pesar de que `Cobra` tenga más puntuación y reputación que las otras dos, su objetivo es crear interfaces CLI, que no es necesario en nuestro caso. Puede integrarse con `Viper` para brindarle soporte para crear configuraciones para sistemas distribuidos, que es lo que buscamos. Como a fin de cuentas, `Cobra` utiliza `Viper` para esto, nuestra elección va a ser esta última, pues el resto de funcionalidades que implementa `Cobra` no las necesitamos.

`Koanf`, como podemos observar se queda un poco atrás respecto a `Viper`.

Elección final: `Viper`