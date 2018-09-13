![anthems](https://raw.githubusercontent.com/plutov/anthems/master/anthem.png)

## Anthems

Anthems is an Action on Google to play a national anthem of a selected country.

### Deploy

This project is deployed to GCP:

```
gcloud app deploy --version v1
```

Production webhook address: https://anthems-19300.appspot.com

### Media

This app uses anthems from https://anthemworld.com

### Conversation flow

- User: Talk to Anthems
- Bot: Hi! I can play a national anthem of any country, which country do you want?
- User: Laplandia
  - Bot: Sorry, I didn't get that. Can you please try again?
- User: Belarus
  - Bot: Cool, Belarus.
  - Bot is playing audio file.
  - Bot: That's it, do you want to listen another country anthem?
- User: Yes.
  - Bot: Great, which country do you want?
- User: No.
  - Bot: Sure, thanks for talking to Anthems. Goodbye!