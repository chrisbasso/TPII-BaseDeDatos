create table cliente(
nrocliente:int,
nombre:text,
apellido:text,
domicilio:text,
telefono:char(12)
);

create table tarjeta(
nrotarjeta:char(12),
nrocliente:int,
validadesde:char(6), -- e.g. 201106
validahasta:char(6),
codseguridad:char(4),
limitecompra:decimal(8,2),
estado:char(10) -- `vigente', `suspendida', `anulada'
);

create table comercio(
nrocomercio:int,
nombre:text,
domicilio:text,
codigopostal:char(8),
telefono:char(12)
):

create table compra(
nrooperacion:int,
nrotarjeta:char(12),
nrocomercio:int,
fecha:timestamp,
monto:decimal(7,2),
pagado:boolean
);

create table rechazo(
nrorechazo:int,
nrotarjeta:char(12),
nrocomercio:int,
fecha:timestamp,
monto:decimal(7,2),
motivo:text
);

create table cierre(
año:int,
mes:int,
terminacion:int,
fechainicio:date,
fechacierre:date,
fechavto:date
);

create table cabecera (
nroresumen:int,
nombre:text,
apellido:text,
domicilio:text,
nrotarjeta:char(12),
desde:date,
hasta:date,
vence:date,
total:decimal(8,2)
);

create table detalle (
nroresumen:int,
nrolinea:int,
fecha:date,
nombrecomercio:text,
monto:decimal(7,2)
);

create table alerta (
nroalerta:int,
nrotarjeta:char(12),
fecha:timestamp,
nrorechazo:int,
codalerta:int, -- 0:rechazo, 1:compra 1min, 5:compra 5min, 32:límite
descripcion:text
);

-- Esta tabla no es parte del modelo de datos, pero se incluye para
-- poder probar las funciones.

create table consumo(
nrotarjeta:char(12),
codseguridad:char(4),
nrocomercio:int,
monto:decimal(7,2)
);

alter table cliente add constraint cliente_pk primary key(nrocliente);

alter table tarjeta add constraint tarjeta_pk primary key(nrotarjeta);

alter table comercio add constraint comercio_pk primary key(nrocomercio);

alter table compra add constraint compra_pk primary key(nrooperacion);

alter table rechazo add constraint rechazo_pk primary key(nrorechazo);

alter table cierre add constraint cierre_pk primary key(año, mes, terminacion);

alter table cabecera add constraint cabecera_pk primary key(nroresumen);

alter table detalle add constraint detalle_pk primary key(nroresumen, nrolinea);

alter table alerta add constraint alerta_pk primary key(nrocliente);

alter table consumo add constraint consumo_pk primary key(nroalerta);


alter table tarjeta add constraint cliente_fk foreign key(nrocliente) references cliente(nrocliente);

alter table compra add constraint tarjeta_fk foreign key(nrotarjeta) references tarjeta(nrotarjeta);
alter table compra add constraint comercio_fk foreign key(nrocomercio) references tarjeta(nrocomercio);

alter table rechazo add constraint tarjeta_fk foreign key(nrotarjeta) references tarjeta(nrotarjeta);
alter table rechazo add constraint comercio_fk foreign key(nrocomercio) references tarjeta(nrocomercio);
alter table rechazo add constraint monto_compra_fk foreign key(monto) references compra(monto);
alter table rechazo add constraint fecha_compra_fk foreign key(fecha) references compra(fecha);

alter table cabecera add constraint nombre_cliente_fk foreign key(nombre) references cliente(nombre);
alter table cabecera add constraint apellido_cliente_fk foreign key(apellido) references cliente(apellido);
alter table cabecera add constraint domicilio_cliente_fk foreign key(domicilio) references cliente(domicilio);
alter table cabecera add constraint tarjeta_fk foreign key(nrotarjeta) references tarjeta(nrotarjeta);

alter table detalle add constraint fecha_compra_fk foreign key(fecha) references compra(fecha);
alter table detalle add constraint comercio_fk foreign key(nombrecomercio) references comercio(nombre);
alter table detalle add constraint monto_compra_fk foreign key(monto) references compra(monto);

alter table alerta add constraint tarjeta_fk foreign key(nrotarjeta) references tarjeta(nrotarjeta);
alter table alerta add constraint fecha_rechazo_fk foreign key(fecha) references rechazo(fecha);
alter table alerta add constraint nrorechazo_rechazo_fk foreign key(nrorechazo) references rechazo(nrorechazo);