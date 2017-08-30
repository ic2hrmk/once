# ONCE
Golang library draft for creating one-time links

## Description
This is simple project that demonstrates usage of _Once_ library.
It's goal is to create disposable short link. 

## Examples
### Create
To create new one-time link:
~~~bash
curl --request POST --url 'http://localhost:4040/?url=google.com.ua'
~~~

As response you will get:
~~~json
{
	"short_link": "me.com/2b586a0e"
}
~~~

### Request value
To be redirected to hidden link:
~~~bash
curl --request GET --url http://localhost:4040/2b586a0e
~~~

In case of second-time link usage 404 error will be provided.