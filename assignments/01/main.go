package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	SUCCESS              = "SUCCESS"
	INVALID_INPUT        = "INVALID_INPUT"
	INVALID_COUNTRY_CODE = "INVALID_COUNTRY_CODE"
)

func solve(name string) (string, string) {
	name = strings.ToUpper(name)
	split := strings.Split(name, " ")
	inputLenth := len(split)
	if inputLenth < 3 {
		return "", INVALID_INPUT
	}
	countryCode := split[inputLenth-1]
	if countryCode != "VN" && countryCode != "US" {
		return "", INVALID_COUNTRY_CODE
	}
	firstName := split[0]
	lastName := split[1]
	if inputLenth == 3 {
		if countryCode == "VN" {
			return lastName + " " + firstName, SUCCESS
		} else {
			return firstName + " " + lastName, SUCCESS
		}
	} else {
		middleName := strings.Join(split[2:inputLenth-1], " ")
		if countryCode == "VN" {
			return lastName + " " + middleName + " " + firstName, SUCCESS
		} else {
			return firstName + " " + middleName + " " + lastName, SUCCESS
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	line = strings.Replace(line, "\n", "", -1)
	name, errorCode := solve(line)
	if errorCode == SUCCESS {
		fmt.Println(name)
	} else {
		fmt.Println(errorCode)
	}
}
