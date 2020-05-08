<html lang="en">
    {{ template "head.tpl" . }}
    <body>
    {{ template "navbar.tpl" . }}
    <div class="container">
        {{ template "jumbotron.tpl" . }}
        {{ template "form.tpl" . }}
        {{ template "flash.tpl" . }}
        {{ template "thumbnails.tpl" . }}
        
    </div>
    <script src="https://ajax.googleapis.com/ajax/libs/jquery/2.2.4/jquery.min.js"></script>
        <script src="static/js/fileinput.min.js" ></script>
        <script src="static/js/bootstrap.min.js" ></script>
        <script src="static/js/bootstrap-gallery.js" ></script>
        <script>
            // bootstrap-fileinput.js
            // initialize form (id='image') with defaults
            $("#image").fileinput();
        </script>
        <script>
            //
            // Get the clicked theme, and save it in a cookie
            //

            $("[id^=theme]").click(function(selector, options){
                var theme = "";

                switch(this.id) {
                    case "theme1":
                        theme = "flatty";
                        break;
                    case "theme2":
                        theme = "solar";
                        break;
                    case "theme3":
                        theme = "darkly";
                        break;
                    case "theme4":
                        theme = "superhero";
                        break;

                    // Should never go here
                    default:
                        theme = "flatty";
                } 

                document.cookie = "theme=" + theme;
                location.reload();
            });
        </script>
 </html>