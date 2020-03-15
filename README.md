# gophercises
a repo for gophercises
This is a test commit

# Setup Go dev environment in Ubuntu
```bash
cd ~
wget https://dl.google.com/go/go1.14.linux-amd64.tar.gz
# or go to https://golang.org/dl/ to find the newest version
sudo tar -C /usr/local -xzf go1.14.linux-amd64.tar.gz
echo "export PATH=\$PATH:/usr/local/go/bin" >> ~/.profile
source ~/.profile
```

## [Gophercise Lessons](https://courses.calhoun.io/courses/cor_gophercises)
1. [HTML LINK PARSER](https://courses.calhoun.io/lessons/les_goph_16)
2. [SITEMAP BUILDER](https://courses.calhoun.io/lessons/les_goph_24)
    - Optimization trick: [the-empty-struct](https://dave.cheney.net/2014/03/25/the-empty-struct)
3. ...