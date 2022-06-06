Pundi X Test Submission
-

Notes:

- App performs SSH to EC2 instance (free tier) to execute fxcored commands

Things to improve:

- Mistakenly built solution around "fxcored query bank total" example in email, only considered:

    - Up to 3 consecutive params
    - Validation for first 2 params
    - Therefore, only works well with "fxcored query bank total" for now
    
- Response can be prettfied more to display results in a better format