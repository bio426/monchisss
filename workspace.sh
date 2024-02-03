#!/bin/bash

session="monchiss"

tmux new-session -d -s $session

window=1
tmux rename-window -t $session:$window 'back'
tmux send-keys -t $session:$window 'cd ~/Desktop/monchisss && nvim .' C-m

window=2
tmux new-window -t $session:$window -n 'front'
tmux send-keys -t $session:$window 'cd ~/Desktop/monchisss/ui && nvim .' C-m

window=3
tmux new-window -t $session:$window -n 'run'
tmux send-keys -t $session:$window 'cd ~/Desktop/monchisss/ui && npm run dev' C-m
tmux split-window -v
tmux send-keys -t $session:$window 'cd ~/Desktop/monchisss && source env.sh && wgo run main.go' C-m

tmux attach-session -t $session
