package main

import "fmt"

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
func (p Profile) updateProfile(modifications string) Profile {
	//modifying only the name here
	p.Name = modifications
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

	fmt.Println(p.updateProfile("andrew"))
}
