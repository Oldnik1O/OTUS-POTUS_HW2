\\test_2.4
\\ Проверка поворотов корабля

package main

import (
  "testing"
)

func TestShipTurn(t *testing.T) {
  ship := &Ship{position: Position{X: 5, Y: 5}, direction: Up}
  ship.Turn(true)

  if ship.direction != Right {
    t.Errorf("Ship did not turn correctly clockwise.")
  }

  ship.Turn(false)

  if ship.direction != Up {
    t.Errorf("Ship did not turn correctly counter-clockwise.")
  }
}