---
name: rent-manager-ui-runtime
description: Work on Angular startup, runtime API configuration, language, navigation, global loading, notifications, and HTTP errors in Rent Manager.
---

# Shared UI flows

Read [coverage](references/coverage.md). UI root: `rent-manager-frontend/projects/rent-manager-frontend/src/app`.

## Startup and configuration

main.ts bootstraps the app with app.config.ts. provideAppInitializer waits for AppConfigService.load(), which fetches `/assets/config.json` through HttpClient and stores apiBaseUrl in a signal. Without loaded configuration, the getter throws; this service has no fallback. ApiUrlService builds business URLs from that base.

In Kubernetes, a ConfigMap mounts config.json with `/v1`. During development, read the actual asset; do not assume environment.ts is the effective source. A frontend-only port-forward does not route `/v1`: Nginx only serves the SPA, while Ingress provides the API proxy.

## HTTP, state, and errors

Interceptors are registered in this order: auth, apiError, loading. Add a token when present; apiError displays error.message from the response body or Unexpected error and rethrows. Some Go errors use `error` instead of `message` and appear as generic errors. There is no automatic renewal on 401.

Read loading.interceptor.ts/loading.service.ts for the counter and finalize; preserve balancing for completed, failed, and canceled requests. Feature services also maintain their own loading/selection state. Do not replace them with global flags without reviewing concurrency. NotificationService supplies the notification component.

## Languages and navigation

ngx-translate loads `/assets/i18n/*.json`, with initial/fallback language es. LanguageService supports es/en and persists `rent_manager_language` in localStorage; initialize retrieves it or uses es. Source assets are in `projects/rent-manager-frontend/public/assets`, copied through angular.json. Follow MainLayout for initialization and language controls. Add keys to existing languages; do not mix hardcoded text with translations without reviewing feature conventions. AuthLayout hosts login; MainLayout hosts administrative navigation, logout, and shared UI.

Check startup with valid/missing config, deep routes through Nginx, concurrent requests, errors/cancellations, API 401 responses, and language changes. Do not interpret index.html returned for configuration or API requests as valid JSON.
