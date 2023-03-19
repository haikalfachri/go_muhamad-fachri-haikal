#!/bin/bash

tgl=$(date)

# Membuat folder parent
mkdir "$1 at $tgl"

# Membuat sub folder about me, my_friends, my_sistem_info
mkdir {"$1 at $tgl"/about_me,"$1 at $tgl"/my_friends,"$1 at $tgl"/my_system_info}

# Membuat sub folder ../about_me/personal, ../about_me/professional
mkdir {"$1 at $tgl"/about_me/personal,"$1 at $tgl"/about_me/professional}

# Membuat file facebook.txt
touch "$1 at $tgl"/about_me/personal/facebook.txt

# Membuat file linkedin.txt
touch "$1 at $tgl"/about_me/professional/linkedin.txt

# Membuat file list_of_my_friends.txt
touch "$1 at $tgl"/my_friends/list_of_my_friends.txt

# Membuat file about_this_laptop.txt
touch "$1 at $tgl"/my_system_info/about_this_laptop.txt

# Membuat file internet_connection
touch "$1 at $tgl"/my_system_info/internet_connection.txt

# Memasukkan argumen ke-2 ke facebook.txt
echo $2 > "$1 at $tgl"/about_me/personal/facebook.txt

# Memasukkan argumen ke-3 ke linkedin.txt
echo $3 > "$1 at $tgl"/about_me/professional/linkedin.txt

# Memasukkan listfriends dari github ke list_of_myfriends.txt
curl https://gist.githubusercontent.com/tegarimansyah/e91f335753ab2c7fb12815779677e914/raw/94864388379fecee450fde26e3e73bfb2bcda194/list%2520of%2520my%2520friends.txt > "$1 at $tgl"/my_friends/list_of_my_friends.txt

# Memasukkan nama user dan uname -a ke about_this_laptop.txt
echo "My username: $(whoami)" > "$1 at $tgl"/my_system_info/about_this_laptop.txt
echo "With host: $(uname -a): $tgl" > "$1 at $tgl"/my_system_info/about_this_laptop.txt 

# Memasukkan ping google ke internet_connection.txt
ping google.com -c 3 > "$1 at $tgl"/my_system_info/internet_connection.txt
