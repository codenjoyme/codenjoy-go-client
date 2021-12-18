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
    "github.com/codenjoyme/codenjoy-go-client/engine/direction"
)

type Solver struct {
    directions direction.Map
}

func NewSolver() (engine.Solver, error) {
    d, err := directions()
    if err != nil {
        return nil, err
    }
    return Solver{
        directions: d,
    }, nil
}

func (s Solver) Answer(message string) string {
    board := newBoard(message)
    fmt.Println("Board \n" + board.String())
    action := s.nextAction(board)
    fmt.Println("\nAnswer: " + action.String())
    fmt.Println("-------------------------------------------------------------")
    return action.String()
}

func (s Solver) nextAction(b *board) direction.Direction {
    // TODO: write your code here
    return s.directions.Get(act)
}
