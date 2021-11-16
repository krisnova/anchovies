// Copyright © 2021 Kris Nóva <kris@nivenly.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package anchovies

import (
	"os"
	"testing"

	"github.com/kris-nova/logger"
)

type TestRecord struct {
	EmbedRecord
	TestString string
	TestInt    int
}

func TestMain(m *testing.M) {
	logger.BitwiseLevel = logger.LogEverything
	os.Exit(m.Run())
}

func TestWriteSimpleRecord(t *testing.T) {
	id := U("mykey")
	m := &TestRecord{
		EmbedRecord: EmbedRecord{
			ID: id,
		},
	}
	err := Write(m)
	if err != nil {
		t.Errorf("unable to write record: %v", err)
	}
}

func TestWriteSimpleRecordMissingID(t *testing.T) {
	m := &TestRecord{}
	err := Write(m)
	if err == nil {
		t.Errorf("possible to write record with missing ID")
	} else {
		logger.Debug(err.Error())
	}
}

func TestWriteReadSimpleRecord(t *testing.T) {
	id := U("anotherkey")
	m := &TestRecord{
		EmbedRecord: EmbedRecord{
			ID: id,
		},
		TestString: "boops",
	}
	err := Write(m)
	if err != nil {
		t.Errorf("unable to write record: %v", err)
	}
	mm := TestRecord{}
	err = Read(id, &mm)
	if err != nil {
		t.Errorf("unable to read record: %v", err)
	}
	if mm.TestString != m.TestString {
		t.Errorf("failed to persist boops")
	}
	if mm.TestInt != m.TestInt {
		t.Errorf("failed to persist boops")
	}
}

func TestWriteReadSimpleRecordDatabaseOverload(t *testing.T) {
	SetDir("/home/nova/.anchoviesboops")
	id := U("anotherkey")
	m := &TestRecord{
		EmbedRecord: EmbedRecord{
			ID: id,
		},
		TestString: "boops",
	}
	err := Write(m)
	if err != nil {
		t.Errorf("unable to write record: %v", err)
	}
	mm := TestRecord{}
	err = Read(id, &mm)
	if err != nil {
		t.Errorf("unable to read record: %v", err)
	}
	if mm.TestString != m.TestString {
		t.Errorf("failed to persist boops")
	}
	if mm.TestInt != m.TestInt {
		t.Errorf("failed to persist boops")
	}
}
