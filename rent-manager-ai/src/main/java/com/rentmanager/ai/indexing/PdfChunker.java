package com.rentmanager.ai.indexing;

import java.io.IOException;
import java.util.List;
import java.util.Map;

import com.rentmanager.ai.document.DocumentException;
import com.rentmanager.ai.document.DocumentMetadata;
import org.springframework.ai.document.Document;
import org.springframework.ai.reader.pdf.PagePdfDocumentReader;
import org.springframework.ai.reader.pdf.config.PdfDocumentReaderConfig;
import org.springframework.ai.transformer.splitter.TokenTextSplitter;
import org.springframework.core.io.ByteArrayResource;
import org.springframework.http.HttpStatus;
import org.springframework.stereotype.Component;

@Component
public final class PdfChunker {
    public List<Document> chunks(final DocumentMetadata document, final byte[] pdf) {
        final List<Document> pages;
        // One page per Document keeps the physical page number accurate after splitting.
        try (final var reader = new CloseablePageReader(pdf)) {
            pages = reader.get();
        } catch (final IOException | RuntimeException exception) {
            throw new DocumentException(HttpStatus.UNPROCESSABLE_CONTENT, "Text could not be extracted from this PDF.", exception);
        }

        final List<Document> scopedPages = pages.stream()
                .filter(page -> page.getText() != null && !page.getText().isBlank())
                .map(page -> new Document(page.getText(), Map.of(
                        "documentId", document.id().toString(),
                        "propertyId", document.propertyId(),
                        "contractId", document.contractId(),
                        "documentType", document.type().name(),
                        "filename", document.originalFilename(),
                        "page", page.getMetadata().get(PagePdfDocumentReader.METADATA_START_PAGE_NUMBER))))
                .toList();

        // The splitter copies each page's metadata onto every resulting chunk.
        final var splitter = TokenTextSplitter.builder().withChunkSize(800)
                .withMinChunkLengthToEmbed(0).build();
        final List<Document> chunks = splitter.apply(scopedPages);
        if (chunks.isEmpty()) {
            throw new DocumentException(HttpStatus.UNPROCESSABLE_CONTENT,
                    "No readable text was found in the PDF. Scanned documents require OCR, which is not supported.");
        }
        return chunks;
    }

    // Spring AI 2.0.1 exposes its PDFBox document as protected but has no close method.
    private static final class CloseablePageReader extends PagePdfDocumentReader implements AutoCloseable {
        private CloseablePageReader(final byte[] pdf) {
            super(new ByteArrayResource(pdf), PdfDocumentReaderConfig.builder().withPagesPerDocument(1).build());
        }

        @Override
        public void close() throws IOException {
            document.close();
        }
    }
}
