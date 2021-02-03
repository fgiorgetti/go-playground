# Bash Completion Example using Cobra Command API

This example demonstrates how to use bash completion along with
[Cobra Command API](https://github.com/spf13/cobra).


## The sample application

A sample application is produced here, named `bashcomp`.

It demonstrates an approach to integrate bash completion with a Go application
using Cobra.

In this example, the complete function is defined in a separate shell script
file named `bash_completion.sh`.

The contents from `bash_completion.sh` will be bound to the `BashCompletionEncoded`
variable at build time, allowing you to write your complete function more easily.


## Building the bashcomp application

To build the application and produce a bash completion shell script, run;

```
make
```

After building it, you can simply source the produced `bashcomp.bash.inc` file, like:

```
source bashcomp.bash.inc
```

## Running and validating auto complete options

Type `bashcomp <tab><tab>` and you will must the list of available options.
If you wanna try a more specific one, type: `./bashcomp thanks <tab><tab>`.


***NOTE:***

    Remember to install bashcomp to a directory that is part of your PATH,
    if you use it in current directory with ./ completion won't work.
