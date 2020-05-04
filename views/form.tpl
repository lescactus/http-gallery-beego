<div class="row">
    <form action="/" method="post" enctype="multipart/form-data" class="form-image">
        <label class="control-label">Select a file to upload</label>
        {{ .xsrfdata }}
        <input name="{{ .htmlInputName }}" id="file" type="file" class="file" data-preview-file-type="text">
        {{/* {{ form.image }} */}}
    </form>
</div>