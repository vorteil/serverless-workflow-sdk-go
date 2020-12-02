// Copyright 2020 The Serverless Workflow Specification Authors
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

// Code generated by Serverless Workflow Go SDK, DO NOT EDIT.

package model

// GetId ...
func (j Delaystate) GetId() string { return *j.Id }

// SetId ...
func (j *Delaystate) SetId(id string) { j.Id = &id }

// GetName ...
func (j Delaystate) GetName() string { return *j.Name }

// GetType ...
func (j Delaystate) GetType() string { return *j.Type }

// GetType ...
func (j Delaystate) GetStart() *Start { return j.Start }

// GetStateDataFilter ...
func (j Delaystate) GetStateDataFilter() Statedatafilter { return *j.StateDataFilter }

// GetDataInputSchema ...
func (j Delaystate) GetDataInputSchema() string { return *j.DataInputSchema }

// GetDataOutputSchema ...
func (j Delaystate) GetDataOutputSchema() string { return *j.DataOutputSchema }

// GetMetadata ...
func (j Delaystate) GetMetadata() Metadata_1 { return j.Metadata }

// GetId ...
func (j Eventstate) GetId() string { return *j.Id }

// SetId ...
func (j *Eventstate) SetId(id string) { j.Id = &id }

// GetName ...
func (j Eventstate) GetName() string { return *j.Name }

// GetType ...
func (j Eventstate) GetType() string { return *j.Type }

// GetType ...
func (j Eventstate) GetStart() *Start { return j.Start }

// GetStateDataFilter ...
func (j Eventstate) GetStateDataFilter() Statedatafilter { return *j.StateDataFilter }

// GetDataInputSchema ...
func (j Eventstate) GetDataInputSchema() string { return *j.DataInputSchema }

// GetDataOutputSchema ...
func (j Eventstate) GetDataOutputSchema() string { return *j.DataOutputSchema }

// GetMetadata ...
func (j Eventstate) GetMetadata() Metadata_1 { return j.Metadata }

// GetId ...
func (j Operationstate) GetId() string { return *j.Id }

// SetId ...
func (j *Operationstate) SetId(id string) { j.Id = &id }

// GetName ...
func (j Operationstate) GetName() string { return *j.Name }

// GetType ...
func (j Operationstate) GetType() string { return *j.Type }

// GetType ...
func (j Operationstate) GetStart() *Start { return j.Start }

// GetStateDataFilter ...
func (j Operationstate) GetStateDataFilter() Statedatafilter { return *j.StateDataFilter }

// GetDataInputSchema ...
func (j Operationstate) GetDataInputSchema() string { return *j.DataInputSchema }

// GetDataOutputSchema ...
func (j Operationstate) GetDataOutputSchema() string { return *j.DataOutputSchema }

// GetMetadata ...
func (j Operationstate) GetMetadata() Metadata_1 { return j.Metadata }

// GetId ...
func (j Parallelstate) GetId() string { return *j.Id }

// GetName ...
func (j Parallelstate) GetName() string { return *j.Name }

// SetId ...
func (j *Parallelstate) SetId(id string) { j.Id = &id }

// GetType ...
func (j Parallelstate) GetType() string { return *j.Type }

// GetType ...
func (j Parallelstate) GetStart() *Start { return j.Start }

// GetStateDataFilter ...
func (j Parallelstate) GetStateDataFilter() Statedatafilter { return *j.StateDataFilter }

// GetDataInputSchema ...
func (j Parallelstate) GetDataInputSchema() string { return *j.DataInputSchema }

// GetDataOutputSchema ...
func (j Parallelstate) GetDataOutputSchema() string { return *j.DataOutputSchema }

// GetMetadata ...
func (j Parallelstate) GetMetadata() Metadata_1 { return j.Metadata }

// GetId ...
func (j Subflowstate) GetId() string { return *j.Id }

// SetId ...
func (j *Subflowstate) SetId(id string) { j.Id = &id }

// GetName ...
func (j Subflowstate) GetName() string { return *j.Name }

// GetType ...
func (j Subflowstate) GetType() string { return *j.Type }

// GetType ...
func (j Subflowstate) GetStart() *Start { return j.Start }

// GetStateDataFilter ...
func (j Subflowstate) GetStateDataFilter() Statedatafilter { return *j.StateDataFilter }

// GetDataInputSchema ...
func (j Subflowstate) GetDataInputSchema() string { return *j.DataInputSchema }

// GetDataOutputSchema ...
func (j Subflowstate) GetDataOutputSchema() string { return *j.DataOutputSchema }

// GetMetadata ...
func (j Subflowstate) GetMetadata() Metadata_1 { return j.Metadata }

// GetId ...
func (j Injectstate) GetId() string { return *j.Id }

// SetId ...
func (j *Injectstate) SetId(id string) { j.Id = &id }

// GetName ...
func (j Injectstate) GetName() string { return *j.Name }

// GetType ...
func (j Injectstate) GetType() string { return *j.Type }

// GetType ...
func (j Injectstate) GetStart() *Start { return j.Start }

// GetStateDataFilter ...
func (j Injectstate) GetStateDataFilter() Statedatafilter { return *j.StateDataFilter }

// GetDataInputSchema ...
func (j Injectstate) GetDataInputSchema() string { return *j.DataInputSchema }

// GetDataOutputSchema ...
func (j Injectstate) GetDataOutputSchema() string { return *j.DataOutputSchema }

// GetMetadata ...
func (j Injectstate) GetMetadata() Metadata_1 { return j.Metadata }

// GetId ...
func (j Foreachstate) GetId() string { return *j.Id }

// SetId ...
func (j *Foreachstate) SetId(id string) { j.Id = &id }

// GetName ...
func (j Foreachstate) GetName() string { return *j.Name }

// GetType ...
func (j Foreachstate) GetType() string { return *j.Type }

// GetType ...
func (j Foreachstate) GetStart() *Start { return j.Start }

// GetStateDataFilter ...
func (j Foreachstate) GetStateDataFilter() Statedatafilter { return *j.StateDataFilter }

// GetDataInputSchema ...
func (j Foreachstate) GetDataInputSchema() string { return *j.DataInputSchema }

// GetDataOutputSchema ...
func (j Foreachstate) GetDataOutputSchema() string { return *j.DataOutputSchema }

// GetMetadata ...
func (j Foreachstate) GetMetadata() Metadata_1 { return j.Metadata }

// GetId ...
func (j Callbackstate) GetId() string { return *j.Id }

// SetId ...
func (j *Callbackstate) SetId(id string) { j.Id = &id }

// GetName ...
func (j Callbackstate) GetName() string { return *j.Name }

// GetType ...
func (j Callbackstate) GetType() string { return *j.Type }

// GetType ...
func (j Callbackstate) GetStart() *Start { return j.Start }

// GetStateDataFilter ...
func (j Callbackstate) GetStateDataFilter() Statedatafilter { return *j.StateDataFilter }

// GetDataInputSchema ...
func (j Callbackstate) GetDataInputSchema() string { return *j.DataInputSchema }

// GetDataOutputSchema ...
func (j Callbackstate) GetDataOutputSchema() string { return *j.DataOutputSchema }

// GetMetadata ...
func (j Callbackstate) GetMetadata() Metadata_1 { return j.Metadata }

// GetId ...
func (j Databasedswitch) GetId() string { return *j.Id }

// SetId ...
func (j *Databasedswitch) SetId(id string) { j.Id = &id }

// GetName ...
func (j Databasedswitch) GetName() string { return *j.Name }

// GetType ...
func (j Databasedswitch) GetType() string { return *j.Type }

// GetType ...
func (j Databasedswitch) GetStart() *Start { return j.Start }

// GetStateDataFilter ...
func (j Databasedswitch) GetStateDataFilter() Statedatafilter { return *j.StateDataFilter }

// GetDataInputSchema ...
func (j Databasedswitch) GetDataInputSchema() string { return *j.DataInputSchema }

// GetDataOutputSchema ...
func (j Databasedswitch) GetDataOutputSchema() string { return *j.DataOutputSchema }

// GetMetadata ...
func (j Databasedswitch) GetMetadata() Metadata_1 { return j.Metadata }

// GetId ...
func (j Eventbasedswitch) GetId() string { return *j.Id }

// SetId ...
func (j *Eventbasedswitch) SetId(id string) { j.Id = &id }

// GetName ...
func (j Eventbasedswitch) GetName() string { return *j.Name }

// GetType ...
func (j Eventbasedswitch) GetType() string { return *j.Type }

// GetType ...
func (j Eventbasedswitch) GetStart() *Start { return j.Start }

// GetStateDataFilter ...
func (j Eventbasedswitch) GetStateDataFilter() Statedatafilter { return *j.StateDataFilter }

// GetDataInputSchema ...
func (j Eventbasedswitch) GetDataInputSchema() string { return *j.DataInputSchema }

// GetDataOutputSchema ...
func (j Eventbasedswitch) GetDataOutputSchema() string { return *j.DataOutputSchema }

// GetMetadata ...
func (j Eventbasedswitch) GetMetadata() Metadata_1 { return j.Metadata }
