* This project is a cli tool "gator" for [boot.dev](https://boot.dev)
To use this tool, you __must__ have postgres, and go installed to run this program.
You can intsall the gator cli with go install github.com/Youg-Otricked/gator\
You run it with go run . [command name] [args]. The commands are: agg: aggregate. constantly agregates a website you pass the url in as args. login sets main user to the usr you pass in as args. register. any command that uses a user must have the user registerd. you regoster with go run . register [username]. reset: clears all tables. users: lists all users. addFeed: basically reigster but for a feed. you must pass in a url link and cnatn use a url unless you have added it for **That user** feeds: lists all feeds ofr curently loged in user follow folllows a feed following listsa ll followed feeds for that usr and unffolw unfollows a feed.
create a file in you home directory (~/) call .gatorconfig.json and inside it put: {"db_url":"postgres://example","current_user_name":""}. this is what he program uses for a lot of important things.

