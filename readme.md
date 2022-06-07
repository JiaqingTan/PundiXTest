Pundi X Test Submission
-



Usage:
- query/bank:
  - http://localhost:3000/query/bank/total
  - http://localhost:3000/query/bank/balances/fx1sp9l4g78veqcsecmm3m77d9t7eaq3nzq60l4s0
  - http://localhost:3000/query/bank/denom-metadata/
  
- query/distribution:
  - http://localhost:3000/query/distribution/commission/fxvaloper1a73plz6w7fc8ydlwxddanc7a239kk45jnl9xwj
  - http://localhost:3000/query/distribution/community-pool
  - http://localhost:3000/query/distribution/params
  - http://localhost:3000/query/distribution/rewards/fx1sp9l4g78veqcsecmm3m77d9t7eaq3nzq60l4s0/fxvaloper1a73plz6w7fc8ydlwxddanc7a239kk45jnl9xwj
  - http://localhost:3000/query/distribution/slashes/fxvaloper1a73plz6w7fc8ydlwxddanc7a239kk45jnl9xwj/1000/2000
  - http://localhost:3000/query/distribution/validator-outstanding-rewards/fxvaloper1a73plz6w7fc8ydlwxddanc7a239kk45jnl9xwj
  

Notes:

- App performs SSH to EC2 instance (free tier) to execute fxcored commands with "--node" flag
- Instance's user and password are in plain text in *config.json* file which is committed to repo for ease of use (not ideal)
- Validator and address validation only has basic prefix validation (no bech32 validation)
- Height validation only checks for positive integers
- Paths are hardcoded in order to easily perform validation on validator, address, heights

Things to improve:
- SSH connection can be more secured, using .pem private key.
- SSH client should be shutdown gracefully on app termination
- *config.json* file can be stored remotely on own servers (for security reasons) and retrieved on VM/K8 deployment
- Improvement to validations can be done: bech32 validation for validator and address
- Validators, addresses and heights can ideally be shifted to be part of request body instead of URL params
- Router paths can be more dynamic instead of having a handler function per
- Output of response can be prettified, raw output is printed as JSON currently

