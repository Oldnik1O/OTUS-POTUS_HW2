\\test_2.3
\\ Проверка движения корабля

package main

import (
  "testing"
)

func TestMoveShip(t *testing.T) {
  server := NewGameServer(10, 10)
  ship := &Ship{position: Position{X: 5, Y: 5}, direction: Up}

  server.AddShip(ship, 1)
  server.MoveShip(ship, Up)

  if server.field[5][5] != nil || server.field[6][5] != ship {
    t.Errorf("Ship was not moved correctly.")
  }
}