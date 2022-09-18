package main

import (
	"bytes"
	"encoding/gob"
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

type ProfileMaker interface {
	UpdateProfilePicture(ProfilePicture)
	CheckDuplicateProfile(Profile) bool
}

func (p *Profile) UpdateProfilePicture(pp ProfilePicture) {
	//here p profile is the profile used to call this particular method
	p.ProfilePicture = pp
}

func (p *Profile) CheckDuplicateProfile(profileToCheck Profile) bool {
	//to make sure that it does not work only on the primitive data types, we can encode it into bytes and then check
	//can use package gobs.encoder which takes care of data types including structs
	var profileBuf bytes.Buffer
	//try encoding it original profile on which method is called into bytes
	err := gob.NewEncoder(&profileBuf).Encode(p)
	if err != nil {
		fmt.Println("Encoding has failed")
		return false
	}
	var profileToCompare bytes.Buffer
	err = gob.NewEncoder(&profileToCompare).Encode(profileToCheck)
	if err != nil {
		fmt.Println("Failed to encode")
		return false
	}

	return profileBuf.String() == profileToCompare.String()
}

func main() {
	name := "Kartik"
	userName := "kartikgarg@infracloud.io"
	designation := "software engineer"
	contactNumber := "12345"
	imageName := "pic1"
	imagePath := "src/pics"

	p1 := completeProfile(name, userName, designation, contactNumber, imageName, imagePath)

	fmt.Println(p1)

	var profileMakerVariable ProfileMaker

	profileMakerVariable = p1

	profileMakerVariable.UpdateProfilePicture(ProfilePicture{
		ImageName: "pic2",
		ImagePath: "newPath",
	})

	fmt.Println(profileMakerVariable)

	//creating a profile to check for duplicacy
	p2 := completeProfile(name, userName, designation, contactNumber, imageName, imagePath)

	profileMakerVariable = p2

	fmt.Println(profileMakerVariable.CheckDuplicateProfile(*p1))
}
