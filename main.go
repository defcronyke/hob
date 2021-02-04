/*Package main is the binary package for this project.

Copyright (c) 2021 Jeremy Carter <jeremy@jeremycarter.ca>

All uses of this project in part or in whole are governed
by the terms of the license contained in the file titled
"LICENSE" that's distributed along with the project, which
can be found in the top-level directory of this project.

If you don't agree to follow those terms or you won't
follow them, you are not allowed to use this project or
anything that's made with parts of it at all. The project
is also	depending on some third-party technologies, and
some of those are governed by their own separate licenses,
so furthermore, whenever legally possible, all license
terms from all of the different technologies apply, with
this project's license terms taking first priority.
*/
package main

import (
	"fmt"
	"os"

	"gitlab.com/defcronyke/libhob"
	"gitlab.com/defcronyke/libhob/src/constant"
	"gitlab.com/defcronyke/libhob/src/root"
)

func main() {
	name := fmt.Sprintf("%v %v", constant.HOB_NAME, constant.HOB_VERSION)
	fmt.Printf("%v starting.\n", name)

	app := root.NewHobRoot(name)
	fmt.Printf("%v\n", app.String())

	res, err, code := libhob.Main(&app)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Exited with error: %v\n%v\n", err, *code)
		os.Exit(*code)
		return
	}

	const CODE = 0
	fmt.Printf("Successful exit: %v\n%v\n", *res, CODE)
	os.Exit(CODE)
}
