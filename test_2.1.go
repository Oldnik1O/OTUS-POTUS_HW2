\\test_2.1
\\ Проверка корректности инициализации сервера
package main

import (
  "testing"
)

func TestNewGameServer(t *testing.T) {
  server := NewGameServer(10, 10)

  if len(server.field) != 10 || len(server.field[0]) != 10 {
    t.Errorf("GameServer was not initialized with the correct dimensions")
  }
}