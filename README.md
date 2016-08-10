# IconVehicleID Check

### This application hits IconMedia's vehicle api endpoint over and over, watching for errors and making lists of successes and errors.

1. ICON_USER=<user> ICON_PASS=<pass> ICON_DOMAIN=<url> go run main.go
2. Goto localhost:8080
3. This will generate error and success csvs. Use for fun and profit -or-
4. Goto localhost:8080/insertsuccess to put successful vehicles in mongoDB, -d icon -c iconSuccess
5. Goto localhost:8080/inserterror to put non-successful vehicles in mongoDB -d icon -c iconErrors

