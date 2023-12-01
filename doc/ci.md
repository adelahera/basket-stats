# Sistemas de integración continua

## Criterios

- Gratuidad del servicio: el servicio debe ser gratuito.
- Integración con Docker y Github: es estrictamente necesario que se pueda integrar con Github Checks. Además, se valorará positivamente si también permite integración con Docker.
- Fiabilidad: como en todas las decisiones que hemos tomado hasta el momento, el sistema debe tener actualizaciones recientes. 
- Simplicidad: con esto básicamente me refiero a que se puede usar el servicio de manera online, y no requiere ninguna instalación local.

## Opciones

### Jenkins

Jenkins, una herramienta de integración continua de código abierto, es altamente configurable y gratuita. Aunque requiere más configuración, se integra con GitHub Checks mediante plugins y ofrece soporte Docker a través del plugin "Docker Pipeline". Con una comunidad activa, Jenkins es una opción sólida y confiable. Punto negativo, está hecho en Java.

### Travis CI

Travis CI es una plataforma de integración continua que destaca por su fácil configuración y se integra sin problemas con GitHub Checks y Docker. Es conocido por su fiabilidad y actualizaciones frecuentes, siendo una opción popular en el desarrollo de proyectos open source. Punto negativo, no es un servicio gratuito.

### Github Actions

GitHub Actions, integrado directamente en GitHub, proporciona minutos ilimitados para repositorios públicos (como es el caso) y es compatible con GitHub Checks. Ofrece soporte nativo para Docker y es altamente fiable, siendo una opción preferida para aquellos que utilizamos GitHub como repositorio principal.

### Circle CI

CircleCI, por su parte, proporciona un plan gratuito con limitaciones (máximo de 6000 minutos) y es reconocido por su fiabilidad. Ofrece una integración fluida con GitHub Checks y tiene soporte nativo para Docker, permitiendo la ejecución de construcciones en contenedores. Su enfoque en la automatización y la personalización lo hace adecuado para diversos flujos de trabajo.

### Drone

Drone es una plataforma de CI/CD gratuita y de código abierto que, aunque puede necesitar configuración adicional para GitHub Checks, destaca por sus actualizaciones frecuentes. Ofrece soporte nativo para Docker, facilitando la ejecución de construcciones en contenedores. Punto negativo, hay que tenerlo instalado, no se puede usar de forma online.

### AppVeyor

AppVeyor es otra opción que nos permite integrar CI en proyectos de open-source, que se utiliza para construir o testear proyectos hosteados en Github. Ofrece un plan gratuito para proyectos de open-source, como es nuestro caso. Además, se puede integrar fácilmente con Github y no requiere instalación en local.

## Decisión

Como hemos podido comprobar, todas los servicios son muy competentes. Sin embargo, hay algunos que no llegar a cumplir los requisitos, como Drone (ya que habría que tenerlo instalado en local), o Travis (que no es gratuito), o Jenkins (que está hecho en Java). Por ello, nos quedan 2 opciones; probar Github Actions, que sería gratuito para nuestro proyecto al tratarse de un repositorio público; y CircleCI, que habría que usar parte de los 6000 minutos gratis que nos ofrecen.

En caso de que nuestro proyecto gastara los 6000 minutos que nos ofree CircleCI, ya no sería gratuito, por lo que la mejor opción para nuestro sería utilizar Github Actions. Sin embargo, como se trata de aprender y probar cosas nuevas, voy a utilizar ambos (al menos probarlos), y a partir de ahí, tomar una decisión.

Una vez probadas ambas opciones, me he decantado por utilizar Github Actions porque, CircleCI, por algún motivo, me dice que es incompatible con Github Checks en mi repositorio.

Una vez comprobado Github Actions, he podido comprobar que, para Windows, ejecutar los tests del proyecto tarda casi 1 minuto y 30 segundos. Es probable que sea por la caché, pues la caché predeterminada para Ubuntu no le vale a Windows. Por ello, voy a intentar configurar AppVeyor, para probar una nueva opción, y hacerlo de tal manera que se ejecute en Windows a través de AppVeyor.