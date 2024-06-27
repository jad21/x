# hgo
Help Go, es una librería para ayudar a extender la sintaxis de golang 

## hgo.IF

```go
    import z "github.com/jad21/hgo"
    ...
    if z.If(true, 1, 0) == 1{
        print("is true")
    }
```

## hgo.Find

```go
    package main
    import (
        "fmt"
        z "github.com/jad21/hgo"
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
        devBook := z.Find(books, func(b Book) bool {
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

## hgo.Reverse

```go
    // Ejemplo con un array de enteros
    numbers := []int{1, 2, 3, 4, 5}
    fmt.Println("Original:", numbers)
    fmt.Println("Reversed:", hgo.Reverse(numbers))

    // Ejemplo con un array de strings
    words := []string{"go", "is", "awesome"}
    fmt.Println("Original:", words)
    fmt.Println("Reversed:", hgo.Reverse(words))

```

## hgo.Contains

```go
     // Ejemplo con un array de enteros
    numbers := []int{1, 2, 3, 4, 5}
    fmt.Println("Contains 3:", hgo.Contains(numbers, 3)) // true

```

## hgo.Map

```go
     // Ejemplo con un array de enteros, duplicando cada valor
    numbers := []int{1, 2, 3, 4, 5}
    double := func(n int) int {
        return n * 2
    }
    doubledNumbers := hgo.Map(numbers, double)
    fmt.Println("Original:", numbers) //     Original: [1 2 3 4 5]
    fmt.Println("Doubled:", doubledNumbers) // Doubled: [2 4 6 8 10]

```

## hgo.Repeat

```go
     
    results := hgo.Repeat("?", 3)
    fmt.Println("Results:", results) //     Original: [? ? ?]
```