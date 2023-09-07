go
package main

import (
  "fmt"
)

// Define a common interface for all space objects
type SpaceObject interface {
  Move(direction Direction)
  Turn(clockwise bool)
  Shoot() bool
  GetPosition() Position
}

type Position struct {
  X, Y int
}

type Direction int

const (
  Up Direction = iota
  Down
  Left
  Right
)
// Интерфейсы и базовые структуры
// Cоздадим простую архитектуру для движения объектов на игровом поле в соответствии с SOLID принципами. 
// Код, который реализует движение игровых объектов на поле
// Define a basic ship that implements the SpaceObject interface
type Ship struct {
  position  Position
  direction Direction
  alive     bool
}

func (s *Ship) Move(direction Direction) {
  switch direction {
  case Up:
    s.position.Y++
  case Down:
    s.position.Y--
  case Left:
    s.position.X--
  case Right:
    s.position.X++
  }
}

func (s *Ship) Turn(clockwise bool) {
  if clockwise {
    s.direction = (s.direction + 1) % 4
  } else {
    s.direction = (s.direction - 1 + 4) % 4
  }
}

func (s *Ship) Shoot() bool {
  // Simplified shooting mechanism. Can be expanded further.
  return true
}

func (s *Ship) GetPosition() Position {
  return s.position
}
