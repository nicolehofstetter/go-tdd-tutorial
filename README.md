#Go Tutorial

This is a repo for the go tutorial: https://quii.gitbook.io/learn-go-with-tests
Following sections contains notes to find concepts later

## Hello, World

Find code in package [main](./main)

- Basic Tests
- ifs
- switch
- const
- named return values
- variable declaration and initialization with ":="

## Integers

Find code in package [integers](./integers)

- integers
- documentation with examples

## Iteration / Repeat

Find code in package [iteration](./iteration)

- for loop with variations
- variable declaration with "var name type"
- benchmarks

## Arrays and Slices

- init array
- for with range ( like a for each)
- slices (arrays without fixed size)
- ```go test -cover``` to see test coverage

- ```reflect.DeepEqual(a,b)``` to compare arrays
- ```make([]int, len(inputNumbers))``` creates a new slice with a starting capacity

- scoped functions ```checkSums := func(...){...}```

- copy arrays

```go 
    x := [3]string{"Лайка", "Белка", "Стрелка"}

	y := x[:] // slice "y" points to the underlying array "x"

	z := make([]string, len(x))
	copy(z, x[:]) // slice "z" is a copy of the slice created from array "x"

```

## Additional Links

[Package Structure in Go](https://dave.cheney.net/2014/12/01/five-suggestions-for-setting-up-a-go-project)
[Additional for loops](https://gobyexample.com/for)
[Blog Post slices](https://go.dev/blog/slices-intro)
