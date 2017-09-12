package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "strings"
  "os"
  "regexp"
  "bytes"
)

func main() {

  if (len(os.Args[1:]) != 2) {
    fmt.Println(" Use the command line specification mentioned in the Assignment\n <executable_name> <keys_file_name> <ciphertext_file_name> \n")

  } else {

  key := os.Args[1]
  if (len(key) > 32) {
    fmt.Println(" Invalid key length \n Key length should be upto 32 characters \n")
  }
  file_name := os.Args[2]
  file, err_data_file := ioutil.ReadFile(file_name)
  /* Error handling if file wasn't opened successfully */
  if (err_data_file != nil) {
    log.Fatal(err_data_file)
  }
  /* Need to perform below operation as ReadFile() returns a byte array */
  str := strings.ToUpper(string(file))
  re := regexp.MustCompile("[^[:alpha:]]")
  s := re.ReplaceAllLiteralString(str,"")
  vignere_decryption(s,key)

}

}

func vignere_decryption(str string, key string) {
	// Define an array to hold ascii characters belonging to the key
	var key_ascii []int
	key_length := len(key)
	key_ascii = make([]int, key_length, key_length)
	for j:=0; j < key_length; j++ {
		key_ascii[j] = int(key[j])
	}

	// Define an array to hold ascii characters belonging to the plaintext
	var ascii []int
  var temp int
	str_length := len(str)
	ascii = make([]int, str_length, str_length)
	for i:=0 ; i < str_length; i++ {
		for z:=0; z < key_length; z++ {

			/* Logic here is I try all commbinations and check if the mod of the
			 index in the plaintext matches the index for the key. If yes, I encode
			 that character */
			if ((i % key_length) == z) {
        temp = (((int(str[i]) - 65) - (key_ascii[z] -65)) % 26)
        if(temp < 0) {
          temp = temp + 26
        }
			ascii[i] = temp + 65
    }
		}
		}


	// Convert the ascii array back to a string and write to a buffer
	var buffer bytes.Buffer
	for k:= 0; k < len(ascii); k++ {
			buffer.WriteString(string(ascii[k]))
	}
	fmt.Println(buffer.String())

  // Find the index for the highest 4 values
}
