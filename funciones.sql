create or replace function autorizarCompra(numtarjeta character, codseguridad character, montoCompra numeric, numcomercio int) returns boolean as $$

declare

    vNroTarjeta char(12);
    vValidaHasta character(6);
    vCodSeguridad char(4);
    vLimiteCompra numeric(8,2);
	vEstado character(10);

    vNroComercio int;

    vMontoNoPagado numeric;

begin
 
    RAISE NOTICE 'Se invoca función autorizacion_compra';

    SELECT t.nrotarjeta, t.validahasta, t.codseguridad, t.limitecompra, t.estado
    INTO vNroTarjeta, vValidaHasta, vCodSeguridad, vLimiteCompra, vEstado
    FROM tarjeta t WHERE numtarjeta = t.nrotarjeta;

    SELECT c.nrocomercio INTO vNroComercio FROM comercio c WHERE numcomercio = c.nrocomercio;


    IF (vNroTarjeta IS NULL) THEN

        RAISE NOTICE 'Tarjeta no existe';
        RETURN FALSE;
        
    END IF;

    IF (vEstado <> 'vigente') THEN
        
        RAISE NOTICE 'Tarjeta no está vigente';
        INSERT INTO rechazo (nrotarjeta,nrocomercio,fecha,monto,motivo) values (vNroTarjeta,vNroComercio,CURRENT_TIMESTAMP,montoCompra,'Tarjeta no vigente');
        RETURN FALSE;

    END IF;

    IF (vCodSeguridad <> codseguridad) THEN

        RAISE NOTICE 'Código de seguridad invalido';
        INSERT INTO rechazo (nrotarjeta,nrocomercio,fecha,monto,motivo) values(vNroTarjeta,vNroComercio,CURRENT_TIMESTAMP,montoCompra,'Código de seguridad invalido'); 
        RETURN FALSE;

    END IF;
     
    vMontoNoPagado := (SELECT SUM(monto) FROM compra WHERE vNroTarjeta = compra.nrotarjeta AND compra.pagado = false);
    
    IF ((vMontoNoPagado + montoCompra) > vLimiteCompra) THEN
        
        RAISE NOTICE 'Supera el límite de compra';
        INSERT INTO rechazo (nrotarjeta,nrocomercio,fecha,monto,motivo) values(vNroTarjeta,vNroComercio,CURRENT_TIMESTAMP,montoCompra,'Supera el límite de compra');
        RETURN FALSE;

    END IF;

    IF (vEstado = 'suspendida') THEN
        
        RAISE NOTICE 'Tarjeta suspendida';
        INSERT INTO rechazo (nrotarjeta,nrocomercio,fecha,monto,motivo) values(vNroTarjeta,vNroComercio,CURRENT_TIMESTAMP,montoCompra,'Tarjeta suspendida');
        RETURN FALSE;

    END IF;
    
    RAISE NOTICE 'Compra aceptada';
    INSERT INTO compra (nrotarjeta,nrocomercio,fecha,monto,pagado) values (vNroTarjeta,vNroComercio,CURRENT_TIMESTAMP,montoCompra,false);
    RETURN TRUE;


end;
$$ language plpgsql;


