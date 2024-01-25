package main

import "fmt"

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

  fmt.Println(activity)
  fmt.Println(l)
  fmt.Println(lString)

  type NoKTP string

  var myKTP NoKTP = "1234567"
  var rikaId string = "3625281"

  // converting rikaKTP to NoKTP data type
  var rikaKTP NoKTP = NoKTP(rikaId)

  fmt.Println(myKTP) // 1234567
  fmt.Println(rikaKTP) // 3625281

  var i = 1
  i++ // i + 1
  fmt.Println(i) // 2

  i++ // i + 1
  fmt.Println(i) // 3

  i-- // i - 1
  fmt.Println(i) // 2

  var word1 = "Learning"
  var word2 = "Golang"

  var result bool = word1 == word2
  fmt.Println(result) // false

  var favNumber = 7
  var hisFavNumber = 7
  fmt.Println(favNumber == hisFavNumber) // true

  var finalGrade = 90
  var attendance = 80

  var passingGrade bool = finalGrade > 80
  var passingAttendance bool = attendance > 80

  var pass bool = passingGrade && passingAttendance
  fmt.Println(pass) // false

  // Creating an array of string than will hold 3 elements
  var names [3]string
  fmt.Println(names)
    
  // Assigning elements to the array
  names[0] = "Siregar"
  names[1] = "Ali"
  names[2] = "Yessica"

  fmt.Println(names)
  fmt.Println(names[0])
  fmt.Println(names[1])
  fmt.Println(names[2])

  var values = [3]int{
    90,
    80,
    95, // <= if we specify elements for an array in a vertical format like this, a comma must be placed after the last element
  }

  fmt.Println(values)
  fmt.Println(values[0])

  // If we specify an array should have 3 elements but we don't assign any value during the variable declration, then the array will have three elements of 0
  var zeroNumbers = [3]int{}

  fmt.Println(zeroNumbers) // [0 0 0]

  // If we specify an array should have 3 elements but we only assign 2 elements during variable declaration, the last index will get a default value 0
  var numbers = [3]int{20, 7}

  fmt.Println(numbers)
  fmt.Println(numbers[2])

  var ages = [...]int{
    20,
    22,
    25,
  }

  fmt.Println(ages) // [20 22 25]
  fmt.Println(len(ages)) // 3
  
  ages[2] = 27

  fmt.Println(ages) // [20 22 27]

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
  var slice1 = months[4:7]

  // A slice with a pointer to index 6, length 3, capacity 6 (the range from index 6 to the last index 11 is 6)
  var slice2 = months[6:9]

  fmt.Println(months)
  fmt.Println(slice1)
  fmt.Println(slice2)

  namese := [...]string{"Eko", "Alana", "Alaia", "Aisha", "Muhammad", "Christopher", "Ketut"}
    slice := namese[4:6]

  fmt.Println(slice) // [Muhammad Christopher]
  fmt.Println(slice[0]) // Muhammad
  fmt.Println(slice[1]) // Christopher

  days := [...]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
  daySlice1 := days[5:]
  daySlice1[0] = "New Friday"
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

  fmt.Println(newSlice) // [Eko Eko]
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

  anArray := [...]int{1, 2, 3, 4, 5}
  aSlice := []int{1, 2, 3, 4, 5}

  fmt.Println(anArray) // [1 2 3 4 5]
  fmt.Println(aSlice) // [1 2 3 4 5]

  // Map
  // First way: creating a Map withh empty key-value pairs
  var aisha map[string]string = map[string]string{}
  aisha["username"] = "aisha"
  aisha["address"] = "Lombok"

  fmt.Println(aisha)
  fmt.Println(aisha["username"])
  fmt.Println(aisha["address"])

  // Second way: creating a Map and directly set the key-value pairs
  person := map[string]string{
    "username": "pierledev",
    "address": "Bandung",
  }

  fmt.Println(person) // map[address:Bandung username:pierledev]
  fmt.Println(person["username"]) // pierledev
  fmt.Println(person["address"]) // Bandung
  fmt.Println(person["hobby"]) // ""

  movie := make(map[string]string)
  movie["title"] = "Hannah Montana"
  movie["actress"] = "Miley Cyrus"
  movie["wrong"] = "Ups"

  delete(movie, "wrong")
  fmt.Println(movie) // map[actress:Miley Cyrus title:Hannah Montana]

  // If expression
  myName := "Pierle"

  if myName == "Pierle" {
    fmt.Println("Hello, Pierle!")
  }

  // Else expression
  money := 100

  if money > 140 {
    fmt.Println("Buy the sneakers!!")
  } else {
    fmt.Println("Save your money!!")
  }

  // Else if expression
  candidateNumber := 1

  if candidateNumber== 1 {
    fmt.Println("Anies Baswedan")
  } else if candidateNumber == 2 {
    fmt.Println("Prabowo Subianto")
  } else {
    fmt.Println("Ganjar Pranowo")
  }

  // If with short statement
  username := "pierledev"

  if length := len(username); length > 10 {
    fmt.Println("Username is too long")
  } else {
    fmt.Println("You can use that username")
  }

  // For loop
  // Manual
  namesCollection := []string{"Andika", "Chris", "Ana"}

  for i := 0; i < len(namesCollection); i++ {
    fmt.Println(namesCollection[i])
  }

  // With for range
  for index, name := range namesCollection {
    fmt.Println("index", index, "name", name)
  }

  // Break
  for i := 0; i < 10; i++ {
    if i == 5 {
      break
    }

    fmt.Println("Loop", i)
  }

  // Continue
  for i := 0; i < 10; i++ {
    if i % 2 == 0 {
      continue
    }

    fmt.Println("Loop", i)
  }

  // Custom error
  err := SaveData("", nil)

    // If else
    if err != nil {
      if validationErr, ok := err.(*validationError); ok {
        fmt.Println("validation error:", validationErr.Message)
      } else if notFoundErr, ok := err.(*notFoundError); ok {
        fmt.Println("not found error:", notFoundErr.Message)
      } else {
        fmt.Println("unknown error:", err.Error())
      }
    } else {
      fmt.Println("success")
    }
}

type validationError struct {
  Message string
}

func (v *validationError) Error() string {
  return v.Message
}

type notFoundError struct {
  Message string
}

func (v *notFoundError) Error() string {
  return v.Message
}

// Using the custom error
func SaveData(id string, data any) error {
  if id == "" {
    return &validationError{Message: "validation error"} // because it is an interface, so we return it as a pointer
  }

  if id != "eko" {
    return &notFoundError{Message: "data not found"}
  }
  
  // ...

  return nil
}