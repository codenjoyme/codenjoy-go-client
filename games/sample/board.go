/*-
 * #%L
 * Codenjoy - it's a dojo-like platform from developers to developers.
 * %%
 * Copyright (C) 2021 Codenjoy
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

package sample

import (
	"fmt"
	"sort"

	"github.com/codenjoyme/codenjoy-go-client/engine"
)

type board struct {
	board *engine.GameBoard
}

func newBoard(message string) *board {
	values := make([]rune, 0, len(elements))
	for _, e := range elements {
		values = append(values, e)
	}
	return &board{engine.NewGameBoard(values, message)}
}

func (b *board) getAt(pt *engine.Point) rune {
	if !pt.IsValid(b.board.Size()) {
		return elements["WALL"]
	}
	return b.board.GetAt(pt)
}

func (b *board) findHero() *engine.Point {
	points := b.board.Find(
		elements["HERO"],
		elements["DEAD_HERO"])

	if len(points) == 0 {
		panic("Hero element has not been found")
	}
	return points[0]
}

func (b *board) isGameOver() bool {
	return len(b.board.Find(elements["DEAD_HERO"])) != 0
}

func (b *board) findOtherHeroes() []*engine.Point {
	return b.board.Find(
		elements["OTHER_HERO"],
		elements["OTHER_DEAD_HERO"])
}

func (b *board) findBarriers() []*engine.Point {
	var points []*engine.Point
	points = appendIfMissing(points, b.findWalls()...)
	points = appendIfMissing(points, b.findBombs()...)
	points = appendIfMissing(points, b.findOtherHeroes()...)
	sort.Sort(engine.SortedPoints(points))
	return points
}

func appendIfMissing(slice []*engine.Point, points ...*engine.Point) []*engine.Point {
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

func (b *board) findWalls() []*engine.Point {
	return b.board.Find(elements["WALL"])
}

func (b *board) findBombs() []*engine.Point {
	return b.board.Find(elements["BOMB"])
}

func (b *board) findGold() []*engine.Point {
	return b.board.Find(elements["GOLD"])
}

func (b *board) String() string {
	return b.board.String() +
		"\nHero at: " + b.findHero().String() +
		"\nOther heroes at: " + fmt.Sprintf("%v", b.findOtherHeroes()) +
		"\nBombs at: " + fmt.Sprintf("%v", b.findBombs()) +
		"\nGold at: " + fmt.Sprintf("%v", b.findGold())
}
