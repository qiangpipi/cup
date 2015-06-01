package tst

import ()

var Testlist []Executer = []Executer{
	&RegisterNormal{},
	&LoginNormal{},
	&RegisterExisting{},
	&LoginWithoutReg{},
}
