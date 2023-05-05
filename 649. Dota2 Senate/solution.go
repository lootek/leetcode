func predictPartyVictory(senate string) string {
    senators := []rune(senate)

    for {
        for i := 0; i < len(senators); i++ {
            ban := 'D'
            if senators[i] == 'D' {
                ban = 'R'
            }

            // fmt.Println(senators)

            for j := i+1; j < len(senators); j++ {
                // fmt.Printf("> %v\n", senators[j])

                if senators[j] == ban {
                    senators = append(senators[:j], senators[j+1:]...)
                    ban = 'X'
                    break
                }
            }

            if ban != 'X' {
                for j := 0; j < i; j++ {
                    // fmt.Printf(">> %v\n", senators[j])

                    if senators[j] == ban {
                        senators = append(senators[:j], senators[j+1:]...)
                        ban = 'X'
                        break
                    }
                }
            }

            // fmt.Println(senators)
            if ban != 'X' {
                if senators[0] == 'R' {
                    return "Radiant"
                }

                return "Dire"
            }
        }
    }
}

