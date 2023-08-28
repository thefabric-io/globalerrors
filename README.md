# Global errors golang package

## Table of Contents
1. [Overview](#overview)
2. [Installation](#installation)
3. [Usage](#usage)
    - [HTTP Status Codes](#http-status-codes)
    - [gRPC Status Codes](#grpc-status-codes)
4. [API](#api)
    - [HTTPStatus](#httpstatus)
    - [GRPCStatus](#grpcstatus)
5. [Custom Errors](#custom-errors)
6. [Contributing](#contributing)
7. [License](#license)

## Overview

The `globalerrors` package provides a convenient way to map custom errors to both HTTP and gRPC status codes. This helps in ensuring consistency in error handling across different service architectures.

## Installation

To install this package, you can run:

```bash
go get -u github.com/thefabric-io/globalerrors
```

## Usage

### HTTP Status Codes

The package offers an `HTTPStatus` function which accepts a custom error as an argument and returns the corresponding HTTP status code.

```go
import (
  "github.com/thefabric-io/globalerrors"
)

func main() {
  err := globalerrors.BadRequest
  statusCode := globalerrors.HTTPStatus(err)
  // statusCode will be http.StatusBadRequest (400)
}
```

### gRPC Status Codes

Similarly, the `GRPCStatus` function accepts a custom error and returns the gRPC status code.

```go
import (
  "github.com/thefabric-io/globalerrors"
)

func main() {
  err := globalerrors.BadRequest
  statusCode := globalerrors.GRPCStatus(err)
  // statusCode will be codes.InvalidArgument
}
```

## API

### HTTPStatus

Syntax:

```go
func HTTPStatus(customError error) int
```

- `customError`: The custom error for which you want to find the HTTP status code.

Returns: Corresponding HTTP status code as an integer.

### GRPCStatus

Syntax:

```go
func GRPCStatus(customError error) codes.Code
```

- `customError`: The custom error for which you want to find the gRPC status code.

Returns: Corresponding gRPC status code.


### Custom Error Example

You can define custom errors that implement the `error` interface and use them with the `globalerrors` package. Here's an example:

```go
import (
	"errors"
	
	"github.com/thefabric-io/globalerrors"
)

func ErrIDIsRequired() error {
	return errIDIsRequired{}
}

type errIDIsRequired struct{}

func (e errIDIsRequired) Error() string {
	return "id is required"
}

func (e errIDIsRequired) Is(target error) bool {
	return errors.Is(target, globalerrors.InternalServerError) || e == target
}

// Usage
func main() {
	err := ErrIDIsRequired()
	httpStatusCode := globalerrors.HTTPStatus(err)
	grpcStatusCode := globalerrors.GRPCStatus(err)

	// Since `ErrIDIsRequired` is defined to be equivalent to `globalerrors.InternalServerError`,
	// httpStatusCode will be http.StatusInternalServerError (500)
	// grpcStatusCode will be codes.Internal
}
```

In this example, we have defined a custom error `ErrIDIsRequired` that implements the `error` interface. We have also defined an `Is` method to make it compatible with `errors.Is`, enabling it to be used with the `globalerrors` package.

---

## Custom Errors Extensions

You can also extend the package to include your custom errors and their corresponding status codes.

1. Add your custom error in the `globalErrors` slice.

```go
var globalErrors = []error{
  ...
  YourCustomError,
  ...
}
```

2. Update `httpStatusCodesMap` and/or `grpcStatusCodesMap` to include the status code mapping for your custom error.

```go
var httpStatusCodesMap = map[error]int{
  ...
  YourCustomError: http.YourStatusCode,
  ...
}

var grpcStatusCodesMap = map[error]codes.Code{
  ...
  YourCustomError: codes.YourStatusCode,
  ...
}
```

## Contributing

We welcome contributions to this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.

---
For further details and queries, please feel free to reach out.