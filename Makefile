deploy:
	gcloud functions deploy DayCount --runtime go111 --trigger-http --env-vars-file .env.yml
	gcloud functions deploy DayCountSchedule --runtime go111 --trigger-http --env-vars-file .env.yml

delete:
	gcloud functions delete DayCount