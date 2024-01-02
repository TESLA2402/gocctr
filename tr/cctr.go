package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	var deleteMode, squeezeMode bool
	flag.BoolVar(&deleteMode, "d", false, "Delete characters specified in the substitution rule from the input.")
	flag.BoolVar(&squeezeMode, "s", false, "Squeeze multiple occurrences of characters into a single instance.")
	flag.Parse()

	args := flag.Args()
	if len(args) < 2 && !deleteMode && !squeezeMode {
		fmt.Println("Invalid number of arguments. Please provide two characters or a character class separated by a space.")
		return
	}
	if deleteMode && len(args) != 1 {
		fmt.Println("Invalid number of arguments. Please provide characters to delete.")
		return
	}

	deleteString := args[0]
	if squeezeMode && len(args) != 1 {
		fmt.Println("Invalid number of arguments. Please provide characters to squeeze.")
		return
	}

	squeezeString := args[0]

	if deleteMode {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := scanner.Text()
			deleted := deleteChars(line, deleteString)
			fmt.Print(deleted)
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading input:", err)
		}
		return

	} else if squeezeMode {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := scanner.Text()
			squeezed := squeezeChars(line, squeezeString)
			fmt.Print(squeezed)
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading input:", err)
		}

		return

	} else {
		parts := args
		if len(parts) != 2 {
			fmt.Println("Invalid substitution rule. Please provide two characters or a character class separated by a space.")
			return
		}
		fromSpec := parts[0]
		toSpec := parts[1]
		fromChars, err := expandSpec(fromSpec)
		if err != nil {
			fmt.Println("Error parsing 'from' specifier:", err)
			return
		}

		toChars, err := expandSpec(toSpec)
		if err != nil {
			fmt.Println("Error parsing 'to' specifier:", err)
			return
		}
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := scanner.Text()
			transliterated := replaceChars(line, fromChars, toChars)
			fmt.Println(transliterated)
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading input:", err)
		}

		return
	}
}

func expandRange(input string) (string, error) {
	re := regexp.MustCompile(`([a-zA-Z0-9])-([a-zA-Z0-9])`)
	matches := re.FindAllStringSubmatch(input, -1)

	var result strings.Builder

	for _, match := range matches {
		if len(match) == 3 {
			startChar := match[1][0]
			endChar := match[2][0]

			for char := startChar; char <= endChar; char++ {
				result.WriteByte(char)
			}
		}
	}

	return result.String(), nil
}
func expandSpec(spec string) (string, error) {
	if spec == "A-Z" || spec == "a-z" {
		return expandClass(spec)
	}
	if strings.HasPrefix(spec, "[:") && strings.HasSuffix(spec, ":]") {
		class := spec[2 : len(spec)-2]
		return expandClass(class)
	}

	return spec, nil
}

func expandClass(class string) (string, error) {
	switch class {
	case "alnum":
		return expandRange("a-zA-Z0-9")
	case "alpha":
		return expandRange("a-zA-Z")
	case "blank":
		return " \t", nil
	case "cntrl":
		return expandRange("\x00-\x1F\x7F")
	case "digit":
		return expandRange("0-9")
	case "lower":
		return expandRange("a-z")
	case "print":
		return expandRange(" -~")
	case "punct":
		return expandRange(`!"#$%&'()*+,-./:;<=>?@[\\]^_` + "`{|}~")
	case "rune":
		return expandRange("\x00-\x10FFFF")
	case "space":
		return expandRange(" \t\n\r\v\f")
	case "special":
		return expandRange(`!@#$%^&*()-=_+[]{}|;:'",.<>?/`)
	case "upper":
		return expandRange("A-Z")
	case "A-Z":
		return expandRange("A-Z")
	case "a-z":
		return expandRange("a-z")
	default:
		return class, nil
	}
}

func replaceChars(input string, fromChars, toChars string) string {
	for i := 0; i < len(fromChars); i++ {
		input = strings.ReplaceAll(input, string(fromChars[i]), string(toChars[i%len(toChars)]))
	}
	return input
}

func deleteChars(input, charsToDelete string) string {
	for _, char := range charsToDelete {
		input = strings.ReplaceAll(input, string(char), "")
	}
	return input
}

func squeezeChars(input, charsToSqueeze string) string {
	squeezeSet := make(map[rune]struct{})

	for _, char := range charsToSqueeze {
		squeezeSet[char] = struct{}{}
	}

	var result strings.Builder
	var prevChar rune

	for _, char := range input {
		if _, shouldSqueeze := squeezeSet[char]; shouldSqueeze {
			if char != prevChar {
				result.WriteRune(char)
			}
		} else {
			result.WriteRune(char)
		}

		prevChar = char
	}

	return result.String()
}
