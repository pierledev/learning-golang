# Go Language (Go-Lang/Golang/Go)

This repo documents my personal notes on learning Go language. I picked up the language by following tutorials from [Pak Eko]('https://www.youtube.com/watch?v=IO_vkyJnMas&list=PL-CtdCApEFH-0i9dzMzLw6FKVrFWv3QvQ'). Since the tutorial is in Indonesian, I've created this English notes for those who want to learn from his teachings but may not speak Indonesian.

## History

- Developed by Google using the C programming language.
- Released to the public as open source in 2009.
- Gained popularity, particularly after being utilized to build Docker in 2011.
- Presently, numerous cutting-edge technologies are developed using Go, surpassing the usage C. Notable projects include Kubernetes, Promotheus, CockroachDB, etc.
- Currently gaining popularity as the preffered language for constructing backend APIa in microservices.

## Why Learn Go?

- Go is straightforward, requiring minimal time to grasp.
- Go supports effective concurrency programming, which aligns well with the current era of multicore processors.
- Go features built-in _garbage collection_, eliminating the need for manual memory management as required in languages like C.
  - _Garbage collection_ refers to the automatic management of computer memory. In languages like Go, the garbage collector is a mechanism that automatically identifies and frees up memory that is no longer in use or needed by the program. This process helps developers avoid the manual handling of memory allocation and deallocation, reducing the risk of memory leaks and making the development process more efficient. Essentially, garbage collection in Go takes care of cleaning up unsused memory, allowing developers to focus more on writing code and less on memory management details.
- It is one of the trending programming languages in contemporary times.

## Go Program Development Process

Our working file, _main.go_ -> is compiled by Go Complier -> producing the binary file _main_

Go supports compilation for various operating systems such as MacOS, Linux, and Windows.

## Creating A Project

- A Project in Go is typically referred to as a _module_.
- To create a _module_, use the following command in the folder where we want to build it:
  ```terminal
    go mod init <module-name>
  ```
- Typically. the folder where we want to build the _module_ shares the same name with the _module_ itself. For example, if our folder is named "learning-golang", we can create the module by executing the following command:
  ```terminal
    go mod init learning-golang
  ```

## Main Function

- Go is quite similar to C/C++ programming languages in that it requires a _main_ function.
- The _main_ function is a function that gets executed when the program is running.
- To define a function, the keyword used is _func_.
- The _main_ function should be placed in a _main package_.
- In Go, semicolons are not mandatory; we can choose whether or not to use them at the end of our code lines.
- Go is case sensitive.

Example:

```go
  // helloworld.go file

  package main

  func main() {
    // ...
  }
```

## Println

- To write and output a sentence to terminal, we first need to import the _fmt_ module.
- This is pretty similar to Java.
- I think it's like _console.log_ in JavaScript

Example:

```go
  // helloworld.go file

  package main

  import "fmt"

  func main() {
    fmt.Println("Hello, world!")
  }
```

## File Compilation

- Run the following command:
  ```go
    go build
  ```
- With that command, this project will attempt to find _main_ function in any files and compile it into a program that matches our current operating system. The output file will have the same name as the project. For example, if we named our project "learning-golang", the the output file will also be named "learning-golang". Subsequently, the output file can be run using terminal. On MacOs or Linux:
  ```terminal
  ./learning-golang
  ```
  On Windows:
  ```terminal
  ./learning-golang.exe
  ```

## Without Compilation

It is used when developers are in the app development process.

```terminal
  go run <file-name>
```

## Multiple Main Function

- In Go, each function within a _module_/_project_ must have a unique name. This means that we cannot create functions with the same name.
- Therefore, if we create a new file, for example, _sample.go_, and then attempt to create a function with the same name, _main_, we won't be able to build the module because the _main_ function duplicates the one already existing in the main function of _helloworld.go_

## Data Type

### Number

There are two data type number:

#### Integer

- Integer (1)
  <table>
    <thead>
      <tr>
        <th>Data Type</th>
        <th>Minimum Value</th>
        <th>Maximum Value</th>
      </tr>
    </thead>
    <tbody>
      <tr>
        <td>int8</td>
        <td>-128</td>
        <td>127</td>
      </tr>
      <tr>
        <td>int16</td>
        <td>-32768</td>
        <td>32767</td>
      </tr>
      <tr>
        <td>int32</td>
        <td>-2147483648</td>
        <td>2147483647</td>
      </tr>
      <tr>
        <td>int64</td>
        <td>-9223372036854775808</td>
        <td>9223372036854775807</td>
      </tr>
    </tbody>
  </table>

- Integer (2)
  <table>
    <thead>
      <tr>
        <th>Data Type</th>
        <th>Minimum Value</th>
        <th>Maximum Value</th>
      </tr>
    </thead>
    <tbody>
      <tr>
        <td>uint8</td>
        <td>0</td>
        <td>255</td>
      </tr>
      <tr>
        <td>uint16</td>
        <td>0</td>
        <td>65535</td>
      </tr>
      <tr>
        <td>uint32</td>
        <td>0</td>
        <td>4294967295</td>
      </tr>
      <tr>
        <td>uint64</td>
        <td>0</td>
        <td>18446744073709551615</td>
      </tr>
    </tbody>
  </table>

  _Note: uint -> unsigned integer_

#### Floating Point

<table>
  <thead>
    <tr>
      <th>Data Type</th>
      <th>Minimum Value</th>
      <th>Maximum Value</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>float32</td>
      <td>1.18x10<sup>-38</sup></td>
      <td>3.410<sup>38</sup></td>
    </tr>
    <tr>
      <td>float64</td>
      <td>2.23x10<sup>-308</sup></td>
      <td>1.80x10<sup>308</sup></td>
    </tr>
    <tr>
      <td>complex64</td>
      <td colspan='2'>complex numbers with float32 real and imaginary parts</td>
    </tr>
    <tr>
      <td>complex128</td>
      <td colspan='2'>complex numbers with float64 real and imaginary parts</td>
    </tr>
  </tbody>
</table>

#### Alias

<table>
  <thead>
    <tr>
      <th>Data Type</th>
      <th>Alias For</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>byte</td>
      <td>uint8</td>
    </tr>
    <tr>
      <td>rune</td>
      <td>int32</td>
    </tr>
    <tr>
      <td>int</td>
      <td>Minimal int32</td>
    </tr>
    <tr>
      <td>uint</td>
      <td>Minimal uint32</td>
    </tr>
  </tbody>
</table>

#### Boolean

- Boolean data type is a data type that has two values: true or false.
- In Go, the Boolean data type is represented by the keyword _bool_.

#### String

- A string is a set of characters.
- The total number of characters in a string can range from 0 to infinity.
- In Go, the string data type is represented by the keyword _string_.
- The value of string in Go always enclosed in quotation marks ("").
- There are some built-in functions for strings, such as:
  <table>
    <thead>
      <tr>
        <th>Function</th>
        <th>Purpose</th>
      </tr>
    </thead>
    <tbody>
      <tr>
        <td>len("string")</td>
        <td>Calculate the length of the string</td>
      </tr>
      <tr>
        <td>"string"[number]</td>
        <td>Take a character at a specific position based on the index (starting from 0)</td>
      </tr>
    </tbody>
  </table>
- Example:

  ```go
  package main

  import "fmt"

  func main() {
    fmt.Println("Learning Golang")
    fmt.Println("Go is Easy!")

    fmt.Println(len("Go is Easy!")) // 11
    fmt.Println("Learning Golang"[0]) // 76 (in byte type), we can convert it to string
  }
  ```

## Variable

- A variable is a storage to save data.
- Variables are used to make data accessible from various parts of the program.
- In Go, a variable can only store data of the same data type. If we want to store data of different types, we need to create multiple variables.
- To create a variable, we use the keyword _var_ followed by the variable name and the its data type.
- It is mandatory to state the data type when creating a variable, unless we directly initialize data to the variable.
- The keyword "var" is not mandatory if we directly initialize data using _:=_ when creating a variable. If we want to change the value of the variable, we need to use the asignment operator _=_ and not _:=_.
- We can make multiple variables in one go, the code will be cleaner.
- All variables created in Go must be used; otherwise, when we run the code, Go will generate an error.
- Example:

  ```go
  func main() {
    // Printing to terminal
    fmt.Println("Learning Golang") // "Learning Golang"
    fmt.Println("Go is Easy!") // "Go is Easy!"

    fmt.Println(len("Go is Easy!")) // 11
    fmt.Println("Learning Golang"[0]) // 76 (in byte type), we can convert it to string

    // Creating a variable
    var name string
    
    name = "Eko"
    fmt.Println(name)
    
    name = "Eko Kurniawan"
    fmt.Println(name)

    // Creating a variable without stating its data type
    var grade = 100
    fmt.Println(grade)

    // Omitting "var" keyword
    country := "Indonesia"
    fmt.Println(country)

    // Re-assigning value to the existing variable
    country = "England"
    fmt.Println(country)

    // Creating multiple variables
    var (
      firstName = "Pierle"
      lastName = "Dev"
    )

    fmt.Println(firstName) // "Pierle"
    fmt.Println(lastName) // "Dev"
  }
  ```

## Constant
- A constant is a variable whose value cannot be changed after the it is assigned for the first time.
- The keyword used is _const_, not _var_.
- Data must be directly assigned when the constant is created.
- We can also create multiple constants in one go.
- Example:
  ```go
  package main

  func main() {
    const day = 7

    // day = 8 // error re-assigning value

    // Creating constants in one go
    const (
      firstName string = "Pierle"
      lastName = "Dev"
    )
  }
  ```

### Data Type Konversion
```go
package main

import "fmt"

func main() {
  var value32 int32 = 32768
  var value64 int64 = int64(value32)

  // Be careful when trying to convert to a data type that can hold data with a smaller size; it can cause a problem
  var value16 int16 = int16(value32)

  fmt.Println(value32) // 32768
  fmt.Println(value64) // 32768
  fmt.Println(value16) // -32768 -> The maximum value of int16 is 32767. When attempting to convert 32768 to the int16 data type, it is set to the lowest number that can be stored in int16, i.e., -32768.

  var num32 int32 = 32770
  var num64 int64 = int64(num32)
  var num16 int16 = int16(num32)

  fmt.Println(num32) // 32770
  fmt.Println(num64) // 32770
  fmt.Println(num16) // -32766

  var activity = "Learning Golang"
  var l = activity[0] // in byte data type
  var lString = string(l) // L

  fmt.Println(activity) // "Learning Golang"
  fmt.Println(l) // 76
  fmt.Println(lString) // "L"
}
```

## Type Declarations
- Type declarations are a capability to create a new data type from an existing data type.
- Type declarations are typically used to create an alias for an existing data type, with the aim of making it easier to understand.
- Example:
  ```go
  package main

  import "fmt"

  func main() {
    type NoKTP string

  var myKTP NoKTP = "1234567"
  var rikaId string = "3625281"

  // converting rikaKTP to NoKTP data type
  var rikaKTP NoKTP = NoKTP(rikaId)

  fmt.Println(myKTP) // 1234567
  fmt.Println(rikaKTP) // 3625281
  }
  ```

## Mathematical Operations
- Addition (+), subtraction (-), multiplication (*), division (/), modulo (%).
- Example:
  ```go
  package main

  import "fmt"

  func main() {
    var a = 10
    var b = 10
    var c = a + b

    fmt.Println(c) // 20
  }
  ```

## Augmented Assignments
<table>
  <thead>
    <tr>
      <th>Maths Operations</th>
      <th>Augmented Assignments</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>a = a + 10</td>
      <td> += 10</td>
    </tr>
    <tr>
      <td>a = a - 10</td>
      <td>a -= 10</td>
    </tr>
    <tr>
      <td>a = a * 10</td>
      <td>a *= 10</td>
    </tr>
    <tr>
      <td>a = a / 10</td>
      <td>a /= 10</td>
    </tr>
    <tr>
      <td>a = a % 10</td>
      <td>a %= 10</td>
    </tr>
  </tbody>
</table>

Example:
```go
import "fmt"

func main() {
  var i = 10
  i += 10

  fmt.Println(i) // 20
}
```

## Unary Operator
<table>
  <thead>
    <tr>
      <th>Operator</th>
      <th>Notes</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>++</td>
      <td>a = a + 1</td>
    </tr>
    <tr>
      <td>--</td>
      <td>a = a - 1</td>
    </tr>
    <tr>
      <td>-</td>
      <td>Negative</td>
    </tr>
    <tr>
      <td>+</td>
      <td>Positive</td>
    </tr>
    <tr>
      <td>!</td>
      <td>Inverse boolean e.g. <i>!true</i> means <i>false</i></td>
    </tr>
  </tbody>
</table>

Example:
```go
package main

import "fmt"

func main() {
  var i = 1
  i++ // i + 1
  fmt.Println(i) // 2

  i++ // i + 1
  fmt.Println(i) // 3

  i-- // i - 1
  fmt.Println(i) // 2
}
```

## Comparison Operation
- Comparison operation is used for comparing two pieces of data.
- It is an operation that yields a boolean value (_true_ or _false_).
- Comparison operators
  <table>
    <thead>
      <tr>
        <th>Operator</th>
        <th>Notes</th>
      </tr>
    </thead>
    <tbody>
      <tr>
        <td>&gt;</td>
        <td>Greater than</td>
      </tr>
      <tr>
        <td>&lt;</td>
        <td>Less than</td>
      </tr>
      <tr>
        <td>&gt;=</td>
        <td>Greater than or equal</td>
      </tr>
      <tr>
        <td>&lt;=</td>
        <td>Less than or equal</td>
      </tr>
      <tr>
        <td>==</td>
        <td>Equal to</td>
      </tr>
      <tr>
        <td>!=</td>
        <td>Not equal to</td>
      </tr>
    </tbody>
  </table>
- Example:
  ```go
  package main

  import "fmt"

  func main() {
    var word1 = "Learning"
    var word2 = "Golang"

    var result bool = word1 == word2
    fmt.Println(result) // false

    var favNumber = 7
    var hisFavNumber = 7
    fmt.Println(favNumber == hisFavNumber) // true
  }
  ```

## Boolean Operation
- The output is either _true_ or _false_
  <table>
    <thead>
      <tr>
        <th>Operator</th>
        <th>Notes</th>
      </tr>
    </thead>
    <tbody>
      <tr>
        <td>&&</td>
        <td>AND</td>
      </tr>
      <tr>
        <td>||</td>
        <td>OR</td>
      </tr>
      <tr>
        <td>!</td>
        <td>NOT</td>
      </tr>
    </tbody>
  </table>
  <table>
    <thead>
      <tr>
        <th>Value 1</th>
        <th>Operator</th>
        <th>Value 2</th>
        <th>Result</th>
      </tr>
    </thead>
    <tbody>
      <tr>
        <td>true</td>
        <td>&&</td>
        <td>true</td>
        <td>true</td>
      </tr>
      <tr>
        <td>true</td>
        <td>&&</td>
        <td>false</td>
        <td>false</td>
      </tr>
      <tr>
        <td>false</td>
        <td>&&</td>
        <td>true</td>
        <td>false</td>
      </tr>
      <tr>
        <td>false</td>
        <td>&&</td>
        <td>false</td>
        <td>false</td>
      </tr>
      <tr>
        <td>true</td>
        <td>||</td>
        <td>true</td>
        <td>true</td>
      </tr>
      <tr>
        <td>true</td>
        <td>||</td>
        <td>false</td>
        <td>true</td>
      </tr>
      <tr>
        <td>false</td>
        <td>||</td>
        <td>true</td>
        <td>true</td>
      </tr>
      <tr>
        <td>false</td>
        <td>||</td>
        <td>false</td>
        <td>false</td>
      </tr>
    </tbody>
  </table>
  <table>
    <thead>
      <tr>
        <th>Operator</th>
        <th>Value</th>
        <th>Result</th>
      </tr>
    </thead>
    <tbody>
      <tr>
        <td>!</td>
        <td>true</td>
        <td>false</td>
      </tr>
      <tr>
        <td>!</td>
        <td>false</td>
        <td>true</td>
      </tr>
    </tbody>
  </table>
- Example:
  ```go
  package main

  import "fmt"

  func main() {
    var finalGrade = 90
    var attendance = 80

    var passingGrade bool = finalGrade > 80
    var passingAttendance bool = attendance > 80

    var pass bool = passingGrade && passingAttendance
    fmt.Println(pass) // false
  }
  ```

## Data Type: Array
- An array is a data type that holds a set of elements of the same type.
- When creating an array, it's necessary to specify the number of elements it can contain.
- The capacity of an array cannot be altered after its creation.
- We can access every element in an array with index starting from 0.
  <table>
    <thead>
      <tr>
        <th>Index</th>
        <th>Element</th>
      </tr>
    </thead>
    <tbody>
      <tr>
        <td>0</td>
        <td>Eko</td>
      </tr>
      <tr>
        <td>1</td>
        <td>Kurniawan</td>
      </tr>
      <tr>
        <td>2</td>
        <td>Khannedy</td>
      </tr>
    </tbody>
  </table>
- Example:
  ```go
  import "fmt"

  func main() {
    // Creating an array of string than will hold 3 elements
    var names [3]string
    fmt.Println(names) // []
    
    // Assigning elements to the array
    names[0] = "Siregar"
    names[1] = "Ali"
    names[2] = "Yessica"

    fmt.Println(names) // [Siregar Ali Yessica]
    fmt.Println(names[0]) // Siregar
    fmt.Println(names[1]) // Ali
    fmt.Println(names[2]) // Yessica
  }
  ```
- We can also create an array directly during variable declaration.
  ```go
  var values = [3]int{
    90,
    80,
    95, // <= if we specify elements for an array in a vertical format like this, a comma must be placed after the last element
  }

  fmt.Println(values) // [90 80 95]
  fmt.Println(values[0]) // 90

  // If we specify an array should have 3 elements but we don't assign any value during the variable declration, then the array will have three elements of 0
  var zeroNumbers = [3]int{}

  fmt.Println(zeroNumbers) // [0 0 0]

  // If we specify an array should have 3 elements but we only assign 2 elements during variable declaration, the last index will get a default value 0
  var numbers = [3]int{20, 7}

  fmt.Println(numbers) // [20 7 0]
  fmt.Println(numbers[2]) // 0
  ```
- Function Array
  <table>
    <thead>
      <tr>
        <th>Operation</th>
        <th>Explanation</th>
      </tr>
    </thead>
    <tbody>
      <tr>
        <td>len(array)</td>
        <td>Get the array length</td>
      </tr>
      <tr>
        <td>array[index]</td>
        <td>Get an element/data at a certain index position</td>
      </tr>
      <tr>
        <td>array[index] = value</td>
        <td>Change an element at a certain index position</td>
      </tr>
    </tbody>
  </table>
- We can also omit specifying the amount of elements/data that should be in an array by using [...] rather than [number of elements]. This way, the program will automatically set the length of the array to the number of elements assigned during variable declaration.
  ```go
  var ages = [...]int{
    20,
    22,
    25,
  }

  fmt.Println(ages) // [20 22 25]
  fmt.Println(len(ages)) // 3
  
  ages[2] = 27

  fmt.Println(ages) // [20 22 27]
  ```
- In Go, there is no way to alter the array length, for example by removing an element as can be done in JavaScript with methods like _.pop()_ or _.shift()_.

## Data Type: Slice
- The data type _slice_ is a segment of an array.
- A slice is similar to an array, with the key difference being that the size of a slice can change.
- Slices and arrays are always connected because a slice is the data that accesses a part or the whole data in an array.
- Slice has three important pieces of data: _pointer_, _length_, and _capacity_.
- _Pointer_ is a reference to the first element in the array that the slice points to.
- _Length_ is the size or length of the slice.
- _Capacity_ is the total capacity of the slice, where the length cannot exceed the capacity.
- Creating slice from an array
  <table>
    <thead>
      <tr>
        <th>Creating Slice</th>
        <th>Explanation</th>
      </tr>
    </thead>
    <tbody>
      <tr>
        <td>array[low:high]</td>
        <td>Create a slice from an array starting from index _low_ up to index before _high_</td>
      </tr>
      <tr>
        <td>array[low:]</td>
        <td>Create a slice from an array starting from index _low_ up to the last index</td>
      </tr>
      <tr>
        <td>array[:high]</td>
        <td>Create a slice from an array starting from index 0 up to index before _high_</td>
      </tr>
      <tr>
        <td>array[:]</td>
        <td>Create a slice from an array starting from index 0 up to the last index</td>
      </tr>
    </tbody>
  </table>
- Example:
  ```go
  func main() {
    var months = [12]string{
      "January",
      "February",
      "March",
      "April",
      "May",
      "June",
      "July",
      "August",
      "September",
      "October",
      "November",
      "December",
    }

    // A slice with a pointer to index 4, length 3, capacity 8 (the range from index 4 to the last index 11 is 8)
    // One of the differences between creating an array and a slice is that when creating a slice, we don't specify the length of the slice
    var slice1 []string = months[4:7]

    // A slice with a pointer to index 6, length 3, capacity 6 (the range from index 6 to the last index 11 is 6)
    var slice2 []string = months[6:9]

    fmt.Println(months) // [January February March April May June July August September October November December]
    fmt.Println(slice1) // [May June July]
    fmt.Println(slice2) // [July August September]

    names := [...]string{"Eko", "Alana", "Alaia", "Aisha", "Muhammad", "Christopher", "Ketut"}
    slice := names[4:6]

    fmt.Println(slice) // [Muhammad Christopher]
    fmt.Println(slice[0]) // Muhammad
    fmt.Println(slice[1]) // Christopher
  }
  ```
- Function in _slice_
  <table>
    <thead>
      <tr>
        <th>Operation</th>
        <th>Explanation</th>
      </tr>
    </thead>
    <tbody>
      <tr>
        <td>len(slice)</td>
        <td>Get the length of the slice, not the array</td>
      </tr>
      <tr>
        <td>cap(slice)</td>
        <td>Get the capacity of the slice</td>
      </tr>
      <tr>
        <td>append(slice, data)</td>
        <td>Create a new slice by adding data to the last position of slice. If the capacity is already full, a new array will be created</td>
      </tr>
      <tr>
        <td>make([]dataType, length, capacity)</td>
        <td>Create a new slice, array will automatically created by the slice</td>
      </tr>
      <tr>
        <td>copy(destination, source)</td>
        <td>Copy a slice from the source to the destination</td>
      </tr>
    </tbody>
  </table>
- Example:
  ```go
  func main() {
    days := [...]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
    daySlice1 := days[5:]
    daySlice1[0] = "New Friday" // the data array has also changed
    daySlice1[1] = "New Saturday"
    fmt.Println(days) // [Sunday, Monday, Tuesday, Wednesday, Thursday, New Friday, New Saturday]

    daySlice2 := append(daySlice1, "Holiday")
    daySlice2[0] = "Ups"
    fmt.Println(daySlice2) // [Ups, New Saturday, Holiday]
    fmt.Println(daySlice1) // [New Friday, New Saturday]
    fmt.Println(days) // [Sunday, Monday, Tuesday, Wednesday, Thursday, New Friday, New Saturday]

    // Create a new slice, array will automatically created by the slice
    // newSlice := make([]string, 2, 5) or
    var newSlice []string = make([]string, 2, 5)
    newSlice[0] = "Eko"
    newSlice[1] = "Eko"
    // newSlice[2] = "Eko" -> error, because we have set the length to be 2. To add another data, we should use the _append()_ function

    fmt.Println(newSlice) // [Eko Eko ]
    fmt.Println(len(newSlice)) // 2
    fmt.Println(cap(newSlice)) // 5

    newSlice2 := append(newSlice, "Adi")
    fmt.Println(newSlice2) // [Eko Eko Adi]
    fmt.Println(len(newSlice2)) // 3
    fmt.Println(cap(newSlice2)) // 5

    // Copy slice
    fromSlice := days[:]
    toSlice := make([]string, len(fromSlice), cap(fromSlice))

    fmt.Println(fromSlice) // [Sunday Monday Tuesday Wednesday Thursday New Friday New Saturday]
    fmt.Println(toSlice) // [      ]

    copy(toSlice, fromSlice)

    fmt.Println(toSlice) // [Sunday Monday Tuesday Wednesday Thursday New Friday New Saturday]
  }
  ```
- Be careful when creating an array, if done incorrectly, you may end up creating a slice instead of an array and vice versa.
  ```go
  import "fmt"

  func main() {
    anArray := [...]int{1, 2, 3, 4, 5}
    aSlice := []int{1, 2, 3, 4, 5} // In a slice, we don't specify the length/amount of data

    fmt.Println(anArray) // [1 2 3 4 5]
    fmt.Println(aSlice) // [1 2 3 4 5]
  }
  ```
- In Go, the use of arrays is less common when developing applications compared to the widespread use of slices.

## Data Type: Map
- In an Array or Slice, we use index numbers/integers from 0 to access data.
- A Map is a data type that holds a collection of homogeneous data, allowing us to specify the data type for each unique index.
- Simply put, a Map is a data type that consists of key-value pairs, where each key must be unique.
- In contrast Arrays and Slices, a Map can accomodate an unlimited amount of data, as long as the keys are distict. If the same key is used, the previous data will be automatically replaced by the new data.
- Function Map:
  <table>
    <thead>
      <tr>
        <th>Operation</th>
        <th>Explanation</th>
      </tr>
    </thead>
    <tbody>
      <tr>
        <td>len(map)</td>
        <td>Retrieve the length of the data in a Map</td>
      </tr>
      <tr>
        <td>map[key]</td>
        <td>Retrieve the data associated with the specified key</td>
      </tr>
      <tr>
        <td>map[key] = value</td>
        <td>Modify data associated with a specific key</td>
      </tr>
      <tr>
        <td>var <i>variable name</i> map[TypeKey]TypeValue = map[TypeKey]TypeValue{}</td>
        <td>Create a new Map</td>
      </tr>
      <tr>
        <td><i>variable name</i> := map[TypeKey]TypeValue{[key]: [value]}</td>
        <td>Create a new Map</td>
      </tr>
      <tr>
        <td>make(map[TypeKey]TypeValue)</td>
        <td>Create a new Map</td>
      </tr>
      <tr>
        <td>delete(map, key)</td>
        <td>Remove data associated with the specified key</td>
      </tr>
    </tbody>
  </table>
- Example:
  ```go
  import "fmt"

  func main() {
    // First way: creating a Map withh empty key-value pairs
    var aisha map[string]string = map[string]string{}
    aisha["username"] = "aisha"
    aisha["address"] = "Lombok"

    fmt.Println(aisha) // map[address:Lombok username:aisha]
    fmt.Println(aisha.username) // aisha
    fmt.Println(aisha.address) // Lombok

    // Second way: creating a Map and directly set the key-value pairs
    person := map[string]string{
      "username": "pierledev",
      "address": "Bandung",
    }

    fmt.Println(person) // map[address:Bandung username:pierledev]
    fmt.Println(person["username"]) // pierledev
    fmt.Println(person["address"]) // Bandung

    // When attempting to access a value using a key that does not exist in our Map, it will return an default value and the type of the value will be depend on the Map type that we have specified. If we have specified the Map to be of string type, the default value will be an empty string
    fmt.Println(person["hobby"]) // ""

    movie := make(map[string]string)
    movie["title"] = "Hannah Montana"
    movie["actress"] = "Miley Cyrus"
    movie["wrong"] = "Ups"

    delete(movie, "wrong")
    fmt.Println(movie) // map[actress:Miley Cyrus title:Hannah Montana]
  }
  ```

## If Expression
- _If_ is one of the keywords used for _branching_ in programming.
- _Branching_ allows the execution of specific code when a certain condition is met.
- the majority of programming languages support the _if expression_.
- Example:
  ```go
  func main() {
    name := "Pierle"

    if name == "Pierle" {
      fmt.Println("Hello, Pierle!") // Hello, Pierle!
    }
  }
  ```

## Else Expression
- The _if_ block is executed when the _if_ condition evaluates to _true_.
- The _else_ expression is used to execute a specific program when the condition evaluates to _false_.
- Example:
  ```go
  func main() {
    money := 100

    if money > 140 {
      fmt.Println("Buy the sneakers!!")
    } else {
      fmt.Println("Save your money!!")
    } // Save your money!!
  }
  ```

## Else If Expression
- We use the _else if_ expression when we need to evaluate multiple conditions.
- Example:
  ```go
  func main() {
    candidateNumber := 1

    if candidateNumber== 1 {
      fmt.Println("Anies Baswedan")
    } else if candidateNumber == 2 {
      fmt.Println("Prabowo Subianto")
    } else {
      fmt.Println("Ganjar Pranowo")
    } // Anies Baswedan
  }
  ```

## If with Short Statement
- The _if_ expression supports a short statement before the condition.
- It is very suitable for creating a simple statement before checking the condition.
- Example:
  ```go
  func main() {
    username := "pierledev"

    if length := len(username); length > 10 {
      fmt.Println("Username is too long")
    } else {
      fmt.Println("You can use that username")
    } // You can use that username
  }
  ```

## Switch Expression
- In addition to the _if_ expression, to create _branching_, we can also use the _switch_ expression.
- The _switch_ expression is simpler compared to the _if_ statement.
- Typically, the _switch_ expression is used to check conditions within a single variable.
- Example:
  ```go
  func main() {
    name := "Kenang"

    switch name {
      case "Kenzo":
        fmt.Println("Halo, Kenzo!")
      case "Kezia":
        fmt.Println("Halo, Kezia!")
      case "Kenang":
        fmt.Println("Halo, Kenang!")
      default:
        fmt.Println("Hi, anonim!")
    } // Halo, Kenang!
  }
  ```
- The _switch_ statement also supports a short statement before the variable's condition that will be checked.
  ```go
  func main() {
    name := "Rina Ilase"

    switch length := len(name); length > 5 {
      case true:
        fmt.Println("Name is too long")
      case false:
        fmt.Println("You can use that name")
    } // Name is too long
  }
  ```
- Conditions in the _switch_ expression are not mandatory.
- If we are not using conditions in the _switch_ expression, we can include conditions in each case.
  ```go
  func main() {
    name := "Pierledev"
    
    length := len(name)

    switch {
      case length > 10:
        fmt.Println("Name is too long")
      case length > 5:
        fmt.Println("Name is quite long")
      default:
        fmt.Println("Perfect!")
    }
  }
  ```
- Based on the provided code, in this case, it might be better to use the _if-else_ expression.

## For Loops
- In programming languages, there is typically a feature called a _loop_.
- One of the _loop_ features is _for loop_.
- Example:
  ```go
  func main() {
    counter := 1

    for counter <= 10 {
      fmt.Println("Loop ", counter)
      counter++
    }

    fmt.Println("Finish")
    /*
      Loop 1
      Loop 2
      Loop 3
      Loop 4
      Loop 5
      Loop 6
      Loop 7
      Loop 8
      Loop 9
      Loop 10
      Finish
    */
  }
  ```
- Inside the _for_ loop, we can add statements, and there are two types of statements that can be included in th _for_ loop:
  - _Init statement_: a statement executed before the _for_ loop begins.
  - _Post statement_: a statement executed at the end of each iteration of the loop.
  ```go
  func main() {
    for counter := 1; counter <= 10; counter++ {
      fmt.Println("Loop", counter)
    }

    // counter := 1 -> init statement
    // counter <= 10 -> condition
    // counter++ -> post statement
  }
  ```
- The _for_ loop can be used to iterate over all elements in a data collection, such as an Array, Slice, or Map.
  ```go
  // Manual
  names := []string{"Andika", "Chris", "Ana"}

  for i := 0; i < len(names); i++ {
    fmt.Println(names[i])
  }
  /*
    Andika
    Chris
    Ana
  */

  // With for range
  for index, name := range names {
    fmt.Println("index", index, "=", name)
  }
  /*
    index 0 = Andika
    index 1 = Chris
    index 2 = Ana
  */

  // With for range, but we don't use the indexes
  for _, name := range names {
    fmt.Println(name)
  }
  /*
    Andika
    Chris
    Ana
  */
  ```

## Break & Continue
- _break_ & _continue_ are keywords that can be used in a _loop_.
- _break_ is used to terminate the entire _loop_.
- _continue_ is used to stop the current iteration and immediately proceed to the next iteration of the _loop_.
- Example:
  ```go
  func main() {
    for i := 0; i < 10; i++ {
      if i == 5 {
        break
      }

      fmt.Println("Loop", i)
    }
    /*
      Loop 0
      Loop 1
      Loop 2
      Loop 3
      Loop 4
    */

    for i := 0; i < 10; i++ {
      if i % 2 == 0 {
        continue
      }

      fmt.Println("Loop", i)
    }
    /*
      Loop 1
      Loop 3
      Loop 5
      Loop 7
      Loop 9
    */
  }
  ```