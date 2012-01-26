package main

import (
	"fmt"
)

type DemoManager struct {
	running Demo
	demos   []Demo
}

func NewDemoManager() *DemoManager {
	return &DemoManager{demos: make([]Demo, 0)}
}

func (m *DemoManager) AddDemos(demos ...Demo) {
	m.demos = append(m.demos, demos...)
}

func (m *DemoManager) Demos() []Demo {
	return m.demos
}

func (m *DemoManager) Load(name string) error {
	for i := range m.demos {
		if m.demos[i].Name() == name {
			if m.running != nil {
				err := m.running.Unload()
				if err != nil {
					return err
				}
			}
			m.running = m.demos[i]
			return m.running.Load()
		}
	}
	return fmt.Errorf("Couldn't find demo %s", name)
}

func (m *DemoManager) Unload() {
	if m.running != nil {
		m.running.Unload()
		m.running = nil
	}
}

func (m *DemoManager) Update() {
	if m.running == nil {
		return
	}
	m.running.Update()
}
