## Gestor de dependencias

Anteriormente, en Go, la gestión de dependencias era menos ágil. Se utilizaban distintos gestores, como son:
- [Dep](https://github.com/golang/dep)
- [Glide](https://github.com/Masterminds/glide)
- [Govendor](https://github.com/kardianos/govendor)

El problema de estos gestores es que no forman parte del proyecto Go oficial y, por tanto, a veces creaban fricciones a la hora de importar nuevas dependencias.

Sin embargo, esto cambió en 2018, con la introducción de Go Modules en la versión 1.11 de Go. Con Go Modules, la gestión de dependencias se hizo más simple y más versátil. Go Modules está integrado directamente en Go, y algunas de sus características son:

- Comandos de gestión de dependencias como, **go get**, **go mod tidy**...

- Generación automática de un archivo de manifiesto, **go.mod**, con información detallada sobre las dependencias.