# x
Help Go, es una librería para ayudar a extender la sintaxis de golang 

## install

```bash
go get github.com/jad21/x
```

## x.IF

```go
import "github.com/jad21/x"
...
if x.If(true, 1, 0) == 1{
    print("is true")
}
```

## x.SwitchMap

```go
cases := map[int]string{
    1: "one",
    2: "two",
    3: "three",
}

// Usando SwitchMap
fmt.Println(x.SwitchMap(1, cases, "default")) // Output: one
fmt.Println(x.SwitchMap(4, cases, "default")) // Output: default
```

## x.Find

```go
package main
import (
    "fmt"
    "github.com/jad21/x"
)

// Book representa un libro con un título y una categoría
type Book struct {
    Title    string
    Category string
}
func main() {
    // Lista de libros
    books := []Book{
        {Title: "The Pragmatic Programmer", Category: "dev"},
        {Title: "Clean Code", Category: "dev"},
        {Title: "The Lord of the Rings", Category: "fiction"},
    }

    // Criterio de búsqueda: encontrar el primer libro de la categoría "dev"
    devBook := x.Find(books, func(b Book) bool {
        return b.Category == "dev"
    })

    // Comprobar si se encontró un libro de la categoría "dev"
    if devBook != nil {
        fmt.Printf("Found book: %s\n", devBook.Title)
    } else {
        fmt.Println("No book found in the 'dev' category")
    }
}

```

## x.Reverse

```go
// Ejemplo con un array de enteros
numbers := []int{1, 2, 3, 4, 5}
fmt.Println("Original:", numbers)
fmt.Println("Reversed:", x.Reverse(numbers))

// Ejemplo con un array de strings
words := []string{"go", "is", "awesome"}
fmt.Println("Original:", words)
fmt.Println("Reversed:", x.Reverse(words))

```

## x.Contains

```go
    // Ejemplo con un array de enteros
numbers := []int{1, 2, 3, 4, 5}
fmt.Println("Contains 3:", x.Contains(numbers, 3)) // true

```

## x.Map

```go
    // Ejemplo con un array de enteros, duplicando cada valor
numbers := []int{1, 2, 3, 4, 5}
double := func(n int) int {
    return n * 2
}
doubledNumbers := x.Map(numbers, double)
fmt.Println("Original:", numbers) //     Original: [1 2 3 4 5]
fmt.Println("Doubled:", doubledNumbers) // Doubled: [2 4 6 8 10]

```

## x.Repeat

```go
results := x.Repeat("?", 3)
fmt.Println("Results:", results) //     Original: [? ? ?]
```

## x.Unique

```go
x.Unique([]int{1, 2, 2, 3, 4, 4, 5}) // [1 2 3 4 5]
```