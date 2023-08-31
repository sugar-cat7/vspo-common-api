# VSPO Common API

This project provides common API components for the unofficial VSPO fan site.

Sites:

- [すぽじゅーる](https://www.vspo-schedule.com/)

other Repositories:

- [すぽじゅーる](https://github.com/sugar-cat7/vspo-schedule-web)

## Project Structure

- `app/`: Application-specific code, including dependency injection and HTTP handlers.
- `constants/`: Constant values used throughout the application.
- `domain/`: Domain-related code, including entities, repositories, and services.
- `infrastructure/`: Code related to external infrastructure, such as Firestore and YouTube.
- `mocks/`: Mock files for testing.
- `usecases/`: Use case definitions for the application.
- `util/`: Utility functions.
- `config/`: Configuration-related code.

### quick start(local)

- setting `.env` file

```bash
YOUTUBE_API_KEY=test-api-key
FIRESTORE_EMULATOR_HOST=localhost:8081
```

- start firestore emulator

```bash
make firestore
```

- start server

```bash
make run
```
