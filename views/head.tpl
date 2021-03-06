<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <!-- The above 3 meta tags *must* come first in the head; any other head content must come *after* these tags -->
    
    <title>Upload your images</title>
    <link href="/static/css/index.css" rel="stylesheet" type="text/css" />
    <!-- <link href="/static/css/bootstrap.min.css" rel="stylesheet" type="text/css" /> -->
    <!-- <link href="/static/css/flatty.min.css" rel="stylesheet" type="text/css" /> -->
    {{ if ne .theme "" }}
        {{ if eq .theme "flatty" }}
            <link href="/static/css/flatty.min.css" rel="stylesheet" type="text/css" />
        {{ end }}
        {{ if eq .theme "solar" }}
            <link href="/static/css/solar.min.css" rel="stylesheet" type="text/css" />
        {{ end }}
        {{ if eq .theme "darkly" }}
            <link href="/static/css/darkly.min.css" rel="stylesheet" type="text/css" />
        {{ end }}
        {{ if eq .theme "superhero" }}
            <link href="/static/css/superhero.min.css" rel="stylesheet" type="text/css" />
        {{ end }}
    {{ else }}
        <link href="/static/css/flatty.min.css" rel="stylesheet" type="text/css" />
    {{ end }}
    <link href="/static/css/fileinput.min.css" rel="stylesheet" type="text/css" />
    <link href="/static/css/bootstrap-gallery.css" rel="stylesheet" type="text/css" />
</head>