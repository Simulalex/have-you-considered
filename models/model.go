package models

type Technology struct {
	Id         int
	Name       string
	Title string
	Summary    string
	Author     *Author
}

type Strength struct {
	Id       int
	Strength string
}

type Weakness struct {
	Id       int
	Weakness string
}

type Author struct {
	Id       int
	Username string
}
