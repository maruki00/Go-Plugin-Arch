# Plugin System Instructions

This document provides instructions for building, adding, and removing plugins in the system.

---

## Building a Plugin

To build a plugin, use the following command in your terminal:

```bash
go build -buildmode=plugin -o plugName.so .
```

- Replace `plugName` with the desired name of your plugin.
- The resulting `.so` file will be the plugin.

---

## Adding a Plugin to the System

1. Build your plugin using the above command.
2. Move the `.so` plugin file to the following folder:
   
   ```
   core/plugins
   ```

3. The plugin will now be active in the system.

---

## Removing a Plugin from the System

1. To deactivate a plugin, move the plugin file from the `core/plugins` folder to the following folder:
   
   ```
   core/trash
   ```

2. The plugin will no longer be active.

---

Ensure the plugin files are handled carefully to avoid issues with the system.


