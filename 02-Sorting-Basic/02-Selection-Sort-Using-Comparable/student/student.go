package student

import (
	"fmt"
)

// 根据原生 sort 的接口实现方法
type Student struct {
	Name  string
	Score int
}

type Students []Student

func New(name string, score int) Student {
	return Student{name, score}
}

func (s Students) Len() int {
	return len(s)
}

func (s Students) Less(i, j int) bool {
	if s[i].Score < s[j].Score {
		return true
	} else if s[i].Score > s[j].Score {
		return false
	} else {
		return s[i].Name < s[i].Name
	}
}

func (s Students) Swap(i, j int) {
	temp := s[i]
	s[i] = s[j]
	s[j] = temp
}

func (s Student) String() string {
	return fmt.Sprintf("Student: %s %v", s.Name, s.Score)
}
