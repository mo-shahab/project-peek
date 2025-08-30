---

## **Core Command Ideas for `ppeek`**

### 1. **`ppeek tree`** *(Already done!)*

* Visualize directory trees with skips (`.git`, `node_modules`, etc.).
* Options: depth limit (`-L`), include/exclude patterns, file size summary.

---

### 2. **`ppeek last-updated`**

* Show when the project (or each major directory) was last updated.
* Could scan `git log` or file modification times.
* Example:

  ```
  ppeek last-updated
  ├── src/       2 days ago
  ├── internal/  1 week ago
  └── README.md  3 hours ago
  ```

---

### 3. **`ppeek size`**

* Show **disk usage by folder/file** (like a mini `du -sh`).
* Could be visual (bars) or tabular.
* Example:

  ```
  ppeek size
  50M  node_modules/
  12M  src/
  1.2M internal/
  ```

---

### 4. **`ppeek deps`**

* List dependencies (Go modules, NPM packages, Pip requirements, etc.).
* Could be language-aware (Go → `go list -m all`, JS → `package.json`).

---

### 5. **`ppeek git-stats`**

* Show Git-related project insights:

  * Last commit
  * Number of contributors
  * Most active files
  * Total commits

---

### 6. **`ppeek find <pattern>`**

* Quick project-wide search (by name or content).
* Could respect `.gitignore` and skip huge dirs automatically.

---

### 7. **`ppeek health`**

* Run a set of **quick sanity checks**:

  * Is `.gitignore` missing key files?
  * Are there untracked/dirty changes?
  * Is `go.mod` tidy? (`go mod tidy -diff`)
  * Any suspicious big files committed?

---

### 8. **`ppeek todo`**

* Collect all TODO/FIXME/NOTE comments across the project.
* Example:

  ```
  ppeek todo
  [src/main.go:42] // TODO: handle error
  [internal/api.go:87] // FIXME: race condition
  ```

---

### 9. **`ppeek license`**

* Show or verify the license of the project.
* Could warn if missing.

---

### 10. **`ppeek init`**

* Project initializer:

  * Creates a standard folder structure (`cmd/`, `internal/`, `pkg/`, `README`, `.gitignore`).
  * Optionally initializes Git.

---

## **Design Philosophy for `ppeek`**

* **Developer-first, scripting-friendly**: all commands should be usable in CI scripts or aliases.
* **One binary, many subcommands**: like `git`, `kubectl`, `docker`.
* **Configurable defaults**: allow `.ppeek.yml` for skip lists, default flags, etc.
* **Cross-language**: not just Go — aim for "project-level" tools, not language-specific ones only.

---

## **Near-term Roadmap**

1. Start with **`tree` + `last-updated` + `size`** — these cover 80% of dev needs.
2. Add **`todo` and `git-stats`** for quick insights.
3. Introduce **config file support** for custom skiplists and patterns.

---


