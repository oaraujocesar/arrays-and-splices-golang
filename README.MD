<div align="center">
  <h1>
    <img src="https://raw.githubusercontent.com/ashleymcnamara/gophers/master/LION_GOPHER.png" width="192" /><br />
    Abstract
  </h1>
</div>

## Arrays

Arrays in go have a predetermined length and are declared like this:

`arr := [N]type{value1, value2, value3, ..., valueN}`

or you can pass `...` instead of the N at the beginning:

`arr := [...]type{value1, value2, value3, ..., valueN}`


## Slices

Slices are declared almost the same way as the Arrays. The only thing that is different is the length number.

When declaring a slice you **do not** pass it:

`arr := []type{value1, value2, value3, ..., valueN}`

```go
  // as simple as you can see
  arr := []int{1, 2, 3, 4}
```

## Blank Identifier 

Go does not allow you to have unused variables inside your code.
So whenever you try to call a function and it returns more than one value, you can ignore one of those returns using the
_blank identifier_

```go
func Add(numbers [5]int) (sum int) {
  for _, number := range numbers {
    sum += number
  }
}
```

Using `range` to iterate over an array, two things are returned: the index (the position of that given number at the time) and the number itself and we're not going to use the index on this piece of code for instance so we can replace the `index` for an underline `_`.

This is us telling Go to ignore whatever is coming on that position.