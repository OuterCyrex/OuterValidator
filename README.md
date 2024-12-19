# OuterValidator

**OuterValidator** provides a better experience for client to check their struct fields when using this it.

Based on *Reflection Package*, **OuterValidator** dynamically inspects the type of each field within the struct, If a field's value does not meet the specified constraints (i.e., Required but is None), **OuterValidator** will detect this and throw an error, alerting the user to the issue.

## Installation

To Install this Package, You need to enter the following command into your terminal.

```shell
go get github.com/outercyrex/outervalidator
```

if this fails, Please Check your GOPROXY to see if it is properly configured.

## Usage

OuterValidator Provides Some Key Constraints:

| Constrains | Usages                                                      |
| ---------- | ----------------------------------------------------------- |
| Required   | the value of this Field must not be None                    |
| Max        | the length of string in this Field must be smaller than Max |
| Min        | the length of string in this Field must be greater than Min |

To add those Constrains, you need to generate a Validator Object first, then Add the Constraints to certain Field.

When you want to check if the values of a struct violates any of their Constraints, just use the method **Check** whose receiver is Validator Object. If it returns False, You can check the error return value to see which field violated which constraint.

## Code Example

Here is a Code Example:

```go
package main

import (
	"fmt"
	"github.com/outercyrex/outervalidator/validate"
)

type User struct {
	Name     string
	Age      int
	Password string
}

func main() {
	v := validate.NewValidator(User{})
	v.Set("Name").Required()
	v.Set("Age").Required()
	v.Set("Password").Required().Min(4).Max(12)

	user := User{
		Name:     "Outer",
		Age:      21,
		Password: "123456",
	}

	ok, err := v.Check(user)
	if !ok {
		fmt.Println(err)
	}
}
```

## Features

**OuterValidator** Provides a New Way for client to check if their Object meet the Constraints.

Through chained calls, client can clearly add constraints to fields.

## Todo

The current Package is still under continuous development and update, and more constraints will be added in the future. Bugs and unreasonable aspects will also be modified.

Thanks for your stars and likes.