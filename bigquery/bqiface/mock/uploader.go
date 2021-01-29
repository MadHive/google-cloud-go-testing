// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"context"
	"github.com/MadHive/google-cloud-go-testing/bigquery/bqiface"
	"sync"
)

// Ensure, that UploaderMock does implement bqiface.Uploader.
// If this is not the case, regenerate this file with moq.
var _ bqiface.Uploader = &UploaderMock{}

// UploaderMock is a mock implementation of bqiface.Uploader.
//
//     func TestSomethingThatUsesUploader(t *testing.T) {
//
//         // make and configure a mocked bqiface.Uploader
//         mockedUploader := &UploaderMock{
//             PutFunc: func(in1 context.Context, in2 interface{}) error {
// 	               panic("mock out the Put method")
//             },
//             SetIgnoreUnknownValuesFunc: func(in1 bool)  {
// 	               panic("mock out the SetIgnoreUnknownValues method")
//             },
//             SetSkipInvalidRowsFunc: func(in1 bool)  {
// 	               panic("mock out the SetSkipInvalidRows method")
//             },
//             SetTableTemplateSuffixFunc: func(in1 string)  {
// 	               panic("mock out the SetTableTemplateSuffix method")
//             },
//         }
//
//         // use mockedUploader in code that requires bqiface.Uploader
//         // and then make assertions.
//
//     }
type UploaderMock struct {
	// PutFunc mocks the Put method.
	PutFunc func(in1 context.Context, in2 interface{}) error

	// SetIgnoreUnknownValuesFunc mocks the SetIgnoreUnknownValues method.
	SetIgnoreUnknownValuesFunc func(in1 bool)

	// SetSkipInvalidRowsFunc mocks the SetSkipInvalidRows method.
	SetSkipInvalidRowsFunc func(in1 bool)

	// SetTableTemplateSuffixFunc mocks the SetTableTemplateSuffix method.
	SetTableTemplateSuffixFunc func(in1 string)

	// calls tracks calls to the methods.
	calls struct {
		// Put holds details about calls to the Put method.
		Put []struct {
			// In1 is the in1 argument value.
			In1 context.Context
			// In2 is the in2 argument value.
			In2 interface{}
		}
		// SetIgnoreUnknownValues holds details about calls to the SetIgnoreUnknownValues method.
		SetIgnoreUnknownValues []struct {
			// In1 is the in1 argument value.
			In1 bool
		}
		// SetSkipInvalidRows holds details about calls to the SetSkipInvalidRows method.
		SetSkipInvalidRows []struct {
			// In1 is the in1 argument value.
			In1 bool
		}
		// SetTableTemplateSuffix holds details about calls to the SetTableTemplateSuffix method.
		SetTableTemplateSuffix []struct {
			// In1 is the in1 argument value.
			In1 string
		}
	}
	lockPut                    sync.RWMutex
	lockSetIgnoreUnknownValues sync.RWMutex
	lockSetSkipInvalidRows     sync.RWMutex
	lockSetTableTemplateSuffix sync.RWMutex
}

// Put calls PutFunc.
func (mock *UploaderMock) Put(in1 context.Context, in2 interface{}) error {
	if mock.PutFunc == nil {
		panic("UploaderMock.PutFunc: method is nil but Uploader.Put was just called")
	}
	callInfo := struct {
		In1 context.Context
		In2 interface{}
	}{
		In1: in1,
		In2: in2,
	}
	mock.lockPut.Lock()
	mock.calls.Put = append(mock.calls.Put, callInfo)
	mock.lockPut.Unlock()
	return mock.PutFunc(in1, in2)
}

// PutCalls gets all the calls that were made to Put.
// Check the length with:
//     len(mockedUploader.PutCalls())
func (mock *UploaderMock) PutCalls() []struct {
	In1 context.Context
	In2 interface{}
} {
	var calls []struct {
		In1 context.Context
		In2 interface{}
	}
	mock.lockPut.RLock()
	calls = mock.calls.Put
	mock.lockPut.RUnlock()
	return calls
}

// SetIgnoreUnknownValues calls SetIgnoreUnknownValuesFunc.
func (mock *UploaderMock) SetIgnoreUnknownValues(in1 bool) {
	if mock.SetIgnoreUnknownValuesFunc == nil {
		panic("UploaderMock.SetIgnoreUnknownValuesFunc: method is nil but Uploader.SetIgnoreUnknownValues was just called")
	}
	callInfo := struct {
		In1 bool
	}{
		In1: in1,
	}
	mock.lockSetIgnoreUnknownValues.Lock()
	mock.calls.SetIgnoreUnknownValues = append(mock.calls.SetIgnoreUnknownValues, callInfo)
	mock.lockSetIgnoreUnknownValues.Unlock()
	mock.SetIgnoreUnknownValuesFunc(in1)
}

// SetIgnoreUnknownValuesCalls gets all the calls that were made to SetIgnoreUnknownValues.
// Check the length with:
//     len(mockedUploader.SetIgnoreUnknownValuesCalls())
func (mock *UploaderMock) SetIgnoreUnknownValuesCalls() []struct {
	In1 bool
} {
	var calls []struct {
		In1 bool
	}
	mock.lockSetIgnoreUnknownValues.RLock()
	calls = mock.calls.SetIgnoreUnknownValues
	mock.lockSetIgnoreUnknownValues.RUnlock()
	return calls
}

// SetSkipInvalidRows calls SetSkipInvalidRowsFunc.
func (mock *UploaderMock) SetSkipInvalidRows(in1 bool) {
	if mock.SetSkipInvalidRowsFunc == nil {
		panic("UploaderMock.SetSkipInvalidRowsFunc: method is nil but Uploader.SetSkipInvalidRows was just called")
	}
	callInfo := struct {
		In1 bool
	}{
		In1: in1,
	}
	mock.lockSetSkipInvalidRows.Lock()
	mock.calls.SetSkipInvalidRows = append(mock.calls.SetSkipInvalidRows, callInfo)
	mock.lockSetSkipInvalidRows.Unlock()
	mock.SetSkipInvalidRowsFunc(in1)
}

// SetSkipInvalidRowsCalls gets all the calls that were made to SetSkipInvalidRows.
// Check the length with:
//     len(mockedUploader.SetSkipInvalidRowsCalls())
func (mock *UploaderMock) SetSkipInvalidRowsCalls() []struct {
	In1 bool
} {
	var calls []struct {
		In1 bool
	}
	mock.lockSetSkipInvalidRows.RLock()
	calls = mock.calls.SetSkipInvalidRows
	mock.lockSetSkipInvalidRows.RUnlock()
	return calls
}

// SetTableTemplateSuffix calls SetTableTemplateSuffixFunc.
func (mock *UploaderMock) SetTableTemplateSuffix(in1 string) {
	if mock.SetTableTemplateSuffixFunc == nil {
		panic("UploaderMock.SetTableTemplateSuffixFunc: method is nil but Uploader.SetTableTemplateSuffix was just called")
	}
	callInfo := struct {
		In1 string
	}{
		In1: in1,
	}
	mock.lockSetTableTemplateSuffix.Lock()
	mock.calls.SetTableTemplateSuffix = append(mock.calls.SetTableTemplateSuffix, callInfo)
	mock.lockSetTableTemplateSuffix.Unlock()
	mock.SetTableTemplateSuffixFunc(in1)
}

// SetTableTemplateSuffixCalls gets all the calls that were made to SetTableTemplateSuffix.
// Check the length with:
//     len(mockedUploader.SetTableTemplateSuffixCalls())
func (mock *UploaderMock) SetTableTemplateSuffixCalls() []struct {
	In1 string
} {
	var calls []struct {
		In1 string
	}
	mock.lockSetTableTemplateSuffix.RLock()
	calls = mock.calls.SetTableTemplateSuffix
	mock.lockSetTableTemplateSuffix.RUnlock()
	return calls
}
