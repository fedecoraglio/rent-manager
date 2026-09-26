package com.rentmanager.ai.document;

import java.io.IOException;

import org.junit.jupiter.api.Test;
import org.springframework.http.HttpStatus;
import org.springframework.mock.web.MockMultipartFile;

import static org.assertj.core.api.Assertions.assertThat;
import static org.assertj.core.api.Assertions.assertThatThrownBy;

final class PdfUploadValidatorTest {
    private final PdfUploadValidator validator = new PdfUploadValidator();

    @Test
    void validatesPdfBytesAndRemovesClientDirectoriesFromFilename() throws IOException {
        final byte[] content = PdfFixtures.pdf(false);
        final var result = validator.validate(new MockMultipartFile("file", "C:\\fakepath\\contrato.pdf", "application/pdf", content));
        assertThat(result.filename()).isEqualTo("contrato.pdf");
        assertThat(result.content()).isEqualTo(content);
    }

    @Test
    void rejectsEmptySpoofedMalformedAndEncryptedPdf() throws IOException {
        for (final byte[] content : new byte[][] {new byte[0], "not a PDF".getBytes(), "%PDF-1.7\ninvalid".getBytes(), PdfFixtures.pdf(true)}) {
            assertThatThrownBy(() -> validator.validate(new MockMultipartFile("file", "contract.pdf", "application/pdf", content)))
                    .isInstanceOfSatisfying(DocumentException.class, error -> assertThat(error.status()).isEqualTo(HttpStatus.BAD_REQUEST));
        }
    }

    @Test
    void rejectsUnsupportedTypeAndOversizedFiles() throws IOException {
        assertThatThrownBy(() -> validator.validate(new MockMultipartFile("file", "contract.txt", "text/plain", PdfFixtures.pdf(false))))
                .isInstanceOfSatisfying(DocumentException.class, error -> assertThat(error.status()).isEqualTo(HttpStatus.UNSUPPORTED_MEDIA_TYPE));
        assertThatThrownBy(() -> validator.validate(new MockMultipartFile("file", "contract.pdf", "application/pdf", new byte[PdfUploadValidator.MAX_BYTES + 1])))
                .isInstanceOfSatisfying(DocumentException.class, error -> assertThat(error.status()).isEqualTo(HttpStatus.CONTENT_TOO_LARGE));
    }
}
