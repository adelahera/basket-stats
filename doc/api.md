# Framework API REST

## Criterios

* Compatibilidad: el framework elegido debe ser compatible con Go.
* Coste: debe ser gratuito.
* Seguridad: se valorará según su puntuación en [Snyk](https://snyk.io/advisor/golang)
* Simplicidad: buscamos un framework liviano y sencillo, cuya deuda técnica sea mínima.
* Rendimiento.

## Opciones

### Gin

[Gin](https://github.com/gin-gonic/gin) es un framework web para Go conocido por su velocidad y rendimiento. Destaca por su sólido soporte para middleware, facilitando la incorporación de funcionalidades adicionales. La sintaxis concisa de Gin simplifica la definición de rutas y el manejo de solicitudes. Implementa renderizado en formato JSON y XML.

### Echo

[Echo](https://github.com/labstack/echo) es otro framework web para Golang que también utiliza el estándar de enrutador HTTP para enrutar solicitudes. Echo tiene una API similar a Gin, pero con algunas diferencias. Echo cuenta con características como soporte de grupos de rutas, inyección de dependencia, recuperación de pánico, compresión de respuestas, entre otras. Echo también tiene documentación más completa y detallada que Gin.

### Fiber

[Fiber](https://github.com/gofiber/fiber) es un framework rápido y eficiente para Go que se inspira en Express.js (Node.js). Diseñado para un rendimiento óptimo, Fiber utiliza una sintaxis similar a Express.js y un sistema de middleware eficiente para manejar solicitudes HTTP.

### Chi

[Chi](https://github.com/go-chi/chi) es un enrutador ligero y rápido para Go. Su diseño se centra en la simplicidad y eficiencia, ofreciendo flexibilidad para construir aplicaciones web. Chi destaca por su contexto rápido que beneficia al rendimiento.

## Decisión

Todas las opciones que se han documentado cumplen de antemano los dos primeros requisitos, son compatibles con Go y son gratuitos. En cuanto a su puntuación en Snyk, podemos compararla en la siguiente tabla:

| Framework | Snyk | 
|-----------|------|
| Gin       | 93   |
| Chi       | 94   |
| Fiber     | 92   |
| Echo      | 94   |

Como podemos observar, todas las opciones son muy sólidas en todos los aspectos, pues es de aclarar que los frameworks son muy parecidos entre sí, por lo que no habría una gran diferencia entre elegir uno u otro. 

No obstante, mi elección va a ser `Fiber`, pues en la lista de frameworks con mejor rendimiento de TechEmpower, (se puede ver [aquí](https://www.techempower.com/benchmarks/#hw=ph&test=fortune&section=data-r22)), `Fiber` se encuentra en la posición 29, siendo el mejor posicionado en cuanto a rendimiento en frameworks cuyo objetivo es el diseño de APIs REST.