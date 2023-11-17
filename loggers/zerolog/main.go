package main

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

var (
	p1 Person
)

func init() {
	p1 = Person{
		Name:  "Fernando",
		Age:   43,
		Email: "fgiorgetti@gmail.com",
	}
}

func main() {
	// UNIX Time is faster and smaller than most timestamps
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	consoleLogging()
	jsonLogging()
}

func jsonLogging() {
	fmt.Println("")
	fmt.Println("JSON logging")
	fmt.Println("")
	log := zerolog.New(os.Stdout).With().Timestamp().Logger().Level(zerolog.InfoLevel)
	logStuff(log)
}

func consoleLogging() {
	fmt.Println("")
	fmt.Println("Console logging")
	fmt.Println("")
	log := zerolog.New(zerolog.NewConsoleWriter()).With().Timestamp().Logger().Level(zerolog.InfoLevel)
	logStuff(log)
}

func logStuff(log zerolog.Logger) {
	log.Info().Msg("info from zerolog")
	log.Warn().Msg("warn from zerolog")
	log.Debug().Msg("debug from zerolog") // hidden as log level == info
	log.Info().Any("person", p1).Msg("Creating person")
	err := fmt.Errorf("unable to write to disk")
	log.Error().Stack().Err(err).
		Any("person", p1).
		Msg("error creating person")
}
