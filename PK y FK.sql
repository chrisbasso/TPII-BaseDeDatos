
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