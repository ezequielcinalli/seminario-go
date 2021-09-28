# Trabajo práctico seminario GoLang-Tudai 2021

## Integrantes del grupo

- Ballone, Gabriel
- Cinalli, Ezequiel

## Resumen del trabajo realizado

Para la resolución del [trabajo propuesto](https://github.com/juanpablopizarro/tudai2021) inicialmente se definió un modulo con *go mod tpe*.

Dentro del paquete *parser* se definio:

- una estructura *Result* la cual almacena el resultado del parseo conteniendo el tipo, valor y longitud
- una funcion publica *GetEstructura* la cual recibe una cadena de caracteres y devuelve dos parametros: un puntero a una estructura de tipo resultado y un error. La misma se encarga del parseo y en caso de encontrar errores en el formato de la cadena recibida. Ya sean caracteres en lugar de números o menos de 5 caracteres de longitud, devolverá nil como primer parametro y el error encontrado como segundo parámetro.

Se realizaron pruebas en la función *main* y ademas se agregó un test unitario utilizando la libreria [testify](https://github.com/stretchr/testify) el cual cubrio el 100% como se puede visualizar de forma gráfica en el [html](https://ezequielcinalli.github.io/seminario-go) generado gracias a la herramienta *go tool cover*.
