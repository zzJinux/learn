#!/usr/bin/expect --
set timeout -1

spawn ./questions.sh

expect "Hello, who are you?\r"
send -- "Im Adam\r"

expect "Can I ask you some questions?\r"
send -- "Sure\r"

expect "What is your favorite topic?\r"
send -- "Technology\r"

expect eof
