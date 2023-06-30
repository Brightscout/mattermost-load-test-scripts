import http from 'k6/http';

const config = JSON.parse(open('../config/config.json'));
<<<<<<< HEAD
const response = JSON.parse(open('../temp_store.json'));
=======
const creds = JSON.parse(open('../temp_store.json'));
>>>>>>> a6f6c9aaefce1279172bb3fe3c31331ad1deb556

export const options = {
    vus: config.PostsConfiguration.Count,
    duration: config.PostsConfiguration.Duration,
}

function getRandomMessage(wordsCount, wordLength) {
	let message = '';
    const characterSet = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
	let words = 0;
    wordsCount = Math.floor(Math.random() * wordsCount) + 1;
    while (words < wordsCount) {
        let count = 0;
        wordLength = Math.floor(Math.random() * wordLength) + 1;
        while (count < wordLength) {
            message += characterSet.charAt(Math.floor(Math.random() * characterSet.length));
<<<<<<< HEAD
            count += 1;
        }

        message += ' ';
        words += 1;
=======
            count++;
        }

        message += ' ';
        words++;
>>>>>>> a6f6c9aaefce1279172bb3fe3c31331ad1deb556
    }

	return message;
}

function getRandomToken() {
<<<<<<< HEAD
    let token = [];
    response.UserResponse.map((u) => token.push(u.token));
    return token[(Math.floor(Math.random() * token.length))];
}

function getRandomChannel() {
    let channelId = [];
    response.ChannelResponse.map((c) => channelId.push(c.id));
    if (response.DMResponse) {
        channelId.push(response.DMResponse.id);
    }

    if (response.GMResponse) {
        channelId.push(response.GMResponse.id);
    }

    return channelId[(Math.floor(Math.random() * channelId.length))];
=======
    let tokens = [];
    creds.UserResponse.map((u) => tokens.push(u.token));
    return tokens[(Math.floor(Math.random() * tokens.length))];
}

function getRandomChannel() {
    let channelIds = [];
    creds.ChannelResponse.map((c) => channelIds.push(c.id));
    if (creds.DMResponse) {
        channelIds.push(creds.DMResponse.id);
    }

    if (creds.GMResponse) {
        channelIds.push(creds.GMResponse.id);
    }

    return channelIds[(Math.floor(Math.random() * channelIds.length))];
>>>>>>> a6f6c9aaefce1279172bb3fe3c31331ad1deb556
}

export default function() {
    const payload = JSON.stringify({
        channel_id: getRandomChannel(),
        message: getRandomMessage(config.PostsConfiguration.MaxWordsCount, config.PostsConfiguration.MaxWordLength)
    })

    const headers = {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${getRandomToken()}`,
    }

    http.post(`${config.ConnectionConfiguration.ServerURL}/api/v4/posts`, payload, {headers});
}
