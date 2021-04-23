# greddit

## About

greddit is a cli utility to interact with the reddit api. I started this mainly because [export-saved-reddit](https://github.com/csu/export-saved-reddit) is kinda unmaintained and I wanted to write something with Go. :)

## Usage

### Create reddit app

Here is a guide to get a client Id and client secret to use with the reddit API: [https://github.com/reddit-archive/reddit/wiki/OAuth2-Quick-Start-Example#first-steps](https://github.com/reddit-archive/reddit/wiki/OAuth2-Quick-Start-Example#first-steps)
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

currently, the following commands are available:

- authenticate  
  authenticates with the reddit api. This command is mainly for testing if the config is set correctly.
- subreddits, sr  
  interact with subreddits
  - list  
    prints a list of all subreddits you are subscribed to
  - subscribe  
    subscribe to subreddit with `NAME`
  - unsubscribe  
    unsubscribe from subreddit with `NAME`
