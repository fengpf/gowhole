// Copyright 2016 The etcd Authors
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

package schedule

import (
	"context"
	"fmt"
	"testing"
)

func TestFIFOSchedule(t *testing.T) {
	jobNum := 3
	s := NewFIFOScheduler()
	defer s.Stop()

	next := 0
	jobCreator := func(i int) Job {
		return func(ctx context.Context) {
			if next != i {
				t.Fatalf("job#%d: got %d, want %d", i, next, i)
			}
			next = i + 1
			// fmt.Println(next)
		}
	}

	var jobs []Job
	for i := 0; i < jobNum; i++ {
		jobs = append(jobs, jobCreator(i))
	}

	for _, j := range jobs {
		i := j
		fmt.Println(i)
		s.Schedule(j)
	}

	s.WaitFinish(jobNum)
	if s.Scheduled() != jobNum {
		t.Errorf("scheduled = %d, want %d", s.Scheduled(), jobNum)
	}
}
