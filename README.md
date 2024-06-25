# hgo
Help Go, es una librería para ayudar a extender la sintaxis de golang 

## hgo.IF

```go
    import h "github.com/jad21/hgo"
    ...
    if h.If(true, 1, 0) == 1{
        print("is true")
    }
```

## hgo.Find

```go
    package main
    import (
        "fmt"
        h "github.com/jad21/hgo"
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
        devBook := h.Find(books, func(b Book) bool {
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