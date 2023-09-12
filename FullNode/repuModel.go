package main

import "math"

func ratings(sender string, receiver string, p int) (int, int) {
	dis, _ := checkDistance(sender, receiver)

	bc := NewBlockchain()
	defer bc.Db.Close()
	bal := getElectorBalance(bc, sender)

	var reviewRating int
	var ratingReward int
	if dis <= disThreshold {
		reviewRating = (disThreshold / dis) * bal * p
		ratingReward = disThreshold / dis
	} else {
		reviewRating = (1 / 2) * bal * p
		ratingReward = 1 / 2
	}
	return reviewRating, ratingReward
}

func reviewReward(rating int, invest int) int {
	reward := rating * invest * roi
	return reward
}

//punishment is an index [0,1], should multiply purchased tokens
func punishment(address string) float64 {

	dis, _ := checkBigDistance(address, punishThreshold)

	var index float64
	if dis <= 10 && dis > 0 {
		index = math.Exp(-0.2 * float64(dis))
	} else {
		index = 0
	}
	return index
}
