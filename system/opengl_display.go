package system

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

type OpenGLDisplay struct {}

func (d OpenGLDisplay) Start(vm *VirtualMachine) {
	err := glfw.Init()

	if err != nil {
		panic(err)
	}

	window, err := glfw.CreateWindow(800, 600, "Chip8", nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()
	window.Focus()

	err = gl.Init()
	if err != nil {
		panic(err)
	}

	gl.Ortho(0, 64, 32, 0, 0, 1)
	gl.ClearColor(0.0, 0.0, 0.0, 1.0)

	for !window.ShouldClose() {
		d.UpdateKeys(window, vm)
		d.Render(vm)
		window.SwapBuffers()
		glfw.PollEvents()
	}

	glfw.Terminate()
}

func (d OpenGLDisplay) Render(vm *VirtualMachine) {
	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.Color3f(0.0, 1.0, 0.0)
	gl.Begin(gl.QUADS)
	for col := 0; col < 64; col++ {
		for row := range vm.Pixels {
			if vm.PixelSetAt(col, row) {
				c := int32(col)
				r := int32(row)
				gl.Vertex2i(c, r)
				gl.Vertex2i(c+1, r)
				gl.Vertex2i(c+1, r+1)
				gl.Vertex2i(c, r+1)
			}
		}
	}
	gl.End()
}

func (d OpenGLDisplay) UpdateKeys(window *glfw.Window, vm *VirtualMachine) {
	vm.Keyboard[0x1] = window.GetKey(glfw.KeyW) == glfw.Press
	vm.Keyboard[0x4] = window.GetKey(glfw.KeyS) == glfw.Press
	vm.Keyboard[0x6] = window.GetKey(glfw.KeyD) == glfw.Press
	vm.Keyboard[0xC] = window.GetKey(glfw.KeyUp) == glfw.Press
	vm.Keyboard[0xD] = window.GetKey(glfw.KeyDown) == glfw.Press
}
