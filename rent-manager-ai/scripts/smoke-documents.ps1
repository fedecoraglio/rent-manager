# Requires PowerShell 7, a Phase 3 service and nomic-embed-text. Uses uniquely scoped documents.
param([string]$BaseUrl = 'http://localhost:8081', [switch]$RestartService,
      [switch]$CheckIndexing, [switch]$CheckEmbeddingFailure)
$ErrorActionPreference = 'Stop'
if (($RestartService -or $CheckIndexing -or $CheckEmbeddingFailure) -and $BaseUrl -ne 'http://localhost:8081') {
    throw 'Docker checks target the local default Compose stack; use the default BaseUrl.'
}
$client = [System.Net.Http.HttpClient]::new()
$client.Timeout = [TimeSpan]::FromSeconds(450)
$scope = 'smoke-' + [Guid]::NewGuid().ToString('N')
$basePath = "$BaseUrl/properties/$scope/contracts/A/documents"
$otherContract = "$BaseUrl/properties/$scope/contracts/B/documents"
$otherProperty = "$BaseUrl/properties/other-$scope/contracts/A/documents"
$created = [System.Collections.Generic.List[string]]::new()

function Assert-Status($Response, [int]$Expected) {
    if ([int]$Response.StatusCode -ne $Expected) {
        throw "Expected HTTP $Expected, received $([int]$Response.StatusCode): $($Response.Content.ReadAsStringAsync().GetAwaiter().GetResult())"
    }
}

function Send-Pdf([string]$Url, [byte[]]$Bytes, [string]$Type = 'RENT_CONTRACT', [string]$Mime = 'application/pdf') {
    $form = [System.Net.Http.MultipartFormDataContent]::new()
    try {
        $part = [System.Net.Http.ByteArrayContent]::new($Bytes)
        $part.Headers.ContentType = [System.Net.Http.Headers.MediaTypeHeaderValue]::new($Mime)
        $form.Add($part, 'file', 'contract.pdf')
        $form.Add([System.Net.Http.StringContent]::new($Type), 'type')
        return $client.PostAsync($Url, $form).GetAwaiter().GetResult()
    } finally {
        $form.Dispose()
    }
}

function Upload-Document([string]$Url, [byte[]]$Bytes) {
    $response = Send-Pdf $Url $Bytes
    try {
        Assert-Status $response 201
        $document = $response.Content.ReadAsStringAsync().GetAwaiter().GetResult() | ConvertFrom-Json
        $created.Add("$Url/$($document.id)")
        if ($document.filename -ne 'contract.pdf' -or $document.size -ne $Bytes.Length -or $document.type -ne 'RENT_CONTRACT') {
            throw 'Upload metadata does not match the PDF'
        }
        if ($document.PSObject.Properties.Name -contains 'storagePath') { throw 'Response exposes internal storage path' }
        if (-not $response.Headers.Location.ToString().EndsWith("/$($document.id)")) { throw 'Missing Location header' }
        return $document
    } finally { $response.Dispose() }
}

function Assert-List([string]$Url, [int]$Count) {
    $response = $client.GetAsync($Url).GetAwaiter().GetResult()
    try {
        Assert-Status $response 200
        $items = @($response.Content.ReadAsStringAsync().GetAwaiter().GetResult() | ConvertFrom-Json)
        if ($items.Count -ne $Count) { throw "Expected $Count documents, received $($items.Count)" }
    } finally { $response.Dispose() }
}

function Invoke-Sql([string]$Sql) {
    $result = docker compose -f "$PSScriptRoot/../compose.yml" exec -T postgres sh -c 'exec psql -U "$POSTGRES_USER" -d "$POSTGRES_DB" -At -v ON_ERROR_STOP=1 -c "$1"' -- $Sql
    if ($LASTEXITCODE -ne 0) { throw 'SQL verification failed' }
    return ($result -join "`n").Trim()
}

function Assert-Chunks($Document, [string]$Contract, [int]$Expected) {
    $id = [Guid]::Parse($Document.id).ToString()
    $count = Invoke-Sql "SELECT count(*) FROM vector_store WHERE document_id = '$id'::uuid AND metadata->>'propertyId' = '$scope' AND metadata->>'contractId' = '$Contract' AND metadata->>'documentType' = 'RENT_CONTRACT' AND metadata->>'filename' = 'contract.pdf' AND metadata->>'page' IN ('1', '2') AND length(content) > 0 AND vector_dims(embedding) = 768;"
    if ([int]$count -ne $Expected) { throw "Expected $Expected scoped 768-dimensional chunks, received $count" }
}

# Valid two-page PDF with searchable text, generated without external PDF utilities.
function New-TestPdf([switch]$Blank) {
$pdfText = "%PDF-1.4`n"
$offsets = [System.Collections.Generic.List[int]]::new()
$firstText = if ($Blank) { '' } else { 'BT /F1 12 Tf 50 700 Td (The rental contract expires on December 31, 2027. Rent is adjusted quarterly.) Tj ET' }
$secondText = if ($Blank) { '' } else { 'BT /F1 12 Tf 50 700 Td (The guarantor is Juan Perez. Pets are allowed with written permission.) Tj ET' }
$objects = @(
    '<< /Type /Catalog /Pages 2 0 R >>',
    '<< /Type /Pages /Kids [3 0 R 6 0 R] /Count 2 >>',
    '<< /Type /Page /Parent 2 0 R /MediaBox [0 0 612 792] /Resources << /Font << /F1 4 0 R >> >> /Contents 5 0 R >>',
    '<< /Type /Font /Subtype /Type1 /BaseFont /Helvetica >>',
    "<< /Length $($firstText.Length) >>`nstream`n$firstText`nendstream",
    '<< /Type /Page /Parent 2 0 R /MediaBox [0 0 612 792] /Resources << /Font << /F1 4 0 R >> >> /Contents 7 0 R >>',
    "<< /Length $($secondText.Length) >>`nstream`n$secondText`nendstream"
)
for ($index = 0; $index -lt $objects.Count; $index++) {
    $offsets.Add($pdfText.Length)
    $pdfText += "$($index + 1) 0 obj`n$($objects[$index])`nendobj`n"
}
$xref = $pdfText.Length
$pdfText += "xref`n0 $($objects.Count + 1)`n0000000000 65535 f `n"
foreach ($offset in $offsets) { $pdfText += $offset.ToString('D10') + " 00000 n `n" }
$pdfText += "trailer`n<< /Size $($objects.Count + 1) /Root 1 0 R >>`nstartxref`n$xref`n%%EOF`n"
return ,([System.Text.Encoding]::ASCII.GetBytes($pdfText))
}
$pdfBytes = New-TestPdf

try {
    Assert-List $basePath 0
    $first = Upload-Document $basePath $pdfBytes
    $second = Upload-Document $basePath $pdfBytes
    $third = Upload-Document $otherContract $pdfBytes
    if ($first.id -eq $second.id) { throw 'Duplicate filenames overwrote one another' }
    if ($RestartService) {
        docker compose -f "$PSScriptRoot/../compose.yml" restart rent-manager-ai
        if ($LASTEXITCODE -ne 0) { throw 'Compose restart failed' }
        $ready = $false
        for ($attempt = 0; $attempt -lt 30; $attempt++) {
            try {
                $probe = $client.GetAsync("$BaseUrl/actuator/health/liveness").GetAwaiter().GetResult()
                try { $ready = [int]$probe.StatusCode -eq 200 } finally { $probe.Dispose() }
            } catch { $ready = $false }
            if ($ready) { break }
            Start-Sleep -Seconds 1
        }
        if (-not $ready) { throw 'Service did not recover after restart' }
    }
    Assert-List $basePath 2
    Assert-List $otherContract 1
    Assert-List $otherProperty 0
    if ($CheckIndexing) {
        Assert-Chunks $first 'A' 2
        Assert-Chunks $second 'A' 2
        Assert-Chunks $third 'B' 2
    }

    $response = $client.GetAsync("$basePath/$($first.id)?download=true").GetAwaiter().GetResult()
    try {
        Assert-Status $response 200
        $download = $response.Content.ReadAsByteArrayAsync().GetAwaiter().GetResult()
        if ([Convert]::ToBase64String($download) -ne [Convert]::ToBase64String($pdfBytes)) { throw 'Downloaded bytes differ from uploaded PDF' }
        if ($response.Content.Headers.ContentType.MediaType -ne 'application/pdf') { throw 'Incorrect download content type' }
        if ($response.Content.Headers.ContentDisposition.DispositionType -ne 'attachment') { throw 'Incorrect download disposition' }
    } finally { $response.Dispose() }

    foreach ($foreign in @($otherContract, $otherProperty)) {
        $response = $client.GetAsync("$foreign/$($first.id)").GetAwaiter().GetResult()
        try { Assert-Status $response 404 } finally { $response.Dispose() }
        $response = $client.DeleteAsync("$foreign/$($first.id)").GetAwaiter().GetResult()
        try { Assert-Status $response 404 } finally { $response.Dispose() }
    }

    foreach ($badBytes in @([byte[]]::new(0), [System.Text.Encoding]::ASCII.GetBytes('%PDF-1.4 invalid'))) {
        $response = Send-Pdf $basePath $badBytes
        try { Assert-Status $response 400 } finally { $response.Dispose() }
    }
    $response = Send-Pdf $basePath $pdfBytes 'UNKNOWN'
    try { Assert-Status $response 400 } finally { $response.Dispose() }
    $response = Send-Pdf $basePath $pdfBytes 'OTHER' 'text/plain'
    try { Assert-Status $response 415 } finally { $response.Dispose() }
    $response = Send-Pdf $basePath ([byte[]]::new(20 * 1024 * 1024 + 1))
    try { Assert-Status $response 413 } finally { $response.Dispose() }
    $response = Send-Pdf $basePath (New-TestPdf -Blank)
    try { Assert-Status $response 422 } finally { $response.Dispose() }
    if ($CheckEmbeddingFailure) {
        docker compose -f "$PSScriptRoot/../compose.yml" stop ollama
        if ($LASTEXITCODE -ne 0) { throw 'Ollama stop failed' }
        try {
            $response = Send-Pdf $basePath $pdfBytes
            try { Assert-Status $response 503 } finally { $response.Dispose() }
        } finally {
            docker compose -f "$PSScriptRoot/../compose.yml" start ollama
            if ($LASTEXITCODE -ne 0) { throw 'Ollama restart failed' }
        }
    }
    Assert-List $basePath 2

    $response = $client.DeleteAsync("$basePath/$($first.id)").GetAwaiter().GetResult()
    try { Assert-Status $response 204 } finally { $response.Dispose() }
    $response = $client.GetAsync("$basePath/$($first.id)").GetAwaiter().GetResult()
    try { Assert-Status $response 404 } finally { $response.Dispose() }
    Assert-List $basePath 1
    Assert-List $otherContract 1
    if ($CheckIndexing) {
        Assert-Chunks $first 'A' 0
        Assert-Chunks $second 'A' 2
        Assert-Chunks $third 'B' 2
        if ([int](Invoke-Sql "SELECT count(*) FROM vector_store WHERE metadata->>'propertyId' = '$scope';") -ne 4) {
            throw 'Unexpected chunks remained after failed uploads or deletion'
        }
        Write-Output 'PASS: real Ollama embeddings, 768 dimensions, document/page metadata and scoped vector cleanup.'
    }
    Write-Output 'PASS: upload/list/download/delete, duplicate filenames, scope isolation, validation and multipart size limit.'
    if ($RestartService) { Write-Output 'PASS: document registry and PDF bytes survive service restart.' }
    if ($CheckEmbeddingFailure) { Write-Output 'PASS: Ollama outage returns 503 and rolls back the upload.' }
} finally {
    foreach ($url in $created) {
        $response = $client.DeleteAsync($url).GetAwaiter().GetResult()
        try {
            if ([int]$response.StatusCode -notin @(204, 404)) { throw "Smoke-test cleanup failed for $url" }
        } finally { $response.Dispose() }
    }
    $client.Dispose()
}
