package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	if len(os.Args) != 2 {
		return fmt.Errorf("exactly one filename must be provided to run post-processing on")
	}
	filename := os.Args[1]
	log.Printf("running post processing on file %v", filename)

	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	file := string(fileBytes)
	file, err = fixDoubleStarProperties(file)
	if err != nil {
		return err
	}

	file, err = fixDoubleBlankProperties(file)
	if err != nil {
		return err
	}

	file, err = fixThumbsUpThumbsDownProperties(file)
	if err != nil {
		return err
	}

	// TODO(kfcampbell): verify file permission is what we want
	err = os.WriteFile(filename, []byte(file), 0666)
	if err != nil {
		return err
	}

	return nil
}

// fixDoubleStarProperties fixes generated enums where "*" and "Star" are
// both webhook events parsed as "Star"
func fixDoubleStarProperties(inputFile string) (string, error) {
	if len(inputFile) < 1 {
		return "", fmt.Errorf("inputFile must not be empty")
	}

	// find instances of "/// Enum Star for value: *" and replace with "Enum All for value: *"
	inputFile = strings.ReplaceAll(inputFile, "/// Enum Star for value: *", "/// Enum All for value: *")

	// find instances of "            Star = 1," and replace with "            All = 1,"
	inputFile = strings.ReplaceAll(inputFile, "            Star = 1,", "            All = 1,")
	return inputFile, nil
}

// fixDoubleBlankProperties removes one of the two BLANK enum properties from Repository11.cs
// TODO(kfcampbell): attempt to figure out where this is defined inside GitHub and PR it
func fixDoubleBlankProperties(inputFile string) (string, error) {
	if len(inputFile) < 1 {
		return "", fmt.Errorf("inputFile must not be empty")
	}

	toReplace := `/// <summary>
            /// Enum BLANK for value: BLANK
            /// </summary>
            [EnumMember(Value = "BLANK")]
            BLANK = 3,

            /// <summary>
            /// Enum BLANK for value: BLANK
            /// </summary>
            [EnumMember(Value = "BLANK")]
            BLANK = 4
`

	// replace duplicate blanks if they exist
	if strings.Contains(inputFile, toReplace) {
		inputFile = strings.ReplaceAll(inputFile, toReplace, `/// <summary>
            /// Enum BLANK for value: BLANK
            /// </summary>
            [EnumMember(Value = "BLANK")]
            BLANK = 3
		`)
	}
	return inputFile, nil
}

func fixThumbsUpThumbsDownProperties(inputFile string) (string, error) {
	if len(inputFile) < 1 {
		return "", fmt.Errorf("inputFile must not be empty")
	}

	// conditions for needing to fix thumbs up/thumbs down issue:
	// filename contains "Reaction"

	toReplace := `/// <summary>
            /// Enum _1 for value: +1
            /// </summary>
            [EnumMember(Value = "+1")]
            _1 = 1,

            /// <summary>
            /// Enum _1 for value: -1
            /// </summary>
            [EnumMember(Value = "-1")]
            _1 = 2,`

	replaceWith := `/// <summary>
            /// Enum ThumbsUp for value: +1
            /// </summary>
            [EnumMember(Value = "+1")]
            ThumbsUp = 1,

            /// <summary>
            /// Enum ThumbsDown for value: -1
            /// </summary>
            [EnumMember(Value = "-1")]
            ThumbsDown = 2,`

	if strings.Contains(inputFile, toReplace) {
		inputFile = strings.ReplaceAll(inputFile, toReplace, replaceWith)
	}

	toReplace = `/// <param name="_1">_1.</param>
        /// <param name="_1">_1.</param>`
	replaceWith = `/// <param name="thumbsUp">thumbsUp.</param>
        /// <param name="thumbsDown">thumbsDown.</param>`

	if strings.Contains(inputFile, toReplace) {
		inputFile = strings.ReplaceAll(inputFile, toReplace, replaceWith)
	}

	toReplace = "int _1 = default(int), int _1 = default(int)"
	replaceWith = "int thumbsUp = default(int), int thumbsDown = default(int)"

	if strings.Contains(inputFile, toReplace) {
		inputFile = strings.ReplaceAll(inputFile, toReplace, replaceWith)
	}

	toReplace = `_1 = _1;
            _1 = _1;`

	replaceWith = `ThumbsUp = thumbsUp;
            ThumbsDown = thumbsDown;`

	if strings.Contains(inputFile, toReplace) {
		inputFile = strings.ReplaceAll(inputFile, toReplace, replaceWith)
	}

	toReplace = `        /// <summary>
        /// Gets and Sets _1
        /// </summary>
        public int _1 { get; private set; }
        /// <summary>
        /// Gets and Sets _1
        /// </summary>
        public int _1 { get; private set; }`

	replaceWith = `        /// <summary>
        /// Gets and Sets ThumbsUp
        /// </summary>
        public int ThumbsUp { get; private set; }
        /// <summary>
        /// Gets and Sets ThumbsDown
        /// </summary>
        public int ThumbsDown { get; private set; }`

	if strings.Contains(inputFile, toReplace) {
		inputFile = strings.ReplaceAll(inputFile, toReplace, replaceWith)
	}

	toReplace = `sb.Append("  _1: ").Append(_1).Append("\n");
            sb.Append("  _1: ").Append(_1).Append("\n");`

	replaceWith = `sb.Append("  ThumbsUp: ").Append(ThumbsUp).Append("\n");
            sb.Append("  ThumbsDown: ").Append(ThumbsDown).Append("\n");`
	return inputFile, nil
}
