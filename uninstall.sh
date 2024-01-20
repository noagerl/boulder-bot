#! /bin/sh

if [ -f /etc/systemd/system/boulder-bot.service ]; then
    sudo systemctl stop boulder-bot.service
    sudo systemctl disable boulder-bot.service
    sudo rm -fv /etc/systemd/system/boulder-bot.service
fi

if [ -f /usr/local/boulder-bot ]; then
    sudo rm -fv /usr/local/boulder-bot
fi

