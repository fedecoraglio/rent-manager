package com.rentmanager.ai.indexing;

import java.io.IOException;
import java.time.Instant;
import java.util.UUID;

import com.rentmanager.ai.document.DocumentException;
import com.rentmanager.ai.document.DocumentMetadata;
import com.rentmanager.ai.document.DocumentType;
import com.rentmanager.ai.document.PdfFixtures;
import org.junit.jupiter.api.Test;

import static org.assertj.core.api.Assertions.assertThat;
import static org.assertj.core.api.Assertions.assertThatThrownBy;

final class PdfChunkerTest {
    private final PdfChunker chunker = new PdfChunker();
    private final UUID id = UUID.randomUUID();
    private final DocumentMetadata metadata = new DocumentMetadata(id, "123", "456", DocumentType.RENT_CONTRACT,
            "contrato.pdf", id + ".pdf", "unused", "application/pdf", 1, Instant.now(), Instant.now());

    @Test
    void splitsLongPagesAndPreservesPhysicalPageNumbersAndAllScopeMetadata() throws IOException {
        final String longPage = "The tenant pays rent monthly. Repairs require the written consent of the owner.\n".repeat(75);
        final var chunks = chunker.chunks(metadata, PdfFixtures.textPdf(longPage, "", "Pets are allowed. El garante es Juan Pérez."));
        assertThat(chunks.stream().filter(chunk -> chunk.getMetadata().get("page").equals(1)).count()).isGreaterThan(1);
        assertThat(chunks.stream().map(chunk -> chunk.getMetadata().get("page"))).containsOnly(1, 3);
        assertThat(chunks).allSatisfy(chunk -> {
            assertThat(chunk.getId()).isNotEqualTo(id.toString());
            assertThat(chunk.getText()).isNotBlank();
            assertThat(chunk.getMetadata()).containsEntry("documentId", id.toString())
                    .containsEntry("propertyId", "123").containsEntry("contractId", "456")
                    .containsEntry("documentType", "RENT_CONTRACT").containsEntry("filename", "contrato.pdf");
        });
        assertThat(chunks.getLast().getText().replaceAll("\\s+", " ")).contains("Pets are allowed", "Juan Pérez");
        assertThat(chunks.stream().map(chunk -> chunk.getId()).distinct().count()).isEqualTo(chunks.size());
    }

    @Test
    void rejectsPdfWithoutReadableTextAndDoesNotDiscardShortClauses() throws IOException {
        assertThatThrownBy(() -> chunker.chunks(metadata, PdfFixtures.textPdf("", "")))
                .isInstanceOf(DocumentException.class).hasMessageContaining("No readable text");
        assertThat(chunker.chunks(metadata, PdfFixtures.textPdf("No."))).hasSize(1);
    }
}
