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
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestEmptyMessage(t *testing.T) {
    board := NewGameBoard([]rune{'a', 'b', 'c'}, "")
    assert.Equal(t, "", board.String())
}

func TestEmptySupportedElements(t *testing.T) {
    assert.Panics(t, func() { NewGameBoard([]rune{}, "aaa"+"bbb"+"ccc") })
}

func TestValidMessageAndSupportedElements(t *testing.T) {
    board := NewGameBoard([]rune{'a', 'b', 'c'}, "aaa"+"bbb"+"ccc")
    assert.Equal(t, "aaa\nbbb\nccc\n", board.String())
}

func TestEraseMessagePrefix(t *testing.T) {
    board := NewGameBoard([]rune{'a', 'b', 'c'}, "board="+"aaa"+"bbb"+"ccc")
    assert.Equal(t, "aaa\nbbb\nccc\n", board.String())
}

func TestMessageWithUnsupportedElements(t *testing.T) {
    assert.Panics(t, func() { NewGameBoard([]rune{'a', 'b', 'c'}, "ab8c") })
}

func TestGetSize(t *testing.T) {
    board := NewGameBoard([]rune{'a', 'b', 'c'}, "aaa"+"bbb"+"ccc")
    assert.Equal(t, 3, board.Size())
}

func TestGetAt(t *testing.T) {
    board := NewGameBoard([]rune{'a', 'b', 'c'}, "aaa"+"bbb"+"ccc")
    assert.Equal(t, 'c', board.GetAt(NewPoint(0, 0)))
    assert.Equal(t, 'c', board.GetAt(NewPoint(1, 0)))
    assert.Equal(t, 'c', board.GetAt(NewPoint(2, 0)))
    assert.Equal(t, 'b', board.GetAt(NewPoint(0, 1)))
    assert.Equal(t, 'b', board.GetAt(NewPoint(1, 1)))
    assert.Equal(t, 'b', board.GetAt(NewPoint(2, 1)))
    assert.Equal(t, 'a', board.GetAt(NewPoint(0, 2)))
    assert.Equal(t, 'a', board.GetAt(NewPoint(1, 2)))
    assert.Equal(t, 'a', board.GetAt(NewPoint(2, 2)))
}

func TestGetAtInvalidPoint(t *testing.T) {
    board := NewGameBoard([]rune{'a', 'b', 'c'}, "aaa"+"bbb"+"ccc")
    assert.Panics(t, func() { board.GetAt(NewPoint(10, 10)) })
}

func TestFind(t *testing.T) {
    board := NewGameBoard([]rune{'a', 'b', 'c'}, "aaa"+"bbb"+"ccc")
    assert.Equal(t, "[[0,2] [1,2] [2,2]]", fmt.Sprintf("%v",
        board.Find('a')))
    assert.Equal(t, "[[0,0] [0,1] [1,0] [1,1] [2,0] [2,1]]", fmt.Sprintf("%v",
        board.Find('b', 'c')))
}

func TestFindNotExistedElement(t *testing.T) {
    board := NewGameBoard([]rune{'a', 'b', 'c'}, "aaa"+"bbb"+"ccc")
    assert.Equal(t, "[]", fmt.Sprintf("%v", board.Find('d')))
}

func TestFindFirst(t *testing.T) {
    board := NewGameBoard([]rune{'a', 'b', 'c'}, "aaa"+"bbb"+"ccc")
    assert.Equal(t, "[0,0]", fmt.Sprintf("%v", board.FindFirst('c')))
    assert.Equal(t, "[0,1]", fmt.Sprintf("%v", board.FindFirst('b', 'c')))
    assert.Equal(t, "[0,1]", fmt.Sprintf("%v", board.FindFirst('c', 'b')))
}

func TestFindFirstNotExistedElement(t *testing.T) {
    board := NewGameBoard([]rune{'a', 'b', 'c'}, "aaa"+"bbb"+"ccc")
    assert.Equal(t, "<nil>", fmt.Sprintf("%v", board.FindFirst('d')))
}

func TestIsAt(t *testing.T) {
    board := NewGameBoard([]rune{'a', 'b', 'c'}, "aaa"+"bbb"+"ccc")
    assert.Equal(t, true, board.IsAt(NewPoint(1, 2), 'a'))
    assert.Equal(t, false, board.IsAt(NewPoint(1, 2), 'b'))
    assert.Equal(t, false, board.IsAt(NewPoint(1, 2), 'c'))
}

func TestIsAtInvalidPoint(t *testing.T) {
    board := NewGameBoard([]rune{'a', 'b', 'c'}, "aaa"+"bbb"+"ccc")
    assert.Equal(t, false, board.IsAt(NewPoint(10, 10), 'b'))
}

func TestFindNear(t *testing.T) {
    board := NewGameBoard([]rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i'}, "abc"+"def"+"ghi")
    assert.Equal(t, []rune{'f', 'd', 'b', 'h'}, board.FindNear(NewPoint(1, 1)))
}

func TestFindNearInvalidPoint(t *testing.T) {
    board := NewGameBoard([]rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i'}, "abc"+"def"+"ghi")
    assert.Equal(t, []rune(nil), board.FindNear(NewPoint(-1, -1)))
}

func TestCountNear(t *testing.T) {
    board := NewGameBoard([]rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i'}, "abc"+"def"+"ghi")
    assert.Equal(t, 2, board.CountNear(NewPoint(1, 1), 'a', 'b', 'c', 'd'))
}

func TestCountNearInvalidPoint(t *testing.T) {
    board := NewGameBoard([]rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i'}, "abc"+"def"+"ghi")
    assert.Equal(t, 0, board.CountNear(NewPoint(-1, -1), 'a', 'b', 'c', 'd'))
}

func TestCountNearNotExistedElement(t *testing.T) {
    board := NewGameBoard([]rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i'}, "abc"+"def"+"ghi")
    assert.Equal(t, 0, board.CountNear(NewPoint(1, 1), 'r'))
    assert.Equal(t, 0, board.CountNear(NewPoint(1, 1), 'x', 'y', 'z'))
}

func TestIsNear(t *testing.T) {
    board := NewGameBoard([]rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i'}, "abc"+"def"+"ghi")
    assert.Equal(t, false, board.IsNear(NewPoint(1, 1), 'a'))
    assert.Equal(t, true, board.IsNear(NewPoint(1, 1), 'b'))
    assert.Equal(t, true, board.IsNear(NewPoint(1, 1), 'c', 'd'))
}

func TestIsNearInvalidPoint(t *testing.T) {
    board := NewGameBoard([]rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i'}, "abc"+"def"+"ghi")
    assert.Equal(t, false, board.IsNear(NewPoint(-1, -1), 'a'))
}

func TestIsNearNotExistedElement(t *testing.T) {
    board := NewGameBoard([]rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i'}, "abc"+"def"+"ghi")
    assert.Equal(t, false, board.IsNear(NewPoint(1, 1), 'r'))
    assert.Equal(t, false, board.IsNear(NewPoint(1, 1), 'x', 'y', 'z'))
}
