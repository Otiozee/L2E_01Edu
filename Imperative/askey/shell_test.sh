#!/bin/bash

echo "=============================="
echo "Default Banner"
echo "=============================="
go run . "Hello World"

echo
echo "=============================="
echo "Shadow Banner"
echo "=============================="
go run . "Hello World" shadow

echo
echo "=============================="
echo "Thinkertoy Banner"
echo "=============================="
go run . "Hello World" thinkertoy

echo
echo "=============================="
echo "Color - Entire Text"
echo "=============================="
go run . --color=red "banana"

echo
echo "=============================="
echo "Color - Target Word"
echo "=============================="
go run . --color=green World "Hello World"

echo
echo "=============================="
echo "RGB Color"
echo "=============================="
go run . --color="rgb(255,0,0)" H "Hello"

echo
echo "=============================="
echo "HEX Color"
echo "=============================="
go run . --color="#ff0000" me "Hummer"

echo
echo "=============================="
echo "HSL Color"
echo "=============================="
go run . --color="hsl(0,100%,50%)" car "Racecar in a car race"

echo
echo "=============================="
echo "Center Alignment"
echo "=============================="
go run . --align=center "Hello World"

echo
echo "=============================="
echo "Right Alignment"
echo "=============================="
go run . --align=right "Hello World"

echo
echo "=============================="
echo "Left Alignment"
echo "=============================="
go run . --align=left "Hello World"

echo
echo "=============================="
echo "Justify Alignment"
echo "=============================="
go run . --align=justify "Hello World from Go"

echo
echo "=============================="
echo "Multiline Input"
echo "=============================="
go run . --color=yellow "Hello\nThere"

echo
echo "=============================="
echo "Multiple Blank Lines"
echo "=============================="
go run . --color=cyan "Hello\n\nThere"

echo
echo "=============================="
echo "Symbols"
echo "=============================="
go run . --color=purple '!"#$%&'\''()*+'

echo
echo "=============================="
echo "Numbers"
echo "=============================="
go run . --color=blue "1 + 1 = 2"

echo
echo "=============================="
echo "Mixed Characters"
echo "=============================="
go run . --color=white '$%this 734Dbf, 123 frake!'

echo
echo "=============================="
echo "Save Output"
echo "=============================="
go run . --output=art.txt "Hello World"

echo
echo "=============================="
echo "Reverse Output"
echo "=============================="
go run . --reverse=art.txt

echo
echo "=============================="
echo "Target Coloring"
echo "=============================="
go run . --color=red kit "a king kitten have kit"

echo
echo "=============================="
echo "Done"
echo "=============================="