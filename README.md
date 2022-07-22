# Google Group Query

Queries Google's Directory Admin SDK for Google Groups.

**WARNING** This is not meant for production use.

## Usage

Query all groups:

```bash
ggquery
```

Query groups by email:

```bash
ggquery email=mygroup@google.com
```

Query groups by name:

```bash
ggquery "name='My group'"
```

Query groups by wildcard:

```bash
ggquery email:aws\*
```
