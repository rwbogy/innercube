// Copyright or whatever, Roger Booth (roger.booth@gmail.com)
// In the unlikely event that you find this code useful,
// feel free to provide attribution :)
package main

import (
	"container/ring"
	"fmt"
)

type Color string

var colors = [...]Color{"white", "blue", "red", "yellow", "orange", "green"}
var edgesForFace = map[Color][]Color{
	"white":  {"red", "green", "orange", "blue"},
	"blue":   {"white", "orange", "yellow", "red"},
	"red":    {"blue", "yellow", "green", "white"},
	"yellow": {"green", "red", "blue", "orange"},
	"orange": {"yellow", "blue", "white", "green"},
	"green":  {"orange", "white", "red", "yellow"},
}

var edgePos = [...]int{0, 7, 6, 4, 3, 2, 6, 5, 4, 2, 1, 0}

type Face [8]Color

type Edge [12]*Color

type Cube struct {
	faceMap map[Color]*Face
	edgeMap map[Color]Edge
}

type Entanglement [8]*Cube

func NewCube() (*Cube, error) {
	newFaceMap := make(map[Color]*Face)
	newEdgeMap := make(map[Color]Edge)
	for _, color := range colors {
		newFaceMap[color] = &Face{color, color, color, color, color, color, color, color}
	}
	i := 0
	for _, faceColor := range colors {
		var newEdge Edge
		for _, edgeColor := range edgesForFace[faceColor] {
		        //fmt.Println(faceColor)
		        //fmt.Println(i)
			newEdge[i] = &newFaceMap[edgeColor][edgePos[i]]
			newEdge[i+1] = &newFaceMap[edgeColor][edgePos[i+1]]
			newEdge[i+2] = &newFaceMap[edgeColor][edgePos[i+2]]
			i += 3
			if i == 12 {
			    i = 0
			}
		}
		newEdgeMap[faceColor] = newEdge
	}
	return &Cube{newFaceMap, newEdgeMap}, nil
}

func NewEntanglement() (*Entanglement, error) {
        var newEntanglement Entanglement
	for i:=0; i<8; i++{
		newEntanglement[i], _ = NewCube();
	}
	return &newEntanglement, nil
}

type ThreeDTransformer struct {
	faceRing *ring.Ring
	edgeRing *ring.Ring
}

func ThreeDRotate(ent *Entanglement, cubeId int, face Color, direction int) error {
        newFaceRing := ring.New(8)
        newEdgeRing := ring.New(12)
        trx := ThreeDTransformer{
		newFaceRing,newEdgeRing}
	for _, faceColor := range ent[cubeId].faceMap[face] {
		trx.faceRing.Value = faceColor
		trx.faceRing = trx.faceRing.Next()
	}
	for _, edgeColorPtr := range ent[cubeId].edgeMap[face] {
		trx.edgeRing.Value = *edgeColorPtr
		trx.edgeRing = trx.edgeRing.Next()
	}
	
	trx.faceRing = trx.faceRing.Move(2*direction)
	trx.edgeRing = trx.edgeRing.Move(3*direction)
	
	for i, _ := range ent[cubeId].faceMap[face] {
	        if v,ok := trx.faceRing.Value.(Color); ok {
		    ent[cubeId].faceMap[face][i] = v
		}
		trx.faceRing = trx.faceRing.Next()
	}
	for i, _ := range ent[cubeId].edgeMap[face] {
	        if v,ok := trx.edgeRing.Value.(Color); ok {
		    *ent[cubeId].edgeMap[face][i] = v
		}	
		trx.edgeRing = trx.edgeRing.Next()
	}

        return nil
}

func main() {
	entanglement1,_ := NewEntanglement()
	fmt.Println(entanglement1[0].faceMap["red"][1])
	fmt.Println(entanglement1[0].faceMap["red"][2])
	fmt.Println(*entanglement1[0].edgeMap["red"][2])
	fmt.Println(*entanglement1[0].edgeMap["red"][3])
	fmt.Println(*entanglement1[0].edgeMap["red"][8])
	fmt.Println(*entanglement1[0].edgeMap["red"][11])
	err := ThreeDRotate(entanglement1, 0, "red", 1)
	fmt.Println(err)
	fmt.Println(entanglement1[0].faceMap["red"][1])
	fmt.Println(entanglement1[0].faceMap["red"][2])
	fmt.Println(*entanglement1[0].edgeMap["red"][2])
	fmt.Println(*entanglement1[0].edgeMap["red"][3])
	fmt.Println(*entanglement1[0].edgeMap["red"][8])
	fmt.Println(*entanglement1[0].edgeMap["red"][11])	
