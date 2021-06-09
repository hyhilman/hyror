# PQ Error

Simple error show stack on print

## Usage

### Printing with stack
```
if err := myAmazingFunction(); err != nil {
    return NewError(err)
}
```


### Printing error without stack
```
var finalError error
if err := myAmazingFunction(); err != nil {
    finalError = NewError(err)
}

finalError.(PQError).ErrorWithoutStack()
```


### Example output

```
with stack // Error message
        pq-error.NewError // stack
        	/home/hyhilman/go/src/pq-error/error.go:25
        pq-error.TestStackPrintWithoutStack
        	/home/hyhilman/go/src/pq-error/error_test.go:88
        testing.tRunner
        	/usr/lib/golang/src/testing/testing.go:1039
        runtime.goexit
        	/usr/lib/golang/src/runtime/asm_amd64.s:1373
```
