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
    "crypto/tls"
    "fmt"
    "log"
    "regexp"
    "strings"

    "github.com/gorilla/websocket"
)

const urlRegex = "(?P<scheme>http|https)://(?P<host>.+)/codenjoy-contest/board/player/(?P<player>\\w+)\\?code=(?P<code>\\d+)"

type WebSocketRunner struct {
    token string
}

func NewWebSocketRunner(url string) *WebSocketRunner {
    return &WebSocketRunner{urlToWsToken(url)}
}

func urlToWsToken(url string) string {
    r := regexp.MustCompile(urlRegex)
    params := r.FindStringSubmatch(url)

    scheme := params[r.SubexpIndex("scheme")]
    if strings.EqualFold(scheme, "https") {
        scheme = "wss"
    } else {
        scheme = "ws"
    }
    return fmt.Sprintf("%s://%s/codenjoy-contest/ws?user=%s&code=%s",
        scheme,
        params[r.SubexpIndex("host")],
        params[r.SubexpIndex("player")],
        params[r.SubexpIndex("code")])
}

func (r *WebSocketRunner) Run(solver Solver) {
    ws := websocket.Dialer{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
    connection, _, err := ws.Dial(r.token, nil)
    if err != nil {
        log.Println("unable to establish websocket connection, error: ", err)
        return
    }
    log.Println("websocket connection successfully established")

    for {
        _, msgFromServer, err := connection.ReadMessage()
        if err != nil {
            log.Println("websocket ReadMessage error: ", err)
            return
        }
        msgToServer := solver.Answer(string(msgFromServer))
        err = connection.WriteMessage(websocket.TextMessage, []byte(msgToServer))
        if err != nil {
            log.Println("websocket WriteMessage error: ", err)
            return
        }
    }
}
