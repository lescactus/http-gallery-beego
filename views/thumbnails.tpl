{{ $uploadDir := .uploadDirectory }}
{{ $thumbnailsDir := .thumbnailDirectory }}
<div class="row">
    {{ range $image, $thumbnail := .images }}
    <div class="col-xs-6 col-sm-3">
        <a href="{{ $uploadDir }}{{ $image }}" class="thumbnail" >
            <img src="{{ $thumbnailsDir }}{{ $thumbnail }}" alt="image">
        </a>
    </div>
    {{ end }}
</div>