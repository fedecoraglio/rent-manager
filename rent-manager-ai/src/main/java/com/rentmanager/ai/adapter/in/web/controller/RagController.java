package com.rentmanager.ai.adapter.in.web.controller;

import com.rentmanager.ai.adapter.in.web.request.RagRequest;
import com.rentmanager.ai.adapter.in.web.response.RagResponse;
import com.rentmanager.ai.application.rag.RagService;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/properties/{propertyId}/contracts/{contractId}/chat")
public final class RagController {

    private final RagService ragService;

    public RagController(final RagService ragService) {
        this.ragService = ragService;
    }

    @PostMapping
    public ResponseEntity<RagResponse> ask(
            @PathVariable final String propertyId,
            @PathVariable final String contractId,
            @RequestBody final RagRequest request) {

        final String answer = ragService.ask(
                propertyId,
                contractId,
                request.question()
        );

        return ResponseEntity.ok(new RagResponse(answer));
    }
}