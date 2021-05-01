# stack
Implementing tiny library of stack with >[golang](https://golang.org/)

## Definition
Una pila ([stack](https://es.wikipedia.org/wiki/Pila_(inform%C3%A1tica))) es una colección ordenada de elementos a los que sólo se puede acceder por un único lugar o extremo, llamado cima o tope de la pila. Las pilas son conocidas como estructuras de datos LIFO (last-in, first-out).

## Metodos Básicos
* Push - insertar elemento
* Pop - eliminar elemento
* Peek - mostrar elemento de la cima 

## Instalar y usar 
Usar comando *get* de golang para agregar el repositorio.
```golang
go get github.com/ing-thejis/stack
```
Para importar la libreria usa la siguiente línea

```golang
import "github.com/ing-thejis/stack"
```

## Ejemplo
```golang
package main

import "github.com/ing-thejis/stack"

func main(){
  pila := stack.NewStack()  //crea una nueva pila
  pila.Push(1)  //agrega un nuevo elemento de valor 1 en la pila
  pila.Push(2)
  pila.Push(3)
  
  fmt.Println(pila.Peek())  //muestra por pantalla la cima de la pila
  fmt.Println(pila.Pop()) //elimina y muestra el valor de la cima
}
``` 
## Autor
**The JIS** <ing.jesithgarcia@gmail.com>
