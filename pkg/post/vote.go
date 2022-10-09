package post

type Vote struct {
	UserId string `json:"user" bson:"user"`
	Vote   int32  `json:"vote" bson:"vote"`
}

func NewVote(userId string, vote int32) *Vote {
	return &Vote{
		UserId: userId,
		Vote:   vote,
	}
}
