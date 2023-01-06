# NOTES

- go test //corre los tests
- go test -cover //muestra el coverage
- go test -coverprofile=coverage //Genera un archivo dificil de leer
- go tool cover -func=coverage //interpreta el archivo creado
- go tool cover -html=coverage //vrea un html legible con el coverage

## Profiling

### El profiling nos permite ver que partes de nuestro codigo ocupan mas recursoso para optimizarlas

- go test -cpuprofile=cpu
- go tool pprof cpu
