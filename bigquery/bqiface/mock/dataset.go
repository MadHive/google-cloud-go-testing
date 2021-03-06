// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"context"
	"github.com/MadHive/google-cloud-go-testing/bigquery/bqiface"
	"sync"
)

// Ensure, that DatasetMock does implement bqiface.Dataset.
// If this is not the case, regenerate this file with moq.
var _ bqiface.Dataset = &DatasetMock{}

// DatasetMock is a mock implementation of bqiface.Dataset.
//
//     func TestSomethingThatUsesDataset(t *testing.T) {
//
//         // make and configure a mocked bqiface.Dataset
//         mockedDataset := &DatasetMock{
//             CreateFunc: func(in1 context.Context, in2 *bqiface.DatasetMetadata) error {
// 	               panic("mock out the Create method")
//             },
//             DatasetIDFunc: func() string {
// 	               panic("mock out the DatasetID method")
//             },
//             DeleteFunc: func(in1 context.Context) error {
// 	               panic("mock out the Delete method")
//             },
//             DeleteWithContentsFunc: func(in1 context.Context) error {
// 	               panic("mock out the DeleteWithContents method")
//             },
//             MetadataFunc: func(in1 context.Context) (*bqiface.DatasetMetadata, error) {
// 	               panic("mock out the Metadata method")
//             },
//             ProjectIDFunc: func() string {
// 	               panic("mock out the ProjectID method")
//             },
//             TableFunc: func(in1 string) bqiface.Table {
// 	               panic("mock out the Table method")
//             },
//             TablesFunc: func(in1 context.Context) bqiface.TableIterator {
// 	               panic("mock out the Tables method")
//             },
//             UpdateFunc: func(in1 context.Context, in2 bqiface.DatasetMetadataToUpdate, in3 string) (*bqiface.DatasetMetadata, error) {
// 	               panic("mock out the Update method")
//             },
//         }
//
//         // use mockedDataset in code that requires bqiface.Dataset
//         // and then make assertions.
//
//     }
type DatasetMock struct {
	// CreateFunc mocks the Create method.
	CreateFunc func(in1 context.Context, in2 *bqiface.DatasetMetadata) error

	// DatasetIDFunc mocks the DatasetID method.
	DatasetIDFunc func() string

	// DeleteFunc mocks the Delete method.
	DeleteFunc func(in1 context.Context) error

	// DeleteWithContentsFunc mocks the DeleteWithContents method.
	DeleteWithContentsFunc func(in1 context.Context) error

	// MetadataFunc mocks the Metadata method.
	MetadataFunc func(in1 context.Context) (*bqiface.DatasetMetadata, error)

	// ProjectIDFunc mocks the ProjectID method.
	ProjectIDFunc func() string

	// TableFunc mocks the Table method.
	TableFunc func(in1 string) bqiface.Table

	// TablesFunc mocks the Tables method.
	TablesFunc func(in1 context.Context) bqiface.TableIterator

	// UpdateFunc mocks the Update method.
	UpdateFunc func(in1 context.Context, in2 bqiface.DatasetMetadataToUpdate, in3 string) (*bqiface.DatasetMetadata, error)

	// calls tracks calls to the methods.
	calls struct {
		// Create holds details about calls to the Create method.
		Create []struct {
			// In1 is the in1 argument value.
			In1 context.Context
			// In2 is the in2 argument value.
			In2 *bqiface.DatasetMetadata
		}
		// DatasetID holds details about calls to the DatasetID method.
		DatasetID []struct {
		}
		// Delete holds details about calls to the Delete method.
		Delete []struct {
			// In1 is the in1 argument value.
			In1 context.Context
		}
		// DeleteWithContents holds details about calls to the DeleteWithContents method.
		DeleteWithContents []struct {
			// In1 is the in1 argument value.
			In1 context.Context
		}
		// Metadata holds details about calls to the Metadata method.
		Metadata []struct {
			// In1 is the in1 argument value.
			In1 context.Context
		}
		// ProjectID holds details about calls to the ProjectID method.
		ProjectID []struct {
		}
		// Table holds details about calls to the Table method.
		Table []struct {
			// In1 is the in1 argument value.
			In1 string
		}
		// Tables holds details about calls to the Tables method.
		Tables []struct {
			// In1 is the in1 argument value.
			In1 context.Context
		}
		// Update holds details about calls to the Update method.
		Update []struct {
			// In1 is the in1 argument value.
			In1 context.Context
			// In2 is the in2 argument value.
			In2 bqiface.DatasetMetadataToUpdate
			// In3 is the in3 argument value.
			In3 string
		}
	}
	lockCreate             sync.RWMutex
	lockDatasetID          sync.RWMutex
	lockDelete             sync.RWMutex
	lockDeleteWithContents sync.RWMutex
	lockMetadata           sync.RWMutex
	lockProjectID          sync.RWMutex
	lockTable              sync.RWMutex
	lockTables             sync.RWMutex
	lockUpdate             sync.RWMutex
}

// Create calls CreateFunc.
func (mock *DatasetMock) Create(in1 context.Context, in2 *bqiface.DatasetMetadata) error {
	if mock.CreateFunc == nil {
		panic("DatasetMock.CreateFunc: method is nil but Dataset.Create was just called")
	}
	callInfo := struct {
		In1 context.Context
		In2 *bqiface.DatasetMetadata
	}{
		In1: in1,
		In2: in2,
	}
	mock.lockCreate.Lock()
	mock.calls.Create = append(mock.calls.Create, callInfo)
	mock.lockCreate.Unlock()
	return mock.CreateFunc(in1, in2)
}

// CreateCalls gets all the calls that were made to Create.
// Check the length with:
//     len(mockedDataset.CreateCalls())
func (mock *DatasetMock) CreateCalls() []struct {
	In1 context.Context
	In2 *bqiface.DatasetMetadata
} {
	var calls []struct {
		In1 context.Context
		In2 *bqiface.DatasetMetadata
	}
	mock.lockCreate.RLock()
	calls = mock.calls.Create
	mock.lockCreate.RUnlock()
	return calls
}

// DatasetID calls DatasetIDFunc.
func (mock *DatasetMock) DatasetID() string {
	if mock.DatasetIDFunc == nil {
		panic("DatasetMock.DatasetIDFunc: method is nil but Dataset.DatasetID was just called")
	}
	callInfo := struct {
	}{}
	mock.lockDatasetID.Lock()
	mock.calls.DatasetID = append(mock.calls.DatasetID, callInfo)
	mock.lockDatasetID.Unlock()
	return mock.DatasetIDFunc()
}

// DatasetIDCalls gets all the calls that were made to DatasetID.
// Check the length with:
//     len(mockedDataset.DatasetIDCalls())
func (mock *DatasetMock) DatasetIDCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockDatasetID.RLock()
	calls = mock.calls.DatasetID
	mock.lockDatasetID.RUnlock()
	return calls
}

// Delete calls DeleteFunc.
func (mock *DatasetMock) Delete(in1 context.Context) error {
	if mock.DeleteFunc == nil {
		panic("DatasetMock.DeleteFunc: method is nil but Dataset.Delete was just called")
	}
	callInfo := struct {
		In1 context.Context
	}{
		In1: in1,
	}
	mock.lockDelete.Lock()
	mock.calls.Delete = append(mock.calls.Delete, callInfo)
	mock.lockDelete.Unlock()
	return mock.DeleteFunc(in1)
}

// DeleteCalls gets all the calls that were made to Delete.
// Check the length with:
//     len(mockedDataset.DeleteCalls())
func (mock *DatasetMock) DeleteCalls() []struct {
	In1 context.Context
} {
	var calls []struct {
		In1 context.Context
	}
	mock.lockDelete.RLock()
	calls = mock.calls.Delete
	mock.lockDelete.RUnlock()
	return calls
}

// DeleteWithContents calls DeleteWithContentsFunc.
func (mock *DatasetMock) DeleteWithContents(in1 context.Context) error {
	if mock.DeleteWithContentsFunc == nil {
		panic("DatasetMock.DeleteWithContentsFunc: method is nil but Dataset.DeleteWithContents was just called")
	}
	callInfo := struct {
		In1 context.Context
	}{
		In1: in1,
	}
	mock.lockDeleteWithContents.Lock()
	mock.calls.DeleteWithContents = append(mock.calls.DeleteWithContents, callInfo)
	mock.lockDeleteWithContents.Unlock()
	return mock.DeleteWithContentsFunc(in1)
}

// DeleteWithContentsCalls gets all the calls that were made to DeleteWithContents.
// Check the length with:
//     len(mockedDataset.DeleteWithContentsCalls())
func (mock *DatasetMock) DeleteWithContentsCalls() []struct {
	In1 context.Context
} {
	var calls []struct {
		In1 context.Context
	}
	mock.lockDeleteWithContents.RLock()
	calls = mock.calls.DeleteWithContents
	mock.lockDeleteWithContents.RUnlock()
	return calls
}

// Metadata calls MetadataFunc.
func (mock *DatasetMock) Metadata(in1 context.Context) (*bqiface.DatasetMetadata, error) {
	if mock.MetadataFunc == nil {
		panic("DatasetMock.MetadataFunc: method is nil but Dataset.Metadata was just called")
	}
	callInfo := struct {
		In1 context.Context
	}{
		In1: in1,
	}
	mock.lockMetadata.Lock()
	mock.calls.Metadata = append(mock.calls.Metadata, callInfo)
	mock.lockMetadata.Unlock()
	return mock.MetadataFunc(in1)
}

// MetadataCalls gets all the calls that were made to Metadata.
// Check the length with:
//     len(mockedDataset.MetadataCalls())
func (mock *DatasetMock) MetadataCalls() []struct {
	In1 context.Context
} {
	var calls []struct {
		In1 context.Context
	}
	mock.lockMetadata.RLock()
	calls = mock.calls.Metadata
	mock.lockMetadata.RUnlock()
	return calls
}

// ProjectID calls ProjectIDFunc.
func (mock *DatasetMock) ProjectID() string {
	if mock.ProjectIDFunc == nil {
		panic("DatasetMock.ProjectIDFunc: method is nil but Dataset.ProjectID was just called")
	}
	callInfo := struct {
	}{}
	mock.lockProjectID.Lock()
	mock.calls.ProjectID = append(mock.calls.ProjectID, callInfo)
	mock.lockProjectID.Unlock()
	return mock.ProjectIDFunc()
}

// ProjectIDCalls gets all the calls that were made to ProjectID.
// Check the length with:
//     len(mockedDataset.ProjectIDCalls())
func (mock *DatasetMock) ProjectIDCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockProjectID.RLock()
	calls = mock.calls.ProjectID
	mock.lockProjectID.RUnlock()
	return calls
}

// Table calls TableFunc.
func (mock *DatasetMock) Table(in1 string) bqiface.Table {
	if mock.TableFunc == nil {
		panic("DatasetMock.TableFunc: method is nil but Dataset.Table was just called")
	}
	callInfo := struct {
		In1 string
	}{
		In1: in1,
	}
	mock.lockTable.Lock()
	mock.calls.Table = append(mock.calls.Table, callInfo)
	mock.lockTable.Unlock()
	return mock.TableFunc(in1)
}

// TableCalls gets all the calls that were made to Table.
// Check the length with:
//     len(mockedDataset.TableCalls())
func (mock *DatasetMock) TableCalls() []struct {
	In1 string
} {
	var calls []struct {
		In1 string
	}
	mock.lockTable.RLock()
	calls = mock.calls.Table
	mock.lockTable.RUnlock()
	return calls
}

// Tables calls TablesFunc.
func (mock *DatasetMock) Tables(in1 context.Context) bqiface.TableIterator {
	if mock.TablesFunc == nil {
		panic("DatasetMock.TablesFunc: method is nil but Dataset.Tables was just called")
	}
	callInfo := struct {
		In1 context.Context
	}{
		In1: in1,
	}
	mock.lockTables.Lock()
	mock.calls.Tables = append(mock.calls.Tables, callInfo)
	mock.lockTables.Unlock()
	return mock.TablesFunc(in1)
}

// TablesCalls gets all the calls that were made to Tables.
// Check the length with:
//     len(mockedDataset.TablesCalls())
func (mock *DatasetMock) TablesCalls() []struct {
	In1 context.Context
} {
	var calls []struct {
		In1 context.Context
	}
	mock.lockTables.RLock()
	calls = mock.calls.Tables
	mock.lockTables.RUnlock()
	return calls
}

// Update calls UpdateFunc.
func (mock *DatasetMock) Update(in1 context.Context, in2 bqiface.DatasetMetadataToUpdate, in3 string) (*bqiface.DatasetMetadata, error) {
	if mock.UpdateFunc == nil {
		panic("DatasetMock.UpdateFunc: method is nil but Dataset.Update was just called")
	}
	callInfo := struct {
		In1 context.Context
		In2 bqiface.DatasetMetadataToUpdate
		In3 string
	}{
		In1: in1,
		In2: in2,
		In3: in3,
	}
	mock.lockUpdate.Lock()
	mock.calls.Update = append(mock.calls.Update, callInfo)
	mock.lockUpdate.Unlock()
	return mock.UpdateFunc(in1, in2, in3)
}

// UpdateCalls gets all the calls that were made to Update.
// Check the length with:
//     len(mockedDataset.UpdateCalls())
func (mock *DatasetMock) UpdateCalls() []struct {
	In1 context.Context
	In2 bqiface.DatasetMetadataToUpdate
	In3 string
} {
	var calls []struct {
		In1 context.Context
		In2 bqiface.DatasetMetadataToUpdate
		In3 string
	}
	mock.lockUpdate.RLock()
	calls = mock.calls.Update
	mock.lockUpdate.RUnlock()
	return calls
}
