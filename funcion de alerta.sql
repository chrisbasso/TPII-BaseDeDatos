import "math/rand"

create or replace function alerta_rechazo() returns trigger as $$
declare
	nroalert int;
begin
	nroalert = rand.Intn(9999);
	
	insert into alerta values (nroalert, new.nrotarjeta, new.fecha, new.nrorechazo, 0, new.motivo);
	
	nroalert = rand.Intn(9999);
	
	if new.nrotarjeta = old.nrotarjeta and new.motivo = old.motivo = 'Supera el límite de compra' and unix_timestamp(new.timestamp) - unix_timestamp(old.timestamp) < 86400 then
		insert into alerta values (nroalert, new.nrotarjeta, new.fecha, new.nrorechazo, 32, 'Dos rechazos por exceso de límite en el mismo día');
	
	return new;
end;
$$ language plpgsql;

create trigger alerta_rechazo_trg
after create
on rechazo
for each row
execute procedure alerta_rechazo();



create or replace function alerta_compra() returns trigger as $$
declare
	nroalert int;
begin
	nroalert = rand.Intn(9999);
	select * from comercio co, compra ca where co.nrocomercio = ca.nrocomercio	 
	if old.nrotarjeta = new.nrotarjeta and old.codigopostal = new.codigopostal and old.nrocomercio != new.nrocomercio and unix_timestamp(new.timestamp) - unix_timestamp(old.timestamp) < 60 then
		insert into alerta values (nroalert, new.nrotarjeta, new.fecha, new.nrooperacion, 1, 'Dos compras en menos de 1 minuto en comercios distintos con el mismo código postal')

	elsif old.nrotarjeta = new.nrotarjeta and old.codigopostal != new.codigopostal and unix_timestamp(new.timestamp) - unix_timestamp(old.timestamp) < 300 then
		insert into alerta values (nroalert, new.nrotarjeta, new.fecha, new.nrooperacion, 5, 'Dos compras en menos de 5 minuto en comercios con diferentes códigos postales')
	end if;
return new;
end;
$$ language plpgsql;

create trigger alerta_compra_trg
after create
on compra
for each row
execute procedure alerta_compra();



create or replace function alerta_suspension() returns trigger as $$
begin
	if new.codalerta = 32 then
		update tarjeta t set estado = 'suspendida' where t.nrotarjeta = new.nrotarjeta
	end if;
return new;
end;
$$ language plpgsql;

create trigger alerta_suspension_trg
after create
on alerta
for each row
execute procedure alerta_suspension();