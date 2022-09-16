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

func completeProfile(name, userName, designation, contactNumber, imageName, imagePath string) *Profile {
	//using pointer return type so the changes do not remain local to this specific function, pointers pass the actual address of the type rather than making a copy of it
	return &Profile{
		Name:          name,
		Username:      userName,
		Designation:   designation,
		ContactNumber: contactNumber,
		ProfilePicture: ProfilePicture{
			ImageName: imageName,
			ImagePath: imagePath,
		},
	}
}

// func completeProfile(name, userName, designation, contactNumber, imageName, imagePath string) Profile {
// 	//using pointer return type so the changes do not remain local to this specific function, pointers pass the actual address of the type rather than making a copy of it
// 	return Profile{
// 		Name:          name,
// 		Username:      userName,
// 		Designation:   designation,
// 		ContactNumber: contactNumber,
// 		ProfilePicture: ProfilePicture{
// 			ImageName: imageName,
// 			ImagePath: imagePath,
// 		},
// 	}
// }

// method to update profile which takes all the modifications
func (p *Profile) UpdateProfile(name, userName, designation, contactNumber, imageName, imagePath string) {
	//taking a pointer here as well so the changes are made to actual profile rather than the local copy of this particular method
	if name != "" {
		p.Name = name
	}
	if userName != "" {
		p.Username = userName
	}
	if designation != "" {
		p.Designation = designation
	}
	if contactNumber != "" {
		p.ContactNumber = contactNumber
	}
	if imageName != "" {
		p.ProfilePicture.ImageName = imageName
	}
	if imagePath != "" {
		p.ProfilePicture.ImagePath = imagePath
	}

}

// func (p Profile) UpdateProfile(name, userName, designation, contactNumber, imageName, imagePath string) {
// 	//taking a pointer here as well so the changes are made to actual profile rather than the local copy of this particular method
// 	if name != "" {
// 		p.Name = name
// 	}
// 	if userName != "" {
// 		p.Username = userName
// 	}
// 	if designation != "" {
// 		p.Designation = designation
// 	}
// 	if contactNumber != "" {
// 		p.ContactNumber = contactNumber
// 	}
// 	if imageName != "" {
// 		p.ProfilePicture.ImageName = imageName
// 	}
// 	if imagePath != "" {
// 		p.ProfilePicture.ImagePath = imagePath
// 	}

// }

func main() {
	name := "Kartik"
	userName := "kartikgarg@infracloud.io"
	designation := "software engineer"
	contactNumber := "12345"
	imageName := "pic1"
	imagePath := "src/pics"

	//creating a new Profile here
	completedProfile := completeProfile(name, userName, designation, contactNumber, imageName, imagePath)
	//fmt.Printf("%p\n", completedProfile)
	fmt.Println(completedProfile)
	name = "Garg"
	//calling method to update the profile
	completedProfile.UpdateProfile(name, userName, designation, contactNumber, imageName, imagePath)
	//printing after updatiion
	fmt.Println(completedProfile)
	//since we have used pointers, we don't have to have a return statement and changes can be made on the same profile, rahter than a copy of it which is local to that specific method

}
