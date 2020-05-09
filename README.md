# slice

It is utils package for slice.

## strs

It is for string slice.

### Examples

```go
Contains([]string{"a", "b", "c"}, "a") // true

Uniq([]string{"a", "b", "a", "b", "c"}) // []string{"a", "b", "c"}

Remove([]string{"a", "b", "c"}, "b") // []string{"a", "c"}

Sub([]string{"a", "b", "c"}, []string{"a", "c"}) // []string{"b"}
```
