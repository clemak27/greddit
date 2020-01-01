# pip install praw --user
import praw

reddit = praw.Reddit(client_id='',
        client_secret='',
        password='',
        user_agent='',
        username='')

print("Authenticated as ")
print(reddit.user.me())

with open('subs.txt') as f:
    lines = f.readlines()

    for subreddit_name in lines:
        subreddit = reddit.subreddit(subreddit_name)
        subreddit.subscribe()
