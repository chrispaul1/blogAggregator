1. Install PostGres and Go to run this program on your device
    PostGres MacOs - brew install postgresql@15
    PostGres Linux - sudo apt update
                     sudo apt install postgresql postgresql-contrib
    Go Install Instruction Link - https://go.dev/doc/install

    For linux only - update the postgress password - make sure you wont forget it and that its simple
        - sudo passwd postgres
    
    Commands to start the postgres server in the background
        - Mac: brew services start postgresql
        - Linux: sudo service postgresql start

    Commands to enter the psql shell for the database
        - Mac: psql postgres
        - Linux: sudo -u postgres psql

2. Next install gator CLI - 
        - go install github.com/chrispaul1/blog@latest

3. Create json config file called .gatorconfig.json in your home directory, so ~/.gatorconfig.json
    It will contain the current user who logged in and the credentials for the postgres database
    Ex.
    {
        "db_url": "connection_string_goes_here",
        "current_user_name": "username_goes_here"
    }
    You will require a database connection string
        - macOS (no password, your username): postgres://johnsmith:@localhost:5432/gator
        - Linux (password, postgres user): postgres://postgres:postgres@localhost:5432/gator 
    Don't worry about the user_name that will be set throught the program

4. Install goose and sqlc for database migrations and to generate go code from queries
        - go install github.com/pressly/goose/v3/cmd/goose@latest

    goose postgres <connection_string> up
    There might be chance goose might not find the folder that contains the migration queries
    So you might have to give the path to the folder
        - goose -dir ./sql/schema postgres "postgres://postgres:postgres@localhost:5432/gator" down
        - goose -dir ./sql/schema postgres "postgres://postgres:postgres@localhost:5432/gator" up

    After running goose down and up, run 
        - sqlc generate
    This will generate the go code from the sql queries that can be used within the code

5. Some of the commands you can perform is
    go run . register <name> 
        - this register a user in the database with the name
        - requires a string after the register word
    go run . login <name>   
        - this will login as the user if they are in the database
        - the name is required
    go run . reset
        - this will reset the database to a blank state
    go run . users
        - this will print out all the users registered in the database
    go run . addfeed "<feed name>" <link>
        - this will add a feed to the database with the link that linked to the current user 
    go run . feeds
        - this will print out all the feeds in the database, link and the user who created it
    go run . agg 30s
        - loops through the feeds and create posts that are saved within the database, the parameter is the time interval where the feeds are parsed and its items are saved into the database.
        It would be wise to give it a large interval not to ddos the sites. like 30s
    go run . follow <feed url>
        - creates a feed follow table for the feed tied to that url for the current user, allows a many to many relationship between the user and feeds. Prints the name of the feed and current user
    go run . following
        - lists all the feeds the current user is following
    go run . unfollow  <feed url>
        - unfollows this feed for the current user, deleted the feed follow struct for this feed and user
    go run . browse (optional parameter)
        - prints the most recent posts based on the number given next to browse, defaults to 2, if nothing is given
            Ex. go run . browse 3 

    