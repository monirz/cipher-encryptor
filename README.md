# cipher-encryptor
## A Caesar Cipher style encryption with extended algorithm.  

An encryptor reads a text file, encrypts the text and creates new file with the ecncypted content.   
Just wrote this for experiment and learning encryption, yet it could be useful to save your sensitive info to a file rather you
used to save it as plain text. 

In plain Caesar Cipher algorithm the text "hello world" with key 3 would look like "khoor zruog", we all know which is kind of guessable 
The algorithm I've come out so far the text "hello world" with key 3 will look like ` d`fegvmdf_V ` 

* Usage *

Build the cipher.go file 

Run the command 

```
  ./cipher hello.txt e

```

Or build the binary file by running `go install`, If your bin path is imported to gopath then  you could just go to any folder
and type 

 ``` 
 cipher hello.txt e
 ``` 

where hello.txt is your text file and `e/d` for encryption/decryption
Then it will ask for a key, enter your key you are done. 

If encryption is successful a file named encrypt.txt file will be created in the same directory with encrypted content or for decryption a file named decrypt.txt wil be created

Not yet user friendly, it'll be updated :)



