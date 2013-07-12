package main

import (
	"html/template"
	"log"
)

func getTemplates() *template.Template {
	t, err := template.New("foo").Parse(`{{define "sign_in.html"}}
<!DOCTYPE html>
<html lang="en" charset="utf-8">
<head>
	<title>Sign In</title>
	<script src="//ajax.googleapis.com/ajax/libs/jquery/1.10.1/jquery.min.js"></script>
    <style type="text/css">
        @import url(http://fonts.googleapis.com/css?family=Roboto:300,400,300italic);
        @import url(http://cdnjs.cloudflare.com/ajax/libs/meyer-reset/2.0/reset.css);
        body {
            color: #565656;
            font-weight: 800;
            letter-spacing: 1px;
        }
        .wrapper {
            margin: 25px auto;
            padding: 20px;
            width: 520px;
            position: relative;
            font-family: 'Roboto';
            font-weight: 300;
        }
        h1 {
            color: #6090c4;
            font-size: 30px;
            font-weight: 400;
            line-height: 42px;
            padding: 0 0 20px 0;
        }
        p {
            font-size: 19px;
            padding: 0 0 20px 0;
            line-height: 24px;
        }
        button {
            border: none;
            display: block;
            background: #b9b9b9;
            height: 40px;
            padding: 0 40px 0 40px;
            color: #ffffff;
            text-align: center;
            font-size: 16px;
            border-radius: 5px;
            margin: 0 0 80px 0;
            -webkit-transition: all 0.15s linear;
            -moz-transition: all 0.15s linear;
            transition: all 0.15s linear;
        }
        button:hover {
          background: #ff8368;
        }
        label {
            color: #6090c4;
            margin-top: -1px;
            border-top: 1px solid #f0f0f0;
            border-bottom: 1px solid #f0f0f0;
            font-size: 20px;
            line-height: 44px;
            padding: 7px 0 10px 0;
        }
        input {
            color: #a2a2a2;
            border: none;
            font-size: 20px;
            font-family: 'Roboto';
            font-style: italic;
            letter-spacing: 1px;
            margin-left: 25px;
        }
        input:focus {
            border: none;
            outline: none;
        }
        span {
            text-transform:capitalize;
        }
        .google {
            margin-top: -7px;
        }
    </style>
</head>
<body>
	<div class="wrapper">
		<h1>Login Required</h1>
		{{ if .Htpasswd }}
		<fieldset>
			<p>Have you been given credentials from <span>{{.Domain}}</span>?</br> Sign in below.</p>
			<form method="POST" action="/oauth2/sign_in">
			<label>Username 
				<input type="text" name="username" size="30" value="John Doe" onblur="if (this.value == '') {this.value = 'John Doe';}" onfocus="if (this.value == 'John Doe') {this.value = '';}">
			</label>
			<br/>
			<label>Password 
				<input type="password" name="password" size="30" value="Required" onblur="if (this.value == '') {this.value = 'Required';}" onfocus="if (this.value == 'Required') {this.value = '';}">
			</label>
			<br/>
			<br/>
			<button type="submit">Sign In</button>
			</form>
		</fieldset>
		{{ end }}
        <form method="GET" action="/oauth2/start">
        <p>Have a <span>{{.Domain}}</span>.com email address? Sign in with Google.</p>
        <button class="google" type="submit">Sign In</button>
        </form>
	</div>
</body>
</html>
{{end}}`)
	if err != nil {
		log.Fatalf("failed parsing template %s", err.Error())
	}

	t, err = t.Parse(`{{define "error.html"}}
<!DOCTYPE html>
<html lang="en" charset="utf-8">
<head><title>{{.Title}}</title></head>
<body>
	<h2>{{.Title}}</h2>
	<p>{{.Message}}</p>
	<hr>
	<p><a href="/oauth2/sign_in">Sign In</a></p>
</body>
</html>{{end}}`)
	if err != nil {
		log.Fatalf("failed parsing template %s", err.Error())
	}
	return t
}
