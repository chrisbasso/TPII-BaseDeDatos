# Base de Datos I: Trabajo Práctico 2
Basso Christian Darìo <christian.d.basso@gmail.com>


## Introducción

El trabajo práctico consiste en la creación de una base de datos relativa a los procesos que son requeridos para realizar compras con tarjeta de crédito.

Se utilizara una base de datos relacional en PostgreSQL y una base de datos NoSQL, con la finalidad de comparar ambos modelos de base de datos.

Para la carga de los datos, las tablas, las funciones y la generación de compras a modo de testear la base de datos se realizará en una aplicación CLI hecha en Go


## Descripción

La base de datos SQL contará con algunas funcionalidades en base a ciertos requisitos:

- A partir de los datos de una compra se puede autorizarla o rechazarla, las compras que son rechazadas se cargan en una tabla de rechazos.
 
 *  La compra puede ser rechazada si la tarjeta no está vigente, sus podibles estados son suspendida, expirada en su fecha de vencimiento o anulada
 *  Es rechazada si al momento de la compra el numero de la tarjeta o su codigo de seguridad es invalido.
 *  Es rechazada si la compra mas el monto adeudado superan el límite de compra de la tarjeta.
 
- Se generan alertas por posibles fraudes, que pueden ser:

  * Si se supera el límite de compra dos veces en un día, en este caso además la tarjeta cambia su estado a suspendida.
  * Si se realizan dos compras en menos de un minuto en dos locales diferentes con mismo codigo postal.
  * Si se realizan dos compras en menos de cinco minutos en dos locales con diferente codigo postal.
  
- Se podrá generar facturas por periodo y por cliente y por cada tarjeta que éste posea:
  
  * Tendra una cabecera, que corresponde a los datos del cliente, la tarjeta y el monto total a pagar.
  * Y un detalle, en el cual se detallan todas las compras pendientes de pago para el periodo especificado.

## Implementación

*Base de datos SQL:*

En archvos '.SQL' se encuentra el cóigo para cada una de las opciones a elegir por el usuario en nuestra aplicación. 
Existen los siguientes archivos

    * TABLAS = la ejecución de CREATE TABLE, una por cada tabla que debe ser creada según nuestro modelo de datos.
    * PK y FK = realiza la ejecución de ALTER TABLE ADD CONSTRAINT para agregar las restricciones de primary key y foreign key requeridas para el modelo de datos.
    * DATOS = la ejecución de INSERT INTO TABLE, los datos para las tablas de clientes, tarjetas, comercios, cierres y compras.
    * FUNCIONES = contiene el código de las funciones para autorizar las compras, generar las facturas y los triggers para generar las alertas luego de hacer una compra.

Desde la aplicación CLI escrita en Go, la cual mediante un FOR que espera el resultado de un Scan del usuario imprime un menu con opciones hasta que el usuario ingrese
una de las opciones númericas. Entonces para cada opción elegida ingresa en distintos condicionales dentro del FOR.

    * Para carga de tablas, constraints, datos y funciones lee cada uno de los archivos .sql y los ejecuta linea por linea.
    * Para testear las funciones y triggers realiza las compras haciendo un Query llamando a cada función con parámetros de compras que fuercen a la base de datos a generar alertas y rechazos.
    * Luego posee Query para mostrar en pantalla los datos cargados en las tablas.

*Base de datos NOSQL:*

Se cargan los datos de clientes, tarjetas, comercios y compras en una base de datos BoltDB mediante documentos JSON y al igual que en nuestra aplicación anterior
se muestra un menu con opciones de cargar y mostrar los datos.

Por cada tipo de dato se crean Struct y se hardcodean los datos en una lista por cada tipo de dato, y luego se recorre la lista con un FOR en el cual por cada elemento lo carga y lo muestra en pantalla.

Se trata de una implementación simple a los fines de comparar la carga de datos de una base de datos relacional y una no relacional.

## Conclusiones

Como corolario al finalizar el trabajo práctico se puede decir que fue de mucha ayuda para aprender en la práctica a realizar un modelo completo en funcionamiento, desde la creación
del modelo de datos, la implementación de Stored Procedures y Triggers, hasta la conexión de la base de datos con una aplicación. 

Además se pueden apreciar las diferencias entre ambas bases de datos. Un detalle notable entre SQL Y NoSQL es que NoSQL no se rige por la normalización como lo haciamos en SQL, por lo que la base de datos 
no se rige por tener los datos relacionados entre si esto puede traer problema de redundancia e integridad de los datos. No tenemos las restricciones de primary key ni foreign key
por lo que realizar la comprobación de las restricciones podría resultar algo engorroso. En NoSQL ademas no hay schemas determinados, por lo que una entidad puede tener mas atributos en para un dato determinado que para otro de la misma entidad.  
