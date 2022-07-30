# GenericPointerList
  
# Add & Remove
  
	playerList := PointerList.NewPointerList[Player]()

	player1 := &Player{ID: 0, Health: 100}
	player2 := &Player{ID: 1, Health: 90}
	player3 := &Player{ID: 2, Health: 10}
	player4 := &Player{ID: 3, Health: 1}
	player5 := &Player{ID: 4, Health: 0}
	player6 := &Player{ID: 5, Health: 15}
	player7 := &Player{ID: 6, Health: 99}
	player8 := &Player{ID: 7, Health: 150}
	player9 := &Player{ID: 8, Health: 60}
	player10 := &Player{ID: 9, Health: 50}

	playerList.Add(player1)
	playerList.Add(player2)
	playerList.Add(player3)
	playerList.Add(player4)
	playerList.Add(player5)
	playerList.Add(player6)
	playerList.Add(player7)
	playerList.Add(player8)
	playerList.Add(player9)
	playerList.Add(player10)

	playerList.Remove(player2)
	playerList.RemoveAt(0)
	playerList.RemoveAt(5)

	for _, currentPlayer := range playerList.ToArray() {
		log.Printf("Player ID:%d, Player Health:%f", currentPlayer.ID, currentPlayer.Health)
	}
	
	![image](https://user-images.githubusercontent.com/59788044/181878662-76e61142-94ac-42ab-8e0e-4af2ac62b1a0.png)


	
