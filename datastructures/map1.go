package main

import "fmt"

func main() {
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"

	fmt.Println(elements["Be"])
}

// elements := map[string]string{
// 	"H":  "Hydrogen",
// 	"He": "Helium",
// 	"Li": "Lithium",
// 	"Be": "Beryllium",
// 	"B":  "Boron",
// 	"C":  "Carbon",
// 	"N":  "Nitrogen",
// 	"O":  "Oxygen",
// 	"F":  "Fluorine",
// 	"Ne": "Neon",
//   }
