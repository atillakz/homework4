package main

import (
	"sort"
)


type Strings struct {

	MyString []string
}

func (st *Strings) Len() int {

	return len(st.MyString)
}

func (st *Strings) Swap(i, j int)   {

	temp := st.MyString[i]

	st.MyString[i] = st.MyString[j]

	st.MyString[j] = temp


}

func (st *Strings) Less(i,j int) bool {

	firstString := st.MyString[i]

	secondString := st.MyString[j]

	newStr := Strings{[]string{secondString, firstString}}

	sort.Strings(newStr.MyString)

	if newStr.MyString[0] == firstString {

		return true
	}
	return false

}



