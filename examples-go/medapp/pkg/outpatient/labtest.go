package outpatient_pkg

import "fmt"

type LabAppt struct {
	LabTestName	string
	PatientName	string
	LabTestStatus	string
}

func (labappt *LabAppt) FinishLabAppt () {
	labappt.LabTestStatus = "done"
}

func ScheduleTest(labappt *LabAppt) {
	fmt.Printf("Scheduling lab test %s for patient %s\n", 
		labappt.LabTestName,
		labappt.PatientName)
}


