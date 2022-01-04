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

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestDirectionChangeX(t *testing.T) {
    assert.Equal(t, 0, Direction{dx: -1}.ChangeX(1))
    assert.Equal(t, 2, Direction{dx: 1}.ChangeX(1))
    assert.Equal(t, 1, Direction{dx: 0}.ChangeX(1))
}

func TestDirectionChangeY(t *testing.T) {
    assert.Equal(t, 1, Direction{dy: 0}.ChangeY(1))
    assert.Equal(t, 2, Direction{dy: 1}.ChangeY(1))
    assert.Equal(t, 0, Direction{dy: -1}.ChangeY(1))
}

func TestDirectionInverted(t *testing.T) {
    m := Map{
        Left:  Direction{name: Left},
        Right: Direction{name: Right},
        Up:    Direction{name: Up},
        Down:  Direction{name: Down},
    }
    assert.Equal(t, Direction{name: Right}, m.Inverted(Left))
    assert.Equal(t, Direction{name: Left}, m.Inverted(Right))
    assert.Equal(t, Direction{name: Down}, m.Inverted(Up))
    assert.Equal(t, Direction{name: Up}, m.Inverted(Down))
}

func TestDirectionInvalidInverted(t *testing.T) {
    m := Map{
        Left:  Direction{name: Left},
        Right: Direction{name: Right},
        Up:    Direction{name: Up},
        Down:  Direction{name: Down},
        Stop:  Direction{name: ""},
    }
    assert.Equal(t, Direction{name: ""}, m.Inverted("Act"))
    assert.Equal(t, Direction{name: ""}, m.Inverted(Stop))
}
