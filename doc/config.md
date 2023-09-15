# Generación de clave pública y privada
- Generar el par de claves con el siguiente comando: ssh-keygen -t rsa -b 4096
- Copiar la clave pública con el comando: cat ~/.ssh/id_rsa.pub
- Copiar la clave pública en nuestro perfil de GitHub, en el apartado "SSH and GPG Keys"

# Configuración correcta del nombre y correo electrónico para que aparezca en los commits correctamente.
- Configurar el nombre de usuario con el comando: git config --global user.name "Tu Nombre"
- Configurar el correo electrónico con el comando: git config --global user.email "tu_email@example.com"
- Comprobar si se ha hecho correctamente con los comandos: git config --global user.name y git config --global user.email

