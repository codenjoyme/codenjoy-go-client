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

package clifford

import (
    "fmt"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestIsGameOver(t *testing.T) {
    board := newBoard("#####►###")
    assert.Equal(t, false, board.isGameOver())

    board = newBoard("###O#####")
    assert.Equal(t, true, board.isGameOver())

    board = newBoard("###o#####")
    assert.Equal(t, true, board.isGameOver())
}

func TestFindHeroNoResult(t *testing.T) {
    board := newBoard("#########")
    assert.Panics(t, func() { fmt.Sprintf("%v", board.findHero()) })
}

func TestFindHero(t *testing.T) {
    board := newBoard("O########")
    assert.Equal(t, "[0,2]", fmt.Sprintf("%v", board.findHero()))

    board = newBoard("###A#####")
    assert.Equal(t, "[0,1]", fmt.Sprintf("%v", board.findHero()))

    board = newBoard("####◄####")
    assert.Equal(t, "[1,1]", fmt.Sprintf("%v", board.findHero()))

    board = newBoard("#####►###")
    assert.Equal(t, "[2,1]", fmt.Sprintf("%v", board.findHero()))

    board = newBoard("######U##")
    assert.Equal(t, "[0,0]", fmt.Sprintf("%v", board.findHero()))

    board = newBoard("########I")
    assert.Equal(t, "[2,0]", fmt.Sprintf("%v", board.findHero()))

    board = newBoard("########E")
    assert.Equal(t, "[2,0]", fmt.Sprintf("%v", board.findHero()))

    board = newBoard("EO A◄►UI ")
    assert.Equal(t, "[0,0]", fmt.Sprintf("%v", board.findHero()))
}

func TestFindHero_Mask(t *testing.T) {
    board := newBoard("o########")
    assert.Equal(t, "[0,2]", fmt.Sprintf("%v", board.findHero()))

    board = newBoard("###a#####")
    assert.Equal(t, "[0,1]", fmt.Sprintf("%v", board.findHero()))

    board = newBoard("####h####")
    assert.Equal(t, "[1,1]", fmt.Sprintf("%v", board.findHero()))

    board = newBoard("#####w###")
    assert.Equal(t, "[2,1]", fmt.Sprintf("%v", board.findHero()))

    board = newBoard("######u##")
    assert.Equal(t, "[0,0]", fmt.Sprintf("%v", board.findHero()))

    board = newBoard("########i")
    assert.Equal(t, "[2,0]", fmt.Sprintf("%v", board.findHero()))

    board = newBoard("########e")
    assert.Equal(t, "[2,0]", fmt.Sprintf("%v", board.findHero()))

    board = newBoard("eo awhui ")
    assert.Equal(t, "[0,0]", fmt.Sprintf("%v", board.findHero()))
}

func TestFindOtherHeroes(t *testing.T) {
    board := newBoard("C» D« KJF")
    assert.Equal(t, "[[0,0] [0,1] [0,2] [1,0] [1,1] [1,2] [2,0]]",
        fmt.Sprintf("%v", board.findOtherHeroes()))

    board = newBoard("cz dZ kjf")
    assert.Equal(t, "[[0,0] [0,1] [0,2] [1,0] [1,1] [1,2] [2,0]]",
        fmt.Sprintf("%v", board.findOtherHeroes()))
}
func TestFindEnemyHeroes(t *testing.T) {
    board := newBoard("LP NQ RTV")
    assert.Equal(t, "[[0,0] [0,1] [0,2] [1,0] [1,1] [1,2] [2,0]]",
        fmt.Sprintf("%v", board.findEnemyHeroes()))

    board = newBoard("lp nq rtv")
    assert.Equal(t, "[[0,0] [0,1] [0,2] [1,0] [1,1] [1,2] [2,0]]",
        fmt.Sprintf("%v", board.findEnemyHeroes()))
}

func TestFindRobbers(t *testing.T) {
    board := newBoard("X) xY (y ")
    assert.Equal(t, "[[0,0] [0,1] [0,2] [1,0] "+
        "[1,1] [1,2]]", fmt.Sprintf("%v", board.findRobbers()))
}

func TestFindBarriers(t *testing.T) {
    board := newBoard("  #  ☼   ")
    assert.Equal(t, "[[2,1] [2,2]]", fmt.Sprintf("%v", board.findBarriers()))
}

func TestFindPits(t *testing.T) {
    board := newBoard("1234**###")
    assert.Equal(t, "[[0,1] [0,2] [1,1] [1,2] [2,1] [2,2]]", fmt.Sprintf("%v", board.findPits()))
}

func TestFindClues(t *testing.T) {
    board := newBoard("##$##&##@")
    assert.Equal(t, "[[2,0] [2,1] [2,2]]", fmt.Sprintf("%v", board.findClues()))
}

func TestFindBackways(t *testing.T) {
    board := newBoard("##W######")
    assert.Equal(t, "[[2,2]]", fmt.Sprintf("%v", board.findBackways()))
}

func TestFindPotions(t *testing.T) {
    board := newBoard("##m######")
    assert.Equal(t, "[[2,2]]", fmt.Sprintf("%v", board.findPotions()))
}

func TestFindDoors(t *testing.T) {
    board := newBoard("gsbGSB###")
    assert.Equal(t, "[[0,1] [0,2] [1,1] [1,2] [2,1] [2,2]]", fmt.Sprintf("%v", board.findDoors()))
}

func TestFindKeys(t *testing.T) {
    board := newBoard("+-!######")
    assert.Equal(t, "[[0,2] [1,2] [2,2]]", fmt.Sprintf("%v", board.findKeys()))
}

func TestReport(t *testing.T) {
    board := newBoard("board=" +
        "☼☼☼☼☼☼☼☼☼" +
        "☼ ►*## $☼" +
        "☼ H pq -☼" +
        "☼ H  1 G☼" +
        "☼m   &  ☼" +
        "☼ + ~~~ ☼" +
        "☼Z3 S   ☼" +
        "☼ @@  X ☼" +
        "☼☼☼☼☼☼☼☼☼")
    assert.Equal(t, ""+
        "☼☼☼☼☼☼☼☼☼\n"+
        "☼ ►*## $☼\n"+
        "☼ H pq -☼\n"+
        "☼ H  1 G☼\n"+
        "☼m   &  ☼\n"+
        "☼ + ~~~ ☼\n"+
        "☼Z3 S   ☼\n"+
        "☼ @@  X ☼\n"+
        "☼☼☼☼☼☼☼☼☼\n"+
        "\n"+
        "Hero at: [2,7]\n"+
        "Other heroes at: [[1,2]]\n"+
        "Enemy heroes at: [[4,6] [5,6]]\n"+
        "Robbers at: [[6,1]]\n"+
        "Mask potions at: [[1,4]]\n"+
        "Keys at: [[2,3] [7,6]]", fmt.Sprintf("%v", board.String()))
}
