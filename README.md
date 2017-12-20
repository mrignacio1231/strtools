# Strtools

A package for manipulate strings 

#### Installation

```bash
$ go get github.com/mrignacio1231/strtools
```


#### Import package in your project
Add following line in your `*.go` file:
```go
import "github.com/mrignacio1231/strtools"
```
If you are unhappy to use long `strtools`, you can do something like this:
```go
import (
  str "github.com/mrignacio1231/strtools"
)
```

#### Examples
Go code with Substr() function
```go

strtools.Substr("Hello World", 0, 4) // Output: Hello

// Support for negative index
strtools.Substr("Hello World", -5, 11) // Output: World
```

Go code with IsString() function
```go
strtools.IsString("I'm a string") // Output: true
strtools.IsString(12.4) // Output: false
```

Generate a random string with Random() function
```go
strtools.Random() // Return a random string
```