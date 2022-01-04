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

package verland

import (
    "fmt"
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
        return elements["PATHLESS"]
    }
    return b.board.GetAt(pt)
}

func (b *board) findHero() *engine.Point {
    points := b.board.Find(
        elements["HERO"],
        elements["HERO_DEAD"])
    if len(points) == 0 {
        panic("Hero element has not been found")
    }
    return points[0]
}

func (b *board) isGameOver() bool {
    return len(b.board.Find(elements["HERO_DEAD"])) != 0
}

func (b *board) isWin() bool {
    return !b.isGameOver() && len(b.board.Find(elements["HERO_HEALING"])) != 0
}

func (b *board) findOtherHeroes() []*engine.Point {
    return b.board.Find(
        elements["OTHER_HERO_DEAD"],
        elements["OTHER_HERO"])
}

func (b *board) findEnemyHeroes() []*engine.Point {
    return b.board.Find(
        elements["ENEMY_HERO_DEAD"],
        elements["ENEMY_HERO"])
}

func (b *board) findWalls() []*engine.Point {
    return b.board.Find(elements["PATHLESS"])
}

func (b *board) findOtherStuff() []*engine.Point {
    return b.board.Find(
        elements["INFECTION"],
        elements["HIDDEN"],
        elements["PATHLESS"])
}

func (b *board) countContagions(pt *engine.Point) int {
    if b.board.IsAt(pt,
        elements["CLEAR"],
        elements["CONTAGION_ONE"],
        elements["CONTAGION_TWO"],
        elements["CONTAGION_THREE"],
        elements["CONTAGION_FOUR"],
        elements["CONTAGION_FIVE"],
        elements["CONTAGION_SIX"],
        elements["CONTAGION_SEVEN"],
        elements["CONTAGION_EIGHT"]) {
        return int(b.board.GetAt(pt) - '0')
    }
    return 0
}

func (b *board) String() string {
    return b.board.String() +
        "\nHero at: " + b.findHero().String() +
        "\nOther heroes at: " + fmt.Sprintf("%v", b.findOtherHeroes()) +
        "\nEnemy heroes at: " + fmt.Sprintf("%v", b.findEnemyHeroes()) +
        "\nOther stuff at: " + fmt.Sprintf("%v", b.findOtherStuff())
}
