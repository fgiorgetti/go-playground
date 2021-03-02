# Bash Completion Example using Cobra Command API (Updated to use go:embed)

This example demonstrates how to use bash completion along with
[Cobra Command API](https://github.com/spf13/cobra).


## The sample application

A sample application is produced here, named `bashcompembed`.

This is just an update to the previous `bashcomp` sample, that uses `go:embed`
available as of `go 1.16`.

It also demonstrates an approach to integrate bash completion with a Go application
using Cobra.

In this example, the complete function is defined in a separate shell script
file named `bash_completion.sh`.

The contents from `bash_completion.sh` will be bound to the `BashCompletionFunction`
variable at build time, allowing you to write your complete function more easily.


## Building the bashcomp application

To build the application and produce a bash completion shell script, run;

```
make
```

After building it, you can simply source the produced `bashcompembed.bash.inc` file, like:

```
source bashcompembed.bash.inc
```

## Running and validating auto complete options

Type `bashcompembed <tab><tab>` and you will see the list of available options.
If you want to try a more specific one, type: `bashcompembed thanks <tab><tab>`.


***NOTE:***

    Remember to install bashcomp to a directory that is part of your PATH,
    if you use it in current directory with ./ completion won't work.
