#! /bin/sh

if [ -f /etc/systemd/system/boulder-bot.service ]; then
    sudo systemctl stop boulder-bot.service
    sudo systemctl disable boulder-bot.service
else
    sudo cp -fv ./boulder-bot.service /etc/systemd/system/boulder-bot.service
    echo Telegram Bot API Token: 
    read api_token
    printf Environment="API_TOKEN=%s" $api_token | sudo tee -a /etc/systemd/system/boulder-bot.service
fi

# TODO copy for the right architecture
sudo cp -fv ./boulder-bot_arm /usr/local/boulder-bot
sudo chmod +x /usr/local/boulder-bot

sudo systemctl enable boulder-bot.service
sudo systemctl start boulder-bot.service