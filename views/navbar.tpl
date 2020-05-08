<!-- Fixed navbar -->
<nav class="navbar navbar-default navbar-fixed-top">
    <div class="container">
        <div class="navbar-header">
           <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#navbar" aria-expanded="false" aria-controls="navbar">
           <span class="sr-only">Toggle navigation</span>
           <span class="icon-bar"></span>
           <span class="icon-bar"></span>
           <span class="icon-bar"></span>
           </button>
           <a class="navbar-brand" href="/">Image gallery</a>
        </div>
        <div id="navbar" class="navbar-collapse collapse">
           <ul class="nav navbar-nav">
              <li class="active"><a href="/">Index</a></li>
              <li class="dropdown">
                 <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">Themes <span class="caret"></span></a>
                 <ul class="dropdown-menu">
                    {{ if ne .theme "" }}
                        {{ if eq .theme "flatty" }}
                            <li><a href="#" id="theme1">Flatty</a></li>
                            <li role="separator" class="divider"></li>
                            <li class="dropdown-header">Availables themes</li>
                            <li><a href="#" id="theme2">Solar</a></li>
                            <li><a href="#" id="theme3">Darkly</a></li>
                            <li><a href="#" id="theme4">Superhero</a></li>
                        {{ end }}
                        {{ if eq .theme "solar" }}
                            <li><a href="#" id="theme2">Solar</a></li>
                            <li role="separator" class="divider"></li>
                            <li class="dropdown-header">Availables themes</li>
                            <li><a href="#" id="theme1">Flatty</a></li>
                            <li><a href="#" id="theme3">Darkly</a></li>
                            <li><a href="#" id="theme4">Superhero</a></li>
                        {{ end }}
                        {{ if eq .theme "darkly" }}
                            <li><a href="#" id="theme3">Darkly</a></li>
                            <li role="separator" class="divider"></li>
                            <li class="dropdown-header">Availables themes</li>
                            <li><a href="#" id="theme1">Flatty</a></li>
                            <li><a href="#" id="theme2">Solar</a></li>
                            <li><a href="#" id="theme4">Superhero</a></li>
                        {{ end }}
                        {{ if eq .theme "superhero" }}
                            <li><a href="#" id="theme4">Superhero</a></li>
                            <li role="separator" class="divider"></li>
                            <li class="dropdown-header">Availables themes</li>
                            <li><a href="#" id="theme1">Flatty</a></li>
                            <li><a href="#" id="theme2">Solar</a></li>
                            <li><a href="#" id="theme3">Darkly</a></li>
                        {{ end }}
                    {{ else }}
                        <li><a href="#" id="theme1">Flatty</a></li>
                        <li><a href="#" id="theme2">Solar</a></li>
                        <li><a href="#" id="theme3">Darkly</a></li>
                        <li><a href="#" id="theme4">Superhero</a></li>
                    {{ end }}
                     {{/* 
                    {% if theme %}
                        {% if theme == "flatty" %}
                            <li><a href="#" id="theme1">Flatty</a></li>
                            <li role="separator" class="divider"></li>
                            <li class="dropdown-header">Availables themes</li>
                            <li><a href="#" id="theme2">Solar</a></li>
                            <li><a href="#" id="theme3">Darkly</a></li>
                            <li><a href="#" id="theme4">Superhero</a></li>
                        {% elif theme == "solar" %}
                            <li><a href="#" id="theme2">Solar</a></li>
                            <li role="separator" class="divider"></li>
                            <li class="dropdown-header">Availables themes</li>
                            <li><a href="#" id="theme1">Flatty</a></li>
                            <li><a href="#" id="theme3">Darkly</a></li>
                            <li><a href="#" id="theme4">Superhero</a></li>
                        {% elif theme == "darkly" %}
                            <li><a href="#" id="theme3">Darkly</a></li>
                            <li role="separator" class="divider"></li>
                            <li class="dropdown-header">Availables themes</li>
                            <li><a href="#" id="theme1">Flatty</a></li>
                            <li><a href="#" id="theme2">Solar</a></li>
                            <li><a href="#" id="theme4">Superhero</a></li>
                        {% elif theme == "superhero" %}
                            <li><a href="#" id="theme4">Superhero</a></li>
                            <li role="separator" class="divider"></li>
                            <li class="dropdown-header">Availables themes</li>
                            <li><a href="#" id="theme1">Flatty</a></li>
                            <li><a href="#" id="theme2">Solar</a></li>
                            <li><a href="#" id="theme3">Darkly</a></li>
                        {% endif %}
                    {% else %} {# Default if no cookie set #}
                        <li><a href="#" id="theme1">Flatty</a></li>
                        <li><a href="#" id="theme2">Solar</a></li>
                        <li><a href="#" id="theme3">Darkly</a></li>
                        <li><a href="#" id="theme4">Superhero</a></li>
                    {% endif %}
                    */}}
                 </ul>
              </li>
           </ul>
        </div>
        <!--/.nav-collapse -->
    </div>
</nav>