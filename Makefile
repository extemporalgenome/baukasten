include $(GOROOT)/src/Make.inc

TARG:=github.com/Agon/baukasten

GOFILES:=engine.go\
drivers.go\
graphicsettings.go\
math.go\
vector.go\
point.go\
size.go\
rectangle.go\
glfwdriver.go\
ogldriver.go\
oglsurface.go

include $(GOROOT)/src/Make.pkg