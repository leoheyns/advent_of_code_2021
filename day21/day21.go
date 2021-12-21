package main

import "fmt"

type gamestate struct {
	p1_pos   int
	p2_pos   int
	p1_score int
	p2_score int
	turn_p1  bool
}

type result struct {
	p1_wins int
	p2_wins int
}

func calcResult(gs gamestate, cache *map[gamestate]result) result {
	res := result{
		p1_wins: 0,
		p2_wins: 0,
	}
	for d1 := 1; d1 < 4; d1++ {
		for d2 := 1; d2 < 4; d2++ {
			for d3 := 1; d3 < 4; d3++ {
				throw := d1 + d2 + d3
				new_gamestate := gs

				if gs.turn_p1 {
					new_gamestate.p1_pos = (gs.p1_pos + throw) % 10
					new_gamestate.p1_score += new_gamestate.p1_pos + 1
				} else {
					new_gamestate.p2_pos = (gs.p2_pos + throw) % 10
					new_gamestate.p2_score += new_gamestate.p2_pos + 1
				}
				if new_gamestate.p1_score >= 21 {
					res.p1_wins++
				} else if new_gamestate.p2_score >= 21 {
					res.p2_wins++
				} else {
					new_gamestate.turn_p1 = !gs.turn_p1
					new_res, ok := (*cache)[new_gamestate]
					if !ok {
						new_res = calcResult(new_gamestate, cache)
					}
					res.p1_wins += new_res.p1_wins
					res.p2_wins += new_res.p2_wins
				}
			}
		}
	}
	(*cache)[gs] = res
	return res
}

func main() {
	p1_pos := 7
	p2_pos := 3

	p1_score := 0
	p2_score := 0

	dice := 0
	dice_count := 0
	for p1_score <= 1000 && p2_score <= 1000 {
		for i := 0; i < 3; i++ {
			dice = (dice % 100) + 1
			p1_pos = (p1_pos + dice) % 10
			dice_count++
		}
		p1_score += p1_pos + 1
		//fmt.Println("p1: ", p1_pos, p1_score)
		if p1_score >= 1000 {
			fmt.Println(p2_score * dice_count)
			break
		}
		for i := 0; i < 3; i++ {
			dice = (dice % 100) + 1
			p2_pos = (p2_pos + dice) % 10
			dice_count++
		}
		p2_score += p2_pos + 1
		//fmt.Println("p2: ", p2_pos, p2_score)

		if p2_score >= 1000 {
			fmt.Println(p1_score * dice_count)
		}

	}
	cache := map[gamestate]result{}

	start := gamestate{
		p1_pos:   7,
		p2_pos:   3,
		p1_score: 0,
		p2_score: 0,
		turn_p1:  true,
	}
	fmt.Println(calcResult(start, &cache))
	fmt.Println(calcResult(start, &cache).p1_wins > calcResult(start, &cache).p2_wins)
}
