include $(GOROOT)/src/Make.inc

TARG:=github.com/Agon/baukasten

GOFILES:=engine.go\
graphicsettings.go\
events.go\
color.go\
math.go\
vector.go\
matrix4x4.go\
point.go\
size.go\
rectangle.go\
gradient.go

include $(GOROOT)/src/Make.pkg