<html lang="en">
    {{ template "head.tpl" . }}
    <body>
    {{ template "navbar.tpl" . }}
    <div class="container">
        {{ template "jumbotron.tpl" . }}
        {{ template "form.tpl" . }}
        {{ template "flash.tpl" . }}
    </div>
    <script src="https://ajax.googleapis.com/ajax/libs/jquery/2.2.4/jquery.min.js"></script>
        <script src="static/js/fileinput.min.js" ></script>
        <script src="static/js/bootstrap.min.js" ></script>
        <script src="static/js/bootstrap-gallery.js" ></script>
    <form action="/" method="post" enctype="multipart/form-data">
        <label for="file">Filename:</label>
        <input type="file" name="file" id="file">
        <input type="submit" name="submit" value="Submit">
        {{ .xsrfdata }}
    </form>
    </body>
 </html>