include $(GOROOT)/src/Make.inc

TARG:=github.com/Agon/baukasten

GOFILES:=windowmanager.go\
graphicsettings.go\
events.go\
vector.go\
color.go\
sdldriver.go

include $(GOROOT)/src/Make.pkg