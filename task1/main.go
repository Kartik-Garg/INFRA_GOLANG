package main

import (
	"fmt"
)

type Profile struct {
	Name           string
	Username       string
	Designation    string
	ContactNumber  string
	ProfilePicture ProfilePicture
}

type ProfilePicture struct {
	ImageName string
	ImagePath string
}

// function to complete the profile and return as struct
func completeProfile(p Profile) Profile {
	return p
}

// method to take profile and modify it by taking modifications as arguments
func (p Profile) updateProfile(modifications int) Profile {
	//writing switch case for modifications
	switch modifications {

	case 1:
		fmt.Printf("Please enter the new name to be modified:")
		var name string
		fmt.Scanln(name)
		p.Name = name
		return p
	case 2:
		fmt.Println("Please enter the new username to be modified:")
		var username string
		fmt.Scan(&username)
		p.Username = username
		return p
	case 3:
		fmt.Println("Please enter the new designation to be modified:")
		var designation string
		fmt.Scan(&designation)
		p.Designation = designation
		return p
	case 4:
		fmt.Println("Enter new contact number to be modified:")
		var contact string
		fmt.Scan(&contact)
		p.ContactNumber = contact
		return p
	case 5:
		fmt.Println("Press 1 to make changes in image name\nPress 2 to make changes in image path\nPress 3 to make changes in both image name and image path: ")
		var profile_changes int
		fmt.Scan(&profile_changes)
		if profile_changes == 1 {
			fmt.Println("Enter the new image name:")
			var image_name string
			fmt.Scan(&image_name)
			p.ProfilePicture.ImageName = image_name
			return p
		}
		if profile_changes == 2 {
			fmt.Println("Enter the new image path:")
			var image_path string
			fmt.Scan(&image_path)
			p.ProfilePicture.ImagePath = image_path
			return p
		}
		if profile_changes == 3 {
			var image_name, image_path string
			fmt.Println("Enter the new image name:")
			fmt.Scan(&image_name)
			fmt.Println("Enter the new image path:")
			fmt.Scan(&image_path)
			p.ProfilePicture.ImageName = image_name
			p.ProfilePicture.ImagePath = image_path
			return p
		}

	}
	//returning profile
	return p
}

func main() {
	p := Profile{
		Name:          "Kartik",
		Username:      "kartikgarg",
		Designation:   "software engineer",
		ContactNumber: "12345",
		ProfilePicture: ProfilePicture{
			ImageName: "pic 1",
			ImagePath: "/src/go/task1",
		},
	}

	fmt.Println(completeProfile(p))

	//taking input from user for modification
	var num int
	fmt.Println("Please enter the number in respect with modifications to be made:\n 1.Name\n 2.Username\n 3.Designation\n 4.ContactNumber\n 5.ProfilePicture")
	fmt.Scanln(&num)

	//output after modifications
	fmt.Println("Profile after modificatoin is: ", p.updateProfile(num))
}
