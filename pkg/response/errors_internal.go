package response

import "fmt"

type InternalError struct {
	Type         string
	Name         string
	FunctionName string
	Description  string
	Trace        interface{}
}

func (in InternalError) Error() error {
	return fmt.Errorf("[%s][%s][%s] %s, trace %v", in.Type, in.Name, in.FunctionName, in.Description, in.Trace)
}

func (in InternalError) String() string {
	return fmt.Sprintf("[%s][%s][%s] %s, trace %v", in.Type, in.Name, in.FunctionName, in.Description, in.Trace)
}

type InternalTestError struct {
	Name         string
	FunctionName string
	Description  string
	Trace        interface{}
}

func (in InternalTestError) Error() error {
	return fmt.Errorf("[Unit Test Error][%s][%s] %s, trace %v", in.Name, in.FunctionName, in.Description, in.Trace)
}
