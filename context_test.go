package main

import (
	"context"
	"fmt"
	"testing"
)

func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")

	fmt.Println(contextA) // => context.Background
	fmt.Println(contextB) // => context.Background.WithValue(b, B)
	fmt.Println(contextC) // => context.Background.WithValue(c, C)
	fmt.Println(contextD) // => context.Background.WithValue(b, B).WithValue(d, D)
	fmt.Println(contextE) // => context.Background.WithValue(b, B).WithValue(e, E)
	fmt.Println(contextF) // => context.Background.WithValue(c, C).WithValue(f, F)

	fmt.Println(contextF.Value("f")) // dapat => F
	fmt.Println(contextF.Value("c")) // dapat milik parent => C
	fmt.Println(contextF.Value("b")) // tidak dapat, beda parent => nil
	fmt.Println(contextA.Value("b")) // tidak bisa mengambil data child => nil
}
