package com.rentmanager.ai.adapter.out.ai;

import java.io.IOException;
import java.nio.charset.StandardCharsets;
import java.util.Locale;

import org.apache.pdfbox.Loader;
import org.springframework.http.HttpStatus;
import org.springframework.stereotype.Component;
import org.springframework.web.multipart.MultipartFile;
import com.rentmanager.ai.domain.document.DocumentException;

@Component
public final class PdfUploadValidator {
    public static final int MAX_BYTES = 20 * 1024 * 1024;

    public ValidatedPdf validate(final MultipartFile file) {
        if (file.isEmpty()) {
            throw DocumentException.invalid("The PDF file is empty.");
        }
        if (file.getSize() > MAX_BYTES) {
            throw new DocumentException(HttpStatus.CONTENT_TOO_LARGE, "PDF files must not exceed 20 MiB.");
        }
        final String suppliedName = file.getOriginalFilename();
        final String normalizedName = suppliedName == null ? "" : suppliedName.replace('\\', '/');
        final String filename = normalizedName.substring(normalizedName.lastIndexOf('/') + 1);
        if (filename.isBlank() || filename.length() > 255 || filename.codePoints().anyMatch(Character::isISOControl)) {
            throw DocumentException.invalid("Provide a filename of at most 255 characters without control characters.");
        }
        if (!filename.toLowerCase(Locale.ROOT).endsWith(".pdf")
                || !"application/pdf".equalsIgnoreCase(file.getContentType())) {
            throw new DocumentException(HttpStatus.UNSUPPORTED_MEDIA_TYPE, "Only application/pdf files with a .pdf filename are supported.");
        }
        try (final var input = file.getInputStream()) {
            final byte[] content = input.readNBytes(MAX_BYTES + 1);
            if (content.length > MAX_BYTES) {
                throw new DocumentException(HttpStatus.CONTENT_TOO_LARGE, "PDF files must not exceed 20 MiB.");
            }
            if (content.length < 5 || !new String(content, 0, 5, StandardCharsets.US_ASCII).equals("%PDF-")) {
                throw DocumentException.invalid("The uploaded file is not a valid PDF.");
            }
            try (final var pdf = Loader.loadPDF(content)) {
                if (pdf.isEncrypted() || pdf.getNumberOfPages() == 0) {
                    throw DocumentException.invalid("The PDF must have at least one page and must not be encrypted.");
                }
            }
            return new ValidatedPdf(filename, content);
        } catch (final IOException exception) {
            throw new DocumentException(HttpStatus.BAD_REQUEST, "The PDF could not be read; upload a valid, unencrypted PDF.", exception);
        }
    }

    public record ValidatedPdf(String filename, byte[] content) {
    }
}
