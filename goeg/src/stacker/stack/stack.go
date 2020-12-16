// Copyright © 2010-12 Qtrac Ltd.
//
// This program or package and any associated files are licensed under the
// Apache License, Version 2.0 (the "License"); you may not use these files
// except in compliance with the License. You can get a copy of the License
// at: http://www.apache.org/licenses/LICENSE-2.0.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// NOTE: code was refactored for better names and to make use of pointers for all methods
// to support large stacks

package stack

import "errors"

type Stack []interface{}

func (ptr *Stack) Pop() (interface{}, error) {
    v := *ptr

    if len( v ) == 0 {
        return nil, errors.New("can't Pop() an empty stack")
    }
    top := len( v ) - 1
    x := v[ top ]

    *ptr = v[ :top ]

    return x, nil
}

func (ptr *Stack) Push(x interface{}) {
    *ptr = append(*ptr, x)
}

func (ptr *Stack) Top() (interface{}, error) {
    v := *ptr
    sz := len( v )
    if sz == 0 {
        return nil, errors.New("can't Top() an empty stack")
    }

    return v[ sz - 1 ], nil
}

func (ptr *Stack) Cap() int {
    return cap(*ptr)
}

func (ptr *Stack) Len() int {
    return len(*ptr)
}

func (ptr *Stack) IsEmpty() bool {
    return len(*ptr) == 0
}
