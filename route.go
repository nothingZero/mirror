package main

import "mirror/framework"

func registerRouter(core *framework.Core) {
	core.Get("foo", FooControllerHandler)
}
