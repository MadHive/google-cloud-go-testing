// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"context"
	"github.com/MadHive/google-cloud-go-testing/bigquery/bqiface"
	"sync"
)

// Ensure, that ClientMock does implement bqiface.Client.
// If this is not the case, regenerate this file with moq.
var _ bqiface.Client = &ClientMock{}

// ClientMock is a mock implementation of bqiface.Client.
//
//     func TestSomethingThatUsesClient(t *testing.T) {
//
//         // make and configure a mocked bqiface.Client
//         mockedClient := &ClientMock{
//             CloseFunc: func() error {
// 	               panic("mock out the Close method")
//             },
//             DatasetFunc: func(in1 string) bqiface.Dataset {
// 	               panic("mock out the Dataset method")
//             },
//             DatasetInProjectFunc: func(in1 string, in2 string) bqiface.Dataset {
// 	               panic("mock out the DatasetInProject method")
//             },
//             DatasetsFunc: func(in1 context.Context) bqiface.DatasetIterator {
// 	               panic("mock out the Datasets method")
//             },
//             DatasetsInProjectFunc: func(in1 context.Context, in2 string) bqiface.DatasetIterator {
// 	               panic("mock out the DatasetsInProject method")
//             },
//             JobFromIDFunc: func(in1 context.Context, in2 string) (bqiface.Job, error) {
// 	               panic("mock out the JobFromID method")
//             },
//             JobFromIDLocationFunc: func(in1 context.Context, in2 string, in3 string) (bqiface.Job, error) {
// 	               panic("mock out the JobFromIDLocation method")
//             },
//             JobsFunc: func(in1 context.Context) bqiface.JobIterator {
// 	               panic("mock out the Jobs method")
//             },
//             LocationFunc: func() string {
// 	               panic("mock out the Location method")
//             },
//             QueryFunc: func(in1 string) bqiface.Query {
// 	               panic("mock out the Query method")
//             },
//             SetLocationFunc: func(in1 string)  {
// 	               panic("mock out the SetLocation method")
//             },
//         }
//
//         // use mockedClient in code that requires bqiface.Client
//         // and then make assertions.
//
//     }
type ClientMock struct {
	// CloseFunc mocks the Close method.
	CloseFunc func() error

	// DatasetFunc mocks the Dataset method.
	DatasetFunc func(in1 string) bqiface.Dataset

	// DatasetInProjectFunc mocks the DatasetInProject method.
	DatasetInProjectFunc func(in1 string, in2 string) bqiface.Dataset

	// DatasetsFunc mocks the Datasets method.
	DatasetsFunc func(in1 context.Context) bqiface.DatasetIterator

	// DatasetsInProjectFunc mocks the DatasetsInProject method.
	DatasetsInProjectFunc func(in1 context.Context, in2 string) bqiface.DatasetIterator

	// JobFromIDFunc mocks the JobFromID method.
	JobFromIDFunc func(in1 context.Context, in2 string) (bqiface.Job, error)

	// JobFromIDLocationFunc mocks the JobFromIDLocation method.
	JobFromIDLocationFunc func(in1 context.Context, in2 string, in3 string) (bqiface.Job, error)

	// JobsFunc mocks the Jobs method.
	JobsFunc func(in1 context.Context) bqiface.JobIterator

	// LocationFunc mocks the Location method.
	LocationFunc func() string

	// QueryFunc mocks the Query method.
	QueryFunc func(in1 string) bqiface.Query

	// SetLocationFunc mocks the SetLocation method.
	SetLocationFunc func(in1 string)

	// calls tracks calls to the methods.
	calls struct {
		// Close holds details about calls to the Close method.
		Close []struct {
		}
		// Dataset holds details about calls to the Dataset method.
		Dataset []struct {
			// In1 is the in1 argument value.
			In1 string
		}
		// DatasetInProject holds details about calls to the DatasetInProject method.
		DatasetInProject []struct {
			// In1 is the in1 argument value.
			In1 string
			// In2 is the in2 argument value.
			In2 string
		}
		// Datasets holds details about calls to the Datasets method.
		Datasets []struct {
			// In1 is the in1 argument value.
			In1 context.Context
		}
		// DatasetsInProject holds details about calls to the DatasetsInProject method.
		DatasetsInProject []struct {
			// In1 is the in1 argument value.
			In1 context.Context
			// In2 is the in2 argument value.
			In2 string
		}
		// JobFromID holds details about calls to the JobFromID method.
		JobFromID []struct {
			// In1 is the in1 argument value.
			In1 context.Context
			// In2 is the in2 argument value.
			In2 string
		}
		// JobFromIDLocation holds details about calls to the JobFromIDLocation method.
		JobFromIDLocation []struct {
			// In1 is the in1 argument value.
			In1 context.Context
			// In2 is the in2 argument value.
			In2 string
			// In3 is the in3 argument value.
			In3 string
		}
		// Jobs holds details about calls to the Jobs method.
		Jobs []struct {
			// In1 is the in1 argument value.
			In1 context.Context
		}
		// Location holds details about calls to the Location method.
		Location []struct {
		}
		// Query holds details about calls to the Query method.
		Query []struct {
			// In1 is the in1 argument value.
			In1 string
		}
		// SetLocation holds details about calls to the SetLocation method.
		SetLocation []struct {
			// In1 is the in1 argument value.
			In1 string
		}
	}
	lockClose             sync.RWMutex
	lockDataset           sync.RWMutex
	lockDatasetInProject  sync.RWMutex
	lockDatasets          sync.RWMutex
	lockDatasetsInProject sync.RWMutex
	lockJobFromID         sync.RWMutex
	lockJobFromIDLocation sync.RWMutex
	lockJobs              sync.RWMutex
	lockLocation          sync.RWMutex
	lockQuery             sync.RWMutex
	lockSetLocation       sync.RWMutex
}

// Close calls CloseFunc.
func (mock *ClientMock) Close() error {
	if mock.CloseFunc == nil {
		panic("ClientMock.CloseFunc: method is nil but Client.Close was just called")
	}
	callInfo := struct {
	}{}
	mock.lockClose.Lock()
	mock.calls.Close = append(mock.calls.Close, callInfo)
	mock.lockClose.Unlock()
	return mock.CloseFunc()
}

// CloseCalls gets all the calls that were made to Close.
// Check the length with:
//     len(mockedClient.CloseCalls())
func (mock *ClientMock) CloseCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockClose.RLock()
	calls = mock.calls.Close
	mock.lockClose.RUnlock()
	return calls
}

// Dataset calls DatasetFunc.
func (mock *ClientMock) Dataset(in1 string) bqiface.Dataset {
	if mock.DatasetFunc == nil {
		panic("ClientMock.DatasetFunc: method is nil but Client.Dataset was just called")
	}
	callInfo := struct {
		In1 string
	}{
		In1: in1,
	}
	mock.lockDataset.Lock()
	mock.calls.Dataset = append(mock.calls.Dataset, callInfo)
	mock.lockDataset.Unlock()
	return mock.DatasetFunc(in1)
}

// DatasetCalls gets all the calls that were made to Dataset.
// Check the length with:
//     len(mockedClient.DatasetCalls())
func (mock *ClientMock) DatasetCalls() []struct {
	In1 string
} {
	var calls []struct {
		In1 string
	}
	mock.lockDataset.RLock()
	calls = mock.calls.Dataset
	mock.lockDataset.RUnlock()
	return calls
}

// DatasetInProject calls DatasetInProjectFunc.
func (mock *ClientMock) DatasetInProject(in1 string, in2 string) bqiface.Dataset {
	if mock.DatasetInProjectFunc == nil {
		panic("ClientMock.DatasetInProjectFunc: method is nil but Client.DatasetInProject was just called")
	}
	callInfo := struct {
		In1 string
		In2 string
	}{
		In1: in1,
		In2: in2,
	}
	mock.lockDatasetInProject.Lock()
	mock.calls.DatasetInProject = append(mock.calls.DatasetInProject, callInfo)
	mock.lockDatasetInProject.Unlock()
	return mock.DatasetInProjectFunc(in1, in2)
}

// DatasetInProjectCalls gets all the calls that were made to DatasetInProject.
// Check the length with:
//     len(mockedClient.DatasetInProjectCalls())
func (mock *ClientMock) DatasetInProjectCalls() []struct {
	In1 string
	In2 string
} {
	var calls []struct {
		In1 string
		In2 string
	}
	mock.lockDatasetInProject.RLock()
	calls = mock.calls.DatasetInProject
	mock.lockDatasetInProject.RUnlock()
	return calls
}

// Datasets calls DatasetsFunc.
func (mock *ClientMock) Datasets(in1 context.Context) bqiface.DatasetIterator {
	if mock.DatasetsFunc == nil {
		panic("ClientMock.DatasetsFunc: method is nil but Client.Datasets was just called")
	}
	callInfo := struct {
		In1 context.Context
	}{
		In1: in1,
	}
	mock.lockDatasets.Lock()
	mock.calls.Datasets = append(mock.calls.Datasets, callInfo)
	mock.lockDatasets.Unlock()
	return mock.DatasetsFunc(in1)
}

// DatasetsCalls gets all the calls that were made to Datasets.
// Check the length with:
//     len(mockedClient.DatasetsCalls())
func (mock *ClientMock) DatasetsCalls() []struct {
	In1 context.Context
} {
	var calls []struct {
		In1 context.Context
	}
	mock.lockDatasets.RLock()
	calls = mock.calls.Datasets
	mock.lockDatasets.RUnlock()
	return calls
}

// DatasetsInProject calls DatasetsInProjectFunc.
func (mock *ClientMock) DatasetsInProject(in1 context.Context, in2 string) bqiface.DatasetIterator {
	if mock.DatasetsInProjectFunc == nil {
		panic("ClientMock.DatasetsInProjectFunc: method is nil but Client.DatasetsInProject was just called")
	}
	callInfo := struct {
		In1 context.Context
		In2 string
	}{
		In1: in1,
		In2: in2,
	}
	mock.lockDatasetsInProject.Lock()
	mock.calls.DatasetsInProject = append(mock.calls.DatasetsInProject, callInfo)
	mock.lockDatasetsInProject.Unlock()
	return mock.DatasetsInProjectFunc(in1, in2)
}

// DatasetsInProjectCalls gets all the calls that were made to DatasetsInProject.
// Check the length with:
//     len(mockedClient.DatasetsInProjectCalls())
func (mock *ClientMock) DatasetsInProjectCalls() []struct {
	In1 context.Context
	In2 string
} {
	var calls []struct {
		In1 context.Context
		In2 string
	}
	mock.lockDatasetsInProject.RLock()
	calls = mock.calls.DatasetsInProject
	mock.lockDatasetsInProject.RUnlock()
	return calls
}

// JobFromID calls JobFromIDFunc.
func (mock *ClientMock) JobFromID(in1 context.Context, in2 string) (bqiface.Job, error) {
	if mock.JobFromIDFunc == nil {
		panic("ClientMock.JobFromIDFunc: method is nil but Client.JobFromID was just called")
	}
	callInfo := struct {
		In1 context.Context
		In2 string
	}{
		In1: in1,
		In2: in2,
	}
	mock.lockJobFromID.Lock()
	mock.calls.JobFromID = append(mock.calls.JobFromID, callInfo)
	mock.lockJobFromID.Unlock()
	return mock.JobFromIDFunc(in1, in2)
}

// JobFromIDCalls gets all the calls that were made to JobFromID.
// Check the length with:
//     len(mockedClient.JobFromIDCalls())
func (mock *ClientMock) JobFromIDCalls() []struct {
	In1 context.Context
	In2 string
} {
	var calls []struct {
		In1 context.Context
		In2 string
	}
	mock.lockJobFromID.RLock()
	calls = mock.calls.JobFromID
	mock.lockJobFromID.RUnlock()
	return calls
}

// JobFromIDLocation calls JobFromIDLocationFunc.
func (mock *ClientMock) JobFromIDLocation(in1 context.Context, in2 string, in3 string) (bqiface.Job, error) {
	if mock.JobFromIDLocationFunc == nil {
		panic("ClientMock.JobFromIDLocationFunc: method is nil but Client.JobFromIDLocation was just called")
	}
	callInfo := struct {
		In1 context.Context
		In2 string
		In3 string
	}{
		In1: in1,
		In2: in2,
		In3: in3,
	}
	mock.lockJobFromIDLocation.Lock()
	mock.calls.JobFromIDLocation = append(mock.calls.JobFromIDLocation, callInfo)
	mock.lockJobFromIDLocation.Unlock()
	return mock.JobFromIDLocationFunc(in1, in2, in3)
}

// JobFromIDLocationCalls gets all the calls that were made to JobFromIDLocation.
// Check the length with:
//     len(mockedClient.JobFromIDLocationCalls())
func (mock *ClientMock) JobFromIDLocationCalls() []struct {
	In1 context.Context
	In2 string
	In3 string
} {
	var calls []struct {
		In1 context.Context
		In2 string
		In3 string
	}
	mock.lockJobFromIDLocation.RLock()
	calls = mock.calls.JobFromIDLocation
	mock.lockJobFromIDLocation.RUnlock()
	return calls
}

// Jobs calls JobsFunc.
func (mock *ClientMock) Jobs(in1 context.Context) bqiface.JobIterator {
	if mock.JobsFunc == nil {
		panic("ClientMock.JobsFunc: method is nil but Client.Jobs was just called")
	}
	callInfo := struct {
		In1 context.Context
	}{
		In1: in1,
	}
	mock.lockJobs.Lock()
	mock.calls.Jobs = append(mock.calls.Jobs, callInfo)
	mock.lockJobs.Unlock()
	return mock.JobsFunc(in1)
}

// JobsCalls gets all the calls that were made to Jobs.
// Check the length with:
//     len(mockedClient.JobsCalls())
func (mock *ClientMock) JobsCalls() []struct {
	In1 context.Context
} {
	var calls []struct {
		In1 context.Context
	}
	mock.lockJobs.RLock()
	calls = mock.calls.Jobs
	mock.lockJobs.RUnlock()
	return calls
}

// Location calls LocationFunc.
func (mock *ClientMock) Location() string {
	if mock.LocationFunc == nil {
		panic("ClientMock.LocationFunc: method is nil but Client.Location was just called")
	}
	callInfo := struct {
	}{}
	mock.lockLocation.Lock()
	mock.calls.Location = append(mock.calls.Location, callInfo)
	mock.lockLocation.Unlock()
	return mock.LocationFunc()
}

// LocationCalls gets all the calls that were made to Location.
// Check the length with:
//     len(mockedClient.LocationCalls())
func (mock *ClientMock) LocationCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockLocation.RLock()
	calls = mock.calls.Location
	mock.lockLocation.RUnlock()
	return calls
}

// Query calls QueryFunc.
func (mock *ClientMock) Query(in1 string) bqiface.Query {
	if mock.QueryFunc == nil {
		panic("ClientMock.QueryFunc: method is nil but Client.Query was just called")
	}
	callInfo := struct {
		In1 string
	}{
		In1: in1,
	}
	mock.lockQuery.Lock()
	mock.calls.Query = append(mock.calls.Query, callInfo)
	mock.lockQuery.Unlock()
	return mock.QueryFunc(in1)
}

// QueryCalls gets all the calls that were made to Query.
// Check the length with:
//     len(mockedClient.QueryCalls())
func (mock *ClientMock) QueryCalls() []struct {
	In1 string
} {
	var calls []struct {
		In1 string
	}
	mock.lockQuery.RLock()
	calls = mock.calls.Query
	mock.lockQuery.RUnlock()
	return calls
}

// SetLocation calls SetLocationFunc.
func (mock *ClientMock) SetLocation(in1 string) {
	if mock.SetLocationFunc == nil {
		panic("ClientMock.SetLocationFunc: method is nil but Client.SetLocation was just called")
	}
	callInfo := struct {
		In1 string
	}{
		In1: in1,
	}
	mock.lockSetLocation.Lock()
	mock.calls.SetLocation = append(mock.calls.SetLocation, callInfo)
	mock.lockSetLocation.Unlock()
	mock.SetLocationFunc(in1)
}

// SetLocationCalls gets all the calls that were made to SetLocation.
// Check the length with:
//     len(mockedClient.SetLocationCalls())
func (mock *ClientMock) SetLocationCalls() []struct {
	In1 string
} {
	var calls []struct {
		In1 string
	}
	mock.lockSetLocation.RLock()
	calls = mock.calls.SetLocation
	mock.lockSetLocation.RUnlock()
	return calls
}
