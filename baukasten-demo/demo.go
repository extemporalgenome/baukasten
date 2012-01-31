package main

type Demo interface {
	Name() string
	Description() string
	Load() error
	Unload() error
	Update()
}
