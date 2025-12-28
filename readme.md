# pixel-font-baker

Tiny Go utility to generate **BMFont-compatible (.fnt)** bitmap fonts from a PNG spritesheet.

This is a one-shot CLI tool. No GUI. No magic.

---

## PNG requirements

- Glyphs are arranged in a **fixed grid**
- Order: **left → right, top → bottom**
- **White glyphs on transparent background**
  (required if you want to color text in your renderer)
- No margins or spacing between cells

---

## Default character order

By default, the font is assumed to start from **space (ASCII 32)** and continue in standard ASCII order:

␣ ! “ # $ % & ’ ( ) * + , - . / 0 1 2 …

If your bitmap uses a different order, you **must provide a custom charset**.

---

## Usage

```bash
go run fontgen.go <image.png> <cell_width> <cell_height> [first_char] [charset]

Examples

Default ASCII font (starting from space):

go run fontgen.go font.png 7 7

Explicit ASCII start:

go run fontgen.go font.png 7 7 32

Custom character order:

go run fontgen.go font.png 7 7 0 "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"