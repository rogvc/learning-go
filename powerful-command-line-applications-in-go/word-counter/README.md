# Word Counter

Simple word counting program. Reads a string from stdin, counts how many words, lines, or bytes (space separated) are in the string.

Good program to learn:
- How [Go modules](https://go.dev/ref/mod#:~:text=Go%20Modules%20Reference%201%20Introduction%20Modules%20are%20how,file%20named%20go.mod%20in%20its%20root%20directory.%20) work
- How to use CLI flags
- Read from stdin
- Parse input
- Output to stdout
- Run tests

# How to use it

Build the executable: 

```bash
# From repository root
cd powerful-command-line-applications-in-go/word-counter
go test -v
go build -o bin
```

Run it:

```bash
# Make sure you're in the word-counter directory
echo "w1 w2 w3 w4 w5" | ./bin/wc
```
> Output should be 5

To count lines, use the `-l` flag

```bash
# Make sure you're in the word-counter directory
echo "w1\n w2\n w3 w4 w5" | ./bin/wc -l
```
> Output should be 3

To count bytes, use the `-b` flag

```bash
# Make sure you're in the word-counter directory
echo "w1 w2 w3" | ./bin/wc -b
```

> Output should be 9
