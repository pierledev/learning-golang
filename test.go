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
}