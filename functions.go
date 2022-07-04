package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/matthewhartstonge/argon2"
)

func verifyPass(c Config) bool {
	ok, err := argon2.VerifyEncoded([]byte(c.pass), []byte(c.encoded))
	if err != nil {
		panic(err)
	}

	return ok
}

func printHeaders(argon argon2.Config) {
	fmt.Printf("HashLength:  %d\n", argon.HashLength)
	fmt.Printf("SaltLength:  %d\n", argon.SaltLength)
	fmt.Printf("TimeCost:    %d\n", argon.TimeCost)
	fmt.Printf("MemoryCost:  %d MB\n", argon.MemoryCost/1024)
	fmt.Printf("Parallelism: %d\n", argon.Parallelism)
	fmt.Printf("Mode:        %s\n", argon.Mode)
	fmt.Printf("Version:     %s\n", argon.Version)
	fmt.Println()
}

func encrypt(pass string, argon argon2.Config) string {

	encodedPass, err := argon.HashEncoded([]byte(pass))
	if err != nil {
		panic(err)
	}
	return string(encodedPass)
}

func start() (argon2.Config, Config) {
	pass := flag.String("pass", "", "The password to encrypt")
	encoded := flag.String("encoded", "", "The encoded password to verify")
	parallelism := flag.Int("paral", 1, "Parallelism specifies the amount of threads to use.")
	memoryCost := flag.Int("mem", 3906, "MemoryCost specifies the amount of memory to use in Kibibytes")
	version := flag.Int("version", 19, "Version specifies the argon2 version to be used.")
	iteration := flag.Int("iteration", 4, "(timeCost)specifies the number of iterations of argon2.")
	verbose := flag.Bool("verbose", false, "print some info about argon 2 parameters when generating a new password")
	versionTool := flag.Bool("v", false, "print the version of the tool and exit")
	flag.Parse()

	if *versionTool {
		fmt.Printf("%s\n", Version)
		os.Exit(0)
	}

	if *pass == "" {
		fmt.Printf("I need a password please?\n\n")
		fmt.Printf("argon2 version: %s\n", Version)
		flag.Usage()
		os.Exit(1)
	}

	argon := argon2.DefaultConfig()
	argon.Parallelism = uint8(*parallelism)
	argon.MemoryCost = uint32(*memoryCost)
	argon.Version = argon2.Version(*version)
	argon.TimeCost = uint32(*iteration)

	return argon, Config{
		pass:    *pass,
		encoded: *encoded,
		verbose: *verbose,
	}
}
