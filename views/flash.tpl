<div class="row">
    {{ if .flash.error }}
    <div class="alert alert-danger">
        <p>{{ .flash.error }}</p>
    </div>
    {{ end }}
    {{ if .flash.success }}
    <div class="alert alert-success">
        <p>{{ .flash.success }} </p>
    </div>
    {{ end }}
</div>