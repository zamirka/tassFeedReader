<html>
    <head>
    <title>Лента новостей ТАСС</title>
    </head>
    <body>
        <div>
        {{range .}}
            <div><a href="{{.Url}}" target="_blank">{{.Title}}<a></div>
        {{end}}        
        </div>
    </body>
</html>