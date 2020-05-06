Command line timed quiz that parses csv for QnA

Build a binary with
  
    go build .

After building the binary it can be used as follows:

```
Usage of ./gophercises_quiz:
  -csv string
    	a csv file in the format of 'question,answer' (default "problems.csv")
  -limit int
    	the time limit for the quiz in seconds (default 30)
  -shuffle
    	display problems in a random order
```

Example output:
```
$ ./gophercises_quiz -limit=15 -shuffle
Problem #1: 7+3 = 10
Problem #2: 5+5 = 9
Problem #3: 1+1 = 
You ran out of time and scored 1 out of 3.
```
