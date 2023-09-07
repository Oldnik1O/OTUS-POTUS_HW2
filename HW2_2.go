// Игровой сервер
// Базовые механизмы движения кораблей на игровом поле
go
type GameServer struct {
  field  [][]SpaceObject
  ships1 []*Ship
  ships2 []*Ship
}

func NewGameServer(width, height int) *GameServer {
  field := make([][]SpaceObject, height)
  for i := range field {
    field[i] = make([]SpaceObject, width)
  }

  return &GameServer{
    field: field,
  }
}

func (gs *GameServer) AddShip(s *Ship, team int) {
  if team == 1 {
    gs.ships1 = append(gs.ships1, s)
  } else {
    gs.ships2 = append(gs.ships2, s)
  }
  gs.field[s.position.Y][s.position.X] = s
}

func (gs *GameServer) MoveShip(s *Ship, direction Direction) {
  // Remove ship from old position
  gs.field[s.position.Y][s.position.X] = nil

  // Update the ship's position
  s.Move(direction)

  // Place ship at new position
  gs.field[s.position.Y][s.position.X] = s
}

func main() {
  server := NewGameServer(10, 10)

  // Add ships for demonstration
  ship1 := &Ship{position: Position{X: 5, Y: 5}, direction: Up}
  server.AddShip(ship1, 1)

  ship2 := &Ship{position: Position{X: 4, Y: 4}, direction: Down}
  server.AddShip(ship2, 2)

  fmt.Println("Ship1 initial position:", ship1.GetPosition())
  server.MoveShip(ship1, Up)
  fmt.Println("Ship1 new position:", ship1.GetPosition())

  fmt.Println("Ship2 initial position:", ship2.GetPosition())
  server.MoveShip(ship2, Down)
  fmt.Println("Ship2 new position:", ship2.GetPosition())
}

