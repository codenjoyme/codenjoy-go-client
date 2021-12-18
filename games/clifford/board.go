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

func (b *board) isGameOver() bool {
    return len(b.board.Find(elements["HERO_DIE"], elements["HERO_MASK_DIE"])) != 0
}

func (b *board) findHero() *engine.Point {
    points := b.board.Find(
        elements["HERO_DIE"],
        elements["HERO_LADDER"],
        elements["HERO_LEFT"],
        elements["HERO_RIGHT"],
        elements["HERO_FALL"],
        elements["HERO_PIPE"],
        elements["HERO_PIT"],

        elements["HERO_MASK_DIE"],
        elements["HERO_MASK_LADDER"],
        elements["HERO_MASK_LEFT"],
        elements["HERO_MASK_RIGHT"],
        elements["HERO_MASK_FALL"],
        elements["HERO_MASK_PIPE"],
        elements["HERO_MASK_PIT"])

    if len(points) == 0 {
        panic("hero element has not been found")
    }
    return points[0]
}

func (b *board) findOtherHeroes() []*engine.Point {
    return b.board.Find(
        elements["OTHER_HERO_DIE"],
        elements["OTHER_HERO_LADDER"],
        elements["OTHER_HERO_LEFT"],
        elements["OTHER_HERO_RIGHT"],
        elements["OTHER_HERO_FALL"],
        elements["OTHER_HERO_PIPE"],
        elements["OTHER_HERO_PIT"],

        elements["OTHER_HERO_MASK_DIE"],
        elements["OTHER_HERO_MASK_LADDER"],
        elements["OTHER_HERO_MASK_LEFT"],
        elements["OTHER_HERO_MASK_RIGHT"],
        elements["OTHER_HERO_MASK_FALL"],
        elements["OTHER_HERO_MASK_PIPE"],
        elements["OTHER_HERO_MASK_PIT"])
}

func (b *board) findEnemyHeroes() []*engine.Point {
    return b.board.Find(
        elements["ENEMY_HERO_DIE"],
        elements["ENEMY_HERO_LADDER"],
        elements["ENEMY_HERO_LEFT"],
        elements["ENEMY_HERO_RIGHT"],
        elements["ENEMY_HERO_FALL"],
        elements["ENEMY_HERO_PIPE"],
        elements["ENEMY_HERO_PIT"],

        elements["ENEMY_HERO_MASK_DIE"],
        elements["ENEMY_HERO_MASK_LADDER"],
        elements["ENEMY_HERO_MASK_LEFT"],
        elements["ENEMY_HERO_MASK_RIGHT"],
        elements["ENEMY_HERO_MASK_FALL"],
        elements["ENEMY_HERO_MASK_PIPE"],
        elements["ENEMY_HERO_MASK_PIT"])
}

func (b *board) findRobbers() []*engine.Point {
    return b.board.Find(
        elements["ROBBER_LADDER"],
        elements["ROBBER_LEFT"],
        elements["ROBBER_RIGHT"],
        elements["ROBBER_FALL"],
        elements["ROBBER_PIPE"],
        elements["ROBBER_PIT"])
}

func (b *board) findBarriers() []*engine.Point {
    return b.board.Find(
        elements["BRICK"],
        elements["STONE"])
}

func (b *board) findPits() []*engine.Point {
    return b.board.Find(
        elements["CRACK_PIT"],
        elements["PIT_FILL_1"],
        elements["PIT_FILL_2"],
        elements["PIT_FILL_3"],
        elements["PIT_FILL_4"])
}

func (b *board) findClues() []*engine.Point {
    return b.board.Find(
        elements["CLUE_KNIFE"],
        elements["CLUE_GLOVE"],
        elements["CLUE_RING"])
}

func (b *board) findBackways() []*engine.Point {
    return b.board.Find(elements["BACKWAY"])
}

func (b *board) findPotions() []*engine.Point {
    return b.board.Find(elements["MASK_POTION"])
}

func (b *board) findDoors() []*engine.Point {
    return b.board.Find(
        elements["OPENED_DOOR_GOLD"],
        elements["OPENED_DOOR_SILVER"],
        elements["OPENED_DOOR_BRONZE"],
        elements["CLOSED_DOOR_GOLD"],
        elements["CLOSED_DOOR_SILVER"],
        elements["CLOSED_DOOR_BRONZE"])
}

func (b *board) findKeys() []*engine.Point {
    return b.board.Find(
        elements["KEY_GOLD"],
        elements["KEY_SILVER"],
        elements["KEY_BRONZE"])
}

func (b *board) String() string {
    return b.board.String() +
        "\nHero at: " + b.findHero().String() +
        "\nOther heroes at: " + fmt.Sprintf("%v", b.findOtherHeroes()) +
        "\nEnemy heroes at: " + fmt.Sprintf("%v", b.findEnemyHeroes()) +
        "\nRobbers at: " + fmt.Sprintf("%v", b.findRobbers()) +
        "\nMask potions at: " + fmt.Sprintf("%v", b.findPotions()) +
        "\nKeys at: " + fmt.Sprintf("%v", b.findKeys())
}
