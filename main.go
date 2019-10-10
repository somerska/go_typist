package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main(){
	reversedStrings := flag.Bool("reversed", false, "Requires the user to type all strings in reverse")
	csvFilePath := flag.String("csv", "strings.csv", "A csv file with format: 'int, string'")
	flag.Parse()

	timedStrings := parseLines(*csvFilePath)
	runTypist(timedStrings, *reversedStrings)
}

type timedString struct {
	time int
	stringTest string
}

//Runs one round of the typist challenge, as soon as user fails (runs out of time or enters incorrect value),
//the game is over.
func runTypist(timedStrings []timedString, reversedStrings bool) {
	numCorrect := 0
	userInputCh := make(chan string)

	reader := bufio.NewReader(os.Stdin)
	go readStdin(*reader, userInputCh)
	for i, ts := range timedStrings {
		fmt.Printf("Problem #%d you have %d seconds to type:%s \n", i+1, ts.time, ts.stringTest)
		correct, reason := answeredCorrectlyInTime(userInputCh, ts, reversedStrings)
		fmt.Println(reason)
		if correct {
			numCorrect++
		} else {
			break
		}
	}
	fmt.Printf("You got %d correct!\n", numCorrect)
}

//read stdin and place users input into channel for consumer to consume
func readStdin(reader bufio.Reader, userInputCh chan string){
	for {
		answer, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading stdin")
		}
		userInputCh <- answer
	}
}

//Selects on a timer or userinput, if the timer returns first then the user took too long to type their answer.
//Otherwise this function compares the expected stringTest with what the user typed in and can also compare reversed
//strings.
func answeredCorrectlyInTime(userInputCh chan string, timedString timedString, reversed bool) (bool, string) {
	timer := time.NewTimer(time.Duration(timedString.time) * time.Second)

	comparison := stringsAreEqual
	if reversed {
		comparison = reversedStringsAreEqual
	}
	select {
	case <-timer.C:
		return false, "\nRan out of time on that one!\n"
	case userInput := <-userInputCh:
		userInput = strings.TrimSpace(userInput)
		if comparison(userInput, timedString.stringTest){
			return true, "Nice job!"
		} else {
			return false, fmt.Sprintf("Got: %s\nExpected: %s", userInput, timedString.stringTest)
		}
	}
}

// opens the file, iterates through line by line parsing out the relevant parts and populates a slice of structs
// to return
func parseLines(csvFilePath string) []timedString {
	file, err := os.Open(csvFilePath)
	if err != nil {
		printThenExit(fmt.Sprintf("Error opening CSV file %s, exiting", csvFilePath))
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	timedStrings := make([]timedString, 0)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		timedStrings = append(timedStrings, getTimedStringFromLine(line))
	}
	return timedStrings
}

//receives a comma separated line, separates out the time from the string and serializes into a struct to return
func getTimedStringFromLine(line string) timedString{
	s := strings.Split(line, ",")
	strTime, stringTest := strings.TrimSpace(s[0]), strings.TrimSpace(s[1])
	stringTime, err := strconv.Atoi(strTime)
	if err != nil {
		printThenExit(fmt.Sprint("Error converting time in csv to an integer, exiting"))
	}
	timedString := timedString{
		time: stringTime,
		stringTest: stringTest,
	}
	return timedString
}

func printThenExit(error string){
	fmt.Println(error)
	os.Exit(1)
}

// Returns string in reverse order
func reverseString(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func stringsAreEqual(s1, s2 string) bool {
	return s1 == s2
}

func reversedStringsAreEqual(s1, s2 string) bool {
	return s1 == reverseString(s2)
}

