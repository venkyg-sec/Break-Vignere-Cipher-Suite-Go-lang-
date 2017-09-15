Student Name  - Venkatesh Gopal
JHED ID - vgopal3

Files in this directory :

1. vigenere_encrypt.go
2. vigenere_decrypt.go
3. vigenere_keylength.go
4. vigenere_cryptanalyze.go
5. plaintext.txt - text file (46KB)
6. rfc.txt - Large text file (for testing - 76KB)

Testing done :
I've tested on files of size upto 76KB and keys upto length 31 (as mentioned
in the assignment), works fine !

To build :

go build <filename.go>

To encrypt a file, replace "VENKY" with the key you are using
./vigenere_encrypt VENKY plaintext.txt > ciphertext.txt

To decrypt a file,
./vigenere_encrypt VENKY plaintext.txt > ciphertext.txt

To find the key length ,

./vigenere_keylength ciphertext.txt

To find the key itself

./vigenere_cryptanalyze ciphertext.txt



Example Execution log :

1. With key = "VENKY" (key size = 31)

venky@venky-potomac:~/go_programming$ ./vigenere_encrypt VENKY plaintext.txt > ciphertext.txt
venky@venky-potomac:~/go_programming$ ./vigenere_decrypt VENKY ciphertext.txt > recovered_plaintext.txt
venky@venky-potomac:~/go_programming$ ./vigenere_keylength ciphertext.txt
 The indices are  [60 65 75 10]
 Predicted key lenth is  5

 venky@venky-potomac:~/go_programming$ ./vigenere_cryptanalyze ciphertext.txt
  The indices are  [60 65 75 10]
  Predicted key lenth is  5
 Printing the key
 VENKY


2. With key = "MATTHEWGREENINSTRUCTORPRACTICAL" (key size = 31)

venky@venky-potomac:~/go_programming$ ./vigenere_encrypt MATTHEWGREENINSTRUCTORPRACTICAL plaintext.txt > ciphertext.txt
venky@venky-potomac:~/go_programming$ ./vigenere_keylength ciphertext.txt
 The indices are  [62 93 31 86]
 Predicted key lenth is  31
venky@venky-potomac:~/go_programming$ ./vigenere_cryptanalyze ciphertext.txt
 The indices are  [62 93 31 86]
 Predicted key lenth is  31
Printing the key
MATTHEWGREENINSTRUCTORPRACTICAL

venky@venky-potomac:~/go_programming$

3. With key = "VENKATESHGOPALGISNOT" (key size = 20)

venky@venky-potomac:~/go_programming$ ./vigenere_encrypt VENKATESHGOPALGISNOT plaintext.txt > ciphertext.txt
venky@venky-potomac:~/go_programming$ ./vigenere_keylength ciphertext.txt
 The indices are  [60 40 20 80]
 Predicted key lenth is  20
venky@venky-potomac:~/go_programming$ ./vigenere_cryptanalyze ciphertext.txt
 The indices are  [60 40 20 80]
 Predicted key lenth is  20
Printing the key
VENKATESHGOPALGISNOT

4. With key = "VENKATESH" (key size = 9)

venky@venky-potomac:~/go_programming$ ./vigenere_encrypt VENKATESH plaintext.txt > ciphertext.txt
venky@venky-potomac:~/go_programming$ ./vigenere_keylength ciphertext.txt
 The indices are  [72 81 9 27]
 Predicted key lenth is  9
venky@venky-potomac:~/go_programming$ ./vigenere_cryptanalyze ciphertext.txt
 The indices are  [72 81 9 27]
 Predicted key lenth is  9
Printing the key
VENKATESH
