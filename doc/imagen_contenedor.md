# Elección de la imagen

## Criterios

Los criterios sobre los que me he basado para tomar la decisión sobre la imagen adecuada son:

* Tamaño: cuanto menos pese la imagen menos va a tardar en arrancar. Obviamente, cuanto más pequeña sea, traerá menos funcionalidades, pero como el proyecto no requiere de demasiadas bibliotecas, no las necesitamos. Es más conveniente que se trate de una imagen sencilla.

* Seguridad: la imagen debe mantenerse actualizada y gozar de un buen soporte.

* Compatibilidad con Golang: sería óptimo que la imagen elegida viniese con Go ya instalado, para ahorrarnos este paso y evitar posibles problemas con las versiones.

## Opciones

### Imagen oficial de Golang

Es la imagen oficial para Docker, mantenida por la propia comunidad. Muy popular. Existen dos variantes, la basada en Debian, y la basada en Alpine:

* `golang:<version>`: basada en Debian. Es la imagen por defecto. Si no se está seguro de las necesidades del proyecto, es la mejor opción. Está diseñada para ser utilizada tanto como un contenedor desechable como la base para construir otras imágenes.

* `golang:<version>-alpine`: Alpine Linux es mucho más pequeño que la mayoría de las imágenes base de distribuciones (aproximadamente 5 MB), lo que resulta en imágenes mucho más simples en general. No incluye herramientas como git, gcc o bash. Esta variante es altamente experimental y no cuenta con soporte oficial por parte del proyecto Go.

Estas dos variantes están ambas muy actualizadas, por lo que no supondría ninguna amenaza en cuanto seguridad, además de, obviamente, ser compatibles con Go.

### Ubuntu

Incluye un conjunto más amplio de utilidades y bibliotecas por defecto. Esto hace que sea una elección adecuada para proyectos que requieren una gama más amplia de herramientas. Además, simplifica la instalación de dependencias adicionales mediante el uso de `apt`. Aunque la imagen es más grande en comparación con opciones más minimalistas como Alpine, su versatilidad y buen soporte a largo plazo la hacen atractiva para proyectos que priorizan la familiaridad y la conveniencia sobre la eficiencia de espacio en disco y red. 

### Debian

Optar por una imagen de Debian para un entorno Go presenta características particulares. Similar a Ubuntu, Debian ofrece una imagen más amplia debido a la inclusión de utilidades y bibliotecas por defecto, lo que resulta en un entorno más cómodo para el desarrollo y la depuración. Al igual que Ubuntu, Debian cuenta con un sistema de gestión de paquetes robusto (apt), permitiendo una fácil instalación de dependencias adicionales. Cuenta con un sólido soporte y actualizaciones regulares.

## Decisión

Leyendo un poco la documentación sobre las diferentes opciones, observamos que el criterio de seguridad es ampliamente cumplido por todas las opciones. Aunque las últimas actualizaciones de Ubuntu y Debian son de octubre de 2023, no hay una gran diferencia con la última de Alpine o la de Golang, que son un mes anterior. 

En cuanto a compatibilidad con Go, tampoco habría una gran diferencia entre usar una u otra, todas las opciones traen Go por defecto y están preparadas para correr aplicaciones en Go.

Donde sí encontramos una gran diferencia es en el tamaño de las imagenes. Ubuntu y Debian traen un paquete de herramientas y bibliotecas demasiado amplio, lo cual hace que su tamaño aumente considerablemente. Sin embargo, con la imagen oficial de Golang, se reduce muchísimo el tamaño, pues prácticamente trae lo esencial para poder ejecutar aplicaciones de Go, y poco más. 

Como ya he dicho anteriormente, para la realización de este proyecto no se requiere de una gran cantidad de herramientas ni se utilizan apenas bibliotecas externas, por lo que, guíandonos por el tamaño, la mejor opción es usar la imagen oficial de Go.

Dentro de la imagen oficial de Go, usaré la variante que trae con Alpine, pues reduce aún más el tamaño de la imagen, lo que debe aumentar la velocidad de ejecución. A pesar de no tener herramientas instaladas como git o bash, al apenas necesitar otras herramientas, no supone un gran trabajo instalar las necesarias manualmente.











