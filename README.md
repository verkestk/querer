# querer&nbsp;[![Build Status](https://travis-ci.org/verkestk/querer.svg?branch=master)](https://travis-ci.org/verkestk/querer)&nbsp;[![godoc reference](https://godoc.org/github.com/verkestk/querer?status.png)](https://godoc.org/github.com/verkestk/querer)

**querer**, _Spanish_, to want, to love, to wish, to like

Querer is a golang package for when you _want_ to populate a struct based on URL query string values.

## Usage

```
type params struct {
	MinDOB   time.Time `query:"min_dob"`
	MaxDOB   time.Time `query:"max_dob"`
	MaxAge   uint      `query:"max_age"`
	LastName string    `query:"last_name"`
}
```

```
params := new(params)
err := UnmarshalQuery(params, req.URL.Query())
if err != nil {
	...
}

...
```

## Supported Types

### Basic Types

This package currently supports:

- bool
- int
- uint
- float64
- string
- time.Time

If you have a need for more types, please open an issue or - even better - a pull request!

### Embedded/Anonymous Structs

Embedded/anonymous structs are also supported. For example.

```
type baseParams struct {
	Param1 int `query:"param_1"`
}

type params struct {
	*baseParams
	Param2 int `query:"param_2"`
}
```