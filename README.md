# GenericPointerList
  
# Add & Remove & RemoveAt
  
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
	
	
![image](https://user-images.githubusercontent.com/59788044/181878728-104924b5-5a58-49b0-87ef-ffe39092d6c4.png)

# Reverse
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

	log.Println("Before")
	log.Println("----------------")

	for _, currentPlayer := range playerList.ToArray() {
		log.Printf("Player ID:%d, Player Health:%f", currentPlayer.ID, currentPlayer.Health)
	}

	log.Println("----------------")
	log.Println("After")
	log.Println("----------------")
	playerList.Reverse()

	for _, currentPlayer := range playerList.ToArray() {
		log.Printf("Player ID:%d, Player Health:%f", currentPlayer.ID, currentPlayer.Health)
	}
	
![image](https://user-images.githubusercontent.com/59788044/181878855-c0aaad7f-408c-4c7b-aaef-309d6d2362df.png)

# Sort
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

	log.Println("Before")
	log.Println("----------------")

	for _, currentPlayer := range playerList.ToArray() {
		log.Printf("Player ID:%d, Player Health:%f", currentPlayer.ID, currentPlayer.Health)
	}

	log.Println("----------------")
	log.Println("After")
	log.Println("----------------")

	playerList.Sort(func(left, right *Player) bool {
		return left.Health > right.Health
	})

	for _, currentPlayer := range playerList.ToArray() {
		log.Printf("Player ID:%d, Player Health:%f", currentPlayer.ID, currentPlayer.Health)
	}
	
![image](https://user-images.githubusercontent.com/59788044/181878946-c4f6bb14-73f6-428c-befd-de664fae296e.png)

# RemoveAll

	playerList := PointerList.NewPointerList[Player]()

	player1 := &Player{ID: 0, Health: 100}
	player2 := &Player{ID: 1, Health: 90}
	player3 := &Player{ID: 2, Health: 10}
	player4 := &Player{ID: 3, Health: 1}
	player5 := &Player{ID: 4, Health: 0}
	player6 := &Player{ID: 5, Health: 0}
	player7 := &Player{ID: 6, Health: 99}
	player8 := &Player{ID: 7, Health: 150}
	player9 := &Player{ID: 8, Health: 60}
	player10 := &Player{ID: 9, Health: 0}

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

	log.Println("Before")
	log.Println("----------------")

	for _, currentPlayer := range playerList.ToArray() {
		log.Printf("Player ID:%d, Player Health:%f", currentPlayer.ID, currentPlayer.Health)
	}

	log.Println("----------------")
	log.Println("After")
	log.Println("----------------")

	playerList.RemoveAll(func(current *Player, index int) bool {
		if current.Health <= 0 {
			log.Println("The player died. ID:", current.ID, " Index:", index)

			return true
		}

		return false
	})

	log.Println("----------------")

	for _, currentPlayer := range playerList.ToArray() {
		log.Printf("Player ID:%d, Player Health:%f", currentPlayer.ID, currentPlayer.Health)
	}
![image](https://user-images.githubusercontent.com/59788044/181879140-96e23e3c-3245-4590-8feb-42818a63c2f4.png)

