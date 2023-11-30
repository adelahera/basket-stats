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

### Imagen de Bitnami

Bitnami ofrece imágenes actualizadas automáticamente, basadas en minideb, una imagen de contenedor Debian minimalista, lo que facilita la transición entre contenedores, máquinas virtuales y imágenes en la nube. Todas las imágenes están firmadas con Docker Content Trust para garantizar su integridad. Bitnami publica regularmente nuevas versiones con las últimas correcciones y características

## Decisión

Leyendo un poco la documentación sobre las diferentes opciones, observamos que el criterio de seguridad es ampliamente cumplido por todas las opciones. Aunque las últimas actualizaciones de Ubuntu y Debian son de octubre de 2023, no hay una gran diferencia con la última de Alpine o la de Golang, que son un mes anterior. En cuanto a Bitnami, su última actualización es de hace sólo 22 días.

En cuanto a compatibilidad con Go, tampoco habría una gran diferencia entre usar una u otra, las imágenes suelen traer Go por defecto, y en caso de que no, instalarlo manualmente no es complicado.

Donde sí encontramos una gran diferencia es en el tamaño de las imagenes. Ubuntu y Debian traen un paquete de herramientas y bibliotecas demasiado amplio, lo cual hace que su tamaño aumente considerablemente. Sin embargo, los contenedores que usan las imágenes oficiales, a pesar de contar con miles de pulls más, suelen incluir un montón de funcionalidades que no vienen al caso, como puede ser Python o Perl, que en nuestro proyecto no es recomendable. 

Quizá por este motivo, utilizar la imagen proporcionada por Bitnami sea la opción más viable. No nos va a incluir bibliotecas ni funcionalidades no necesarias, lo cual va a mantener el tamaño de la imagen más reducido. 












