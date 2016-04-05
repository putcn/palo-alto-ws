package main

import (
	humm "fmt"
	rand "math/rand"
	"strconv"
)

type TextIndex struct {
	docId    int
	position int
}

var (
	pop           string  = "中文测试"
	fNumber       float64 = 3.142
	invertedIndex map[string][]*TextIndex
	terms                    = []string{"this", "is", "some", "test", "text"}
	testt         *TextIndex = &TextIndex{docId: 1, position: 2}
)

func init() {
	createInvertIndex()
}

func main() {
	searchKeyWords := []string{"this", "wow", "ok"}

	report1 := " - doc %v at position %v \n"
	report2 := " - the doc content is : %v"

	for j := 0; j < len(searchKeyWords); j++ {
		docIndexs, matchNum := Search(searchKeyWords[j])
		humm.Printf("%v matches \n", matchNum)
		for i := 0; i < len(docIndexs); i++ {
			humm.Printf(report1, docIndexs[i].docId, docIndexs[i].position)
			humm.Printf(report2, getDocById(docIndexs[i].docId))
		}
	}
}

func test() (x int, y string) {
	return 1, "foo"
}

func createInvertIndex() {
	invertedIndex = make(map[string][]*TextIndex)
	for i := 0; i < len(terms); i++ {
		addTerm(terms[i])
	}
	humm.Println("finished creating index")
}

func getDocById(docId int) string {
	return "some random text for id " + strconv.Itoa(docId) + "\n"
}

func Search(t string) (docIndexs []*TextIndex, matchLength int) {
	humm.Printf("searching \"%v\" \n", t)
	_, ok := invertedIndex[t]
	if !ok {
		addTerm(t)
	}
	humm.Printf("length of index now is %v \n", len(invertedIndex))
	return invertedIndex[t], len(invertedIndex[t])
}

func addTerm(term string) {
	//add to index
	//sacn docs to update mapping
	for j := 0; j < rand.Intn(10); j++ {
		invertedIndex[term] = append(invertedIndex[term], &TextIndex{rand.Intn(100), rand.Intn(100)})
	}
}

func AddDoc(doc string) {
	//scan terms
	//update invertIndex
}
