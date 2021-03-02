package main

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	//go:embed bash_completion.sh
	BashCompletionFunction string
	rootCmd                *cobra.Command
)

func main() {
	// Initializing cobra
	rootCmd = &cobra.Command{
		Use:                    "bashcompembed",
		Aliases:                nil,
		SuggestFor:             nil,
		Short:                  "A demonstration on how to use bash completion using Cobra Command API",
		BashCompletionFunction: BashCompletionFunction,
	}

	// Hello command
	rootCmd.AddCommand(helloCommand())
	rootCmd.AddCommand(goodbyeCommand())
	rootCmd.AddCommand(thanksCommand())
	rootCmd.AddCommand(completionCommand())

	err := rootCmd.Execute()
	if err != nil {
		fmt.Printf("Error executing rootCmd: %s\n", err)
	}
}

func completionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "completion",
		Short: "Print out bash completion functions to help using bashcomp",
		Long: `You can output the content from the 'completion' command into a ${HOME}/bashcomp.bash.inc
source it or source it from ${HOME}/.bash_profile`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return rootCmd.GenBashCompletion(os.Stdout)
		},
	}
}

func thanksCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "thanks [from|to] name",
		Short: "Send or receive a thanks to/from someone",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 2 {
				return fmt.Errorf(cmd.Use)
			}
			if args[0] != "from" && args[0] != "to" {
				return fmt.Errorf("you have to specify 'from' or 'to'")
			}
			return nil
		},
		BashCompletionFunction: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			fromTo := args[0]
			name := args[1]

			if fromTo == "from" {
				fmt.Printf("%s says Thank you!\n", name)
			} else {
				fmt.Printf("Thank you very much fellow %s!\n", name)
			}
			return nil
		},
	}
}

func goodbyeCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "goodbye name",
		Short: "Say goodbye to a given person",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("%s command expects a name to say goodbye", cmd.Name())
			}
			return nil
		},
		BashCompletionFunction: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]
			fmt.Printf("Goodbye fellow %s\n", name)
			return nil
		},
	}
}

func helloCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "hello name",
		Short: "Say hello to a given person",

		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("%s command expects a name to salute", cmd.Name())
			}
			return nil
		},
		BashCompletionFunction: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]
			fmt.Printf("Hello my dear %s\n", name)
			return nil
		},
	}
}
