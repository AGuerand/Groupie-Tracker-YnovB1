package main

import (
	model "Groupie-Tracker/src/models"
	"fmt"
)

func main() {
	Artists := model.LoadData()
	for index, data := range Artists {
		fmt.Println("Artist n:", index+1, " Nom:", data.Name, " Membre:", data.Members)
		fmt.Println("Date de création:", data.CreationDate, " Date du premier album:", data.FirstAlbum)
		fmt.Println("image du groupe:", data.Image)
		fmt.Println("Lieu des concerts:", data.Locations)
		fmt.Println("Dates des concerts:", data.Dates)
		fmt.Println()
	}
}
