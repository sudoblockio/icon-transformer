### Routines 

Routines are repeatable functions that can be run as either:
- init sequences 
  - Run as soon as the steady state worker runs but not after 
- cron jobs
  - On cron that update things that need to be update regularly but not streamed like stats tables 
- recovery 
  - Only run when things go sideways (ie loosing redis cache)
