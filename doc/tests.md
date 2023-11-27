## Biblioteca de aserciones

En Go, no existe una biblioteca de aserciones estándar como en algunos otros lenguajes. En caso de incluir una, habría que utilizar una biblioteca externa. Para Go, las bibliotecas de aserciones más utilizadas son:

* [Testify](https://github.com/stretchr/testify): es una biblioteca de aserciones popular para Go. Proporciona una amplia variedad de funciones de aserción que pueden hacer que los tests sean más legibles y expresivas.

* [Go-cmp](https://pkg.go.dev/github.com/google/go-cmp/cmp): go-Cmp es una biblioteca que se utiliza para comparar estructuras y realizar aserciones en Go de manera más sofisticada.

* [Gomega](https://github.com/onsi/gomega): es una biblioteca que funciona bien con Ginkgo y proporciona expresiones de aserción más legibles. Aunque no es un ejecutor de pruebas en sí, a menudo se utiliza junto con Ginkgo.

## Test runners

En la librería estándar de Go se incluye, en el paquete testing, un ejecutor de pruebas llamado `go test`. Es fácil de usar y adecuado para la mayoría de proyectos. Además tiene la ventaja de que no habría que incluir ninguna dependencia, al estar incluido en la librería estándar de Go.

Otras opciones, todas externas a Go, sería utilizar:

* [Ginkgo](https://github.com/onsi/ginkgo): es un framework BDD para Go que te permite escribir pruebas de manera más estructurada y legible. Ginkgo incluye su propio task runner.

* [Goconvey](https://github.com/smartystreets/goconvey): es una herramienta que proporciona una interfaz web para ejecutar y visualizar las pruebas en tiempo real a medida que desarrollas. Esto puede ser útil para una retroalimentación rápida durante el desarrollo.

* [Testify](https://github.com/stretchr/testify): además de utilizarse como biblioteca de aserciones, puede utilizarse también como test runner. Además, puede usarse en combinación con `go test` 

## Decisión

El primer criterio que hay que seguir siempre es el de las mejores prácticas, o estándares en caso de que los haya. Aquí nos encontramos en una situación un poco peliaguda, pues en Go, uno de sus proverbios, (puedes ver todos [aquí](https://go-proverbs.github.io/)), es: *A little copying is better than a little dependency*. Con lo cuál, comprobamos que en el entorno de Go, no están a favor de añadir dependencias si no son estrictamente necesarias. 

Sin embargo, parece ser que los tests frameworks son una excepción a esta práctica, y son considerados lo suficientemente importantes como para incluirlos como dependencia.

Teniendo en cuenta esto, he decidido que la opción que voy a utilizar es la de `Testify`, únicamente como biblioteca de aserciones, a pesar de poder usarse como test runner también. 

En cuanto al test runner, utilizaré la opción que nos ofrece la biblioteca estándar de Go, `go test`. 

Esta combinación me ha parecido lo más adecuado, tanto en el sentido de que no genera una gran deuda técnica, como en el cumplimiento de las buenas prácticas en Go, pues, a pesar de tratarse de una biblioteca externa, se considera prácticamente una prolongación de Go, incluyendo las aserciones. Me parece la opción más *light*, ya que otros frameworks como Ginkgo, sí que generarían mayor deuda técnica, además de incumplir de manera más directa las buenas prácticas de Go.