**Problem Statement**

To make an API to fetch latest videos sorted in reverse chronological order of their publishing date-time from YouTube for a given tag/search query in a paginated response.

The app should perform the following tasks:
* Server should call the YouTube API continuously in background (async) with some interval (say 10 seconds) for fetching the latest videos for a predefined search query and should store the data of videos (specifically these fields - Video title, description, publishing datetime, thumbnails URLs and any other fields you require) in a database with proper indexes.
* A GET API which returns the stored video data in a paginated response sorted in descending order of published datetime.
* A basic search API to search the stored videos using their title and description.
* Dockerize the project.
* It should be scalable and optimised.
* Add support for supplying multiple API keys so that if quota is exhausted on one, it automatically uses the next available key.
* Optimise search api, so that it's able to search videos containing partial match for the search query in either video title or description.

***Bonus Features***

* Add support for supplying multiple API keys so that if quota is exhausted on one, it automatically uses the next available key.
* Make a dashboard to view the stored videos with filters and sorting options (optional)
* Optimise search api, so that it's able to search videos containing partial match for the search query in either video title or description.
    - Ex 1: A video with title *`How to make tea?`* should match for the search query `tea how`

**Build Instructions**

Clone project using following command,
> git clone https://github.com/Ayush-Walia/Fampay-Youtube.git

Move to the project directory using,
> cd Fampay-Youtube

In order to run the main app, run the following command,
> docker compose up

In order to stop the app, run the following command,
> docker compose down

Note: db takes 30 sec to start as docker-compose check db healthy status after 30s

**Input commands**

To get the stored video data:
> curl -v -X GET 'http://localhost:8080/api/v1/videos?pageNo=1&pageSize=100'

To search for the videos title or description in stored videos(supports partial match):
> curl -v -X GET 'http://localhost:8080/api/v1/search?query=query&pageNo=1&pageSize=10'

Note: 
* APIKeys can be set in docker-compose env var, multiple keys can passed comma separated.
* Postman collection is also present in the repo