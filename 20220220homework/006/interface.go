package main

type Inerface interface {
	PostStatus(pi *PersonalInformation) error
	DeleteStatus(id string) error
	GetStatus() ([]*CircleWithFRDInfo, error)
	DeleteAll() error
}
