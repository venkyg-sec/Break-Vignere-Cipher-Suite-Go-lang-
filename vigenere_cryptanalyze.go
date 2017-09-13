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

    predicted_key_length := FindKeyLength(s)
    fmt.Println( " Predicted key lenth is ", predicted_key_length)
    vigenere_cryptanalyze(s, predicted_key_length)
  }
}

func FindKeyLength(str string) (int) {

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
  
  return predicted_key_length

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

func vigenere_cryptanalyze(ciphertext string, predicted_key_length int) {

  s := make([]string, predicted_key_length)
  english_alphabet_upper_case := "ABCDEFGHIGHLMNOPQRSTUVWXYZ"
  english_alphabet_upper_case_length := len(english_alphabet_upper_case)
  ciphertext_length := len(ciphertext)

  english_language_alphabet_frequency :=
  [26]float64{.082, .015, .028, .043, .127, .022, .020, .061, .070, .002, .008,
     .040, .024, .067, .075, .019, .001, .060, .063, .091, .028, .010, .023,
     .001, .020, .001}
  // TODO Remove print statement
  //fmt.Println(english_language_alphabet_frequency)

  /* Extract all individual characters for the ciphertext and place them in an
  sliced array of strings */

  for j := 0; j < predicted_key_length; j++ {

    for i := 0; (i + predicted_key_length) < ciphertext_length ; i = i +
    predicted_key_length {

      s[j] += string(ciphertext[i + j])
    }
  }

  // // TODO Remove the print statements
  // fmt.Println(s[0])
  // fmt.Println(s[1])
  // fmt.Println(s[2])
  // fmt.Println(s[3])

  /* Declare a multi-dimensional array with the row holding the key index and
  the column holding the index for the alphabet */

  frequency_array := make([][]float64, predicted_key_length)

  for i := 0; i < predicted_key_length; i++ {
		frequency_array[i] = make([]float64, 26)
	}
  for j := 0; j < predicted_key_length; j++ {

  for i := 0; i < english_alphabet_upper_case_length; i++ {
    frequency_array[j][i] = float64(strings.Count(s[j],
      string(english_alphabet_upper_case[i])))

  }
}
  // // TODO Remove the print statements
  // fmt.Println(frequency_array)

  frequency_match := make([][]float64, predicted_key_length)
  for i := 0; i < predicted_key_length; i++ {
		frequency_match[i] = make([]float64, 26)
	}

  for i := 0; i < predicted_key_length; i++ {
    for j :=0; j < 26; j++ {
      for k := 0; k < 26; k++ {
      frequency_match[i][j] += float64(english_language_alphabet_frequency[k] *
      frequency_array[i][(j + k) % 26])
    }
    }

  }
  // TODO Remove print statements
  //fmt.Println(frequency_match)

  var max_values []float64
  max_values = make([]float64, predicted_key_length, predicted_key_length)

  var index []int
  index = make([]int, predicted_key_length, predicted_key_length)

  for i := 0; i < predicted_key_length; i++ {
    max_values[i] = findMaximum(frequency_match[i])
    index[i] = findIndexFloat(frequency_match[i],max_values[i])
  }

  // TODO Remove print statements
  fmt.Println("The maximum values are ", max_values)
  fmt.Println(" Corresponding indices are ", index)

  // Call the function to convert the key representation from numeric to string
  indexToKeyString(index)
}

/* Function to find the maximum value in the frequency array */

func findMaximum(individual_character_array []float64) (float64) {

  var sorted_counter []float64
  sorted_counter = make([]float64, 26, 26)
  for i := 0; i < 26; i++ {
    sorted_counter[i] = individual_character_array[i]
  }
  sort.Float64s(sorted_counter)
  // Sorting and returning the last value
  return sorted_counter[25]

}

func findIndexFloat(individual_character_array []float64, maximum_value float64)(int) {
  i := 0
  for i = 0; i < 26; i++ {
    if(individual_character_array[i] == maximum_value) {
      return i
    }
  }
  return i
}

func indexToKeyString(index []int) {

  fmt.Println("Printing the key")
  index_length := len(index)
  for i := 0; i < index_length; i++ {
    fmt.Printf("%s", string(index[i] + 65))
  }
  fmt.Println("\n")
}
