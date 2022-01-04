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

package direction

import "errors"

type Direction struct {
    name Base
    dx   int
    dy   int
}

type Base string

type Map map[Base]Direction

const (
    Left  Base = "LEFT"
    Right Base = "RIGHT"
    Up    Base = "UP"
    Down  Base = "DOWN"
    Stop  Base = "STOP"
)

func NewMap(directions ...Direction) (Map, error) {
    m := make(Map, len(directions))
    var n Base
    for _, d := range directions {
        n = d.name
        if d.name == Stop {
            d.name = ""
        }
        m[n] = d
    }
    if m.valid() {
        return m, nil
    }
    return nil, errors.New("map doesn't contain all required basic directions: Left, Right, Up, Down, Stop")
}

func New(dx, dy int, name Base) Direction {
    return Direction{
        dx:   dx,
        dy:   dy,
        name: name,
    }
}

func (d Direction) ChangeX(x int) int {
    return x + d.dx
}

func (d Direction) ChangeY(y int) int {
    return y + d.dy
}

func (m Map) Get(name Base) Direction {
    if m == nil {
        return Direction{}
    }
    return m[name]
}

func (m Map) Inverted(name Base) Direction {
    if m == nil {
        return Direction{}
    }
    switch name {
    case Left:
        return m[Right]
    case Right:
        return m[Left]
    case Down:
        return m[Up]
    case Up:
        return m[Down]
    default:
        return m[Stop]
    }
}

func (m Map) Clockwise(name Base) Direction {
    if m == nil {
        return Direction{}
    }
    switch name {
    case Up:
        return m[Left]
    case Left:
        return m[Down]
    case Down:
        return m[Right]
    case Right:
        return m[Up]
    default:
        return m[Stop]
    }
}

func (m Map) ContrClockwise(name Base) Direction {
    if m == nil {
        return Direction{}
    }
    switch name {
    case Up:
        return m[Right]
    case Right:
        return m[Down]
    case Down:
        return m[Left]
    case Left:
        return m[Up]
    default:
        return m[Stop]
    }
}

func (m Map) MirrorTopBottom(name Base) Direction {
    if m == nil {
        return Direction{}
    }
    switch name {
    case Up:
        return m[Left]
    case Right:
        return m[Down]
    case Down:
        return m[Right]
    case Left:
        return m[Up]
    default:
        return m[Stop]
    }
}

func (m Map) MirrorBottomTop(name Base) Direction {
    if m == nil {
        return Direction{}
    }
    switch name {
    case Up:
        return m[Right]
    case Right:
        return m[Up]
    case Down:
        return m[Left]
    case Left:
        return m[Down]
    default:
        return m[Stop]
    }
}

func (m Map) valid() bool {
    return m[Left].name != "" && m[Right].name != "" && m[Up].name != "" && m[Down].name != ""
}

func (d Direction) String() string {
    return string(d.name)
}
