package com.rentmanager.ai.document;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.UUID;

import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.condition.EnabledOnOs;
import org.junit.jupiter.api.condition.OS;
import org.junit.jupiter.api.io.TempDir;

import static org.assertj.core.api.Assertions.assertThat;
import static org.assertj.core.api.Assertions.assertThatThrownBy;

final class FileSystemDocumentStorageTest {
    @TempDir
    private Path root;

    @Test
    void savesDistinctFilesAndCanRestoreOrFinishDeletion() throws IOException {
        final var storage = new FileSystemDocumentStorage(root.toString());
        final var scope = new DocumentScope("123", "456");
        final byte[] bytes = PdfFixtures.pdf(false);
        final String first = storage.save(scope, UUID.randomUUID(), bytes);
        final String second = storage.save(scope, UUID.randomUUID(), bytes);
        assertThat(first).isNotEqualTo(second).startsWith("properties/123/contracts/456/");
        assertThat(storage.read(first)).isEqualTo(bytes);
        assertThat(storage.stageDeletion(first)).isTrue();
        assertThat(root.resolve(first)).doesNotExist();
        storage.restoreDeletion(first);
        assertThat(storage.read(first)).isEqualTo(bytes);
        storage.stageDeletion(first);
        storage.completeDeletion(first);
        assertThat(root.resolve(first + ".deleting")).doesNotExist();
        assertThat(storage.read(second)).isEqualTo(bytes);
        assertThat(storage.stageDeletion(first)).isFalse();
    }

    @Test
    void rejectsTraversalAndNeverOverwritesExistingDocuments() throws IOException {
        final var storage = new FileSystemDocumentStorage(root.toString());
        assertThatThrownBy(() -> new DocumentScope("..", "456")).isInstanceOf(DocumentException.class);
        assertThatThrownBy(() -> storage.read("../outside.pdf")).isInstanceOf(DocumentException.class);
        assertThatThrownBy(() -> storage.remove(".")).isInstanceOf(DocumentException.class);
        final UUID id = UUID.randomUUID();
        final var scope = new DocumentScope("123", "456");
        final byte[] original = PdfFixtures.pdf(false);
        final String path = storage.save(scope, id, original);
        assertThatThrownBy(() -> storage.save(scope, id, new byte[] {1, 2})).isInstanceOf(DocumentException.class);
        assertThat(storage.read(path)).isEqualTo(original);
        try (final var files = Files.walk(root)) {
            assertThat(files.filter(Files::isRegularFile).toList()).containsExactly(root.resolve(path));
        }
    }

    @Test
    @EnabledOnOs({OS.LINUX, OS.MAC}) // Windows requires optional symlink privileges.
    void rejectsSymbolicLinksOutsideStorage() throws IOException {
        final var storage = new FileSystemDocumentStorage(root.toString());
        final Path outside = Files.createTempDirectory(root.getParent(), "outside-");
        try {
            Files.createSymbolicLink(root.resolve("properties"), outside);
            assertThatThrownBy(() -> storage.save(new DocumentScope("123", "456"), UUID.randomUUID(), new byte[] {1}))
                    .isInstanceOf(DocumentException.class);
            try (final var files = Files.list(outside)) {
                assertThat(files.count()).isZero();
            }
        } finally {
            Files.deleteIfExists(root.resolve("properties"));
            Files.deleteIfExists(outside);
        }
    }
}
