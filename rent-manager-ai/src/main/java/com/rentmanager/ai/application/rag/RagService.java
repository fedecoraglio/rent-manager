package com.rentmanager.ai.application.rag;

import java.util.stream.Collectors;

import com.rentmanager.ai.port.out.DocumentRetriever;
import org.springframework.ai.chat.client.ChatClient;
import org.springframework.ai.chat.model.ChatModel;
import org.springframework.ai.ollama.OllamaChatModel;
import org.springframework.ai.ollama.api.OllamaChatOptions;
import org.springframework.stereotype.Service;

@Service
public final class RagService {

    private final DocumentRetriever documentRetriever;
    private final ChatClient chatClient;

    public RagService(
            final DocumentRetriever documentRetriever,
            final ChatClient chatClient) {
        this.documentRetriever = documentRetriever;
        this.chatClient = chatClient;
    }

    public String ask(
            final String propertyId,
            final String contractId,
            final String question) {

        final var chunks = documentRetriever.retrieve(
                propertyId,
                contractId,
                question
        );

        if (chunks.isEmpty()) {
            return "No relevant information was found in the contract documents.";
        }

        final String context = chunks.stream()
                .map(chunk -> """
                        Source: %s
                        Page: %s
                        Content:
                        %s
                        """.formatted(
                        chunk.filename(),
                        chunk.page() != null ? chunk.page() : "unknown",
                        chunk.text()
                ))
                .collect(Collectors.joining("\n---\n"));

        return chatClient.prompt()
                .system("""
                        You are an assistant for a rental property management application.
                        Answer only using the information provided in the document context.
                        Do not invent dates, amounts, names, clauses or contractual conditions.
                        If the answer cannot be found in the context, say:
                        "I could not find that information in the contract documents."
                        Keep the answer concise and factual.
                        """)
                .user("""
                        DOCUMENT CONTEXT:
                        %s
                        QUESTION:
                        %s
                        """.formatted(context, question))
                .call()
                .content();
    }
}