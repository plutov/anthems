## Anthems

Anthems is an Action on Google to play a national anthem of selected country.

Source: [List of national anthems](https://en.wikipedia.org/wiki/List_of_national_anthems)

### Deploy

This project is deployed to GCP:

```
gcloud app deploy
```

Production webhook address: https://anthems-19300.appspot.com

### Media

All anthems are uploaded to this S3 bucket: https://s3.amazonaws.com/actions-on-google-anthems/

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