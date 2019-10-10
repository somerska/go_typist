  # go_typist 

### Description

A command line typing game, to test your typing abilities.  The typing tests are pulled in from a csv and the user must
type the line correctly in the allotted time.

### Installation

To run ensure you have golang 1.11 or later installed https://golang.org/doc/install

### Compatability
Currently this program is written with the expectation of windows line endings, so for now this only works on windows.

### Configuration

Then program accepts command line arguments to augment behavior.  

```cmd
C:\Users\Kevin\go\src\go_typist>go_typist.exe --help
Usage of go_typist.exe:
  -csv string
        A csv file with format: 'int, string' (default "strings.csv")
  -reversed
        Requires the user to type all strings in reverse
```

### Instructions to play

Create a csv file or borrow the one from this repo.  The program expects the csv to be of the 

format: integer,string

Where the integer represents the amount of time the player has to type in the string. For example, you might have a
csv with the following entries:

```
5,can you type this in 5 seconds?
3,quick one, type fast!
8,wow you have 8 seconds to type this, plenty of time.
```

Once the csv is created, direct the program to the csv like this inside your terminal:
```cmd
C:\Users\Kevin\go\src\go_typist>go_typist.exe --csv=strings.csv
Problem #1 you have 5 seconds to type:abcde
ab
Ran out of time on that one!

You got 0 correct!
```

If you're really looking for a challenge, run the program with reversed mode on.

```
C:\Users\Kevin\go\src\go_typist>go_typist.exe --csv=strings.csv --reversed
Problem #1 you have 7 seconds to type:short string
gnirts trohs
Nice job!
Problem #2 you have 12 seconds to type:this is a long string
gnirts gnol a si si
Ran out of time on that one!

You got 1 correct!
```








 