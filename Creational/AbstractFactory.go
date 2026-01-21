package main

import "fmt"

// Abstract products

type iLipStick interface {
	getShade() string
	getBrand() string
	getTexture() string
}

type iLipLiner interface {
	getShade() string
	getBrand() string
	getWidthInMm() int
}

type lipStick struct {
	shade   string
	brand   string
	texture string
}

func (l *lipStick) getShade() string {
	return l.shade
}

func (l *lipStick) getBrand() string {
	return l.brand
}

func (l *lipStick) getTexture() string {
	return l.texture
}

type lipLiner struct {
	shade     string
	brand     string
	widthInMm int
}

func (l *lipLiner) getShade() string {
	return l.shade
}

func (l *lipLiner) getBrand() string {
	return l.brand
}

func (l *lipLiner) getWidthInMm() int {
	return l.widthInMm
}

// Concrete products

type maybellineLipstick struct {
	lipStick
}

type maybellineLipliner struct {
	lipLiner
}

type revlonLipstick struct {
	lipStick
}

type revlonLipliner struct {
	lipLiner
}

// Abstract factory interface
type makeupFactory interface {
	makeLipstick(string, string) iLipStick
	makeLipliner(string, int) iLipLiner
}

// Concrete factory 1
type maybellineFactory struct {
}

func (m *maybellineFactory) makeLipstick(shade, texture string) iLipStick {
	return &maybellineLipstick{
		lipStick: lipStick{
			shade:   shade,
			texture: texture,
			brand:   "maybelline New York",
		},
	}
}

func (m *maybellineFactory) makeLipliner(shade string, widthInMm int) iLipLiner {
	return &maybellineLipliner{
		lipLiner: lipLiner{
			shade:     shade,
			widthInMm: widthInMm,
			brand:     "maybelline New York",
		},
	}
}

// Concrete factory 2
type revlonFactory struct {
}

func (m *revlonFactory) makeLipstick(shade, texture string) iLipStick {
	return &revlonLipstick{
		lipStick: lipStick{
			shade:   shade,
			texture: texture,
			brand:   "revlon",
		},
	}
}

func (m *revlonFactory) makeLipliner(shade string, widthInMm int) iLipLiner {
	return &revlonLipliner{
		lipLiner: lipLiner{
			shade:     shade,
			widthInMm: widthInMm,
			brand:     "revlon",
		},
	}
}

// method to create and fetch respective factory
func getFactory(s string) makeupFactory {
	if s == "maybelline" {
		return &maybellineFactory{}
	} else if s == "revlon" {
		return &revlonFactory{}
	} else {
		fmt.Println("Err creating factory of", s)
		return nil
	}
}

// client code from here-

// helper method to print Lipstick Deets
func printLipstickDeets(l iLipStick) {
	fmt.Println("lipStick brand is", l.getBrand())
	fmt.Println("lipStick shade is", l.getShade())
	fmt.Println("lipStick texture is", l.getTexture())
}

// helper method to print Lipliner Deets
func printLipslinerDeets(l iLipLiner) {
	fmt.Println("lipLiner brand is", l.getBrand())
	fmt.Println("lipLiner shade is", l.getShade())
	fmt.Println("lipLiner width in mm is", l.getWidthInMm())
}

func main() {
	f1 := getFactory("maybelline")
	f2 := getFactory("revlon")

	f1lipstick := f1.makeLipstick("pink", "bar")
	f1lipliner := f1.makeLipliner("violet", 5)

	f2lipstick := f2.makeLipstick("brown nude", "liquid")
	f2lipliner := f2.makeLipliner("dark brown", 10)

	fmt.Println("\nf1 deets:")
	printLipstickDeets(f1lipstick)
	printLipslinerDeets(f1lipliner)

	fmt.Println("\nf2 deets:")
	printLipstickDeets(f2lipstick)
	printLipslinerDeets(f2lipliner)
}

/*
Sample response:

saitejomurtula@Sais-MacBook-Pro Creational % go run AbstractFactory.go

f1 deets:
lipStick brand is maybelline New York
lipStick shade is pink
lipStick texture is bar
lipLiner brand is maybelline New York
lipLiner shade is violet
lipLiner width in mm is 5

f2 deets:
lipStick brand is revlon
lipStick shade is brown nude
lipStick texture is liquid
lipLiner brand is revlon
lipLiner shade is dark brown
lipLiner width in mm is 10
saitejomurtula@Sais-MacBook-Pro Creational %

*/
