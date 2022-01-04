/*-
 * #%L
 * Codenjoy - it's a dojo-like platform from developers to developers.
 * %%
 * Copyright (C) 2012 - 2022 Codenjoy
 * %%
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public
 * License along with this program.  If not, see
 * <http://www.gnu.org/licenses/gpl-3.0.html>.
 * #L%
 */

package engine

import (
    "fmt"
    "math"
    "sort"
    "strings"
)

type GameBoard struct {
    elements []rune
    len      int
    size     int
}

func NewGameBoard(supportedElements []rune, message string) *GameBoard {
    board := GameBoard{}
    message = strings.Replace(message, "board=", "", 1)
    board.initElementsArray(supportedElements, []rune(message))
    board.len = len(board.elements)
    board.size = int(math.Sqrt(float64(board.len)))
    return &board
}

func (b *GameBoard) initElementsArray(supportedElements []rune, rawBoard []rune) {
    b.elements = make([]rune, len(rawBoard))
    for i := 0; i < len(b.elements); i++ {
        nextElement := rawBoard[i]
        for _, v := range supportedElements {
            if nextElement == v {
                b.elements[i] = nextElement
                break
            }
        }
        if b.elements[i] == 0 {
            panic(fmt.Sprintf("invalid element: %v", nextElement))
        }
    }
}

func (b *GameBoard) AppendIfMissing(slice []*Point, points ...*Point) []*Point {
    for _, p := range points {
        existed := false
        for _, ele := range slice {
            if ele.Equal(p) {
                existed = true
                break
            }
        }
        if !existed {
            slice = append(slice, p)
        }
    }
    return slice
}

func (b *GameBoard) Size() int {
    return b.size
}

func (b *GameBoard) GetAt(pt *Point) rune {
    if !pt.IsValid(b.size) {
        panic(fmt.Sprintf("invalid point %s" + pt.String()))
    }
    return b.elements[b.pointToIndex(pt.X(), pt.Y())]
}

func (b *GameBoard) Find(wanted ...rune) []*Point {
    var points []*Point
    for i, el := range b.elements {
        for _, w := range wanted {
            if w == el {
                points = append(points, b.indexToPoint(i))
            }
        }
    }
    sort.Sort(SortedPoints(points))
    return points
}

func (b *GameBoard) FindFirst(wanted ...rune) *Point {
    for i, el := range b.elements {
        for _, w := range wanted {
            if w == el {
                return b.indexToPoint(i)
            }
        }
    }
    return nil
}

func (b *GameBoard) IsAt(pt *Point, wanted ...rune) bool {
    if !pt.IsValid(b.size) {
        return false
    }
    el := b.GetAt(pt)
    for _, w := range wanted {
        if w == el {
            return true
        }
    }
    return false
}

func (b *GameBoard) FindNear(pt *Point) []rune {
    var elements []rune

    right := NewPoint(pt.X()+1, pt.Y())
    if right.IsValid(b.size) {
        elements = append(elements, b.GetAt(right))
    }
    left := NewPoint(pt.X()-1, pt.Y())
    if left.IsValid(b.size) {
        elements = append(elements, b.GetAt(left))
    }
    up := NewPoint(pt.X(), pt.Y()+1)
    if up.IsValid(b.size) {
        elements = append(elements, b.GetAt(up))
    }
    down := NewPoint(pt.X(), pt.Y()-1)
    if down.IsValid(b.size) {
        elements = append(elements, b.GetAt(down))
    }

    return elements
}

func (b *GameBoard) CountNear(pt *Point, wanted ...rune) int {
    counter := 0
    for _, el := range b.FindNear(pt) {
        for _, w := range wanted {
            if w == el {
                counter++
            }
        }
    }
    return counter
}

func (b *GameBoard) IsNear(pt *Point, wanted ...rune) bool {
    for _, w := range wanted {
        if b.CountNear(pt, w) != 0 {
            return true
        }
    }
    return false
}

func (b *GameBoard) pointToIndex(x int, y int) int {
    return (b.size-1-y)*b.size + x
}

func (b *GameBoard) indexToPoint(index int) *Point {
    x := index % b.size
    y := int(math.Ceil(float64(b.size - 1 - index/b.size)))
    return NewPoint(x, y)
}

func (b *GameBoard) String() string {
    builder := strings.Builder{}
    for y := b.size - 1; y >= 0; y-- {
        for x := 0; x < b.size; x++ {
            builder.WriteString(string(b.elements[b.pointToIndex(x, y)]))
        }
        builder.WriteString("\n")
    }
    return builder.String()
}
