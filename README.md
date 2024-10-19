# nextpermutation
## Description
Go package that adds function that will create a next lexicographical permutation of given slice or string

## Features
Get next lexicographical permutation of numbers or runes
Get next lexicographical permutation of a string

## Usage
Calling `nextpermutation.Next` with parameters `slice address, start position, end position`

Permutation edits the original slice

Returns `true` if next lexicographicaly higher permutation exists, false if it does not

`end` parameter is set to the element if it's larger than the slice that is being edited
```go
nextpermutation.Next(&slice, start, end)
```
for example:
```go
slice := []int{1, 2, 3, 4}
fmt.Println(slice)
nextpermutation.Next(&slice, 0, len(slice))
fmt.Println(slice)
```
outputs:
```
[1 2 3 4]
[1 2 4 3]
```

```go
slice := []int{1, 2, 3, 4}
fmt.Println(slice)
for nextpermutation.Next(&slice, 0, len(slice)) {
	fmt.Println(slice)
}
```
outputs:
```
[1 2 3 4]
[1 2 4 3]
[1 3 2 4]
[1 3 4 2]
[1 4 2 3]
[1 4 3 2]
[2 1 3 4]
[2 1 4 3]
[2 3 1 4]
[2 3 4 1]
[2 4 1 3]
[2 4 3 1]
[3 1 2 4]
[3 1 4 2]
[3 2 1 4]
[3 2 4 1]
[3 4 1 2]
[3 4 2 1]
[4 1 2 3]
[4 1 3 2]
[4 2 1 3]
[4 2 3 1]
[4 3 1 2]
[4 3 2 1]
```

and for strings
```go
slice := []int{1, 2, 3, 4}
fmt.Println(slice)
nextpermutation.Next(&slice, 0, len(slice))
fmt.Println(slice)
```
outputs:
```
[1 2 3 4]
[1 2 4 3]
```
and for strings:
```go
str := "1234"
fmt.Println(str)
for nextpermutation.NextString(&str, 0, len(str)) {
	fmt.Println(str)
}
```
outputs:
```
1234
1243
1324
1342
1423
1432
2134
2143
2314
2341
2413
2431
3124
3142
3214
3241
3412
3421
4123
4132
4213
4231
4312
4321
```
