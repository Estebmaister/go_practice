package str

import (
	"cmp"
	"reflect"
	"slices"
	"testing"
)

func Test_canConstruct_and_memo(t *testing.T) {
	t.Parallel()
	type args struct {
		s        string
		wordBank []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"abcdef",
			args{
				"abcdef",
				[]string{"ab", "abc", "cd", "def", "abcd"},
			},
			true,
		},
		{
			"skateboard",
			args{
				"skateboard",
				[]string{"bo", "rd", "ate", "t", "ska", "sk", "boar"},
			},
			false,
		},
		{
			"enterapotentpot",
			args{
				"enterapotentpot",
				[]string{"a", "p", "ent", "enter", "ot", "o", "t"},
			},
			true,
		},
		{
			"very large iteration",
			args{
				"eeeeeeeeeeeeeeeeeeeeeeeeeeeef", // time complexity grows exponentially with length
				[]string{"e", "ee", "eee", "eeee", "eeeee", "eeeeee"},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canConstruct(tt.args.s, tt.args.wordBank); got != tt.want {
				t.Errorf("canConstruct() = %v, want %v", got, tt.want)
			}
		})
	}
	for _, tt := range tests {
		t.Run(tt.name+" memoized", func(t *testing.T) {
			if got := canConstructMemo(tt.args.s, tt.args.wordBank, map[string]bool{}); got != tt.want {
				t.Errorf("canConstructMemo() = %v, want %v", got, tt.want)
			}
		})
	}
	for _, tt := range tests {
		t.Run(tt.name+" tabulated", func(t *testing.T) {
			if got := canConstructTab(tt.args.s, tt.args.wordBank); got != tt.want {
				t.Errorf("canConstructTab() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countConstruct(t *testing.T) {
	t.Parallel()
	type args struct {
		s        string
		wordBank []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test1",
			args{
				"purple",
				[]string{"purp", "p", "ur", "le", "purpl"},
			},
			2,
		},
		{
			"test2",
			args{
				"abcdef",
				[]string{"ab", "abc", "cd", "def", "abcd"},
			},
			1,
		},
		{
			"test3",
			args{
				"skateboard",
				[]string{"bo", "rd", "ate", "t", "ska", "sk", "boar"},
			},
			0,
		},
		{
			"test4",
			args{
				"enterapotentpot",
				[]string{"a", "p", "ent", "enter", "ot", "o", "t"},
			},
			4,
		},
		{
			"test5",
			args{
				"",
				[]string{"cat", "dog", "mouse"},
			},
			1,
		},
		{
			"test6",
			args{
				"eeeeeeeeeeeeeeeeeeeeeeeeeeef",
				[]string{"e", "ee", "eee", "eeee", "eeeee", "eeeeee"},
			},
			0,
		},
		{
			"test7",
			args{
				"eeeeeeeeeeeeeeeeeeeeeeeeeeef",
				[]string{"e", "ee", "eee", "eeee", "eeeee", "eeeeee", "f"},
			},
			56058368,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countConstruct(tt.args.s, tt.args.wordBank); got != tt.want {
				t.Errorf("countConstruct() = %v, want %v", got, tt.want)
			}
		})
	}
	for _, tt := range tests {
		t.Run(tt.name+" memo", func(t *testing.T) {
			if got := countConstructMemo(tt.args.s, tt.args.wordBank, map[string]int{}); got != tt.want {
				t.Errorf("countConstructMemo() = %v, want %v", got, tt.want)
			}
		})
	}
	for _, tt := range tests {
		t.Run(tt.name+" tabulation", func(t *testing.T) {
			if got := countConstructTab(tt.args.s, tt.args.wordBank); got != tt.want {
				t.Errorf("countConstructTab() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_allConstruct(t *testing.T) {
	t.Parallel()
	var emptyDobleSlice [][]string
	type args struct {
		s        string
		wordBank []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			"purple",
			args{
				"purple",
				[]string{"purp", "p", "ur", "le", "purpl"},
			},
			[][]string{
				{"purp", "le"},
				{"p", "ur", "p", "le"},
			},
		},
		{
			"abcdef",
			args{
				"abcdef",
				[]string{"ab", "abc", "cd", "def", "abcd", "ef", "c"},
			},
			[][]string{
				{"ab", "cd", "ef"},
				{"ab", "c", "def"},
				{"abc", "def"},
				{"abcd", "ef"},
			},
		},
		{
			"skateboard",
			args{
				"skateboard",
				[]string{"bo", "rd", "ate", "t", "ska", "sk", "boar"},
			},
			emptyDobleSlice,
		},
		{
			"enterapotentpot",
			args{
				"enterapotentpot",
				[]string{"a", "p", "ent", "enter", "ot", "o", "t"},
			},
			[][]string{
				{"enter", "a", "p", "ot", "ent", "p", "ot"},
				{"enter", "a", "p", "ot", "ent", "p", "o", "t"},
				{"enter", "a", "p", "o", "t", "ent", "p", "ot"},
				{"enter", "a", "p", "o", "t", "ent", "p", "o", "t"},
			},
		},
		{
			"empty slice",
			args{
				"",
				[]string{"cat", "dog", "mouse"},
			},
			[][]string{{}},
		},
		// {
		// 	"very long iteration",
		// 	args{
		// 		"eeeeeeeeeeeeeeeeeeeeeeeeeeeef",
		// 		[]string{"e", "ee", "eee", "eeee", "eeeee", "eeeeee"},
		// 	},
		// 	emptyDobleSlice,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := allConstruct(tt.args.s, tt.args.wordBank); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("allConstruct() = %v, want %v", got, tt.want)
			}
		})
	}
	for _, tt := range tests {
		t.Run(tt.name+" memo", func(t *testing.T) {
			if got := allConstructMemo(tt.args.s, tt.args.wordBank, map[string][][]string{}); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("allConstructMemo() = %v, want %v", got, tt.want)
			}
		})
	}
	for _, tt := range tests {
		t.Run(tt.name+" tabulation "+tt.args.s, func(t *testing.T) {
			got := allConstructTab(tt.args.s, tt.args.wordBank)
			t.Log(got)

			slices.SortFunc(got, sortArrayOfStringArrays)
			slices.SortFunc(tt.want, sortArrayOfStringArrays)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("allConstructTab() = %v, want %v", got, tt.want)
			}
		})
	}
}

func sortArrayOfStringArrays(a, b []string) int {
	if n := cmp.Compare(len(a), len(b)); n != 0 {
		return n
	}
	// If arr length are equal, go deep, order by internal elem length
	return deepSortHelper(a, b)
}

func deepSortHelper(a, b []string) int {
	for i := 0; i < len(a); i++ {
		if n := cmp.Compare(len(a[i]), len(b[i])); n != 0 {
			return n
		}
	}
	return cmp.Compare(a[len(a)-1], b[len(b)-1])
}
