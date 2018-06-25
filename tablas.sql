create table cliente(
nrocliente int,
nombre text,
apellido text,
domicilio text,
telefono char(12)
);

create table tarjeta(
nrotarjeta char(12),
nrocliente int,
validadesde char(6), -- e.g. 201106
validahasta char(6),
codseguridad char(4),
limitecompra decimal(8,2),
estado char(10) -- `vigente', `suspendida', `anulada'
);

create table comercio(
nrocomercio int,
nombre text,
domicilio text,
codigopostal char(8),
telefono char(12)
);

create table compra(
nrooperacion serial,
nrotarjeta char(12),
nrocomercio int,
fecha timestamp,
monto decimal(7,2),
pagado boolean

);

create table rechazo(
nrorechazo serial,
nrotarjeta char(12),
nrocomercio int,
fecha timestamp,
monto decimal(7,2),
motivo text
);

create table cierre(
año int,
mes int,
terminacion int,
fechainicio date,
fechacierre date,
fechavto date
);

create table cabecera (
nroresumen serial,
nombre text,
apellido text,
domicilio text,
nrotarjeta char(12),
desde date,
hasta date,
vence date,
total decimal(8,2)
);

create table detalle (
nroresumen int,
nrolinea int,
fecha date,
nombrecomercio text,
monto decimal(7,2)
);

create table alerta (
nroalerta int,
nrotarjeta char(12),
fecha timestamp,
nrorechazo int,
codalerta int, -- 0:rechazo, 1:compra 1min, 5:compra 5min, 32:límite
descripcion text
);

-- Esta tabla no es parte del modelo de datos, pero se incluye para
-- poder probar las funciones.

create table consumo(
nrotarjeta char(12),
codseguridad char(4),
nrocomercio int,
monto decimal(7,2)
);
