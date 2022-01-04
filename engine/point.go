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
)

type Point struct {
    x, y int
}

func NewPoint(x int, y int) *Point {
    return &Point{x, y}
}

func (p *Point) X() int {
    return p.x
}

func (p *Point) Y() int {
    return p.y
}

func (p *Point) IsValid(boardSize int) bool {
    return (p.x >= 0 && p.x < boardSize) && (p.y >= 0 && p.y < boardSize)
}

func (p *Point) String() string {
    return fmt.Sprintf("[%d,%d]", p.x, p.y)
}

func StepRight(pt *Point) *Point {
    return NewPoint(pt.X()+1, pt.Y())
}

func StepLeft(pt *Point) *Point {
    return NewPoint(pt.X()-1, pt.Y())
}

func StepUp(pt *Point) *Point {
    return NewPoint(pt.X(), pt.Y()+1)
}

func StepDown(pt *Point) *Point {
    return NewPoint(pt.X(), pt.Y()-1)
}

func (lhs *Point) Equal(rhs *Point) bool {
    return lhs.X() == rhs.X() && lhs.Y() == rhs.Y()
}

type SortedPoints []*Point

func (p SortedPoints) Len() int           { return len(p) }
func (p SortedPoints) Less(i, j int) bool { return p[i].String() < p[j].String() }
func (p SortedPoints) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
