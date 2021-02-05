
# PasteBin
#### Purpose of this application is solely for learning, praticing distributed systems, and new technologies.

## What
Application that offers the same functionality as [pastebin](https://pastebin.pl/) or [dpaste](https://dpaste.com/).
Usecase is to share a plain text (**Paste**) and get a short link/identifier (**8-character**) to access it publicly.

## How
Application will be consist of following **needs**. These "needs" can be implemented as seperate microservices.
- **Key generator**, which will generate unique keys upon requested. It should also have *"batch generate"* functionality.
- **Storage**, which will keep the *Pastes*. Text size of pastes can be different in terms of size, so, we might have multiple storages optimized for different usecases.