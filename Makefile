include $(GOROOT)/src/Make.inc

TARG:=github.com/Agon/baukasten

GOFILES:=windowmanager.go\
displaymanager.go\
graphicsettings.go\
sdldriver.go\
ogldriver.go\
events.go\
color.go\
vector.go


include $(GOROOT)/src/Make.pkg