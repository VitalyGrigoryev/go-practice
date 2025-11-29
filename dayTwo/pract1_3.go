package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	// Создаём огромный срез аргументов (миллион слов)
	args := make([]string, 1_000_000)
	for i := range args {
		args[i] = "слово"
	}

	// Замеряем echo1 (конкатенация в цикле)
	start := time.Now()
	s, sep := "", ""
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
	fmt.Printf("echo1 (цикл): %v\n", time.Since(start))

	// Замеряем echo3 (strings.Join)
	start = time.Now()
	_ = strings.Join(args, " ")
	fmt.Printf("echo3 (Join): %v\n", time.Since(start))
}
