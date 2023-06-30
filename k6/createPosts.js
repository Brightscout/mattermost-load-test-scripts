import http from 'k6/http';

const config = JSON.parse(open('../config/config.json'));
const response = JSON.parse(open('../temp_store.json'));

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
            count += 1;
        }

        message += ' ';
        words += 1;
    }

	return message;
}

function getRandomToken() {
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
