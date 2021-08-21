# greddit

## About

greddit is a cli utility to interact with the reddit api.

## Usage

### Create reddit app

Here is a guide to get a client Id and client secret to use with the reddit API:
[https://github.com/reddit-archive/reddit/wiki/OAuth2-Quick-Start-Example#first-steps](https://github.com/reddit-archive/reddit/wiki/OAuth2-Quick-Start-Example#first-steps)

The necessary data must be provided in a config file in this format:

```json
{
  "ID": "id",
  "Secret": "secret",
  "Username": "username",
  "Password": "password"
}
```

### Running greddit

`greddit [global options] command [command options] [arguments...]`

The following commands are available:

- authenticate  
  authenticates with the reddit api.
  This command is mainly for testing if the config is set correctly.
- subreddits, sr  
  interact with subreddits
  - list  
    prints a list of all subreddits you are subscribed to
  - subscribe  
    subscribe to subreddit with `NAME`
  - unsubscribe  
    unsubscribe from subreddit with `NAME`
- upvoted  
  interact with upvoted posts
  - list  
    prints a list of all posts you upvoted
- saved  
  interact with posts you saved
  - list  
    prints a list of all posts you saved
- saved-comments  
  interact with comments you saved
  - list  
    prints a list of all posts you posted a saved comment to
- downvoted  
  interact with downvoted posts
  - list  
    prints a list of all posts you downvoted
- submitted  
  interact with posts you submitted
  - list  
    prints a list of all posts you submitted
- submitted-comments  
  interact with comments you submitted
  - list  
    prints a list of all posts you posted a submitted comment to
- export  
  export content of your reddit account
  - one of: `upvoted`, `saved`, `saved-comments`, `downvoted`, `submitted`, `submitted-comments`
  - you can choose the output format with an `-f` flag:
    - Exporting posts is supported in markdown, html and text.
    - Exporting comments is supported in md.
