package main

import (
	"fmt"
	"image"
	_ "image/png"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run fontgen.go <image.png> <cell_width> <cell_height> [first_char] [charset]")
		fmt.Println("")
		fmt.Println("Examples:")
		fmt.Println("  go run fontgen.go font.png 8 12")
		fmt.Println("  go run fontgen.go font.png 8 12 32")
		fmt.Println("  go run fontgen.go font.png 8 12 32 \" !\\\"#$%&'()*+,-./0123456789\"")
		fmt.Println("")
		fmt.Println("Arguments:")
		fmt.Println("  image.png   - path to font spritesheet")
		fmt.Println("  cell_width  - width of each character cell in pixels")
		fmt.Println("  cell_height - height of each character cell in pixels")
		fmt.Println("  first_char  - ASCII code of first character (default: 32 = space)")
		fmt.Println("  charset     - custom character order (overrides first_char)")
		os.Exit(1)
	}

	imagePath := os.Args[1]
	var cellW, cellH int
	fmt.Sscanf(os.Args[2], "%d", &cellW)
	fmt.Sscanf(os.Args[3], "%d", &cellH)

	firstChar := 32
	if len(os.Args) > 4 {
		fmt.Sscanf(os.Args[4], "%d", &firstChar)
	}

	charset := ""
	if len(os.Args) > 5 {
		charset = os.Args[5]
	}

	file, err := os.Open(imagePath)
	if err != nil {
		fmt.Printf("Error opening image: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Printf("Error decoding image: %v\n", err)
		os.Exit(1)
	}

	bounds := img.Bounds()
	imgW := bounds.Max.X
	imgH := bounds.Max.Y

	cols := imgW / cellW
	rows := imgH / cellH
	totalChars := cols * rows

	fmt.Printf("Image: %dx%d, Cell: %dx%d, Grid: %dx%d (%d chars)\n", imgW, imgH, cellW, cellH, cols, rows, totalChars)

	baseName := strings.TrimSuffix(filepath.Base(imagePath), filepath.Ext(imagePath))
	outputPath := filepath.Join(filepath.Dir(imagePath), baseName+".fnt")

	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("info face=\"%s\" size=%d bold=0 italic=0 charset=\"\" unicode=1 stretchH=100 smooth=0 aa=1 padding=0,0,0,0 spacing=1,1 outline=0\n", baseName, cellH))
	sb.WriteString(fmt.Sprintf("common lineHeight=%d base=%d scaleW=%d scaleH=%d pages=1 packed=0 alphaChnl=0 redChnl=0 greenChnl=0 blueChnl=0\n", cellH, cellH-2, imgW, imgH))
	sb.WriteString(fmt.Sprintf("page id=0 file=\"%s\"\n", filepath.Base(imagePath)))
	sb.WriteString(fmt.Sprintf("chars count=%d\n", totalChars))

	for i := 0; i < totalChars; i++ {
		col := i % cols
		row := i / cols
		x := col * cellW
		y := row * cellH

		var charID int
		if charset != "" && i < len(charset) {
			charID = int(charset[i])
		} else {
			charID = firstChar + i
		}

		sb.WriteString(fmt.Sprintf("char id=%d x=%d y=%d width=%d height=%d xoffset=0 yoffset=0 xadvance=%d page=0 chnl=15\n",
			charID, x, y, cellW, cellH, cellW))
	}

	err = os.WriteFile(outputPath, []byte(sb.String()), 0644)
	if err != nil {
		fmt.Printf("Error writing .fnt file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Generated: %s\n", outputPath)
	fmt.Println("\nCharacter mapping:")
	for i := 0; i < totalChars && i < 96; i++ {
		var charID int
		if charset != "" && i < len(charset) {
			charID = int(charset[i])
		} else {
			charID = firstChar + i
		}
		if charID >= 32 && charID < 127 {
			fmt.Printf("  %d: '%c'\n", i, charID)
		} else {
			fmt.Printf("  %d: (code %d)\n", i, charID)
		}
	}
	if totalChars > 96 {
		fmt.Printf("  ... and %d more\n", totalChars-96)
	}
}