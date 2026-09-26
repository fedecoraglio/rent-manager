package com.rentmanager.ai.document;

import java.net.URI;
import java.nio.charset.StandardCharsets;
import java.util.List;
import java.util.UUID;

import org.springframework.http.ContentDisposition;
import org.springframework.http.HttpHeaders;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.multipart.MultipartFile;

@RestController
@RequestMapping("/properties/{propertyId}/contracts/{contractId}/documents")
public final class DocumentController {
    private final DocumentService service;

    public DocumentController(final DocumentService service) {
        this.service = service;
    }

    @PostMapping(consumes = MediaType.MULTIPART_FORM_DATA_VALUE)
    public ResponseEntity<DocumentResponse> upload(@PathVariable final String propertyId,
            @PathVariable final String contractId, @RequestParam("type") final DocumentType type,
            @RequestParam("file") final MultipartFile file) {
        final var document = service.upload(new DocumentScope(propertyId, contractId), type, file);
        final URI location = URI.create("/properties/" + propertyId + "/contracts/" + contractId + "/documents/" + document.id());
        return ResponseEntity.created(location).body(DocumentResponse.from(document));
    }

    @GetMapping
    public List<DocumentResponse> list(@PathVariable final String propertyId, @PathVariable final String contractId) {
        return service.list(new DocumentScope(propertyId, contractId)).stream().map(DocumentResponse::from).toList();
    }

    @GetMapping("/{documentId}")
    public ResponseEntity<byte[]> download(@PathVariable final String propertyId, @PathVariable final String contractId,
            @PathVariable final UUID documentId, @RequestParam(defaultValue = "false") final boolean download) {
        final var result = service.download(new DocumentScope(propertyId, contractId), documentId);
        final var disposition = (download ? ContentDisposition.attachment() : ContentDisposition.inline())
                .filename(result.document().originalFilename(), StandardCharsets.UTF_8).build();
        return ResponseEntity.ok().contentType(MediaType.APPLICATION_PDF).contentLength(result.content().length)
                .header(HttpHeaders.CONTENT_DISPOSITION, disposition.toString())
                .header(HttpHeaders.CACHE_CONTROL, "no-store")
                .header("X-Content-Type-Options", "nosniff")
                .body(result.content());
    }

    @DeleteMapping("/{documentId}")
    public ResponseEntity<Void> delete(@PathVariable final String propertyId, @PathVariable final String contractId,
            @PathVariable final UUID documentId) {
        service.delete(new DocumentScope(propertyId, contractId), documentId);
        return ResponseEntity.noContent().build();
    }
}
