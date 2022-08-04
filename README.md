# itutilsbot
Telegram bot made up to provide basic IT utilities like dns queries, man search and more

## Auth policies

Access policies can be applied to the bot using rego modules. Messages are parsed to json and provided to the modules as input. A default policy could be found [here](https://github.com/Eryalito/itutilsbot/tree/main/auth).

Path to policies folder may be changed with the flag `-a`, the default value is `auth`. This feature can be disabled witth the flag `-d`, allowing all incoming messages.
