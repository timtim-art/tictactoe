package gameplay

import (
	"fmt"
	"math"
	"math/rand"
)

const (
	battlefieldX = 470
	battlefieldY = 670
)

func CreateWarriorsRandomly() []*Warrior {
	warriors := make([]*Warrior, 10)
	for i := 0; i < 10; i++ {
		symbols := []string{"scissor", "stone", "paper"}
		warrior := &Warrior{
			X:      rand.Intn(battlefieldX),
			Y:      rand.Intn(battlefieldY),
			Symbol: symbols[rand.Intn(3)],
		}
		warriors[i] = warrior
	}
	return warriors
}

func CalcPositions(warriors []*Warrior) {
	for _, warrior := range warriors {
		enemy, err := identifyClosestEnemy(warrior, warriors)
		if err != nil {
			continue
		}
		move, err := decide(warrior, enemy)
		if err != nil {
			continue
		}
		move(warrior, enemy)
	}
}

func Finished(warriors []*Warrior) bool {
	prevSymbol := warriors[0].Symbol
	for i := 1; i < 10; i++ {
		if warriors[i].Symbol != prevSymbol {
			return false
		}
		prevSymbol = warriors[i].Symbol
	}
	return true
}

func identifyClosestEnemy(warrior *Warrior, warriors []*Warrior) (*Warrior, error) {
	var closestEnemy *Warrior
	shortestDistance := -1.
	for _, otherWarrior := range warriors {
		if otherWarrior.Symbol == warrior.Symbol {
			continue
		}

		distance := calcDistance(warrior, otherWarrior)

		if shortestDistance == -1 || distance < shortestDistance {
			closestEnemy = otherWarrior
			shortestDistance = distance
		}
	}

	if closestEnemy == nil {
		return nil, fmt.Errorf("could not identify closest enemy")
	}

	return closestEnemy, nil
}

func calcDistance(warrior, enemy *Warrior) float64 {
	return math.Sqrt(math.Pow(float64(enemy.X-warrior.X), 2.0) + math.Pow(float64(enemy.Y-warrior.Y), 2.0))
}

func decide(warrior, enemy *Warrior) (func(*Warrior, *Warrior), error) {
	switch {
	case warrior.Symbol == "scissor" && enemy.Symbol == "stone":
		return runAway, nil
	case warrior.Symbol == "scissor" && enemy.Symbol == "paper":
		return attack, nil
	case warrior.Symbol == "paper" && enemy.Symbol == "scissor":
		return runAway, nil
	case warrior.Symbol == "paper" && enemy.Symbol == "stone":
		return attack, nil
	case warrior.Symbol == "stone" && enemy.Symbol == "paper":
		return runAway, nil
	case warrior.Symbol == "stone" && enemy.Symbol == "scissor":
		return attack, nil
	}
	return nil, fmt.Errorf("strategy does not exist")
}

func runAway(warrior, enemy *Warrior) {
	if warrior.Y == enemy.Y && warrior.X == enemy.X {
		warrior.Symbol = enemy.Symbol
		return
	}

	if warrior.Y >= enemy.Y && warrior.Y < battlefieldY {
		warrior.Y++
	} else if warrior.Y < enemy.Y && warrior.Y > 0 {
		warrior.Y--
	}

	if warrior.X >= enemy.X && warrior.X < battlefieldX {
		warrior.X++
	} else if warrior.X < enemy.X && warrior.X > 0 {
		warrior.X--
	}
}

func attack(warrior, enemy *Warrior) {
	if warrior.Y == enemy.Y && warrior.X == enemy.X {
		enemy.Symbol = warrior.Symbol
		return
	}

	if warrior.Y > enemy.Y && warrior.Y > 0 {
		warrior.Y--
	} else if warrior.Y < enemy.Y && warrior.Y < battlefieldY {
		warrior.Y++
	}

	if warrior.X > enemy.X && warrior.X > 0 {
		warrior.X--
	} else if warrior.X < enemy.X && warrior.X < battlefieldX {
		warrior.X++
	}
}
