# Backend project

## Set in motion
- Download and install the search engine [ZincSearch](https://docs.zinc.dev/installation/)
- Download the [database](http://www.cs.cmu.edu/~enron/enron_mail_20110402.tgz) where Enron Corp emails are located

- The main file to run the project is located in 'email-indexer/server/cmd'.
```sh
cd email-indexer/server/cmd
```

- Then the environment variables 'ZINC_FIRST_ADMIN_USER' and 'ZINC_FIRST_ADMIN_PASSWORD' must be set. These variables are used to do a basic authentication with ZinchSearch.
```sh
set ZINC_FIRST_ADMIN_USER=admin
set ZINC_FIRST_ADMIN_PASSWORD=Complexpass#123
```

- And finally
```sh
go run main.go
```

