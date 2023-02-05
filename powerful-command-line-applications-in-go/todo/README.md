# Todo

Simple todo list program.

Good program to learn:
- More in-depth flags
- Creating and consuming files
- Working JSON with Go
- Testing

# How to use it

Build the executable: 

```bash
# From repository root
cd powerful-command-line-applications-in-go/todo/cli/todo
go test -v
go build -o bin/
```

Run it:

To see usage

```bash
# Make sure you're in the todo/cli/todo directory
./bin/todo --help
```

To add a task, use the `-[a|add]` flag

```bash
# Make sure you're in the todo/cli/todo directory
# ./bin/todo -a "<task>"
./bin/todo -a "Something to do"
```

To list saved tasks, use the `-[l|list]` flag

```bash
# Make sure you're in the todo/cli/todo directory
./bin/todo -l
```

To complete a task, use the `-[c|complete]` flag

```bash
# Make sure you're in the todo/cli/todo directory
# ./bin/todo -c <index-of-task> 
./bin/todo -c 0 
```

To delete a task, use the `-[d|delete]` flag

```bash
# Make sure you're in the todo/cli/todo directory
# ./bin/todo -d <index-of-task> 
./bin/todo -d 0 
```