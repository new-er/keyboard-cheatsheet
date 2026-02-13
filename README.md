# keyboard-cheatsheet

A small Go CLI tool that prints a live keyboard cheatsheet in your terminal. It watches your current key presses and active window (via showmethekey and hyprctl) and prioritizes relevant shortcuts, so you always see what matters right now.

## Features
- Shows active keyboard shortcuts based on current keys and application
- Orders shortcuts by context using showmethekey and hyprctl
- Works in Hyprland Wayland environments
- Lightweight terminal UI

## Prerequisites
Before installing, make sure you have:
- [Hyprland](https://github.com/hyprwm/Hyprland)
- [showmethekey](https://github.com/AlynxZhou/showmethekey)

## Installation
You can install the latest release with:
```bash
go install github.com/new-er/keyboard-cheatsheet@latest
```

## Usage

```bash
keyboard-cheatsheet
```

## Customization

### Key mapping
You can change how key codes map to key names by editing `keymap.json`.

### Shortcuts
You can define or adjust shortcuts in `shortcuts.json`.

This lets you personalize the cheatsheet to your workflow.

## Contributing
Contributions are welcome! Please feel free to open an issue or submit a pull request.
