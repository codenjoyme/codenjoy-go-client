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

package verland

import (
	"fmt"
	"github.com/codenjoyme/codenjoy-go-client/engine"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAtInvalidPoint(t *testing.T) {
	board := newBoard("☼☼☼" + "☼☼☼" + "☼☼☼")
	assert.Equal(t, elements["PATHLESS"], board.getAt(engine.NewPoint(-1, -1)))
}

func TestFindHero(t *testing.T) {
	board := newBoard("☼♥☼" + "☼☼☼" + "☼☼☼")
	assert.Equal(t, "[1,2]", board.findHero().String())

	board = newBoard("☼☼☼" + "☼♥☼" + "☼☼☼")
	assert.Equal(t, "[1,1]", board.findHero().String())

	board = newBoard("☼☼☼" + "☼☼☼" + "☼♥☼")
	assert.Equal(t, "[1,0]", board.findHero().String())

	board = newBoard("☼☼♥" + "☼☼☼" + "☼☼☼")
	assert.Equal(t, "[2,2]", board.findHero().String())
}

func TestFindHeroNoResult(t *testing.T) {
	board := newBoard("☼☼☼" + "☼☼☼" + "☼☼☼")
	assert.Panics(t, func() { board.findHero() })
}

func TestIsGameOver(t *testing.T) {
	board := newBoard("☼☼☼" + "☼☼♥" + "☼☼☼")
	assert.Equal(t, false, board.isGameOver())

	board = newBoard("☼☼☼" + "X☼☼" + "☼☼☼")
	assert.Equal(t, true, board.isGameOver())
}

func TestFindOtherHeroes(t *testing.T) {
	board := newBoard("☼Y☼" + "☼♠☼" + "☼☼☼")
	assert.Equal(t, "[[1,1] [1,2]]", fmt.Sprintf("%v", board.findOtherHeroes()))
}

func TestFindEnemyHeroes(t *testing.T) {
	board := newBoard("☼Z☼" + "☼♣☼" + "☼☼☼")
	assert.Equal(t, "[[1,1] [1,2]]", fmt.Sprintf("%v", board.findEnemyHeroes()))
}

func TestFindWalls(t *testing.T) {
	board := newBoard("***" + "☼**" + "☼**")
	assert.Equal(t, "[[0,0] [0,1]]", fmt.Sprintf("%v", board.findWalls()))
}

func TestBoardCountContagions(t *testing.T) {
	board := newBoard("***" + "***" + "8**")
	assert.Equal(t, 8, board.countContagions(engine.NewPoint(0, 0)))
}
func TestReport(t *testing.T) {
	board := newBoard("board=" +
		"☼☼☼☼☼☼☼☼☼" +
		"☼1 Y   y☼" +
		"☼*2  x  ☼" +
		"☼o 3 ♠ +☼" +
		"☼♥  4   ☼" +
		"☼   Z  ♣☼" +
		"☼z  5678☼" +
		"☼  !  X ☼" +
		"☼☼☼☼☼☼☼☼☼")
	assert.Equal(t,
		"☼☼☼☼☼☼☼☼☼\n"+
			"☼1 Y   y☼\n"+
			"☼*2  x  ☼\n"+
			"☼o 3 ♠ +☼\n"+
			"☼♥  4   ☼\n"+
			"☼   Z  ♣☼\n"+
			"☼z  5678☼\n"+
			"☼  !  X ☼\n"+
			"☼☼☼☼☼☼☼☼☼\n"+
			"\nHero at: [1,4]"+
			"\nOther heroes at: [[3,7] [5,5]]"+
			"\nEnemy heroes at: [[4,3] [7,3]]"+
			"\nOther stuff at: [[0,0] [0,1] [0,2] [0,3] [0,4] [0,5] [0,6] [0,7] [0,8] [1,0] [1,5] [1,6] [1,8] [2,0] [2,8] "+
			"[3,0] [3,8] [4,0] [4,8] [5,0] [5,8] [6,0] [6,8] [7,0] [7,8] [8,0] [8,1] [8,2] [8,3] [8,4] [8,5] [8,6] "+
			"[8,7] [8,8]]", board.String())
}
