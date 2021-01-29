// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package bqiface

import (
	"context"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/iterator"
)

//go:generate moq -out=mock/client.go -pkg=mock . Client
type Client interface {
	Location() string
	SetLocation(string)
	Close() error
	Dataset(string) Dataset
	DatasetInProject(string, string) Dataset
	Datasets(context.Context) DatasetIterator
	DatasetsInProject(context.Context, string) DatasetIterator
	Query(string) Query
	JobFromID(context.Context, string) (Job, error)
	JobFromIDLocation(context.Context, string, string) (Job, error)
	Jobs(context.Context) JobIterator
}

//go:generate moq -out=mock/copier.go -pkg=mock . Copier
type Copier interface {
	JobIDConfig() *bigquery.JobIDConfig
	SetCopyConfig(CopyConfig)
	Run(context.Context) (Job, error)
}

//go:generate moq -out=mock/dataset.go -pkg=mock . Dataset
type Dataset interface {
	ProjectID() string
	DatasetID() string
	Create(context.Context, *DatasetMetadata) error
	Delete(context.Context) error
	DeleteWithContents(context.Context) error
	Metadata(context.Context) (*DatasetMetadata, error)
	Update(context.Context, DatasetMetadataToUpdate, string) (*DatasetMetadata, error)
	Table(string) Table
	Tables(context.Context) TableIterator
}

//go:generate moq -out=mock/dataset_iterator.go -pkg=mock . DatasetIterator
type DatasetIterator interface {
	SetListHidden(bool)
	SetFilter(string)
	SetProjectID(string)
	Next() (Dataset, error)
	PageInfo() *iterator.PageInfo
}

//go:generate moq -out=mock/extractor.go -pkg=mock . Extractor
type Extractor interface {
	JobIDConfig() *bigquery.JobIDConfig
	SetExtractConfig(ExtractConfig)
	Run(context.Context) (Job, error)
}

//go:generate moq -out=mock/loader.go -pkg=mock . Loader
type Loader interface {
	JobIDConfig() *bigquery.JobIDConfig
	SetLoadConfig(LoadConfig)
	Run(context.Context) (Job, error)
}

//go:generate moq -out=mock/job.go -pkg=mock . Job
type Job interface {
	ID() string
	Location() string
	Config() (bigquery.JobConfig, error)
	Status(context.Context) (*bigquery.JobStatus, error)
	LastStatus() *bigquery.JobStatus
	Cancel(context.Context) error
	Wait(context.Context) (*bigquery.JobStatus, error)
	Read(context.Context) (RowIterator, error)
}

//go:generate moq -out=mock/job_iterator.go -pkg=mock . JobIterator
type JobIterator interface {
	SetProjectID(string)
	SetAllUsers(bool)
	SetState(bigquery.State)
	Next() (Job, error)
	PageInfo() *iterator.PageInfo
}

//go:generate moq -out=mock/query.go -pkg=mock . Query
type Query interface {
	JobIDConfig() *bigquery.JobIDConfig
	SetQueryConfig(QueryConfig)
	Run(context.Context) (Job, error)
	Read(context.Context) (RowIterator, error)
}

//go:generate moq -out=mock/row_iterator.go -pkg=mock . RowIterator
type RowIterator interface {
	SetStartIndex(uint64)
	Schema() bigquery.Schema
	TotalRows() uint64
	Next(interface{}) error
	PageInfo() *iterator.PageInfo
}

//go:generate moq -out=mock/table.go -pkg=mock . Table
type Table interface {
	CopierFrom(...Table) Copier
	Create(context.Context, *bigquery.TableMetadata) error
	DatasetID() string
	Delete(context.Context) error
	ExtractorTo(dst *bigquery.GCSReference) Extractor
	FullyQualifiedName() string
	LoaderFrom(bigquery.LoadSource) Loader
	Metadata(context.Context) (*bigquery.TableMetadata, error)
	ProjectID() string
	Read(ctx context.Context) RowIterator
	TableID() string
	Update(context.Context, bigquery.TableMetadataToUpdate, string) (*bigquery.TableMetadata, error)
	Uploader() Uploader
}

//go:generate moq -out=mock/table_iterator.go -pkg=mock . TableIterator
type TableIterator interface {
	Next() (Table, error)
	PageInfo() *iterator.PageInfo
}

//go:generate moq -out=mock/uploader.go -pkg=mock . Uploader
type Uploader interface {
	SetSkipInvalidRows(bool)
	SetIgnoreUnknownValues(bool)
	SetTableTemplateSuffix(string)
	Put(context.Context, interface{}) error
}
