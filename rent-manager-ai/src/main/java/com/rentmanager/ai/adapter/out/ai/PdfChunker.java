package com.rentmanager.ai.adapter.out.ai;

import java.io.IOException;
import java.util.ArrayList;
import java.util.List;
import java.util.Map;
import java.util.regex.Pattern;

import com.rentmanager.ai.domain.document.DocumentException;
import com.rentmanager.ai.domain.document.DocumentMetadata;
import org.springframework.ai.document.Document;
import org.springframework.ai.reader.pdf.PagePdfDocumentReader;
import org.springframework.ai.reader.pdf.config.PdfDocumentReaderConfig;
import org.springframework.ai.transformer.splitter.TokenTextSplitter;
import org.springframework.core.io.ByteArrayResource;
import org.springframework.http.HttpStatus;
import org.springframework.stereotype.Component;

@Component
public final class PdfChunker {

    private static final Pattern CLAUSE_START = Pattern.compile(
            "(?=\\b[\\p{Lu}ÁÉÍÓÚÜÑ]{3,}(?:\\s+[\\p{Lu}ÁÉÍÓÚÜÑ]{3,}){0,3}\\s*:)"
    );

    public List<Document> chunks(final DocumentMetadata document, final byte[] pdf) {
        final List<Document> pages;

        try (final var reader = new CloseablePageReader(pdf)) {
            pages = reader.get();
        } catch (final IOException | RuntimeException exception) {
            throw new DocumentException(
                    HttpStatus.UNPROCESSABLE_CONTENT,
                    "Text could not be extracted from this PDF.",
                    exception
            );
        }

        final List<Document> clauseDocuments = new ArrayList<>();

        for (final Document page : pages) {
            if (page.getText() == null || page.getText().isBlank()) {
                continue;
            }

            final String text = normalize(page.getText());

            final Map<String, Object> metadata = Map.of(
                    "documentId", document.id().toString(),
                    "propertyId", document.propertyId(),
                    "contractId", document.contractId(),
                    "documentType", document.type().name(),
                    "filename", document.originalFilename(),
                    "page", page.getMetadata().get(
                            PagePdfDocumentReader.METADATA_START_PAGE_NUMBER
                    )
            );

            for (final String clause : splitClauses(text)) {
                if (!clause.isBlank()) {
                    clauseDocuments.add(new Document(clause, metadata));
                }
            }
        }

        /*
         * Clause boundaries are preserved first.
         * Token splitting is only needed when a clause itself is too large.
         */
        final var splitter = TokenTextSplitter.builder()
                .withChunkSize(500)
                .withMinChunkSizeChars(100)
                .withMinChunkLengthToEmbed(5)
                .withKeepSeparator(true)
                .build();

        final List<Document> chunks = splitter.apply(clauseDocuments);

        if (chunks.isEmpty()) {
            throw new DocumentException(
                    HttpStatus.UNPROCESSABLE_CONTENT,
                    "No readable text was found in the PDF. Scanned documents require OCR, which is not supported."
            );
        }

        return chunks;
    }

    private static List<String> splitClauses(final String text) {
        return CLAUSE_START.splitAsStream(text).map(String::trim).filter(value -> !value.isBlank()).toList();
    }

    private static String normalize(final String text) {
        return text
                .replace('\u00A0', ' ')
                .replaceAll("\\s+", " ")
                .trim();
    }

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