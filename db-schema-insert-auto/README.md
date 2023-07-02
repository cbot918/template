# Bulk User to DB
this is a project for research/practice
- makefile script
- automation something
- postgres cli with make and docker

# what's inside main.go
its a tool for batch copy data into postgres db
it will download a dataset from IMDB, then copy to db table
the dataset content 12M rows

# run
run 
```bash
make run
```
verify
```
make select
```
psql-cli
```
make psql
```
<br>

# Annoncement

code in main.go is from

https://mariocarrion.com/2020/08/27/go-implementing-complex-pipelines-part-5.html

for research use temporary
# refs:
https://mariocarrion.com/2020/08/27/go-implementing-complex-pipelines-part-5.html
https://github.com/MarioCarrion/complex-pipelines/blob/master/part5/main.go

