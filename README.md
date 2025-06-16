# Pixl - Pixel Art Editor

Pixl is a modern pixel art editor built with Go and the Fyne toolkit. It provides a simple yet powerful interface for creating and editing pixel art.

## Features

- 🎨 Pixel-perfect drawing with customizable brush colors
- 🖌️ Color picker with hue-based color selection
- 🎯 Color swatches for quick color access
- 📁 File operations (New, Open, Save, Save As)
- 🖼️ Support for PNG image format
- 🔍 Canvas zooming with mouse wheel
- 🖱️ Pan canvas with right mouse button
- 📏 Pixel grid visualization
- 🎯 Precise pixel placement with cursor guides

## Requirements
Note: Currently, this project is configured to run in WSL Ubuntu environment. For running on native Windows or other platforms, additional setup and configuration may be required.

### WSL Ubuntu Requirements
- Windows 10/11 with WSL2 installed
- Ubuntu on WSL2
- Go 1.24 or later
- Required system dependencies:
  ```bash
  sudo apt-get update
  sudo apt-get install gcc libgl1-mesa-dev xorg-dev
  ```

## Run Project (require WSL Ubuntu)

1. Clone the repository:
```bash
git clone https://github.com/teeratem/pixl.git
cd pixl
```

2. Install dependencies:
```bash
go mod download
```

3. Run the application:
```bash
LANG=en_US.UTF-8 go run ./pixl 
```

## Usage

### Basic Controls

- **Left Mouse Button**: Draw pixels
- **Right Mouse Button**: Pan canvas
- **Mouse Wheel**: Zoom in/out
- **Color Picker**: Select colors using the hue-based picker on the right
- **Color Swatches**: Quick access to frequently used colors

### File Operations

- **New**: Create a new canvas with custom dimensions
- **Open**: Load an existing PNG image
- **Save**: Save the current image (prompts for filename if new)
- **Save As**: Save the current image with a new filename

### Canvas Features

- The canvas is centered in the main window
- Color swatches are displayed at the bottom
- Color picker is available on the right side
- The canvas can be panned and zoomed for detailed work
- Pixel grid is visible for precise placement

## Project Structure

```
pixl/
├── apptype/         # Application types and interfaces
├── pxcanvas/        # Canvas implementation
│   └── brush/       # Brush tools and cursor
├── swatch/          # Color swatch implementation
├── ui/              # User interface components
├── util/            # Utility functions
└── main.go          # Application entry point
```

## Dependencies

- [fyne.io/fyne/v2](https://github.com/fyne-io/fyne) - GUI toolkit
- [github.com/lusingander/colorpicker](https://github.com/lusingander/colorpicker) - Color picker component

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
