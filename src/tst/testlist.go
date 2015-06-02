package tst

import ()

var Testlist []Executer = []Executer{
	&RegisterNormal{},
	&RegisterExisting{},
	&RegisterNulUsername{},
	&LoginNormal{},
	&LoginWithoutReg{},
}
