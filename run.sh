go build -o instana
env --debug $(cat .env | grep -v '^#') ./instana
