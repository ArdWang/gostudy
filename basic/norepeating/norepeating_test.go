package main

import "testing"

func TestSubstr(t *testing.T)  {
	tests := [] struct{
		s string
		ans int
	}{
		// Normal cases
		{"abcabcbb", 4},
		{"pwwkew",4},

		// Edge cases
		{"",0},
		{"b",1},
		{"bbbbbbbbbbbb",2},
		{"abcabcabcd",5},

		//Chinese support
		{"这里是慕课网",6},
		{"一二三三二一",4},
		{"黑化肥挥发发会飞灰发合肥发黑会飞起来",9},
	}

	for _, tt := range tests{
		actual := lengthOfNonRepeatingSubStr(tt.s)
		if actual != tt.ans{
			t.Errorf("got %d for input %s;"+ "expected %d",
				actual, tt.s, tt.ans)
		}
	}

}

func BenchmarkSubstr(b *testing.B){
	s := "黑化肥挥发发会飞灰发合肥发黑会飞起来"
	for i := 0; i<13; i++ {
		s = s + s
	}
	b.Logf("len(s) = %d", len(s))
	ans := 9
	b.ResetTimer()

	for i := 0; i < b.N; i++{
		actual := lengthOfNonRepeatingSubStr(s)
		if actual != ans{
			b.Errorf("got %d for input %s;"+ "expected %d",
				actual, s, ans)
		}
	}
}