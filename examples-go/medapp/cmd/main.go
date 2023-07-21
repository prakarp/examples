package main

// Golang tutorial by an example, for programmers proficient in 
// C or Python or Java or a related procedural programming
// language

// import a few directories, each directory contains a package.
// the package name and the directory name may be different, but
// by convention they are usually the same.

// all directories are relative to 
// $GOROOT/src or $GOPATH/src for each directory under $GOPATH

import (
	"fmt"
	"github.com/prakarp/examples/examples-go/medapp/pkg/outpatient"
)

// always use a main function and keep the core activities in 
// other libraries, in a golang recommended directory structure.

func main() {
	fmt.Println("starting main program...")

	// On purpose, this example uses outpatient_pkg as the package name
	// in a directory named outpatient imported above.

	t := outpatient_pkg.LabAppt { 
		LabTestName: "standard panel blood", 
		PatientName: "first last ssn-id",  
		LabTestStatus: "scheduled" ,
	}

	outpatient_pkg.ScheduleTest(&t)

	t.FinishLabAppt()

	// useful to dump the entire structure in a single line, it is not
	// a pretty-printing method, i.e. there are no indentations.
	fmt.Printf("%+v\n", t)

}


