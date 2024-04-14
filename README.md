# RETO FINANCIERA COMPARTAMOS

API CRUD de clientes asociados a ciudades.

### Demo

- Debe iniciar sesion en [Link dominio](https://explore.swaggerhub.com/catalog?owner=MIGUELSR1084&api=reto-financiera_compartamos&version=1.0.0)

*Nota*:
- Al hacer la petición la primera vez puede demorar un poco, ya que el servicio estará levantándose.

*Precaución*
- El editor online no es tan fiable, ya que puede mostrar resultados cuando no le pasas datos a los __Param o Queryparam__,
  se recomienda hacer las pruebas con Postman.


---
### Documentación

**Codigo**

El proyecto tiene uns estructura de arquitectura hexagonal, donde se separa la logica de negocio de la logica de persistencia.
ademas del control de errores, tambien se manejo una salida standar de respuesta cuando hay un error o cuando la peticion fue todo un exito.

Se uso _Gin_ como framework para el manejo de rutas y ORM(basico) basado en https://github.com/AJRDRGZ/db-query-builder para el manejo de la base de datos.

Se uso un logger para el manejo de errores y mensajes de salida, se uso _zap_ como logger. Esto nos ayuda para la trazabilidad(Basica) de donde se origino el error.


**APIs(Open api)**

En la carpeta _.encora/apim/openapi-spec.yaml_ se encuentra la documentación detallada de los endpoints, el tipo de dato de cada campo,
para qué sirve cada campo, etc. Esto es muy útil para que el equipo de front-end pueda consumir los servicios de manera correcta.

Para visualizar la documentación de los endpoints puede abrir el proyecto en su editor(debe tener instalado una extension de openAPI o swagger)
o puede utilizar el siguiente link: [Doc Endpoints](https://app.swaggerhub.com/apis/MIGUELSR1084/reto-financiera_compartamos/1.0.0)

**Base de datos**

Se uso _PostgreSQL_ en la base de datos, en el siguiente link([Dock db](https://dbdiagram.io/d/reto-financiera-6619eed803593b6b61e2a3ea)) se encuentra el diagrama de la base de datos,
con las tablas y las relaciones entre ellas. Ademas de que campos tiene cada tabla y la descripcion de cada columna.

---
_Requerimiento funcional_
```azure
CRUD de Clientes:
Listar: Mostrar la lista de todos los clientes registrados.
Agregar: Permitir la creación de nuevos clientes.
Modificar: Actualizar la información de un cliente existente.
Eliminar: Eliminar un cliente de la base de datos.
Campos de Cliente:
DNI: Campo clave único para identificar a cada cliente. Debe ser numérico y tener 8 dígitos.
Nombres: Nombre(s) del cliente.
Apellidos: Apellido(s) del cliente.
Fecha de Nacimiento: En formato dd/mm/yyyy.
Sexo: Puede ser M (masculino) o F (femenino).
Listado de Ciudades: El cliente puede pertenecer a cualquier ciudad. Las ciudades se obtendrán de una base de datos.
Restricciones:
Solo se podrá crear clientes mayores de 18 años.
Solo se podrá eliminar clientes mayores de 80 años.
```


**FIN.**
