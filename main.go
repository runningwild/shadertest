package main

import (
	"fmt"
	gl "github.com/chsc/gogl/gl21"
	"github.com/runningwild/glop/gin"
	"github.com/runningwild/glop/gos"
	"github.com/runningwild/glop/gui"
	"github.com/runningwild/glop/render"
	"github.com/runningwild/glop/system"
	"github.com/runningwild/shadertest/base"
	// "github.com/runningwild/shadertest/texture"
	// "math"
	"os"
	"path/filepath"
	"runtime"
	// "runtime/pprof"
)

var (
	sys          system.System
	datadir      string
	ui           *gui.Gui
	wdx, wdy     int
	device_index gin.DeviceIndex
	// key_map  base.KeyMap
)

func init() {
	runtime.LockOSThread()
	sys = system.Make(gos.GetSystemInterface())

	datadir = filepath.Join(os.Args[0], "..", "..")
	base.SetDatadir(datadir)
	base.Log().Printf("Setting datadir: %s", datadir)
	wdx = 1024
	wdy = 768

	// var key_binds base.KeyBinds
	// base.LoadJson(filepath.Join(datadir, "key_binds.json"), &key_binds)
	// key_map = key_binds.MakeKeyMap()
	// base.SetDefaultKeyMap(key_map)
}

const dy = 100

type Listener struct {
	s string
}

func (l *Listener) Think(t int64) {
}
func (l *Listener) HandleEventGroup(group gin.EventGroup) {
	k := group.Events[0].Key
	if k.Id().Device.Type == gin.DeviceTypeController && k.Id().Index == gin.ControllerAxis0Positive+1 {
		device_index = k.Id().Device.Index
	}
	l.s = fmt.Sprintf("%v: %v", k, k.CurPressAmt())
}
func (l *Listener) GetLast() string {
	return l.s
}

func main() {
	{
		f, err := os.Create("/Users/jwills/code/src/github.com/runningwild/shadertest/log.err")
		if err != nil {
			panic("shoot")
		}
		os.Stderr = f
		f, err = os.Create("/Users/jwills/code/src/github.com/runningwild/shadertest/log.out")
		if err != nil {
			panic("shoot")
		}
		os.Stdout = f
	}
	sys.Startup()
	err := gl.Init()
	if err != nil {
		panic(err)
	}
	fmt.Printf("RAWR!!!\n")
	render.Init()
	render.Queue(func() {
		sys.CreateWindow(10, 10, wdx, wdy)
		sys.EnableVSync(true)
		err := gl.Init()
		if err != nil {
			panic(err)
		}
	})
	base.InitShaders()
	runtime.GOMAXPROCS(2)
	ui, err = gui.Make(gin.In(), gui.Dims{wdx, wdy}, filepath.Join(datadir, "fonts", "skia.ttf"))
	if err != nil {
		panic(err)
	}

	anchor := gui.MakeAnchorBox(gui.Dims{wdx, wdy})
	ui.AddChild(anchor)
	var v float64
	// var profile_output *os.File
	// var num_mem_profiles int
	// ui.AddChild(base.MakeConsole())
	size := 19.0
	base.InitShaders()
	x := gl.Double(0.0)
	// y := 0.0
	// tex := texture.LoadFromPath(filepath.Join(base.GetDataDir(), "test/out.dff.small.png"))
	fmt.Printf("RAWR!\n")
	listener := Listener{}
	gin.In().RegisterEventListener(&listener)
	button := gin.In().GetKeyFlat(gin.ControllerButton0+6, gin.DeviceTypeController, gin.DeviceIndexAny)
	fmt.Printf("RAWR!\n")
	for button.FramePressCount() == 0 {
		sys.Think()
		// dsize := gin.In().GetKey(gin.MouseWheelVertical).FramePressAmt()
		// size += dsize
		// x -= float64(tex.Dx()) * dsize / 2
		// y -= float64(tex.Dy()) * dsize / 2
		// if gin.In().GetKey(gin.Down).FramePressAmt() > 0 {
		// 	y += 10
		// }
		// if gin.In().GetKey(gin.Up).FramePressAmt() > 0 {
		// 	y -= 10
		// }
		// if gin.In().GetKey(gin.Left).FramePressAmt() > 0 {
		// 	x += 10
		// }
		// if gin.In().GetKey(gin.Right).FramePressAmt() > 0 {
		// 	x -= 10
		// }
		render.Queue(func() {
			ui.Draw()
			gl.Enable(gl.BLEND)
			gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
			gl.Disable(gl.TEXTURE_2D)
			gl.Color4ub(255, 0, 0, 255)
			gl.Begin(gl.QUADS)
			gl.Vertex2d(100+x, 20)
			gl.Vertex2d(100+x, gl.Double(size+20))
			gl.Vertex2d(200+x, gl.Double(size+20))
			gl.Vertex2d(200+x, 20)
			x += 1
			gl.End()
			gl.Enable(gl.TEXTURE_2D)
			gl.Color4ub(255, 255, 255, 255)
			// // str := "!@#$%^&*"
			// diff := 5.0 / (math.Log(size) + math.Pow(size, 0.7))

			// // Works for 1200
			// diff = 50 * math.Pow(base.GetDictionary("skia").Scale(), 2) / math.Pow(size, 1.0)

			// // Works for 3000
			// diff = 50 * math.Pow(base.GetDictionary("skia").Scale(), 1.5) / math.Pow(size, 0.8)
			// //0.340637
			// //0.159241
			// diff = 75 * math.Pow(base.GetDictionary("skia").Scale(), 1.0) / math.Pow(size, 1.0)
			// diff = 10 / math.Pow(size, 1.0)
			// diff = 20/math.Pow(size, 1.0) + 5*math.Pow(base.GetDictionary("skia").Scale(), 1.0)/math.Pow(size, 1.0)
			// if diff > 0.45 {
			//   diff = 0.45
			// }
			// base.EnableShader("distance_field")
			// base.SetUniformF("distance_field", "dist_min", float32(0.5-diff))
			// base.SetUniformF("distance_field", "dist_max", float32(0.5+diff))
			// base.GetDictionary("skia").RenderString(str, 100, 20, 0, dy, gui.Left)
			// base.GetDictionary("skia").RenderString(str, 100, 20+2*dy, 0, dy/4, gui.Left)
			// base.GetDictionary("skia").RenderString(str, 100, 20, 0, size, gui.Left)
			lorem := "Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum"
			kk := gin.In().GetKeyFlat(gin.ControllerAxis0Positive+1, gin.DeviceTypeController, device_index)
			kl := gin.In().GetKeyFlat(gin.ControllerAxis0Positive+1, gin.DeviceTypeController, gin.DeviceIndexAny)
			s := fmt.Sprintf("%1.2f %1.2f - %1.2f %1.2f", kk.FramePressAvg(), kk.FramePressAmt(), kl.FramePressAvg(), kl.FramePressAmt())
			devices := sys.GetActiveDevices()
			y := 500.0
			for _, t := range []gin.DeviceType{gin.DeviceTypeController, gin.DeviceTypeKeyboard, gin.DeviceTypeMouse} {
				for _, d := range devices[t] {
					var s string
					switch t {
					case gin.DeviceTypeController:
						s = "controller"
					case gin.DeviceTypeKeyboard:
						s = "keyboard"
					case gin.DeviceTypeMouse:
						s = "mouse"
					}
					base.GetDictionary("skia").RenderString(fmt.Sprintf("%s: %d", s, d), 100, y, 0, 45, gui.Left)
					y -= 50
				}
			}
			base.GetDictionary("luxisr").RenderString(s, 50, 50, 0, size, gui.Left)
			// base.GetDictionary("luxisr").RenderString(lorem, 50, 50+size, 0, size, gui.Left)
			base.GetDictionary("skia").RenderString(lorem, 50, 50+2*size, 0, size, gui.Left)
			base.Log().Printf("Foo")
			//      base.EnableShader("")
			// gl.Enable(gl.ALPHA_TEST)
			// gl.AlphaFunc(gl.GREATER, 0.5)
			// tex := texture.LoadFromPath(filepath.Join(base.GetDataDir(), "ships/ship.png"))
			// tex.Bind()
			// tex.RenderAdvanced(x, y, float64(tex.Dx())*size, float64(tex.Dy())*size, 0, true)
			// tex.RenderNatural(300, 100)
			// gl.Disable(gl.ALPHA_TEST)
		})
		render.Queue(func() {
			sys.SwapBuffers()
		})
		render.Purge()

		// if key_map["cpu profile"].FramePressCount() > 0 {
		// 	if profile_output == nil {
		// 		profile_output, err = os.Create(filepath.Join(datadir, "cpu.prof"))
		// 		if err == nil {
		// 			err = pprof.StartCPUProfile(profile_output)
		// 			if err != nil {
		// 				fmt.Printf("Unable to start CPU profile: %v\n", err)
		// 				profile_output.Close()
		// 				profile_output = nil
		// 			}
		// 			fmt.Printf("profout: %v\n", profile_output)
		// 		} else {
		// 			fmt.Printf("Unable to start CPU profile: %v\n", err)
		// 		}
		// 	} else {
		// 		pprof.StopCPUProfile()
		// 		profile_output.Close()
		// 		profile_output = nil
		// 	}
		// }

		// if key_map["mem profile"].FramePressCount() > 0 {
		// 	f, err := os.Create(filepath.Join(datadir, fmt.Sprintf("mem.%d.prof", num_mem_profiles)))
		// 	if err != nil {
		// 		base.Error().Printf("Unable to write mem profile: %v", err)
		// 	}
		// 	pprof.WriteHeapProfile(f)
		// 	f.Close()
		// 	num_mem_profiles++
		// }

		v += 0.01
	}
}
