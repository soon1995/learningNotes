```bash
$ wget https://raw.githubusercontent.com/jamtur01/dockerbook-code/master/code/5/sample/website/index.html
$ cd ..
$ docker run -d -p 80 --name website -v $PWD/website:/var/www/html/website jamtur01/nginx nginx
```
