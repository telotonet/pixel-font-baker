fontgen — Bitmap Font Generator (.fnt)

Minimal Go utility to generate BMFont-compatible bitmap fonts from a PNG spritesheet.

⸻

Requirements for PNG
	•	Glyphs are laid out in a fixed grid
	•	Order: left → right, top → bottom
	•	White glyphs on transparent background (required for colored text rendering)
	•	No margins or spacing between cells

⸻

Default character order

By default, fonts are assumed to start from space (ASCII 32) and continue in standard ASCII order:

␣ ! " # $ % & ' ( ) * + , - . / 0 1 2 ...

If your bitmap uses a different order, you must explicitly specify it via charset.

⸻

Usage

go run fontgen.go <image.png> <cell_width> <cell_height> [first_char] [charset]

Examples

Default ASCII (starting from space):

go run fontgen.go font.png 7 7

Explicit ASCII start (space = 32):

go run fontgen.go font.png 7 7 32

Custom character order:

go run fontgen.go font.png 7 7 0 "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"


⸻

Output
	•	Generates <image>.fnt next to the PNG
	•	Compatible with standard BMFont loaders