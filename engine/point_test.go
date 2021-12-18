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

package engine

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestPointsIsValid(t *testing.T) {
    t.Run("valid points", func(t *testing.T) {
        assert.Equal(t, true, NewPoint(0, 0).IsValid(10))
        assert.Equal(t, true, NewPoint(5, 5).IsValid(10))
        assert.Equal(t, true, NewPoint(9, 9).IsValid(10))
        assert.Equal(t, true, NewPoint(0, 9).IsValid(10))
        assert.Equal(t, true, NewPoint(9, 0).IsValid(10))
    })
    t.Run("invalid points", func(t *testing.T) {
        assert.Equal(t, false, NewPoint(-1, 9).IsValid(10))
        assert.Equal(t, false, NewPoint(9, -1).IsValid(10))
        assert.Equal(t, false, NewPoint(11, 9).IsValid(10))
        assert.Equal(t, false, NewPoint(9, 11).IsValid(10))
    })
}
