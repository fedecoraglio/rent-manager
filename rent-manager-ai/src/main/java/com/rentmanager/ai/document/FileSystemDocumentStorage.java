package com.rentmanager.ai.document;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.LinkOption;
import java.nio.file.Path;
import java.util.UUID;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.http.HttpStatus;
import org.springframework.stereotype.Component;

@Component
public final class FileSystemDocumentStorage implements DocumentStorage {
    private final Path root;

    public FileSystemDocumentStorage(@Value("${app.documents.storage-path}") final String root) throws IOException {
        this.root = Files.createDirectories(Path.of(root).toAbsolutePath().normalize()).toRealPath();
    }

    @Override
    public String save(final DocumentScope scope, final UUID id, final byte[] content) {
        final String relative = "properties/" + scope.propertyId() + "/contracts/" + scope.contractId() + "/" + id + ".pdf";
        final Path target = resolve(relative);
        try {
            Files.createDirectories(target.getParent());
            resolve(relative);
            final Path temporary = Files.createTempFile(target.getParent(), ".upload-", ".tmp");
            try {
                Files.write(temporary, content);
                Files.move(temporary, target);
            } catch (final IOException exception) {
                try {
                    Files.deleteIfExists(temporary);
                } catch (final IOException cleanup) {
                    exception.addSuppressed(cleanup);
                }
                throw exception;
            }
            return relative;
        } catch (final IOException exception) {
            throw unavailable(exception);
        }
    }

    @Override
    public byte[] read(final String path) {
        final Path file = resolve(path);
        try {
            if (!Files.isRegularFile(file, LinkOption.NOFOLLOW_LINKS)) {
                throw new DocumentException(HttpStatus.NOT_FOUND, "The stored document file is missing.");
            }
            return Files.readAllBytes(file);
        } catch (final IOException exception) {
            throw unavailable(exception);
        }
    }

    @Override
    public void remove(final String path) {
        try {
            Files.deleteIfExists(resolve(path));
        } catch (final IOException exception) {
            throw unavailable(exception);
        }
    }

    @Override
    public boolean stageDeletion(final String path) {
        final Path source = resolve(path);
        final Path staged = resolve(path + ".deleting");
        try {
            // An earlier interrupted deletion can be retried safely using its retained file.
            if (Files.exists(staged, LinkOption.NOFOLLOW_LINKS)) {
                if (Files.exists(source, LinkOption.NOFOLLOW_LINKS)) {
                    throw new IOException("Both original and staged deletion exist");
                }
                return true;
            }
            if (!Files.exists(source, LinkOption.NOFOLLOW_LINKS)) {
                return false; // Allow removal of a registry row whose file is already absent.
            }
            Files.move(source, staged);
            return true;
        } catch (final IOException exception) {
            throw unavailable(exception);
        }
    }

    @Override
    public void restoreDeletion(final String path) {
        try {
            Files.move(resolve(path + ".deleting"), resolve(path));
        } catch (final IOException exception) {
            throw unavailable(exception);
        }
    }

    @Override
    public void completeDeletion(final String path) {
        remove(path + ".deleting");
    }

    private Path resolve(final String relative) {
        final Path path = root.resolve(relative).normalize();
        if (!path.startsWith(root) || path.equals(root)) {
            throw DocumentException.invalid("Invalid document storage path.");
        }
        // Storage is service-owned. Refuse symlinks instead of following them outside the volume.
        for (Path current = path; !current.equals(root); current = current.getParent()) {
            if (Files.isSymbolicLink(current)) {
                throw DocumentException.invalid("Symbolic links are not allowed in document storage.");
            }
        }
        return path;
    }

    private static DocumentException unavailable(final IOException cause) {
        return new DocumentException(HttpStatus.SERVICE_UNAVAILABLE, "Document storage is unavailable.", cause);
    }
}
