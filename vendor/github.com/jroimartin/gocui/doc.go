// Copyright 2014 The gocui Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package gocui allows to create console user interfaces.

Create a new GUI:

	g := gocui.NewGui()
	if err := g.Init(); err != nil {
		// handle error
	}
	defer g.Close()

	// Set layout and key bindings
	// ...

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		// handle error
	}

Set the layout function:

	g.SetLayout(fcn)

On each iteration of the GUI's main loop, the "layout function" is executed.
These layout functions can be used to set-up and update the application's main
views, being possible to freely switch between them. Also, it is important to
mention that a main loop iteration is executed on each reported event
(key-press, mouse event, window resize, etc).

GUIs are composed by Views, you can think of it as buffers. Views implement the
io.ReadWriter interface, so you can just write to them if you want to modify
their content. The same is valid for reading.

Create and initialize a view with absolute coordinates:

	if v, err := g.SetView("viewname", 2, 2, 22, 7); err != nil {
		if err != gocui.ErrUnknownView {
			// handle error
		}
		fmt.Fprintln(v, "This is a new view")
		// ...
	}

Views can also be created using relative coordinates:

	maxX, maxY := g.Size()
	if v, err := g.SetView("viewname", maxX/2-30, maxY/2, maxX/2+30, maxY/2+2); err != nil {
		// ...
	}

Configure keybindings:

	if err := g.SetKeybinding("viewname", gocui.KeyEnter, gocui.ModNone, fcn); err != nil {
		// handle error
	}

gocui implements full mouse support that can be enabled with:

	g.Mouse = true

Mouse events are handled like any other keybinding:

	if err := g.SetKeybinding("viewname", gocui.MouseLeft, gocui.ModNone, fcn); err != nil {
		// handle error
	}

IMPORTANT: Views can only be created, destroyed or updated in three ways: from
layout funcions, from keybinding callbacks or via *Gui.Execute(). The reason
for this is that it allows gocui to be conccurent-safe. So, if you want to
update your GUI from a goroutine, you must use *Gui.Execute(). For example:

	g.Execute(func(g *gocui.Gui) error {
		v, err := g.View("viewname")
		if err != nil {
			// handle error
		}
		v.Clear()
		fmt.Fprintln(v, "Writing from different goroutines")
		return nil
	})

By default, gocui provides a basic edition mode. This mode can be extended
and customized creating a new Editor and assigning it to *Gui.Editor:

	type Editor interface {
		Edit(v *View, key Key, ch rune, mod Modifier)
	}

DefaultEditor can be taken as example to create your own custom Editor:

	var DefaultEditor Editor = EditorFunc(simpleEditor)

	func simpleEditor(v *View, key Key, ch rune, mod Modifier) {
		switch {
		case ch != 0 && mod == 0:
			v.EditWrite(ch)
		case key == KeySpace:
			v.EditWrite(' ')
		case key == KeyBackspace || key == KeyBackspace2:
			v.EditDelete(true)
		// ...
		}
	}

For more information, see the examples in folder "_examples/".
*/
package gocui