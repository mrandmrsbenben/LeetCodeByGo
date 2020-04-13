package main

import "fmt"

//执行用时 :60 ms, 在所有 Go 提交中击败了25.00%的用户
//内存消耗 :10.8 MB, 在所有 Go 提交中击败了100.00%的用户
// Tweet Tweet
type Tweet struct {
	userId int
	postId int
}

// Account Twitter Account
type Account struct {
	id    int
	foees map[int]int
}

func createAccount(userId int) *Account {
	return &Account{userId, map[int]int{userId: 1}}
}

// Twitter Twitter Struct
type Twitter struct {
	accounts map[int]*Account
	tweets   []*Tweet
}

// Constructor Initialize your data structure here.
func Constructor() Twitter {
	return Twitter{map[int]*Account{}, []*Tweet{}}
}

// PostTweet Compose a new tweet.
func (this *Twitter) PostTweet(userId int, tweetId int) {
	if _, ok := this.accounts[userId]; !ok {
		this.accounts[userId] = createAccount(userId)
	}
	this.tweets = append(this.tweets, &Tweet{userId, tweetId})
}

// GetNewsFeed Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent.
func (this *Twitter) GetNewsFeed(userId int) []int {
	tweets := make([]int, 0)
	if acc, ok := this.accounts[userId]; ok {
		for i := len(this.tweets) - 1; i >= 0 && len(tweets) < 10; i-- {
			if _, following := acc.foees[this.tweets[i].userId]; following {
				tweets = append(tweets, this.tweets[i].postId)
			}
		}
	}
	return tweets
}

// Follow Follower follows a followee. If the operation is invalid, it should be a no-op.
func (this *Twitter) Follow(followerId int, followeeId int) {
	foer := this.accounts[followerId]
	if foer == nil {
		foer = createAccount(followerId)
		this.accounts[followerId] = foer
	}
	foer.foees[followeeId] = 1
}

// Unfollow Follower unfollows a followee. If the operation is invalid, it should be a no-op.
func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if followerId == followeeId {
		return
	}
	foer := this.accounts[followerId]
	if foer != nil {
		delete(foer.foees, followeeId)
	}
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
/**
Twitter twitter = new Twitter();

// User 1 posts a new tweet (id = 5).
twitter.postTweet(1, 5);

// User 1's news feed should return a list with 1 tweet id -> [5].
twitter.getNewsFeed(1);

// User 1 follows user 2.
twitter.follow(1, 2);

// User 2 posts a new tweet (id = 6).
twitter.postTweet(2, 6);

// User 1's news feed should return a list with 2 tweet ids -> [6, 5].
// Tweet id 6 should precede tweet id 5 because it is posted after tweet id 5.
twitter.getNewsFeed(1);

// User 1 unfollows user 2.
twitter.unfollow(1, 2);

// User 1's news feed should return a list with 1 tweet id -> [5],
// since user 1 is no longer following user 2.
twitter.getNewsFeed(1);
*/
func main() {
	twitter := Constructor()
	twitter.PostTweet(1, 5)
	fmt.Println("Output: ", twitter.GetNewsFeed(1)) // -> [5]
	twitter.Follow(1, 2)
	twitter.PostTweet(2, 6)
	fmt.Println("Output: ", twitter.GetNewsFeed(1)) // -> [6, 5]
	twitter.Unfollow(1, 2)
	fmt.Println("Output: ", twitter.GetNewsFeed(1)) // -> [5]
}
