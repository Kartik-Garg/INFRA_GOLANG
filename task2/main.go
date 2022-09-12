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

func completeProfile(p Profile) Profile {
	return p
}

func (p Profile) updateProfile(modifications string) Profile {
	p.Name = modifications
	return p
}

//Task2: crearing interfaces

type ProfileMaker interface {
	UpdateProfilePicture(ProfilePicture)
	CheckDuplicateProfile(Profile) bool
}

func (p ProfilePicture) UpdateProfilePicture(pp ProfilePicture) {
	p = pp
	fmt.Println(p)
}

func (p1 Profile) CheckDuplicateProfile(p2 Profile) bool {
	if p1 == p2 {
		return true
	}
	return false
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

	pp := ProfilePicture{
		"image1",
		"imagePath",
	}

	pp.UpdateProfilePicture(pp)

	p2 := Profile{
		"Ananth",
		"dayakar",
		"SWE",
		"123",
		ProfilePicture{
			"pic1",
			"/src",
		},
	}

	p3 := Profile{
		Name:          "Kartik",
		Username:      "kartikgarg",
		Designation:   "software engineer",
		ContactNumber: "12345",
		ProfilePicture: ProfilePicture{
			ImageName: "pic 1",
			ImagePath: "/src/go/task1",
		},
	}

	//checking second method of interface
	fmt.Println(p.CheckDuplicateProfile(p2))

	//comparing p3 with p
	fmt.Println(p.CheckDuplicateProfile(p3))

	//comparing p with p
	fmt.Println(p.CheckDuplicateProfile(p))
}
