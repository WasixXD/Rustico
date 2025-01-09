# Rustico

![image](https://github.com/user-attachments/assets/21553e6f-bd29-47a6-9b9f-25ea0d79fd2b)

A self-hosted project storer for those who don't want to forget unicorns.


# How to install 🚀
Create a `docker-compose.yml` file and paste the content:

```
services:
  rustico:
    image: lucaswasilewski/rustico
    environment:
      - SERVER_URL=http://localhost:3001
    ports:
      - 3000:3000
      - 3001:3001
```

Then run with

```bash
docker compose up -d 
```


# Goals 🏆
- [ x ] Create a idea
- [ x ] Get ideas from your pool of ideas
- [ x ] Mark an idea as completed

