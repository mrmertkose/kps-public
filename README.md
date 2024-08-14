### **Usage**

This guide explains how to use the Kps GetService function and what behavior to expect.

###  **Installation**
`go get github.com/mrmertkose/kps-public`

###  **Parameters**

* identityNumber (int64): The identity number of the individual. Should be a valid integer value.
* name (string): The first name of the individual. This is a required field.
* surname (string): The last name of the individual. This is a required field.
* year (int32): The birth year of the individual. Should be a valid four-digit year.

###  **Return Values**
* bool: Returns true if the service operation is successful, otherwise false.
* error: Returns an error if there is a problem with the input parameters or the service operation.

```go
package main

import (
	"fmt"
	"github.com/mrmertkose/kps-public"
)

func main() {
	var identityNumber int64 = 1234567890
	var year int32 = 1990
	name := "John"
	surname := "Doe"

	service, err := kps.GetService(identityNumber, name, surname, year)
	if err != nil {
		fmt.Println("Error:", err)
	}
	
	fmt.Println(service)
}
```