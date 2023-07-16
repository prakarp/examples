package main

import (
	"fmt"
	"medapp/pkg/outpatient"
)

func main() {
	fmt.Println("starting main program...")

	t := outpatient_pkg.LabAppt { 
		LabTestName: "standard panel blood", 
		PatientName: "first last ssn-id",  
		LabTestStatus: "scheduled" ,
	}

	outpatient_pkg.ScheduleTest(&t)

	t.FinishLabAppt()

	

	fmt.Printf("%+v\n", t)

}




	
