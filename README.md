# MyFitnesGoLearningApp
Yes, I created this go program. Most of it I guess. Design is from a paid website that sold the components to me (u cant buy it anymore, soo its mine now). As mentioned it was ment for me to learn Go, TailWind, Simple auth, SqLite and soo on. You will probably encounter some auth stuff in the code, yea its useless, it was not safe from the beginning. I Still implemented it, but scrached it later (yea .env will be included lol). 



## Idea
Long and short answer: for learning of course. And as you expected I did learn a lot even tho its fucked up XD. Also its far from complete with a lot of weird code and bugs (just so you know).

Soo you are interested in the features. Yea every feature is implemented in atleas 2 ways (because of learning). I went out of my way to implement the most ridiculous UX on the planed, but I dont care HAHA. 

## Features And Pictures
Features include:
- Register via CLI (well plan was big but in the end i implemented for ony one user)
- Login page
- Stopwatch that can beep every 1min
- Tracking of KG total for every exercise
- Weight wracking with a fancy line graph
- You can add random YT link (can be modified)
- You get a lot of daily challanges (can be modified)
- And a fancy Exercise tracker (you can add/remove/update, it does not track history only the current KG that you can lift)
-- Go to https://strengthlevel.com/strength-standards if you want to insert the base data for exercises.


## Installation and usage

Well its nice and simple. As mentioned before you will need to create a user via CLI (only one time). From here you just go to the localhost:8080 page and you can start using it.

- If you want to run it with the go compiler you do:
```
go run . -username=<insert> -pass=<insert>
```
- and next time you run it with:
```
go run . -username=<insert> -pass=<insert>
```


**Or if you want to build and use it that way:**
```
go build .
.\GYM-App -username=<insert> -pass=<insert>
```
And later you can just double click on the exe

## Want to reset the app
Just remove api.db and start fresh.
