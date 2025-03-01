package main

import (
	"fmt"
	services2 "github.com/johnfercher/chaos/deco/services"
	chaos2 "github.com/johnfercher/chaos/deco/template/chaos"
	"github.com/johnfercher/chaos/struct/structservices"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "deco",
	Short: "Deco is a decorator deco for go code",
	Run:   Command,
}

func Command(cmd *cobra.Command, args []string) {
	_type, _ := cmd.Flags().GetString("type")

	input, err := cmd.Flags().GetString("input")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if input == "" {
		fmt.Fprintln(os.Stderr, "error: input is empty")
		os.Exit(1)
	}

	_interface, err := cmd.Flags().GetString("interface")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Printf("Generate %s decoratorGenerator for %s in %s\n", _type, _interface, input)

	file := structservices.NewFile()
	interpreter := structservices.NewInterfaceInterpreter()
	decoratorGenerator := services2.NewDecoratorGenerator("Chaos", chaos2.Decorator, chaos2.Method)
	orchestrator := services2.NewGenerationOrchestrator(file, interpreter, decoratorGenerator)

	err = orchestrator.Generate(input, _interface)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func main() {
	rootCmd.PersistentFlags().String("type", "chaos", "The decorator type generation")
	rootCmd.PersistentFlags().String("input", "", "Input file")
	rootCmd.PersistentFlags().String("interface", "", "Interface to generate decorator")
	//rootCmd.SetArgs([]string{"--type=chaos", "--input=docs/examples/interfaces.go", "--interface=SingleParameterWithTwoReturns"})
	rootCmd.SetArgs([]string{"--type=chaos", "--input=docs/examples/medium/productapi/internal/api/httphandler.go"})

	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
