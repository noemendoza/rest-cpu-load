# rest-cpu-load



### Makefile tarjets

- make run: ejecuta el projecto.

- make build: Crea el binario dentro del directorio `build` , el nombre del binario es `rest-cpu-load`



Para ejecutar a mano el binario (luego de haber ejecutado `make build`):  `./build/rest-cpu-load`



### Endpoint:

El binario abre el puerto 8080 y queda a la escucha en todas las intefaces del host, el path expuesto para el servicio es `/stress/:secs` donde `:secs` es un parametro numérico que respresenta el número de segundos que se desea estresar el host que se encuentra ejecutando el servicio.

Ejemplo:

```bash
curl -XGET http://127.0.0.1:8080/stress/10
```



El programa estresa la cpu del servidor entre un 40 y 60 por ciento.

