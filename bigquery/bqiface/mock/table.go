// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"cloud.google.com/go/bigquery"
	"context"
	"github.com/MadHive/google-cloud-go-testing/bigquery/bqiface"
	"sync"
)

// Ensure, that TableMock does implement bqiface.Table.
// If this is not the case, regenerate this file with moq.
var _ bqiface.Table = &TableMock{}

// TableMock is a mock implementation of bqiface.Table.
//
//     func TestSomethingThatUsesTable(t *testing.T) {
//
//         // make and configure a mocked bqiface.Table
//         mockedTable := &TableMock{
//             CopierFromFunc: func(in1 ...bqiface.Table) bqiface.Copier {
// 	               panic("mock out the CopierFrom method")
//             },
//             CreateFunc: func(in1 context.Context, in2 *bigquery.TableMetadata) error {
// 	               panic("mock out the Create method")
//             },
//             DatasetIDFunc: func() string {
// 	               panic("mock out the DatasetID method")
//             },
//             DeleteFunc: func(in1 context.Context) error {
// 	               panic("mock out the Delete method")
//             },
//             ExtractorToFunc: func(dst *bigquery.GCSReference) bqiface.Extractor {
// 	               panic("mock out the ExtractorTo method")
//             },
//             FullyQualifiedNameFunc: func() string {
// 	               panic("mock out the FullyQualifiedName method")
//             },
//             LoaderFromFunc: func(in1 bigquery.LoadSource) bqiface.Loader {
// 	               panic("mock out the LoaderFrom method")
//             },
//             MetadataFunc: func(in1 context.Context) (*bigquery.TableMetadata, error) {
// 	               panic("mock out the Metadata method")
//             },
//             ProjectIDFunc: func() string {
// 	               panic("mock out the ProjectID method")
//             },
//             ReadFunc: func(ctx context.Context) bqiface.RowIterator {
// 	               panic("mock out the Read method")
//             },
//             TableIDFunc: func() string {
// 	               panic("mock out the TableID method")
//             },
//             UpdateFunc: func(in1 context.Context, in2 bigquery.TableMetadataToUpdate, in3 string) (*bigquery.TableMetadata, error) {
// 	               panic("mock out the Update method")
//             },
//             UploaderFunc: func() bqiface.Uploader {
// 	               panic("mock out the Uploader method")
//             },
//         }
//
//         // use mockedTable in code that requires bqiface.Table
//         // and then make assertions.
//
//     }
type TableMock struct {
	// CopierFromFunc mocks the CopierFrom method.
	CopierFromFunc func(in1 ...bqiface.Table) bqiface.Copier

	// CreateFunc mocks the Create method.
	CreateFunc func(in1 context.Context, in2 *bigquery.TableMetadata) error

	// DatasetIDFunc mocks the DatasetID method.
	DatasetIDFunc func() string

	// DeleteFunc mocks the Delete method.
	DeleteFunc func(in1 context.Context) error

	// ExtractorToFunc mocks the ExtractorTo method.
	ExtractorToFunc func(dst *bigquery.GCSReference) bqiface.Extractor

	// FullyQualifiedNameFunc mocks the FullyQualifiedName method.
	FullyQualifiedNameFunc func() string

	// LoaderFromFunc mocks the LoaderFrom method.
	LoaderFromFunc func(in1 bigquery.LoadSource) bqiface.Loader

	// MetadataFunc mocks the Metadata method.
	MetadataFunc func(in1 context.Context) (*bigquery.TableMetadata, error)

	// ProjectIDFunc mocks the ProjectID method.
	ProjectIDFunc func() string

	// ReadFunc mocks the Read method.
	ReadFunc func(ctx context.Context) bqiface.RowIterator

	// TableIDFunc mocks the TableID method.
	TableIDFunc func() string

	// UpdateFunc mocks the Update method.
	UpdateFunc func(in1 context.Context, in2 bigquery.TableMetadataToUpdate, in3 string) (*bigquery.TableMetadata, error)

	// UploaderFunc mocks the Uploader method.
	UploaderFunc func() bqiface.Uploader

	// calls tracks calls to the methods.
	calls struct {
		// CopierFrom holds details about calls to the CopierFrom method.
		CopierFrom []struct {
			// In1 is the in1 argument value.
			In1 []bqiface.Table
		}
		// Create holds details about calls to the Create method.
		Create []struct {
			// In1 is the in1 argument value.
			In1 context.Context
			// In2 is the in2 argument value.
			In2 *bigquery.TableMetadata
		}
		// DatasetID holds details about calls to the DatasetID method.
		DatasetID []struct {
		}
		// Delete holds details about calls to the Delete method.
		Delete []struct {
			// In1 is the in1 argument value.
			In1 context.Context
		}
		// ExtractorTo holds details about calls to the ExtractorTo method.
		ExtractorTo []struct {
			// Dst is the dst argument value.
			Dst *bigquery.GCSReference
		}
		// FullyQualifiedName holds details about calls to the FullyQualifiedName method.
		FullyQualifiedName []struct {
		}
		// LoaderFrom holds details about calls to the LoaderFrom method.
		LoaderFrom []struct {
			// In1 is the in1 argument value.
			In1 bigquery.LoadSource
		}
		// Metadata holds details about calls to the Metadata method.
		Metadata []struct {
			// In1 is the in1 argument value.
			In1 context.Context
		}
		// ProjectID holds details about calls to the ProjectID method.
		ProjectID []struct {
		}
		// Read holds details about calls to the Read method.
		Read []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// TableID holds details about calls to the TableID method.
		TableID []struct {
		}
		// Update holds details about calls to the Update method.
		Update []struct {
			// In1 is the in1 argument value.
			In1 context.Context
			// In2 is the in2 argument value.
			In2 bigquery.TableMetadataToUpdate
			// In3 is the in3 argument value.
			In3 string
		}
		// Uploader holds details about calls to the Uploader method.
		Uploader []struct {
		}
	}
	lockCopierFrom         sync.RWMutex
	lockCreate             sync.RWMutex
	lockDatasetID          sync.RWMutex
	lockDelete             sync.RWMutex
	lockExtractorTo        sync.RWMutex
	lockFullyQualifiedName sync.RWMutex
	lockLoaderFrom         sync.RWMutex
	lockMetadata           sync.RWMutex
	lockProjectID          sync.RWMutex
	lockRead               sync.RWMutex
	lockTableID            sync.RWMutex
	lockUpdate             sync.RWMutex
	lockUploader           sync.RWMutex
}

// CopierFrom calls CopierFromFunc.
func (mock *TableMock) CopierFrom(in1 ...bqiface.Table) bqiface.Copier {
	if mock.CopierFromFunc == nil {
		panic("TableMock.CopierFromFunc: method is nil but Table.CopierFrom was just called")
	}
	callInfo := struct {
		In1 []bqiface.Table
	}{
		In1: in1,
	}
	mock.lockCopierFrom.Lock()
	mock.calls.CopierFrom = append(mock.calls.CopierFrom, callInfo)
	mock.lockCopierFrom.Unlock()
	return mock.CopierFromFunc(in1...)
}

// CopierFromCalls gets all the calls that were made to CopierFrom.
// Check the length with:
//     len(mockedTable.CopierFromCalls())
func (mock *TableMock) CopierFromCalls() []struct {
	In1 []bqiface.Table
} {
	var calls []struct {
		In1 []bqiface.Table
	}
	mock.lockCopierFrom.RLock()
	calls = mock.calls.CopierFrom
	mock.lockCopierFrom.RUnlock()
	return calls
}

// Create calls CreateFunc.
func (mock *TableMock) Create(in1 context.Context, in2 *bigquery.TableMetadata) error {
	if mock.CreateFunc == nil {
		panic("TableMock.CreateFunc: method is nil but Table.Create was just called")
	}
	callInfo := struct {
		In1 context.Context
		In2 *bigquery.TableMetadata
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
//     len(mockedTable.CreateCalls())
func (mock *TableMock) CreateCalls() []struct {
	In1 context.Context
	In2 *bigquery.TableMetadata
} {
	var calls []struct {
		In1 context.Context
		In2 *bigquery.TableMetadata
	}
	mock.lockCreate.RLock()
	calls = mock.calls.Create
	mock.lockCreate.RUnlock()
	return calls
}

// DatasetID calls DatasetIDFunc.
func (mock *TableMock) DatasetID() string {
	if mock.DatasetIDFunc == nil {
		panic("TableMock.DatasetIDFunc: method is nil but Table.DatasetID was just called")
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
//     len(mockedTable.DatasetIDCalls())
func (mock *TableMock) DatasetIDCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockDatasetID.RLock()
	calls = mock.calls.DatasetID
	mock.lockDatasetID.RUnlock()
	return calls
}

// Delete calls DeleteFunc.
func (mock *TableMock) Delete(in1 context.Context) error {
	if mock.DeleteFunc == nil {
		panic("TableMock.DeleteFunc: method is nil but Table.Delete was just called")
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
//     len(mockedTable.DeleteCalls())
func (mock *TableMock) DeleteCalls() []struct {
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

// ExtractorTo calls ExtractorToFunc.
func (mock *TableMock) ExtractorTo(dst *bigquery.GCSReference) bqiface.Extractor {
	if mock.ExtractorToFunc == nil {
		panic("TableMock.ExtractorToFunc: method is nil but Table.ExtractorTo was just called")
	}
	callInfo := struct {
		Dst *bigquery.GCSReference
	}{
		Dst: dst,
	}
	mock.lockExtractorTo.Lock()
	mock.calls.ExtractorTo = append(mock.calls.ExtractorTo, callInfo)
	mock.lockExtractorTo.Unlock()
	return mock.ExtractorToFunc(dst)
}

// ExtractorToCalls gets all the calls that were made to ExtractorTo.
// Check the length with:
//     len(mockedTable.ExtractorToCalls())
func (mock *TableMock) ExtractorToCalls() []struct {
	Dst *bigquery.GCSReference
} {
	var calls []struct {
		Dst *bigquery.GCSReference
	}
	mock.lockExtractorTo.RLock()
	calls = mock.calls.ExtractorTo
	mock.lockExtractorTo.RUnlock()
	return calls
}

// FullyQualifiedName calls FullyQualifiedNameFunc.
func (mock *TableMock) FullyQualifiedName() string {
	if mock.FullyQualifiedNameFunc == nil {
		panic("TableMock.FullyQualifiedNameFunc: method is nil but Table.FullyQualifiedName was just called")
	}
	callInfo := struct {
	}{}
	mock.lockFullyQualifiedName.Lock()
	mock.calls.FullyQualifiedName = append(mock.calls.FullyQualifiedName, callInfo)
	mock.lockFullyQualifiedName.Unlock()
	return mock.FullyQualifiedNameFunc()
}

// FullyQualifiedNameCalls gets all the calls that were made to FullyQualifiedName.
// Check the length with:
//     len(mockedTable.FullyQualifiedNameCalls())
func (mock *TableMock) FullyQualifiedNameCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockFullyQualifiedName.RLock()
	calls = mock.calls.FullyQualifiedName
	mock.lockFullyQualifiedName.RUnlock()
	return calls
}

// LoaderFrom calls LoaderFromFunc.
func (mock *TableMock) LoaderFrom(in1 bigquery.LoadSource) bqiface.Loader {
	if mock.LoaderFromFunc == nil {
		panic("TableMock.LoaderFromFunc: method is nil but Table.LoaderFrom was just called")
	}
	callInfo := struct {
		In1 bigquery.LoadSource
	}{
		In1: in1,
	}
	mock.lockLoaderFrom.Lock()
	mock.calls.LoaderFrom = append(mock.calls.LoaderFrom, callInfo)
	mock.lockLoaderFrom.Unlock()
	return mock.LoaderFromFunc(in1)
}

// LoaderFromCalls gets all the calls that were made to LoaderFrom.
// Check the length with:
//     len(mockedTable.LoaderFromCalls())
func (mock *TableMock) LoaderFromCalls() []struct {
	In1 bigquery.LoadSource
} {
	var calls []struct {
		In1 bigquery.LoadSource
	}
	mock.lockLoaderFrom.RLock()
	calls = mock.calls.LoaderFrom
	mock.lockLoaderFrom.RUnlock()
	return calls
}

// Metadata calls MetadataFunc.
func (mock *TableMock) Metadata(in1 context.Context) (*bigquery.TableMetadata, error) {
	if mock.MetadataFunc == nil {
		panic("TableMock.MetadataFunc: method is nil but Table.Metadata was just called")
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
//     len(mockedTable.MetadataCalls())
func (mock *TableMock) MetadataCalls() []struct {
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
func (mock *TableMock) ProjectID() string {
	if mock.ProjectIDFunc == nil {
		panic("TableMock.ProjectIDFunc: method is nil but Table.ProjectID was just called")
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
//     len(mockedTable.ProjectIDCalls())
func (mock *TableMock) ProjectIDCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockProjectID.RLock()
	calls = mock.calls.ProjectID
	mock.lockProjectID.RUnlock()
	return calls
}

// Read calls ReadFunc.
func (mock *TableMock) Read(ctx context.Context) bqiface.RowIterator {
	if mock.ReadFunc == nil {
		panic("TableMock.ReadFunc: method is nil but Table.Read was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockRead.Lock()
	mock.calls.Read = append(mock.calls.Read, callInfo)
	mock.lockRead.Unlock()
	return mock.ReadFunc(ctx)
}

// ReadCalls gets all the calls that were made to Read.
// Check the length with:
//     len(mockedTable.ReadCalls())
func (mock *TableMock) ReadCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockRead.RLock()
	calls = mock.calls.Read
	mock.lockRead.RUnlock()
	return calls
}

// TableID calls TableIDFunc.
func (mock *TableMock) TableID() string {
	if mock.TableIDFunc == nil {
		panic("TableMock.TableIDFunc: method is nil but Table.TableID was just called")
	}
	callInfo := struct {
	}{}
	mock.lockTableID.Lock()
	mock.calls.TableID = append(mock.calls.TableID, callInfo)
	mock.lockTableID.Unlock()
	return mock.TableIDFunc()
}

// TableIDCalls gets all the calls that were made to TableID.
// Check the length with:
//     len(mockedTable.TableIDCalls())
func (mock *TableMock) TableIDCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockTableID.RLock()
	calls = mock.calls.TableID
	mock.lockTableID.RUnlock()
	return calls
}

// Update calls UpdateFunc.
func (mock *TableMock) Update(in1 context.Context, in2 bigquery.TableMetadataToUpdate, in3 string) (*bigquery.TableMetadata, error) {
	if mock.UpdateFunc == nil {
		panic("TableMock.UpdateFunc: method is nil but Table.Update was just called")
	}
	callInfo := struct {
		In1 context.Context
		In2 bigquery.TableMetadataToUpdate
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
//     len(mockedTable.UpdateCalls())
func (mock *TableMock) UpdateCalls() []struct {
	In1 context.Context
	In2 bigquery.TableMetadataToUpdate
	In3 string
} {
	var calls []struct {
		In1 context.Context
		In2 bigquery.TableMetadataToUpdate
		In3 string
	}
	mock.lockUpdate.RLock()
	calls = mock.calls.Update
	mock.lockUpdate.RUnlock()
	return calls
}

// Uploader calls UploaderFunc.
func (mock *TableMock) Uploader() bqiface.Uploader {
	if mock.UploaderFunc == nil {
		panic("TableMock.UploaderFunc: method is nil but Table.Uploader was just called")
	}
	callInfo := struct {
	}{}
	mock.lockUploader.Lock()
	mock.calls.Uploader = append(mock.calls.Uploader, callInfo)
	mock.lockUploader.Unlock()
	return mock.UploaderFunc()
}

// UploaderCalls gets all the calls that were made to Uploader.
// Check the length with:
//     len(mockedTable.UploaderCalls())
func (mock *TableMock) UploaderCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockUploader.RLock()
	calls = mock.calls.Uploader
	mock.lockUploader.RUnlock()
	return calls
}