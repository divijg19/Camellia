# Frontend Architecture (Jaspr + Flutter-ready)

This folder is now a Dart workspace that supports:
- Jaspr web app (`apps/jaspr_web`)
- Flutter app shell (`apps/flutter_app`)
- Shared domain layer for both (`packages/shared_core`)

## Why this structure
- Keep product logic in pure Dart (`shared_core`) so both Jaspr and Flutter can reuse it.
- Let server-rendered Go (`backend/webui`) handle SSR + HTMX interactions now.
- Allow a progressive shift to richer Jaspr/Flutter clients later without duplicating models/use-cases.

## Quick start
1. Install Dart SDK and Jaspr CLI.
2. From `frontend/dart_workspace/apps/jaspr_web`, run `dart pub get` then `jaspr serve`.
3. From `frontend/dart_workspace/apps/flutter_app`, run `flutter pub get` then `flutter run`.
