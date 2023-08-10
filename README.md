# About

This application is only test a better performance (multi-thread) to send e-mail notifications by SMTP utilization Goroutines and Channels.

### Tested on GLPi version 10.0.9

# Tasks
- [x] Check queue every 15 seconds.
- [x] Pushed queued emails via SMTP.
- [x] Unsent emails remain in the queue
- [x] Sent emails are logged
- [x] Generated action log in crontasklogs.

# Build and run
 - Rename file ..env_copy_and_rename_this_envfile to .env
 - Set your config data in .env
 - Execute `go build`

# Performance without microservice
<img src='https://i.imgur.com/EE5Zpmv.png'>

# Performance with microservice
<img src='https://i.imgur.com/Psf2coi.png'>