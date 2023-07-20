import http from 'k6/http';

const config = JSON.parse(open('../config/config.json'));
const creds = JSON.parse(open('../temp_store.json'));

export var options = {};
if (config.LoadTestConfiguration.RPS) {
    options = {
        discardResponseBodies: true,
        scenarios: {
            contacts: {
            executor: config.LoadTestConfiguration.Executor,
            duration: config.LoadTestConfiguration.Duration,
            rate: config.LoadTestConfiguration.Rate,
            timeUnit: config.LoadTestConfiguration.TimeUnit,
            preAllocatedVUs: config.LoadTestConfiguration.VirtualUserCount,
            },
        }
    }
} else {
    options = {
        vus: config.LoadTestConfiguration.VirtualUserCount,
        duration: config.LoadTestConfiguration.Duration,
    }
}

export function setup() {
	if (config.PostsConfiguration.MaxWordsCount <= 0) {
        console.error("Error in validating the posts configuration:", "max word count should be greater than 0");
		return;
	}

	if (config.PostsConfiguration.MaxWordLength <= 0) {
        console.error("Error in validating the posts configuration:", "max word length should be greater than 0");
		return;
	}
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
            count++;
        }

        message += ' ';
        words++;
    }

	return message;
}

function getRandomToken() {
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

    const resp = http.post(`${config.ConnectionConfiguration.ServerURL}/api/v4/posts`, payload, {headers});
    check(resp, {
        'Post status is 201': (r) => resp.status === 201,
    });
}
