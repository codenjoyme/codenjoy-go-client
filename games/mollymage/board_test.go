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

package mollymage

import (
    "fmt"
    "github.com/codenjoyme/codenjoy-go-client/engine"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestGetAtInvalidPoint(t *testing.T) {
    board := newBoard("###" + "###" + "###")
    assert.Equal(t, elements["WALL"], board.getAt(engine.NewPoint(-1, -1)))
}

func TestFindHero(t *testing.T) {
    board := newBoard("#☺#" + "###" + "###")
    assert.Equal(t, "[1,2]", board.findHero().String())

    board = newBoard("###" + "#☻#" + "###")
    assert.Equal(t, "[1,1]", board.findHero().String())

    board = newBoard("###" + "###" + "#Ѡ#")
    assert.Equal(t, "[1,0]", board.findHero().String())

    board = newBoard("Ѡ☺☻" + "###" + "###")
    assert.Equal(t, "[0,2]", board.findHero().String())
}

func TestFindHeroNoResult(t *testing.T) {
    board := newBoard("###" + "###" + "###")
    assert.Panics(t, func() { board.findHero() })
}

func TestIsGameOver(t *testing.T) {
    board := newBoard("###" + "##☺" + "###")
    assert.Equal(t, false, board.isGameOver())

    board = newBoard("###" + "Ѡ##" + "###")
    assert.Equal(t, true, board.isGameOver())
}

func TestFindOtherHeroes(t *testing.T) {
    board := newBoard("#♥#" + "#♠#" + "#♣#")
    assert.Equal(t, "[[1,0] [1,1] [1,2]]", fmt.Sprintf("%v", board.findOtherHeroes()))
}

func TestFindEnemyHeroes(t *testing.T) {
    board := newBoard("#ö#" + "#Ö#" + "#ø#")
    assert.Equal(t, "[[1,0] [1,1] [1,2]]", fmt.Sprintf("%v", board.findEnemyHeroes()))
}

func TestFindBarriers(t *testing.T) {
    board := newBoard("☼&#" + "123" + "♥♠♣")
    assert.Equal(t, "[[0,0] [0,1] [0,2] [1,0] [1,1] [1,2] [2,0] [2,1] [2,2]]",
        fmt.Sprintf("%v", board.findBarriers()))
}

func TestFindWalls(t *testing.T) {
    board := newBoard("###" + "☼##" + "☼##")
    assert.Equal(t, "[[0,0] [0,1]]", fmt.Sprintf("%v", board.findWalls()))
}

func TestFindGhosts(t *testing.T) {
    board := newBoard("##&" + "##&" + "###")
    assert.Equal(t, "[[2,1] [2,2]]", fmt.Sprintf("%v", board.findGhosts()))
}

func TestFindTreasureBoxes(t *testing.T) {
    board := newBoard("҉#҉" + "҉҉҉" + "҉#҉")
    assert.Equal(t, "[[1,0] [1,2]]", fmt.Sprintf("%v", board.findTreasureBoxes()))
}

func TestFindPotions(t *testing.T) {
    board := newBoard("123" + "45#" + "☻♠#")
    assert.Equal(t, "[[0,0] [0,1] [0,2] [1,0] [1,1] [1,2] [2,2]]",
        fmt.Sprintf("%v", board.findPotions()))
}

func TestFindBlasts(t *testing.T) {
    board := newBoard("###" + "###" + "##҉")
    assert.Equal(t, "[[2,0]]", fmt.Sprintf("%v", board.findBlasts()))
}

func TestFindPerks(t *testing.T) {
    board := newBoard("#cr" + "#i+" + "#TA")
    assert.Equal(t, "[[1,0] [1,1] [1,2] [2,0] [2,1] [2,2]]", fmt.Sprintf("%v", board.findPerks()))
}

func TestReport(t *testing.T) {
    board := newBoard("board=" +
        "☼☼☼☼☼☼☼☼☼" +
        "☼1 ♣   ♠☼" +
        "☼#2  &  ☼" +
        "☼# 3 ♣ ♠☼" +
        "☼☺  4   ☼" +
        "☼   ö H☻☼" +
        "☼x H ҉҉҉☼" +
        "☼& &    ☼" +
        "☼☼☼☼☼☼☼☼☼")
    assert.Equal(t, "☼☼☼☼☼☼☼☼☼\n"+
        "☼1 ♣   ♠☼\n"+
        "☼#2  &  ☼\n"+
        "☼# 3 ♣ ♠☼\n"+
        "☼☺  4   ☼\n"+
        "☼   ö H☻☼\n"+
        "☼x H ҉҉҉☼\n"+
        "☼& &    ☼\n"+
        "☼☼☼☼☼☼☼☼☼\n"+
        "\n"+
        "Hero at: [1,4]\n"+
        "Other heroes at: [[3,7] [5,5] [7,5] [7,7]]\n"+
        "Enemy heroes at: [[4,3]]\n"+
        "Ghosts at: [[1,1] [3,1] [5,6]]\n"+
        "Potions at: [[1,7] [2,6] [3,5] [4,4] [7,3] [7,5] [7,7]]\n"+
        "Blasts at: [[5,2] [6,2] [7,2]]\n"+
        "Expected blasts at: [[2,7]]", board.String())
}
