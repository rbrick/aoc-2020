package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Bag struct {
	Name string

	HeldQuantity map[string]int

	// go through this
	Holds []*Bag
}

func (b *Bag) BagCapacity() int {
	return len(b.Holds)
}

func (b *Bag) TotalCapacity() int {
	cap := 0

	for _, sb := range b.Holds {
		cap += b.HeldQuantity[sb.Name] * sb.TotalCapacity()
	}

	return cap
}

func (b *Bag) CanHold(s string) bool {
	if b.Name == s {
		return false
	}

	for _, held := range b.Holds {
		if held.Name == s || held.CanHold(s) {
			return true
		}

	}

	return false
}

type BagGraph struct {
	Graph map[string]*Bag
}

func (g *BagGraph) Insert(b *Bag) {
	existing, ok := g.Graph[b.Name]

	if !ok {
		g.Graph[b.Name] = b
	} else {
		existing.HeldQuantity = b.HeldQuantity
		existing.Holds = append(existing.Holds, b.Holds...)
	}

	for _, other := range b.Holds {
		if _, ok := g.Graph[other.Name]; !ok {
			g.Graph[other.Name] = other
		}
	}

}

func (g *BagGraph) Find(name string) *Bag {
	return g.Graph[name]
}

func parseBag(s string, graph *BagGraph) *Bag {

	// "vibrant purple bags contain 3 shiny lavender bags, 1 mirrored gray bag, 4 muted bronze bags."

	// vibrant brown bags contain no other bags.

	s = strings.Replace(s, "bags", "", -1)
	s = strings.Replace(s, "bag", "", -1)
	s = strings.Replace(s, ".", "", -1)

	// split on contain
	split := strings.Split(s, "contain")

	// "vibrant purple"
	parent := strings.TrimSpace(split[0])
	rawBags := strings.Split(split[1], ",")

	parentBag := &Bag{
		Name:         parent,
		HeldQuantity: make(map[string]int),
	}

	for _, bag := range rawBags {
		b := strings.TrimSpace(bag)

		if b == "no other" {
			break
		}

		// parse the bag

		parsed := strings.SplitN(b, " ", 2)
		quanity, _ := strconv.Atoi(strings.TrimSpace(parsed[0]))
		name := strings.TrimSpace(parsed[1])

		subBag := graph.Find(name)

		if subBag == nil {
			subBag = &Bag{
				Name: name,
			}
		}

		parentBag.HeldQuantity[name] = quanity
		parentBag.Holds = append(parentBag.Holds, subBag)
	}

	graph.Insert(parentBag)

	return parentBag
}

func main() {

	graph := &BagGraph{
		Graph: make(map[string]*Bag),
	}

	f, err := os.Open("input")

	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		parseBag(line, graph)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	totalCanBeHeld := 0
	target := "shiny gold"

	for _, v := range graph.Graph {

		if v.CanHold(target) {
			totalCanBeHeld += 1
		}
	}

	log.Println(totalCanBeHeld)
	log.Println(graph.Find(target).TotalCapacity() - 1)
}
