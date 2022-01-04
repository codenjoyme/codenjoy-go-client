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
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestValidUrlToWstoken(t *testing.T) {
    tests := []struct {
        name    string
        url     string
        wstoken string
    }{
        {
            name:    "urlToWsToken with valid HTTP URL",
            url:     "http://127.0.0.1:8080/codenjoy-contest/board/player/793wdxskw521spo4mn1y?code=531459153668826800",
            wstoken: "ws://127.0.0.1:8080/codenjoy-contest/ws?user=793wdxskw521spo4mn1y&code=531459153668826800",
        },
        {
            name:    "urlToWsToken with valid HTTPS URL",
            url:     "https://dojorena.io/codenjoy-contest/board/player/793wdxskw521spo4mn1y?code=531459153668826800",
            wstoken: "wss://dojorena.io/codenjoy-contest/ws?user=793wdxskw521spo4mn1y&code=531459153668826800",
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := urlToWsToken(tt.url); got != tt.wstoken {
                t.Errorf("urlToWsToken() = %v, want %v", got, tt.wstoken)
            }
        })
    }
}

func Test_Panic_UrlToWsToken(t *testing.T) {
    tests := []struct {
        name string
        url  string
    }{
        {
            name: "urlToWsToken with unsupported scheme",
            url:  "ftp://dojorena.io/codenjoy-contest/board/player/793wdxskw521spo4mn1y?code=531459153668826800",
        },
        {
            name: "urlToWsToken with invalid host",
            url:  "https://codenjoy-contest/board/player/793wdxskw521spo4mn1y?code=531459153668826800",
        },
        {
            name: "urlToWsToken with playerId that contains non word characters",
            url:  "https://dojorena.io/codenjoy-contest/board/player/7**wdxskw521spo4mn1y?code=531459153668826800",
        },
        {
            name: "urlToWsToken with code that contains non word characters",
            url:  "https://dojorena.io/codenjoy-contest/board/player/793wdxskw521spo4mn1y?code=AA5459153668826800",
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            assert.Panics(t, func() { urlToWsToken(tt.url) })
        })
    }
}
