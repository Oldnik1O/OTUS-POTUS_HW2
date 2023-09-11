\\test_2.2
\\ Проверка добавления корабля на сервер
package main

import (
  "testing"
)
func TestAddShip(t *testing.T) {
  server := NewGameServer(10, 10)
  ship := &Ship{position: Position{X: 5, Y: 5}, direction: Up}

  server.AddShip(ship, 1)

  if server.field[5][5] != ship {
    t.Errorf("Ship was not added at the correct position")
  }
}