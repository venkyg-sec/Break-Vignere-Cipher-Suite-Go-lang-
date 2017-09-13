package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "strings"
  "os"
  "regexp"
  "sort"
  //"bytes"
)

func main() {

  if (len(os.Args[1:]) != 1) {
    fmt.Println(" Use the command line specification mentioned in the Assignment\n <executable_name> <ciphertext_file_name> \n")
  } else {
    file_name := os.Args[1]
    file, err_data_file := ioutil.ReadFile(file_name)
    /* Error handling if file wasn't opened successfully */
    if (err_data_file != nil) {
      log.Fatal(err_data_file)
    }
    str := strings.ToUpper(string(file))
    re := regexp.MustCompile("[^[:alpha:]]")
    s := re.ReplaceAllLiteralString(str,"")

    FindKeyLength(s)
  }
}

func FindKeyLength(str string) {

  str_length := len(str)
  var counter []int
  maximum_key_size := 100
  counter = make([]int, maximum_key_size, maximum_key_size)

  for shift := 1; shift < maximum_key_size; shift++ {

    for i := 0; i < str_length ; i++ {

    // Case when we want to ignore the characters past the string length
    // post shifting
      if((i + shift) >= str_length) {
        break
      }
    c := string(str[i])
    d := string(str[i + shift])

    if(c == d) {
      counter[shift] += 1
    }
  }

}
  /*for a:=0; a< maximum_key_size; a++ {
    fmt.Println(a , " : ", counter[a])

  } */

  // Define a slice for having a new copy of array to sort
  var sorted_counter []int
  counter_length := len(counter)
  sorted_counter = make([]int, counter_length, counter_length)

  for i := 0; i < counter_length; i++ {
    sorted_counter[i] = counter[i]
  }
  sort.Ints(sorted_counter)
  //fmt.Println(sorted_counter)

  // Maintain the indices which have the maximum value
  var index []int
  index = make([]int, 4,4)
  for i:=0; i < 4; i++ {
    index[i] = findIndex(counter, counter_length, sorted_counter[counter_length -(i + 1)])
}
  fmt.Println(" The indices are ", index)
  gcd_1 := gcd(index[0], index[1])
  gcd_2 := gcd(index[2], index[3])
  predicted_key_length := gcd(gcd_1, gcd_2)

  if predicted_key_length == 1 {

    sort.Ints(index)
    predicted_key_length = index[0] + 1
  }
  fmt.Println( " Predicted key lenth is ", predicted_key_length)
}

func findIndex(counter []int, length int, value int) int {
  i := 0
  for i = 0; i < length; i++ {
    if(counter[i] == value) {
      return i
    }
  }
  return i
  }

func gcd(temp1 int,temp2 int) (int){
       var gcdnum int
    /* Use of And operator in For Loop */
    for i := 1; i <=temp1 && i <=temp2 ; i++ {
            if(temp1%i==0 && temp2%i==0) {
                gcdnum=i
            }
        }
        return gcdnum
}
