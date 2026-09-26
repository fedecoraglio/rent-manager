package com.rentmanager.ai.document;

import java.io.ByteArrayOutputStream;
import java.io.IOException;

import org.apache.pdfbox.pdmodel.PDDocument;
import org.apache.pdfbox.pdmodel.PDPage;
import org.apache.pdfbox.pdmodel.PDPageContentStream;
import org.apache.pdfbox.pdmodel.font.PDType1Font;
import org.apache.pdfbox.pdmodel.font.Standard14Fonts;
import org.apache.pdfbox.pdmodel.encryption.AccessPermission;
import org.apache.pdfbox.pdmodel.encryption.StandardProtectionPolicy;
import org.springframework.mock.web.MockMultipartFile;

public final class PdfFixtures {
    private PdfFixtures() {
    }

    static byte[] pdf(final boolean encrypted) throws IOException {
        try (final var pdf = new PDDocument(); final var output = new ByteArrayOutputStream()) {
            pdf.addPage(new PDPage());
            if (encrypted) {
                pdf.protect(new StandardProtectionPolicy("owner-password", "user-password", new AccessPermission()));
            }
            pdf.save(output);
            return output.toByteArray();
        }
    }

    static MockMultipartFile upload() throws IOException {
        return new MockMultipartFile("file", "contract.pdf", "application/pdf", pdf(false));
    }

    public static byte[] textPdf(final String... pages) throws IOException {
        try (final var pdf = new PDDocument(); final var output = new ByteArrayOutputStream()) {
            for (final String text : pages) {
                final var page = new PDPage();
                pdf.addPage(page);
                try (final var stream = new PDPageContentStream(pdf, page)) {
                    stream.beginText();
                    stream.setFont(new PDType1Font(Standard14Fonts.FontName.HELVETICA), 8);
                    stream.setLeading(9);
                    stream.newLineAtOffset(50, 740);
                    for (final String line : text.split("\n")) {
                        stream.showText(line);
                        stream.newLine();
                    }
                    stream.endText();
                }
            }
            pdf.save(output);
            return output.toByteArray();
        }
    }
}
