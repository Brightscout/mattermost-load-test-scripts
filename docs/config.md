# Loadtest configurations

## Connection configurations

### ServerURL

*string*

The URL to direct the load. It should be the public facing URL of the target Mattermost instance.

### AdminEmail

*string*

The email for the system administrator of the target Mattermost instance.

### AdminPassword

*string*

The password for the system administrator of the target Mattermost instance.

## User Configurations

### Username

*string*

The username for the new user to be created.

### Email

*string*

The email for the new user to be created.

### Password

*string*

The password for the new user to be created.

## Channel Configurations

### DisplayName

*string*

The display name for the new channel to be created.

### Name

*string*

The name for the new channel to be created.

### Type

*string*

The type of new channel to be created. Channel types can be `O` and `P`, which denote open channel and private channel, respectively.

### MMTeamName

*string*

The Mattermost team name for the new channel to be created.

### MSTeamsTeamID

*string*

The MS Teams Team ID having the MS Teams channel to which the Mattermost channel is to be linked.

### MSTeamsChannelID

*string*

The MS Teams Channel ID to which the Mattermost channel is to be linked.

## Post Configurations

### Count

*int*

The count of virtual users running concurrently and creating posts in the Mattermost channels, DMs, and GMs.

### MaxWordsCount

*int*

The maximum number of words in a sentence in a post.

### MaxWordLength

*int*

The maximum length of each word in a sentence of a post.

### Duration

*string*

The duration in seconds specifying the total duration of the test run.
