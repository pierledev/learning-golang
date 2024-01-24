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

## Function
- Earlier, we learned about a mandatory function to be created for our program to run, the _main_ function.
- A function is a deliberately created code block in a program that can be used multiple times.
- To create a function, we use the _func_ keyword, followed by the function name and the code block inside the function.
- After creating a function, we can execute it by calling it using the function name followed by parentheses.
- When creating a function, sometimes we need data from outside, which we refer to as arguments passed into function parameters.
- We can set more than 1 parameter.
- Parameter is not mandatory.
- Parameters are not mandatory, but if we define a function with parameters, we must provide data (arguments) for those parameters when calling the function.
- Functions can return a value or multiple values.
- To indicate that a function returns a value/multiple values, we should specify the data type expected to be returned by the function one by one
- If we declare a function with a return data type, it is mandatory to return data wthin the function.
- To return a value from a function, we can use the _return_ keyword followed by the value/data.
- Multiple return values must be captured for all the values.
- If we want to ignore the return values, we can use the underscore *_* symbol.
- Named return values. Typically, when we indicate that a function returns a value, we declare the data type of the returned value in the function. However, we can also create variables directly in the function's return data type.
- Example:
  ```go
  // Without parameter
  func sayHello() {
    fmt.Println("Hello")
  }

  // With parameters
  func sayHelloTo(firstName string, lastName string) {
    fmt.Println("Hello", firstName, lastName, "!")
  }

  // Return a value
  func getHello(name string) string {
    return "Hello " + name
  }

  // Return multiple values
  func getFullName() (string, string) {
    return "Komang", "Ida"
  }

  func getCompleteName() (username, email, phone string) {
    username = "chika"
    email = "chika@gmail.com"
    phone = "08726241233"

    return username, email, phone
  }

  func getGrade() (name string, grade int) {
    name = "Ario"
    grade = 90

    return ario, grade
  }

  func main() {
    sayHello() // Hello
    sayHelloTo("Robert", "Agung") // Hello Robert Agung !
    
    result := getHello("Rika")
    fmt.Println(result) // Hello Rika

    firstName, lastName := getFullName()
    fmt.Println(firstName, lastName) // Komang, Ida

    name, _ := getFullName()
    fmt.Println(name) // Komang

    username, email, phone := getCompleteName()
    fmt.Println(username, email, phone) // chika, chika@gmail.com, 08726241233

    _, grade := getGrade()
    fmt.Println(grade) // 90
  }
  ```
- The last parameter in a function has the ability to be a varargs.
- Varargs means the parameter can receive more than one input or can be considered similar to an Array
- The difference between regular parameters and the array data is:
  - For the parameter with an array type, it is madatory to create an array before sending it to the function.
  - For a parameter using varargs, we can directly send the data, and if there is more than one value, we can use a colon.
- Example:
  ```go
  func sumAll(numbers ...int) int { // numbers here is actually a slice of integers []int
    total := 0

    for _, number := range numbers {
      total += number
    }

    return total
  }

  func main() {
    total := sumAll(10, 10, 10, 10, 10, 10)
    fmt.Println(total) // 60
  }
  ```
- Sometimes, there are some cases where we use a Variadic Function, but we have a variable in the form of a slice. We can convert the slice to be a vararg parameter.
  ```go
  func sumAll(numbers ...int) int { // numbers here is actually a slice of integers []int
    total := 0

    for _, number := range numbers {
      total += number
    }

    return total
  }

  func main() {
    total := sumAll(10, 10, 10, 10, 10, 10)
    ft.Println(total) // 60

    numbers := []int{10, 10, 10} // slice
    total = sumAll(numbers...) // convert slice data to varargs
    fmt.Println(total) // 30
  }
  ```
- A function is a first-class citizen.
- Additionnally, a function is a data type and can be stored in a variable.
  ```go
  func getGoobBye(name string) string {
    return "Good Bye " + name
  }

  func main() {
    goodbye := getGoodBye
    fmt.Println(goodbye("Eko")) // Good Bye Eko
  }
  ```
- A function cannot only be stored in a variable as a value, but can also be used as a parameter for aother function, _function as parameter_.
  ```go
  func sayHelloWithFilter(name string, filter func(string) string) {
    fmt.Println("Hello ", filter(name))
  }

  func spamFilter(name string) string {
    if name === "Bitchie" {
      return "..."
    } else {
      return name
    }
  }

  func main() {
    sayHelloWithFilter("Bitchie", spamFilter) // Hello ...
  }
  ```
- Type declarations can be used to create an alias for a function, making it easier to use the function as a parameter.
  ```go
  type Filter func(string) string

  func sayHelloWithFilter(name string, filter Filter) {
    fmt.Println("Hello ", filter(name))
  }

  func spamFilter(name string) string {
    if name == "Bitchie" {
      return "..."
    } else {
      return name
    }
  }
  ```
- An _anonymous function_, also known as a function without a name, allows us to create a function directly in a variable or parameter without the need to define a function separately. Earlier, every time we created a function, a name was always assigned to it. However, sometimes it is easier to use a function directly without the need for a predefined name.
  ```go
  type Blacklist func(string) bool

  func registerUser(name string, blacklist Blacklist) {
    if(blacklist(name) {
      fmt.Println("Your are blocked", name)
    } else {
      fmt.Println("Welcome", name)
    })
  }

  func main() {
    // The first way
    blacklist := func(name string) bool {
      return name === "Bitchie"
    }

    registerUser("Andini", blacklist) // Welcome Andini

    // The second way
    registerUser("Bitchie", func(name string) bool {
      return name == "Bitchie"
    }) // you are blocked Bitchie
  }
  ```
- A _recursive function_ is a function that calls itself.  One illustrative example is the calculation of a _factorial_.
  ```go
  // Factorial with loop
  func factorialLoop(value int) int {
    result := 1
    for i := value; i > 0; i-- {
      result *= i
    }
    
    return result
  }

  // Factorial with recursive
  func factorialRecursive(value int) int {
    if value == 1 {
      return 1
    } else {
      return value * factorialRecursive(value - 1)
    }
  }

  func main() {
    fmt.Println(factorialLoop(10)) // 10 * 9 * 8 * 7 * 6 * 5 * 4 * 3 * 2 * 1 = 3628800

    fmt.Println(factorialRecursive(10)) // 3628800
  }
  ```

## Closures
- Closure refers to a function's capability to interact with data within the same scope. It's crucial to use this feature wisely when developing applications.
  ```go
  func main() {
    counter := 0
    increment := func() { // anynomous function
      fmt.Println("Increment")
      counter++
    }

    increment()
    increment()
    fmt.Println(counter) // 2
  }
  ```

## Defer, Panic, and Recover
- This concept is somewhat akin to a try-catch structure in other programming languages.

### Defer
- A _defer_ function is a function that can be scheduled to be executed after another function has finished executing. 
- The _defer_ function will always be executed even if an error occurs in the executed function.
- This is a useful pattern for tasks like cleanup operations or logging that you want to ensure happen even if an error occurs or in a specific order.
  ```go
  func logging() {
    fmt.Println("Finished calling function")
  }

  func runApplication() {
    defer logging() // It will be called after all the block code below run.  It schedules the logging() function to be executed after the runApplication() function finishes. 
    fmt.Println("Run application")
    // Any other code in the runApplication function would be executed before the deferred function.

    /*
      Run application
      Finished calling function
    */
  }

  func main() {
    runApplication()
  }
  ```

### Panic
- The _panic_ function is used to abruptly stop a program.
- The _panic_ function is typically invoked when a panic occurs during the execution of a program.
- When the _panic_ function is called, the program is halted, but any _deferred_ function will still be executed.
  ```go
  func endApp() {
    fmt.Println("End App")
  }

  func runApp(error bool) {
    defer endApp()

    if error {
      panic("Error")
    }

    endApp() // If this regular function/a defer function is placed below the error-checking block and an error occurs, the code will not be triggered.
  }

  func main() {
    runApp(true)

    /*
      End App
      panic: Error
    */
  }
  ```

### Recover
- _Recover_ is a function used to capture data from a _panic_.
- Through the _recover_ process, a _panic_ is halted, allowing the program to continue running.
  ```go
  func endApp() {
    fmt.Println("End App")
  }

  // Incorrect recover program code
  func runApp(error bool) {
    defer endApp()

    if(error) {
      panic("ERROR")
    }

    // The following codes won't be executed once the program run and the panic() function is executed
    message := recover()
    fmt.Println("Error happened", message)
  }
  ```
  ```go
  // Correct recover program code
  // Place the recover function and get the message from the panic() function inside the defer function because even if the program encounters an error, the defer function will still be executed
  func endApp() {
    fmt.Println("End App")
    message := recover()
    ft.Println("Error happened", message)
  }

  func runApp(error bool) {
    defer endApp()

    if(error) {
      panic("ERROR")
    }
  }

  func main() {
    runApp(true)
    fmt.Println("Printed!")

    /*
      Output:
      End app
      Error happened ERROR
      Printed!
    */
  }
  ```

## Comment
- The best comment for a code is the code itself.
- Write code for maximum readability, keeping it clear and straightforward.
- Add comments only when necessary for additional context.
  ```go
  // This is a single line comment

  /*
    This is
    multiline comment
  */
  ```

## Struct
- _Struct_ is a data template/data prototype used to combine zero or more data types into a unified entity.
- Typically, _structs_ represent data in an application program that we create.
- Data in a _struct_ is stored in a _field_.
- In simple terms, a _struct_ is a collection of _fields_.
- Struct cannot be used directly.
- However, we can create data or objects from the struct that we have defined.
  ```go
  type Customer struct {
    Name, Address string
    Age           int
  }

  func main() {
    var loyalCustomer1 Customer
    fmt.Println(loyalCustomer1) // {"" "" 0} // default values for data created with the Customer struct

    loyalCustomer1.Name = "Jacob"
    loyalCustomer1.Address = "SCBD"
    loyalCustomer1.Age = 24

    fmt.Println(loyalCustomer1) // {Jacob SCBD 24}
    fmt.Println(localCustomer1.Name) // Jacob
  }
  ```
- Earlier, we have already created data using a _struct_, but actually, there are various methods for creating data from a _struct_, one of them is struct literals.
  ```go
  type Customer struct {
    Name, Address string
    Age           int
  }

  func main() {
    joko := Customer{
      Name: "Joko",
      Address: "Indonesia",
      Age: 30,
    }
    fm.Println(joko) // {Joko Indonesia 30}

    budi := Customer{"Budi", "Indonesia", 30}
    fmt.Println(budi) // {Budi Indonesia 30}
  }
  ```
- A _struct_ is a data, similar to other data types, it can be used as a parameter for functions.
- _Structs_ can have methods, just like how functions can be associated with a _struct_.
- A method is essentially a function.
  ```go
  type Customer struct {
    Name, Address string
    Age   int
  }

  func (customer Customer) sayHello(name string) {
    fmt.Println("Hello", name, "my name is", customer.Name)
  }

  func main() {
    rully := Customer{Name: "Rully"}
    rully.sayHello("Radit") // Hello Radit my name is Rully
  }
  ```

## Interface
- An _interface_ is an Abstract data type without direct implementation.
- It consists of method definitions.
- Generally, an _interface_ is utilized as a form of _contract_. A contract implies that there must be code that implements it.
- Any data type that conforms to the contract of the interface is considered to be of that interface type, thus, there is no need to manually implement the interface.
- This is somewhat different from other programming languages where, when creating an interface, we must explicitly specify which interface is being implemented.
  ```go
  type HasName interface {
    GetName() string // the code that implements this interface must have method GetName that returns string data type
  }

  func SayHello(value HasName) {
    fmt.Println("Hello", value.GetName())
  }

  type Person struct {
    Name string
  }

  // In general, a person is an implementation of the HasName interface.
  func (person Person) GetName() string {
    return person.Name
  }

  func main() {
    person := Person{Name: "Eko"}
    SayHello(person) // Hello Eko
  }

  type Animal struct {
    Name string
  }

  func (animal Animal) GetName() string {
    return animal.Name
  }

  animal := Animal{Name: "Cat"}
  SayHello(animal) // Hello Cat
  ```
- Go is not an object oriented programming language.
- Typically, in object-oriented programming languages, there is a single parent data type at the top that can be considered as the root for all data implementations in that programming language; for example, in Java, there is java.lang.Object.
- To handle such cases in Go, we can use an empty interface.
- An empty interface is an interface that does not have any method declarations, making automatically every data type an implementation of it.
- The empty interface is also given a type alias named _any_. type any = interface{}.
- There are examples of the use of th empty interface in Go:
  - fmt.Println(a ...interface{})
  - panic(v interface{})
  - recover() interface{}
  - etc.
- So, when we use methods like fmt.Println, panic, or recover in Go, the parameters are of type empty interface (interface{}). This allows these functions to accept arguments of any type. The empty interface is a powerful feature in Go that provides flexibility in handling various types of data without explicitly specifying their types.
  ```go
  func Ups() interface{} { // iterface{} or any
    // return 1 it can returns int
    // return true it can also return bool
    return "Ups" // it can also return string and any other types
  }

  func main() {
    empty := Ups()
    fmt.Println(empty)
  }
  ```

## Nil
- Typically, in other programming languages, an object that has not been initialized will automatically have _null_ or _nil_ as its value.
- In Go, when we create a variable with a certain data type, it automatically gets a default value.
- However, in Go, there is special value called _nil_, representing an empty data.
- _Nil_ itself can only be used with certain data types, such as _interface_, _function_, _map_, _slice_, _pointer_, and _channel_.
  ```go
  // Invalid because nil cannot be used with data type string
  func check(name string) string {
    if name == "" {
      return nill
    } else {
      return name
    }
  }

  // Valid
  func NewMap(name string) map[string]string {
    if name == "" {
      return nil
    } else {
      return map[string]string {
        "name": name,
      }
    }
  }

  func main() {
    data := NewMap("")

    if(data == nil) {
      fmt.Println("Empty data")
    } else {
      fmt.Println(data)
    }

    // Output: Empty data
  }
  ```

## Type Assertions
- _Type Assertions_ is a capability to convert a data type to a desired data type.
- This feature is often used when we dealing an empty data interface.
  ```go
  func random() interface{} {
    return "OK"
  }

  func main() {
    /*
    result := random()
    resultString := result.(string) // convert from any (interface{} type is any) to string
    fmt.Println(resultString) // OK

    resultInt := result.(int)
    fmt.Println(resultInt) // panic
    */

    // OR
    var result any = random()
    var resultString string = result.(string)
    fmt.Println(resultString) // OK

    var resultInt int = result.(int)
    fmt.Println(resultInt) // panic
  }
  ```
- When we incorrectly use _type assertions_, it can lead to a panic in our application. If a _panic_ occurs and is not recovered, the program will terminate automatically. To ensure safety, a better approach for type assertion is to first check the data type in use and then perform _type assertion_ using a _switch_ expression.
  ```go
  func main() {
    result := random()

    switch value := result.(type) {
      case string:
        fmt.Println("String", value)
      case int:
        fmt.Println("Int", value)
      default:
        fmt.Println("Unknown")
    }
  }
  ```

## Pointer
- By default, all variables in Go are passed by value and not by reference. This means that when we send a variable to a function, method, or other variables, what is actually sent is a duplication of its value. Therefore, when we modify the data, the original data remains safe and unaffected because the modification does not use a reference or point to the same memory location.
- In other programming language like in JavaScript, objects (including arrays and functions) are assigned and passed by reference, while primitive types (like numbers and strings) are assigned and passed by value. In JavaScript, when we assign an object to another variable, both variables reference the same object in memory. Therefore, modifications to the object through one variable will affect the other. This is different from Go's behavior, where variables hold independent copies of their values.
  ```go
  // Go
  package main

  import "fmt"

  type Address struct {
    City, Province, Country string
  }

  func main() {
    // Creating an instance of the Address struct
    address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
    
     // Creating a new variable and copying the values from address1
    address2 := address1 // Copy value

    // Modifying the City field in the new variable
    address2.City = "Bandung"

    // Printing the original and modified variables
    fmt.Println(address1) // City remains as "Subang"
    fmt.Println(address2) // City is now "Bandung"
  }
  ```
  ```js
  // JavaScript
  let address1 = { city: "Subang", province: "Jawa Barat", country: "Indonesia" };

  // Creating a new variable and referencing the same object
  let address2 = address1;

  // Modifying the city property in the new variable
  address2.city = "Bandung";

  // Printing the original and modified variables
  console.log("Original Address:", address1); // City is now "Bandung"
  console.log("Modified Address:", address2); // City is also "Bandung"
  ```
- A pointer is a capability to create a reference to the data location in the same memory without duplicating the existing data.
- In simple words, with the ability of pointers, we can achieve pass by reference.
- To assign a variable with the pointer value to another variable, we can use operator _&_ followed by the variable name.
  ```go
  type Address struct {
    City, Province, Country string
  }

  func main() {
    address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
    address2 := &address1 // address2 is set to have the data that has the same location with address1 // &address1 -> pointer to address1

    /*
    OR
    var address1 Address = Address{"Subang", "Jawa Barat", "Indonesia"}
    var address2 *Address = &address2
    */
    
    address2.City = "Bandung"

    fmt.Println(address1) // City is now Bandung
    fmt.Println(address2) // City is also Bandung 
  }
  ```

## Asterisk Operator *
- When we modify a variable through a pointer, only that specific variable is affected.
- Other variables referencing the same data will remain unchanged.
  ```go
  type Address struct {
    City, Province, Country string
  }

  func main() {
    address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
    address2 := &address1

    address2.City = "Bandung"

    address2 = &Address{"Jakarta", "DKI Jakarta", "Indonesia"} // address2 is a pointer to address1, so we need to put '&'

    fmt.Println(address1) // {Bandung Jawa Barat Indonesia}
    fmt.Println(address2) // {Jakarta DKI Jakarta Indonesia}
  }
  ```
- If we want to update all variables referencing that data, we can use the _*_ operator.
  ```go
  type Address struct {
    City, Province, Country string
  }
  
  func main() {
    address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
    address2 := &address1

    address2.City = "Bandung"

    *address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}

    fmt.Println(address1) // {Jakarta DKI Jakarta Indonesia}
    fmt.Println(address2) // {Jakarta DKI Jakarta Indonesia}
  }
  ```

## New Operator
- Earlier, we created pointers using the _&_ operator.
- Go also provides a _new_ function that can be used to create a _pointer_.
- However, the _new_ function can only return a _pointer_ to an empty data, meaning there is no initial data assigned.
  ```go
  type Address struct {
    City, Province, Country string
  }

  func main() {
    // Creating pointer not with the 'new' function
    var address1 *Address = &Address{}
    var address2 *Address = address1

    // Creating pointer with the 'new' function
    address1 := new(Address) // pointer to Address
    address2 := adddress1 // also pointer to address1

    address2.Country = "Indonesia" // the result is, all pointers to the data Country changed to Indonesia

    fmt.Println(address1) // &{"" "" "Indonesia"}
    fmt.Println(address2) // &{"" "" "Indonesia"}
  }
  ```

## Pointer in Function
- When we create a parameter in a function, by default, it is pass by value, meaning the data will be copied and then sent to the function.
- Therefore, if we change the data in a function, the original data will never be altered.
- This makes variables secure, as they cannot be modified.
- However, sometimes we want to create a function that can change the original data parameter.
- To achieve this, we can use a pointer in a function's parameter.
- To make a parameter a pointer, we can use the _*_ operator in the parameter.
  ```go
  // Example for code that is not yet using a pointer
  type Address struct {
    City, Province, Country string
  }

  func ChangeAddressToIndonesia(address Address) {
    address.Country = "Indonesia"
  }

  func main() {
    address := Address{"Subang", "Jawa Barat", ""}
    ChangeAddressToIndonesia(address)

    fmt.Println(address) // not changed {"Subang" "Jawa Barat" ""}
  }
  ```
  ```go
  // Example for code that is  using a pointer
  type Address struct {
    City, Province, Country string
  }

  func ChangeAddressToIndonesia(address *Address) {
    address.Country = "Indonesia"
  }

  func main() {
    var address *Address := &Address{"Subang", "Jawa Barat", ""}
    ChangeAddressToIndonesia(address)

    fmt.Println(address) // changed to {"Subang" "Jawa Barat" "Indonesia"}
  }
  ```
  ```go
  // If we've already created it as a non-pointer, we can simply add '&' directly to the function argument.
  type Address struct {
    City, Province, Country string
  }

  func ChangeAddressToIndonesia(address *Address) {
    address.Country = "Indonesia"
  }

  func main() {
    address := Address{"Subang", "Jawa Barat", ""}
    ChangeAddressToIndonesia(&address) // <= here

    fmt.Println(address) // changed to {"Subang" "Jawa Barat" "Indonesia"}
  }
  ```
- Although methods are attached to a struct, the data struct accessed in a method is actually passed by value.
- **It's recommended to use pointers in methods to avoid unnecessary memory duplication when calling methods.**
  ```go
  // Without pointer
  type Man struct {
    Name string
  }

  func (man Man) Married() {
    man.Name = "Mr. " + man.Name
  }

  func main() {
    eko := Man{"Eko"}
    eko.Married()

    fmt.Println(eko.Name) // Eko, not Mr.Eko because the Married function accepts a parameter by value, creating a copy and not using a reference to the original variable/argument
  }
  ```
  ```go
  // With pointer
  type man struct {
    Name string
  }

  func (man *Man) Married() {
    man.Name = "Mr. " + man.Name
  }

  func main() {
    eko := Man{"Eko"}
    eko.Married()

    fmt.Println(eko.Name) // Mr. Eko
  }
  ```

## Package & Import
- A package is a container used to organize program code in Go.
- By utilizing packages, we can structure and organize the code we create.
- In essence, a package corresponds to a directory or folder in our operating system.
- The name of a Go package should match the name of the folder or directory in which the Go files of the package are stored. This is a convention in Go. This helps Go tools and developers maintain a consistent and predictable structure for organizing code. For example, if we have our existing folder "learning-golang" and we make another folder inside it called _helper_, create a Go file inside it also named _helper_. At the very beginning of that file, write:
  ```go
  package helper

  func SayHello(name string) string {
    return "Hello " + name
  }
  ```
- To use the package, we use _import_.
- In standard practice, a Go file can only access other Go files that belong to the same package.
-  If we want to access Go files outside of the package, we can use _import_. For example if we want to ue _helper_ package inside _main.go_ file in the root folder a.k.a "learning-golang", then we should write the following code:
  ```go
  // main.go
  import (
    "learning-golang/helper"
    "fmt"
  )

  func main() {
    result := helper.SayHello("Eko")
    fmt.Println(result) // Hello Eko
  }
  ```