create or replace function autorizarCompra(numtarjeta char(12), codseguridad char(4), montoCompra numeric, numcomercio int) returns boolean as $$

declare

    vNroTarjeta char(12);
    vValidaHasta character(8);
    vCodSeguridad char(4);
    vLimiteCompra numeric(8,2);
	vEstado character(10);

    vNroComercio int;

    vMontoNoPagado numeric;

begin
 
    

    SELECT t.nrotarjeta, t.validahasta, t.codseguridad, t.limitecompra, t.estado
    INTO vNroTarjeta, vValidaHasta, vCodSeguridad, vLimiteCompra, vEstado
    FROM tarjeta t WHERE numtarjeta = t.nrotarjeta;

    SELECT c.nrocomercio INTO vNroComercio FROM comercio c WHERE numcomercio = c.nrocomercio;


    IF (vNroTarjeta IS NULL) THEN

        RAISE NOTICE 'Tarjeta no existe';
        RETURN FALSE;
        
    END IF;

    IF (vEstado = 'anulada') THEN
        
        RAISE NOTICE 'Tarjeta está anulada';
        INSERT INTO rechazo (nrotarjeta,nrocomercio,fecha,monto,motivo) values (vNroTarjeta,vNroComercio,CURRENT_TIMESTAMP,montoCompra,'Tarjeta anulada');
        RETURN FALSE;

    END IF;

    IF (vEstado = 'suspendida') THEN
        
        RAISE NOTICE 'Tarjeta suspendida';
        INSERT INTO rechazo (nrotarjeta,nrocomercio,fecha,monto,motivo) values(vNroTarjeta,vNroComercio,CURRENT_TIMESTAMP,montoCompra,'Tarjeta suspendida');
        RETURN FALSE;

    END IF;

    vValidaHasta := vValidaHasta || '01';
    

    IF ((cast(vValidaHasta as date)) < CURRENT_DATE) THEN

        RAISE NOTICE 'Tarjeta vencida';
        INSERT INTO rechazo (nrotarjeta,nrocomercio,fecha,monto,motivo) values(vNroTarjeta,vNroComercio,CURRENT_TIMESTAMP,montoCompra,'Tarjeta vencida');
        RETURN FALSE;

    END IF;
     

    IF (vCodSeguridad <> codseguridad) THEN

        RAISE NOTICE 'Código de seguridad invalido';
        INSERT INTO rechazo (nrotarjeta,nrocomercio,fecha,monto,motivo) values(vNroTarjeta,vNroComercio,CURRENT_TIMESTAMP,montoCompra,'Código de seguridad invalido'); 
        RETURN FALSE;

    END IF;
    
    vMontoNoPagado := (SELECT SUM(monto) FROM compra WHERE vNroTarjeta = compra.nrotarjeta AND compra.pagado = false);
    
    IF ((montoCompra > vLimiteCompra) OR ((montoCompra + vMontoNoPagado) > vLimiteCompra)) THEN
        
        RAISE NOTICE 'Supera el límite de compra';
        INSERT INTO rechazo (nrotarjeta,nrocomercio,fecha,monto,motivo) values(vNroTarjeta,vNroComercio,CURRENT_TIMESTAMP,montoCompra,'Supera el límite de compra');
        RETURN FALSE;

    END IF;

   
    
    RAISE NOTICE 'Compra aceptada';
    INSERT INTO compra (nrotarjeta,nrocomercio,fecha,monto,pagado) values (vNroTarjeta,vNroComercio,CURRENT_TIMESTAMP,montoCompra,false);
    RETURN TRUE;


end;
$$ language plpgsql;


create or replace function generarFactura(numcliente int, periodo date) returns void as $$

declare
    


    vNombre text;
    vApellido text;
    vDomicilio text;
    vNroTarjeta char(12);
    vTerminacion int;
    vFechaInicio date;
    vFechaCierre date;
    vFechaVto date;

    vConsumoTotal decimal(8,2);


    lista_tarjetas char(12)[] = (SELECT ARRAY (SELECT nrotarjeta FROM tarjeta WHERE nrocliente = numcliente));
    numeroT char(12);
    vNumeroLinea int = 0;

    vCompra record;
    vNroResumen int;

    vFechaCompra date;
    vNombreComercio text;
        

begin

    SELECT c.nombre, c.apellido, c.domicilio 
    INTO vNombre, vApellido, vDomicilio
    FROM cliente c WHERE numcliente = c.nrocliente;


    
    FOREACH numeroT IN ARRAY lista_tarjetas LOOP
    
        vNroTarjeta := numeroT;

        vTerminacion:= substring(vNroTarjeta from 12 for 1)::int;
    
        SELECT c.fechainicio, c.fechacierre, c.fechavto
        INTO vFechaInicio, vFechaCierre, vFechaVto
        FROM cierre c 
        WHERE c.terminacion = vTerminacion 
        AND EXTRACT(MONTH FROM periodo) = c.mes
		AND EXTRACT(YEAR FROM periodo) = c.año;

        vConsumoTotal := (SELECT SUM(monto) FROM compra WHERE vNroTarjeta = compra.nrotarjeta 
                            AND EXTRACT(MONTH from vFechaInicio) = EXTRACT(MONTH from compra.fecha)
                            AND EXTRACT(YEAR from vFechaInicio) = EXTRACT(YEAR from compra.fecha));
                        
        

        INSERT INTO cabecera (nombre, apellido, domicilio, nrotarjeta, desde, hasta, vence,total) 
               values (vNombre, vApellido, vDomicilio, vNroTarjeta,vFechaInicio,vFechaCierre,vFechaVto,vConsumoTotal);

        
       

        SELECT cabecera.nroresumen INTO vNroResumen FROM cabecera 
        WHERE cabecera.nrotarjeta = vNroTarjeta
        AND EXTRACT(MONTH from vFechaInicio) = EXTRACT(MONTH from cabecera.desde)
        AND EXTRACT(YEAR from vFechaInicio) = EXTRACT(YEAR from cabecera.desde);

        

        FOR vCompra IN SELECT c.nrotarjeta, c.nrocomercio, fecha, monto 
        FROM compra c WHERE vNroTarjeta = c.nrotarjeta
        AND EXTRACT(MONTH from vFechaInicio) = EXTRACT(MONTH from c.fecha)
        AND EXTRACT(YEAR from vFechaInicio) = EXTRACT(YEAR from c.fecha) LOOP

            SELECT comercio.nombre INTO vNombreComercio FROM comercio WHERE comercio.nrocomercio = vCompra.nrocomercio;                  

            INSERT INTO detalle (nroresumen,nrolinea,fecha,nombrecomercio,monto) 
            values (vNroResumen, vNumeroLinea, vCompra.fecha,vNombreComercio, vCompra.monto);
            
            vNumeroLinea := vNumeroLinea + 1;            
    
        end loop;


    end loop;



end;
$$ language plpgsql;

create or replace function alerta_rechazo() returns trigger as $$

declare

    alertaAnt record;

begin
	
    SELECT * INTO alertaAnt FROM alerta
    WHERE alerta.nrotarjeta = new.nrotarjeta
    AND cast(new.fecha as date) = cast(alerta.fecha as date)
    AND alerta.descripcion = 'Supera el límite de compra'
    AND new.motivo = 'Supera el límite de compra';

   
    IF FOUND THEN
		
        insert into alerta (nrotarjeta,fecha, nrorechazo, codalerta, descripcion) 
        values (new.nrotarjeta, new.fecha, new.nrorechazo, 32, 'Tarjeta suspendida por exceso de límite en el mismo día');

        update tarjeta set estado='suspendida' where nrotarjeta = new.nrotarjeta;

        ELSE

        insert into alerta (nrotarjeta,fecha, nrorechazo, codalerta, descripcion) values (new.nrotarjeta, new.fecha, new.nrorechazo, 0, new.motivo);
	
	END IF;

    
    return new;

end;

$$ language plpgsql;

create trigger alerta_rechazo_trg
after insert
on rechazo
for each row
execute procedure alerta_rechazo();

create or replace function alerta_compra() returns trigger as $$

declare

    compraAnt record;
    compraAnt2 record;

begin

    SELECT * INTO compraAnt FROM compra, comercio WHERE compra.nrocomercio = comercio.nrocomercio
    AND compra.fecha > CURRENT_TIMESTAMP - (1 * interval '1 minute')
	AND compra.nrotarjeta = new.nrotarjeta
	AND compra.nrocomercio <> new.nrocomercio
	AND comercio.codigopostal = (SELECT codigopostal FROM comercio WHERE nrocomercio = new.nrocomercio);

    IF FOUND THEN

        insert into alerta (nrotarjeta,fecha, codalerta, descripcion) 
        values (new.nrotarjeta, new.fecha, 1, 'Dos compras en menos de un minuto');

    END IF;

    SELECT * INTO compraAnt2 FROM compra, comercio WHERE compra.nrocomercio = comercio.nrocomercio
    AND compra.fecha > CURRENT_TIMESTAMP - (5 * interval '1 minute')
	AND compra.nrotarjeta = new.nrotarjeta
	AND compra.nrocomercio <> new.nrocomercio
	AND comercio.codigopostal <> (SELECT codigopostal FROM comercio WHERE nrocomercio = new.nrocomercio);

    IF FOUND THEN

        insert into alerta (nrotarjeta,fecha, codalerta, descripcion) 
        values (new.nrotarjeta, new.fecha, 5, 'Dos compras en CP diferentes');

    END IF;

    
    return new;

end;

$$ language plpgsql;

create trigger alerta_compra_trg
after insert
on compra
for each row
execute procedure alerta_compra();


create or replace function realizarConsumosTest() returns void as $$

declare

    unConsumo record;
    cur_consumos CURSOR FOR SELECT * FROM consumo;
        
begin

    OPEN cur_consumos;
    LOOP
        FETCH cur_consumos INTO unConsumo;
        EXIT WHEN NOT FOUND;

        PERFORM autorizarCompra(unConsumo.nrotarjeta,unConsumo.codseguridad,unConsumo.monto, unConsumo.nrocomercio);

    END LOOP;

    RETURN;
end;
$$ language plpgsql;







