package models

//点赞+关注+收藏        一键三连
type LikeTable struct { //点赞
	LikeId      int `JSON:"likeid"`
	UserAccount int `JSON:"useraccount"`
	LikeNoteId  int `JSON:"likenoteid"`
}
type FollowTable struct { //关注
	FollowId      int `JSON:"followid"`
	UserAccount   int `JSON:"useraccount"`
	FollowAccount int `JSON:"followaccount"`
}
type CollectTable struct { //收藏
	CollectId     int `JSON:"collectid"`
	UserAccount   int `JSON:"useraccount"`
	CollectNoteId int `JSON:"collectnoteid"`
}
